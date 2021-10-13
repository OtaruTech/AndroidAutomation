package automation

import (
	"log"
	"reflect"
	"strings"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/automation"
	"github.com/flipped-aurora/gin-vue-admin/server/rpc/message"
	"github.com/flipped-aurora/gin-vue-admin/server/rpc/messagedaemon"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/robfig/cron/v3"
)

type TestResult struct {
	Id             int      `json:"id"`
	SerialNo       string   `json:"serialNo"`
	Result         []string `json:"result"`
	LogcatFileName string   `json:"logcat"`
	BuildId        string   `json:"buildId"`
}

type TestRunner struct {
	Id        int      `json:"id"`
	SerialNo  string   `json:"serialNo"`
	Testcases []string `json:"testcases"`
	Completed []string `json:"completed"`
}

type RuntimeState struct {
	SerialNo     string `json:"serialNo"`
	State        int    `json:"state"`
	TestcaseName string `json:"testcaseName"`
	Duration     int    `json:"duration"`
}

type RemoteHost struct {
	remoteService *message.RemoteService
	devices       []string
}

type JobScheduler struct {
	cron *cron.Cron
	job  autocode.Job
}

func (scheduler *JobScheduler) timerFunc() {
	lock.Lock()
	log.Println("automation: daily job", scheduler.job.Name)
	var otaUrl string
	var serialNo string
	var testcases []string
	var timeout int
	if scheduler.job.OtaUrl == "" {
		var req = automation.LatestOtaRequest{
			OtaPath:    scheduler.job.OtaPath,
			OtaFormat:  scheduler.job.OtaFormat,
			FileFormat: scheduler.job.FileFormat,
		}
		_, response := androidService.GetLatestOTA(req)
		otaUrl = response.LatestOtaUrl
	} else {
		otaUrl = scheduler.job.OtaUrl
	}
	testcases = strings.Split(scheduler.job.Testcases, ",")
	_, deviceList := getDeviceList()
	var findDevice = false
	for _, dev := range deviceList {
		ret, runtimeState := AndroidGetRuntimeState(dev.SerialNo)
		log.Println("require:", scheduler.job.Product)
		log.Println("serialNo:", dev.SerialNo, "product:", dev.Product, "state:", runtimeState.State, "ret:", ret)
		if scheduler.job.Product == dev.Product && runtimeState.State == 0 && ret == 0 {
			serialNo = dev.SerialNo
			findDevice = true
			break
		}
	}
	if !findDevice {
		log.Println("automation: cannot find product", scheduler.job.Product)
	} else {
		timeout = 0
		id := AndroidRunTestcase(serialNo, testcases, timeout, otaUrl)
		log.Println("automation: testId", id)
		var testRunner autocode.TestRunner
		testRunner.TestId = &id
		testRunner.Testcases = scheduler.job.Testcases
		testRunner.Name = scheduler.job.Name
		testRunner.Owner = scheduler.job.Owner
		testRunner.SerialNo = serialNo
		if err := testRunnerService.CreateTestRunner(testRunner); err != nil {
			log.Println("automation: failed to create test runner", err)
		}
	}
	lock.Unlock()
}

var (
	remoteHostMap   map[string]*RemoteHost
	localService    *message.LocalService
	listener        *message.MessageBusListener
	testIdMap       map[int]string
	jobSchedulerMap map[string]*JobScheduler
	lock            sync.Mutex
)

var deviceService = service.ServiceGroupApp.AutoCodeServiceGroup.DeviceService
var reportService = service.ServiceGroupApp.AutoCodeServiceGroup.ReportService
var consoleService = service.ServiceGroupApp.AutoCodeServiceGroup.DeviceConsoleService
var jobService = service.ServiceGroupApp.AutoCodeServiceGroup.JobService
var androidService = service.ServiceGroupApp.AutomationServiceGroup.AndroidService
var testRunnerService = service.ServiceGroupApp.AutoCodeServiceGroup.TestRunnerService

var clientAvailableHandle message.ServiceAvailableHandle = func(clientId string, service string) {
	log.Println("automation: active", clientId, "->", service)
	if strings.HasPrefix(service, "Android-") && remoteHostMap[service] == nil {
		go func() {
			log.Println("automation: add host", service)
			remoteService := message.GetService(service)
			remoteHost := &RemoteHost{remoteService: remoteService, devices: make([]string, 0)}
			remoteHostMap[service] = remoteHost
		}()
	}
}

