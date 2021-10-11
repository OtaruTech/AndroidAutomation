package automation

import (
	"log"
	"reflect"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/rpc/message"
	"github.com/flipped-aurora/gin-vue-admin/server/rpc/messagedaemon"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type TestResult struct {
	Id             int      `json:"id"`
	SerialNo       string   `json:"serialNo"`
	Result         []string `json:"result"`
	LogcatFileName string   `json:"logcat"`
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

var (
	remoteHostMap map[string]*RemoteHost
	localService  *message.LocalService
	listener      *message.MessageBusListener
	testIdMap     map[int]string
)

var deviceService = service.ServiceGroupApp.AutoCodeServiceGroup.DeviceService
var reportService = service.ServiceGroupApp.AutoCodeServiceGroup.ReportService
var consoleService = service.ServiceGroupApp.AutoCodeServiceGroup.DeviceConsoleService

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
	var id uint
	if len(list) > 0 {
		for _, v := range list {
			log.Println(v.SerialNo, "-", v.BuildId)
			if v.SerialNo == m.GetString("serialNo", "Unknown") {
				log.Println("automation: device already exist")
				exist = true
				id = v.ID
				break
			}
		}
	}
	if !exist {
		//add device into database
		addDevice(&m)
	} else {
		updateDevice(id, &m)
	}

	return message.Message{}
}

var uploadTestResultHandle message.MethodHandle = func(m message.Message) message.Message {
	id := m.GetInt("id", -1)
	serialNo := m.GetString("serialNo", "")
	__results := m.GetStringArray("result")
	logcat := m.GetString("logcat", "")
	results := make([]string, 0)
	for _, v := range __results {
		if len(v) > 10 {
			results = append(results, v)
		}
	}
	testResult := &TestResult{Id: id, SerialNo: serialNo, Result: results, LogcatFileName: logcat}
	// log.Println("automation:", testResult)
	//store the test result
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
		if(v.SerialNo == serialNo) {
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

func Initialize() {
	log.Println("automation: Android Automation Start!!")
	remoteHostMap = make(map[string]*RemoteHost)
	testIdMap = make(map[int]string)
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

func updateDevice(id uint, msg *message.Message) {
	var device autocode.Device
	device.ID = id
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
	if err := deviceService.UpdateDevice(device); err != nil {
		log.Println("automation: failed to update device", err)
	}
}

func addResult(result *TestResult) {
	var report autocode.Report
	report.TestId = &result.Id
	report.Result = strings.Join(result.Result, "|||")
	report.Logcat = result.LogcatFileName
	if err := reportService.CreateReport(report); err != nil {
		log.Println("automation: failed to create report", err)
	}
	log.Println("automation: addResult", report)
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
	log.Println("automation: getRuntimeState", result)
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
	log.Println("automation: getRuntimeState", result)
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