var clientLostHandle message.ServiceLostHandle = func(clientId string, services []string) {
	log.Println("automation: lost", clientId, "->", services)
	for _, v := range services {
		if strings.HasPrefix(v, "Android-") {
			log.Println("automation: remove host", v)
			if remoteHostMap[v] != nil {
				delete(remoteHostMap, v)
			}
		}
	}
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

var deviceStateChangeHandle message.MethodHandle = func(m message.Message) message.Message {
	log.Printf("%s device: [%s]%s - %s - %s", If(m.GetBool("state", false), "New", "Remove"),
		m.GetString("host", ""),
		m.GetString("serialNo", "Unknown"),
		m.GetString("model", ""), m.GetString("product", ""))
	host := m.GetString("host", "")
	serialNo := m.GetString("serialNo", "")
	remoteHost := remoteHostMap[host]
	if remoteHost == nil {
		log.Println("automation: invalid host", host)
		return message.Message{}
	}
	if !m.GetBool("state", false) {
		//remote serialNo in devices
		var exist = false
		var index int
		for k, v := range remoteHost.devices {
			if v == serialNo {
				exist = true
				index = k
				break
			}
		}
		if exist {
			remoteHost.devices = append(remoteHost.devices[:index], remoteHost.devices[index+1:]...)
		} else {
			log.Println("automation: serialNo", serialNo, "not exist in host", host)
		}
		return message.Message{}
	} else {
		//add serialNo in devices
		var exist = false
		for _, v := range remoteHost.devices {
			if v == serialNo {
				exist = true
				break
			}
		}
		if !exist {
			remoteHost.devices = append(remoteHost.devices, serialNo)
		} else {
			log.Println("automation: serialNo", serialNo, "already exist in host", host)
		}
	}
	//query the database
	err, list := getDeviceList()
	if err != nil {
		log.Println("automation: failed to get device list", err)
		return message.Message{}
	}
	var exist bool
	exist = false
	var device *autocode.Device
	if len(list) > 0 {
		for _, v := range list {
			log.Println(v.SerialNo, "-", v.BuildId)
			if v.SerialNo == m.GetString("serialNo", "Unknown") {
				log.Println("automation: device already exist")
				exist = true
				device = &v
				break
			}
		}
	}
	if !exist {
		//add device into database
		addDevice(&m)
	} else {
		updateDevice(device, &m)
	}

	return message.Message{}
}

var uploadTestResultHandle message.MethodHandle = func(m message.Message) message.Message {
	id := m.GetInt("id", -1)
	serialNo := m.GetString("serialNo", "")
	__results := m.GetStringArray("result")
	logcat := m.GetString("logcat", "")
	buildId := m.GetString("buildId", "")
	results := make([]string, 0)
	for _, v := range __results {
		if len(v) > 10 {
			results = append(results, v)
		}
	}
	testResult := &TestResult{
		Id:             id,
		SerialNo:       serialNo,
		Result:         results,
		LogcatFileName: logcat,
		BuildId:        buildId}
	addResult(testResult)
	return message.Message{}
}

var uploadConsoleHandle message.MethodHandle = func(m message.Message) message.Message {
	serialNo := m.GetString("serialNo", "")
	console := m.GetString("console", "")
	_, consoleList := getConsoleList()
	exist := false
	var deviceConsole *autocode.DeviceConsole
	deviceConsole = nil
	for _, v := range consoleList {
		if v.SerialNo == serialNo {
			exist = true
			deviceConsole = &v
		}
	}
	if !exist {
		//add console log
		addDeviceConsole(serialNo, console)
	} else {
		//update console log
		updateDeviceConsole(deviceConsole, console)
	}
	// log.Println(serialNo + ":", console)
	//save console log
	return message.Message{}
}

func destroyJobs() {
	for key, scheduler := range jobSchedulerMap {
		if scheduler != nil {
			scheduler.cron.Stop()
		}
		delete(jobSchedulerMap, key)
	}
}

func createJobs() {
	err, jobList := getJobList()
	if err != nil {
		log.Println("automation: failed to get job list")
		return
	}
	log.Println("automation: job list", jobList)
	for _, job := range jobList {
		jobSchedulerMap[job.Name] = &JobScheduler{
			job: job,
			cron: cron.New(),
		}
		// spec := fmt.Sprintf("0 %d * * ?", *job.Hour)
		spec := "@every 10m"
		jobSchedulerMap[job.Name].cron.AddFunc(spec, jobSchedulerMap[job.Name].timerFunc)
		if *job.Enable {
			log.Println("automation: start job", job.Name)
			jobSchedulerMap[job.Name].cron.Start()
		}
	}
}

func Initialize() {
	log.Println("automation: Android Automation Start!!")
	remoteHostMap = make(map[string]*RemoteHost)
	testIdMap = make(map[int]string)
	jobSchedulerMap = make(map[string]*JobScheduler)
	localService = nil
	messagedaemon.InitDaemon()
	message.InitMessage()
	listener = &message.MessageBusListener{
		AvailableHandle: clientAvailableHandle, LostHandle: clientLostHandle}
	message.RegisterBusListener(*listener)
	//initialize services
	localService = message.CreateLocalService("server")
	if localService != nil {
		localService.RegisterMethod("deviceStateChange", deviceStateChangeHandle)
		localService.RegisterMethod("uploadResult", uploadTestResultHandle)
		localService.RegisterMethod("uploadConsole", uploadConsoleHandle)
		message.AddService(localService)
	} else {
		log.Fatalln("automation:", "cannot create local service")
	}
	//start crontab to schedule the job
	createJobs()
}

func getConsoleList() (err error, consoleList []autocode.DeviceConsole) {
	var pageInfo autocodeReq.DeviceConsoleSearch
	pageInfo.Page = 0
	pageInfo.PageSize = 100
	err, list, _ := consoleService.GetDeviceConsoleInfoList(pageInfo)
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		s := reflect.ValueOf(list)
		for i := 0; i < s.Len(); i++ {
			ele := s.Index(i)
			consoleList = append(consoleList, ele.Interface().(autocode.DeviceConsole))
		}
	}
	return
}

func getDeviceList() (err error, deviceList []autocode.Device) {
	var pageInfo autocodeReq.DeviceSearch
	pageInfo.Page = 0
	pageInfo.PageSize = 100
	err, list, _ := deviceService.GetDeviceInfoList(pageInfo)
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		s := reflect.ValueOf(list)
		for i := 0; i < s.Len(); i++ {
			ele := s.Index(i)
			deviceList = append(deviceList, ele.Interface().(autocode.Device))
		}
	}
	return
}

func getJobList() (err error, jobList []autocode.Job) {
	var pageInfo autocodeReq.JobSearch
	pageInfo.Page = 0
	pageInfo.PageSize = 100
	err, list, _ := jobService.GetJobInfoList(pageInfo)
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		s := reflect.ValueOf(list)
		for i := 0; i < s.Len(); i++ {
			ele := s.Index(i)
			jobList = append(jobList, ele.Interface().(autocode.Job))
		}
	}
	return
}

func addDeviceConsole(serialNo string, console string) {
	var deviceConsole autocode.DeviceConsole
	deviceConsole.SerialNo = serialNo
	deviceConsole.Console = console
	if err := consoleService.CreateDeviceConsole(deviceConsole); err != nil {
		log.Println("automation: failed to create device console", err)
	}
}

func updateDeviceConsole(deviceConsole *autocode.DeviceConsole, console string) {
	var newDeviceConsole autocode.DeviceConsole
	newDeviceConsole = *deviceConsole
	newDeviceConsole.SerialNo = deviceConsole.SerialNo
	newDeviceConsole.Console = console
	newDeviceConsole.ID = deviceConsole.ID
	if err := consoleService.UpdateDeviceConsole(newDeviceConsole); err != nil {
		log.Println("automation: failed to update device console", err)
	}
}

func addDevice(msg *message.Message) {
	var device autocode.Device
	device.SerialNo = msg.GetString("serialNo", "")
	device.Model = msg.GetString("model", "")
	device.Product = msg.GetString("product", "")
	device.Version = msg.GetString("version", "")
	if len(msg.GetString("buildId", "")) > 32 {
		device.BuildId = string([]byte(msg.GetString("buildId", ""))[:32])
	} else {
		device.BuildId = msg.GetString("buildId", "")
	}
	device.BuildType = msg.GetString("buildType", "")
	device.Sku = msg.GetString("sku", "")
	device.Sdk = msg.GetString("sdk", "")
	device.Security = msg.GetString("security", "")
	device.Api = msg.GetString("api", "")
	device.Config = msg.GetString("config", "")
	device.Camera = msg.GetString("camera", "")
	device.Scanner = msg.GetString("scanner", "")
	device.Wwan = msg.GetString("wwan", "")
	if err := deviceService.CreateDevice(device); err != nil {
		log.Println("automation: failed to create device", err)
	}
}

func updateDevice(device *autocode.Device, msg *message.Message) {
	var newDevice autocode.Device
	newDevice = *device
	newDevice.SerialNo = msg.GetString("serialNo", "")
	newDevice.Model = msg.GetString("model", "")
	newDevice.Product = msg.GetString("product", "")
	newDevice.Version = msg.GetString("version", "")
	if len(msg.GetString("buildId", "")) > 32 {
		newDevice.BuildId = string([]byte(msg.GetString("buildId", ""))[:32])
	} else {
		newDevice.BuildId = msg.GetString("buildId", "")
	}
	newDevice.BuildType = msg.GetString("buildType", "")
	newDevice.Sku = msg.GetString("sku", "")
	newDevice.Sdk = msg.GetString("sdk", "")
	newDevice.Security = msg.GetString("security", "")
	newDevice.Api = msg.GetString("api", "")
	newDevice.Config = msg.GetString("config", "")
	newDevice.Camera = msg.GetString("camera", "")
	newDevice.Scanner = msg.GetString("scanner", "")
	newDevice.Wwan = msg.GetString("wwan", "")
	if err := deviceService.UpdateDevice(newDevice); err != nil {
		log.Println("automation: failed to update device", err)
	}
}

func addResult(result *TestResult) {
	var report autocode.Report
	report.TestId = &result.Id
	report.Result = strings.Join(result.Result, "|||")
	report.Logcat = result.LogcatFileName
	report.BuildId = result.BuildId
	if err := reportService.CreateReport(report); err != nil {
		log.Println("automation: failed to create report", err)
	}
}

func AndroidRunTestcase(serialNo string, testcases []string, timeout int, otaUrl string) int {
	log.Println("automation:", serialNo, "->", testcases)
	remoteService := getService(serialNo)
	if remoteService == nil {
		return -1
	}
	var msg message.Message
	msg.SetString("serialNo", serialNo)
	msg.SetStringArray("testcases", testcases)
	msg.SetInt("timeout", timeout)
	msg.SetString("otaUrl", otaUrl)
	result := remoteService.CallMethod("runTestcase", msg)
	log.Println("automation: runTestcase", result)
	testId := result.GetInt("testId", -1)
	testIdMap[testId] = serialNo
	if result.Empty() {
		return -1
	}
	return result.GetInt("testId", -1)
}

func AndroidGetRuntimeState(serialNo string) (ret int, runtimeState RuntimeState) {
	remoteService := getService(serialNo)
	if remoteService == nil {
		ret = -1
		return
	}
	var msg message.Message
	msg.SetString("serialNo", serialNo)
	result := remoteService.CallMethod("getRuntimeState", msg)
	if result.Empty() {
		ret = -1
		return
	}
	ret = 0
	runtimeState.SerialNo = result.GetString("serialNo", "")
	runtimeState.State = result.GetInt("state", 2)
	runtimeState.TestcaseName = result.GetString("testcaseName", "")
	runtimeState.Duration = result.GetInt("duration", 0)
	return
}

func AndroidGetTestRunnerState(id int) (ret int, runner TestRunner) {
	remoteService := getService(testIdMap[id])
	if remoteService == nil {
		ret = -1
		return
	}
	var msg message.Message
	msg.SetInt("testId", id)
	result := remoteService.CallMethod("getTestRunnerState", msg)
	runner.Id = result.GetInt("id", -1)
	runner.SerialNo = result.GetString("serialNo", "")
	runner.Testcases = result.GetStringArray("testcases")
	runner.Completed = result.GetStringArray("completed")
	return
}

func getService(serialNo string) *message.RemoteService {
	for _, host := range remoteHostMap {
		for _, sn := range host.devices {
			if sn == serialNo {
				return host.remoteService
			}
		}
	}
	return nil
}

func JobChanged(request *automation.JobChangedReq) error {
	log.Println("automation:", request)
	var newJob autocode.Job
	_, job := jobService.GetJob(uint(request.Id))
	newJob = job
	newJob.Enable = &request.Enable
	if err := jobService.UpdateJob(newJob); err != nil {
		log.Println("automation: failed to update job", err)
		return err
	}
	//restart the jobs
	destroyJobs()
	createJobs()
	// cron := schedulerMap[newJob.Name]
	// if cron == nil {
	// 	return errors.New("cannot find scheduler")
	// }
	// if request.Enable {
	// 	log.Println("automation: start scheduler", newJob.Name)
	// 	cron.Start()
	// } else {
	// 	log.Println("automation: stop scheduler", newJob.Name)
	// 	cron.Stop()
	// }
	return nil
}
