package bridge

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"machine/pkg/message"
	"machine/pkg/utils"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/go-basic/uuid"
	adb "github.com/zach-klippenstein/goadb"
)

type Device struct {
	device       *adb.Device
	console      string
	jobChan      chan *TestRunner
	busy         bool
	startTime    int64
	testcaseName string
	serialNo     string
	model        string
	product      string
	DeviceProp
}

type DeviceProp struct {
	version   string
	buildId   string
	buildType string
	sku       string
	sdk       string
	security  string
	api       string
	config    string
	camera    string
	scanner   string
	wwan      string
}

type TestRunner struct {
	Id        int      `json:"id"`
	SerialNo  string   `json:"serialNo"`
	Testcases []string `json:"testcases"`
	Completed []string `json:"completed"`
	OtaUrl    string   `json:"otaUrl"`
}

type TestResult struct {
	id             int
	serialNo       string
	result         []string
	logcatFileName string
	buildId        string
}

var (
	port            = flag.Int("p", adb.AdbPort, "")
	client          *adb.Adb
	androidService  *message.LocalService
	deviceMap       map[string]*Device
	remoteService   *message.RemoteService
	testRunnerMap   map[int]*TestRunner
	testResultMap   map[int]*TestResult
	forgetDeviceMap map[string]*Device
	uid             string
	remoteHost      string
)

func Initialize(host string) {
	remoteHost = host
	var err error
	client, err = adb.NewWithConfig(adb.ServerConfig{
		Port: *port,
	})
	if err != nil {
		log.Fatal(err)
	}
	client.StartServer()

	deviceMap = make(map[string]*Device)
	forgetDeviceMap = make(map[string]*Device)
	testRunnerMap = make(map[int]*TestRunner)
	testResultMap = make(map[int]*TestResult)
	go listenDevices()
}

func listenDevices() {
	log.Println("Watching for device state changes.")
	watcher := client.NewDeviceWatcher()
	for event := range watcher.C() {
		serial := event.Serial
		state := event.NewState
		if forgetDeviceMap[serial] != nil {
			log.Println("forget the device:", serial)
			continue
		}
		// log.Printf("\t[%s] %s - %s\n", time.Now(), serial, state)
		if state == adb.StateOnline {
			device := client.Device(adb.DeviceWithSerial(serial))
			devInfo, err := device.DeviceInfo()
			if err != nil {
				log.Println("failed to get device info:", err.Error())
				continue
			}
			var dev Device
			dev.device = device
			dev.jobChan = make(chan *TestRunner, 0)
			dev.serialNo = serial
			dev.model = devInfo.Model
			dev.product = devInfo.Product
			dev.busy = false
			deviceMap[serial] = &dev
			dev.getDeviceProp()
			log.Printf("New device %s - %s - %s", dev.serialNo, dev.model, dev.product)
			//create thread to process the testcases
			go processTestcase(&dev)
			if remoteService != nil {
				var param message.Message
				param.SetBool("state", true)
				param.SetString("host", "Android-"+uid)
				param.SetString("serialNo", dev.serialNo)
				param.SetString("model", dev.model)
				param.SetString("product", dev.product)
				param.SetString("version", dev.version)
				param.SetString("buildId", dev.buildId)
				param.SetString("buildType", dev.buildType)
				param.SetString("sku", dev.sku)
				param.SetString("sdk", dev.sdk)
				param.SetString("security", dev.security)
				param.SetString("api", dev.api)
				param.SetString("config", dev.config)
				param.SetString("camera", dev.camera)
				param.SetString("scanner", dev.scanner)
				param.SetString("wwan", dev.wwan)
				remoteService.CallMethod("deviceStateChange", param)
			}
			// push(true, "./runtime/logs/log20210907.log", "/sdcard/log.log", device)
			// pull(true, "/sdcard/log.log", "./log.log", device)
			// logcat(device)
			// installApk("C:\\Users\\H151136\\Desktop\\CPU Z_v1.40_apkpure.com.apk", device)
			// getlogcat("./logcat.log", device)
			// log.Println(dev)
		} else if state == adb.StateDisconnected {
			dev := deviceMap[serial]
			if dev == nil {
				log.Println("cannot find device:", serial)
				return
			}
			log.Printf("Remove device %s - %s - %s", dev.serialNo, dev.model, dev.product)
			close(dev.jobChan)
			var param message.Message
			param.SetBool("state", false)
			param.SetString("host", "Android-"+uid)
			param.SetString("serialNo", dev.serialNo)
			param.SetString("model", dev.model)
			param.SetString("product", dev.product)
			if remoteService != nil {
				remoteService.CallMethod("deviceStateChange", param)
			}
			delete(deviceMap, serial)
		}
	}
	if watcher.Err() != nil {
		log.Fatal(watcher.Err())
	}
	log.Println("Watching for device state changes done")
}

func processTestcase(device *Device) {
	for job := range device.jobChan {
		results := make([]string, len(job.Testcases))
		device.busy = true
		device.console = ""
		device.startTime = time.Now().Unix()
		var err error
		err = nil
		//install test apk
		device.installApk("automation.apk")
		if job.OtaUrl != "" {
			device.testcaseName = "OTA Downloading..."
			// log.Println("Download OTA:" + job.OtaUrl)
			otaName := "OTA-" + job.SerialNo + ".zip"
			device.console += getFormatTime() + ": OTA Downloading: " + job.OtaUrl + "\n"
			uploadConsoleLog(device)
			err = utils.DownloadOTA(job.OtaUrl, otaName)
			if err == nil {
				device.console += getFormatTime() + ": OTA Download done\n"
				device.console += getFormatTime() + ": Push OTA to /sdcard/Download/\n"
				uploadConsoleLog(device)
				device.push(true, otaName, "/sdcard/Download/"+otaName)
				device.console += getFormatTime() + ": Push OTA to /sdcard/Download/ done\n"
				//run upgrade to update the OTA
				device.testcaseName = "OTA Upgrading..."
				device.console += getFormatTime() + ": OTA Upgrading...\n"
				uploadConsoleLog(device)
				ret := device.runTestcase("com.honeywell.test.Tools#updateOTA")
				if strings.Contains(ret, "FAILURES") {
					// log.Println("OTA Failed")
					err = errors.New("OTA Failed")
					device.console += getFormatTime() + ": OTA failed\n"
					uploadConsoleLog(device)
				} else {
					device.console += getFormatTime() + ": OTA done\n"
					forgetDeviceMap[device.serialNo] = device
					//reboot the device
					device.testcaseName = "Device rebooting..."
					device.console += getFormatTime() + ": Device rebooting...\n"
					uploadConsoleLog(device)
					device.reboot()
					time.Sleep(time.Minute * 2)
					uploadDeviceProp(device)
					device.console += getFormatTime() + ": Device done\n"
					delete(forgetDeviceMap, device.serialNo)
					uploadConsoleLog(device)
				}
			} else {
				device.console += getFormatTime() + ": OTA Download failed\n"
				uploadConsoleLog(device)
			}
		}
		device.logcat()
		logFileName := fmt.Sprintf("target-%s-%04d%02d%02d-%02d%02d%02d.log",
			device.serialNo,
			time.Now().Year(),
			time.Now().Month(),
			time.Now().Day(),
			time.Now().Hour(),
			time.Now().Minute(),
			time.Now().Second())
		if err == nil {
			for i, v := range job.Testcases {
				if v != "" {
					// log.Println(device.serialNo, "-> run:", v)
					device.console += getFormatTime() + ": Run testcase: " + v + "\n"
					uploadConsoleLog(device)
					shortName := strings.Split(v, "#")[1]
					device.testcaseName = fmt.Sprintf("(%d/%d)%s", i+1, len(job.Testcases), shortName)
					result := device.runTestcase(v)
					testRunnerMap[job.Id].Completed[i] = "true"
					results = append(results, result)
					// log.Println("  result:", result)
					device.console += getFormatTime() + ": " + result + "\n"
					uploadConsoleLog(device)
				} else {
					log.Println("Invalid testcase:", v, "!")
				}
			}
		} else {
			//ota download failed
			device.testcaseName = "OTA Upgrade failed"
		}
		device.getlogcat(logFileName)

		dir, _ := os.Getwd()
		var fullName string
		if runtime.GOOS == "windows" {
			fullName = dir + "\\" + logFileName
		} else if runtime.GOOS == "linux" {
			fullName = dir + "/" + logFileName
		}
		remotePath := "root@" + remoteHost + ":/root/logcat/"
		utils.Scp(fullName, remotePath, "2222")
		testResult := &TestResult{
			id:             job.Id,
			serialNo:       device.serialNo,
			result:         results,
			logcatFileName: logFileName,
			buildId:        device.buildId}
		testResultMap[job.Id] = testResult
		device.busy = false
		device.testcaseName = ""
		// log.Println("test result:", testResult)
		//upload the test result
		uploadTestResult(testResult)
	}
	log.Println("Device", device.serialNo, "thread end")
}

const StdIoFilename = "-"

func (dev *Device) pull(showProgress bool, remotePath, localPath string) int {
	if remotePath == "" {
		log.Println("error: must specify remote file")
		return -1
	}

	if localPath == "" {
		localPath = filepath.Base(remotePath)
	}

	info, err := dev.device.Stat(remotePath)
	if adb.HasErrCode(err, adb.ErrCode(adb.FileNoExistError)) {
		log.Println("remote file does not exist:", remotePath)
		return -1
	} else if err != nil {
		log.Printf("error reading remote file %s: %s\n", remotePath, err)
		return -1
	}

	remoteFile, err := dev.device.OpenRead(remotePath)
	if err != nil {
		log.Printf("error opening remote file %s: %s\n", remotePath, adb.ErrorWithCauseChain(err))
		return -1
	}
	defer remoteFile.Close()

	var localFile io.WriteCloser
	if localPath == StdIoFilename {
		localFile = os.Stdout
	} else {
		localFile, err = os.Create(localPath)
		if err != nil {
			log.Printf("error opening local file %s: %s\n", localPath, err)
			return -1
		}
	}
	defer localFile.Close()

	if err := copyWithProgressAndStats(localFile, remoteFile, int(info.Size), showProgress); err != nil {
		log.Println("error pulling file:", err)
		return -1
	}
	return 0
}

func (dev *Device) push(showProgress bool, localPath, remotePath string) int {
	if remotePath == "" {
		log.Println("error: must specify remote file")
		return -1
	}

	var (
		localFile io.ReadCloser
		size      int
		perms     os.FileMode
		mtime     time.Time
	)
	if localPath == "" || localPath == StdIoFilename {
		localFile = os.Stdin
		// 0 size will hide the progress bar.
		perms = os.FileMode(0660)
		mtime = adb.MtimeOfClose
	} else {
		var err error
		localFile, err = os.Open(localPath)
		if err != nil {
			log.Printf("error opening local file %s: %s\n", localPath, err)
			return -1
		}
		info, err := os.Stat(localPath)
		if err != nil {
			log.Printf("error reading local file %s: %s\n", localPath, err)
			return -1
		}
		size = int(info.Size())
		perms = info.Mode().Perm()
		mtime = info.ModTime()
	}
	defer localFile.Close()

	writer, err := dev.device.OpenWrite(remotePath, perms, mtime)
	if err != nil {
		log.Printf("error opening remote file %s: %s\n", remotePath, err)
		return -1
	}
	defer writer.Close()

	if err := copyWithProgressAndStats(writer, localFile, size, showProgress); err != nil {
		log.Println("error pushing file:", err)
		return -1
	}
	return 0
}

func (dev *Device) installApk(apkPath string) int {
	tmpApk := "/data/local/tmp/tmp.apk"
	rc := dev.push(false, apkPath, tmpApk)
	if rc != 0 {
		return rc
	}
	cmd := "pm install /data/local/tmp/tmp.apk"
	result, err := dev.device.RunCommand(cmd)
	if err != nil {
		rc = -1
	}
	log.Println("install", apkPath, result)
	return rc
}

func copyWithProgressAndStats(dst io.Writer, src io.Reader, size int, showProgress bool) error {
	var progress *pb.ProgressBar
	if showProgress && size > 0 {
		progress = pb.New(size)
		// Write to stderr in case dst is stdout.
		progress.Output = os.Stderr
		progress.ShowSpeed = true
		progress.ShowPercent = true
		progress.ShowTimeLeft = true
		progress.SetUnits(pb.U_BYTES)
		progress.Start()
		dst = io.MultiWriter(dst, progress)
	}

	startTime := time.Now()
	copied, err := io.Copy(dst, src)

	if progress != nil {
		progress.Finish()
	}

	if pathErr, ok := err.(*os.PathError); ok {
		if errno, ok := pathErr.Err.(syscall.Errno); ok && errno == syscall.EPIPE {
			// Pipe closed. Handle this like an EOF.
			err = nil
		}
	}
	if err != nil {
		return err
	}

	duration := time.Now().Sub(startTime)
	rate := int64(float64(copied) / duration.Seconds())
	log.Printf("%d B/s (%d bytes in %s)\n", rate, copied, duration)

	return nil
}

func (dev *Device) logcat() {
	dev.device.RunCommand("logcat -c")
	go func() {
		dev.device.RunCommand("logcat > /data/local/tmp/logcat.log")
	}()
}

func (dev *Device) getlogcat(localPath string) {
	result, err := dev.device.RunCommand("ps -A | grep \"logcat\" | awk '{print $2}'")
	if err != nil {
		log.Println("failed to run grep logcat command")
		return
	}
	dev.device.RunCommand("kill", result)
	dev.pull(false, "/data/local/tmp/logcat.log", localPath)
}

func (dev *Device) getprop(prop string) string {
	result, err := dev.device.RunCommand("getprop", prop)
	if err != nil {
		log.Println("failed to getprop", prop)
		return ""
	}
	return result
}

func (dev *Device) reboot() error {
	_, err := dev.device.RunCommand("reboot")
	return err
}

func (dev *Device) getDeviceProp() {
	dev.version = strings.Trim(dev.getprop("ro.build.version.release_or_codename"), "\n")
	dev.buildId = strings.Trim(dev.getprop("ro.build.display.id"), "\n")
	dev.buildType = strings.Trim(dev.getprop("ro.build.type"), "\n")
	dev.sku = strings.Trim(dev.getprop("ro.boot.product.hardware.sku"), "\n")
	dev.sdk = strings.Trim(dev.getprop("ro.build.version.sdk"), "\n")
	dev.security = strings.Trim(dev.getprop("ro.build.version.security_patch"), "\n")
	dev.api = strings.Trim(dev.getprop("ro.product.first_api_level"), "\n")
	dev.config = strings.Trim(dev.getprop("ro.vendor.hon.extconf.num"), "\n")
	dev.camera = strings.Trim(dev.getprop("ro.vendor.hon.plat.camera"), "\n")
	dev.scanner = strings.Trim(dev.getprop("ro.vendor.hon.plat.imager.sensor"), "\n")
	dev.wwan = strings.Trim(dev.getprop("ro.vendor.hon.plat.wwan"), "\n")
}

func (dev *Device) runTestcase(name string) string {
	cmd := fmt.Sprintf("am instrument -w -r -e class '%s' com.honeywell.test.test/androidx.test.runner.AndroidJUnitRunner",
		name)
	result, err := dev.device.RunCommand(cmd)
	if err != nil {
		return ""
	}
	return result
}

var runTestcaseHandle message.MethodHandle = func(input message.Message) (output message.Message) {
	serialNo := input.GetString("serialNo", "")
	_testcases := input.GetStringArray("testcases")
	testcases := make([]string, 0)
	for _, v := range _testcases {
		if len(v) > 10 {
			testcases = append(testcases, v)
		}
	}
	otaUrl := input.GetString("otaUrl", "")

	if _, exist := deviceMap[serialNo]; !exist {
		log.Println("Device", serialNo, "not available")
		return message.Message{}
	}
	device := deviceMap[serialNo]
	if device.busy {
		log.Println("Device", serialNo, "is busy")
		var out message.Message
		out.SetInt("testId", -1)
		return out
	}
	var runner TestRunner
	runner.Id = int(time.Now().Unix())
	runner.SerialNo = serialNo
	runner.Testcases = make([]string, len(testcases))
	runner.Completed = make([]string, len(testcases))
	runner.OtaUrl = otaUrl
	for k, v := range testcases {
		runner.Testcases[k] = v
		runner.Completed[k] = "false"
	}
	//add test runner to process
	log.Println("run:", runner)
	device.jobChan <- &runner

	testRunnerMap[runner.Id] = &runner
	var out message.Message
	out.SetInt("testId", runner.Id)
	return out
}

var shellCmdHandle message.MethodHandle = func(input message.Message) (output message.Message) {
	serialNo := input.GetString("serialNo", "")
	cmd := input.GetString("cmd", "")
	args := input.GetString("args", "")

	if _, exist := deviceMap[serialNo]; !exist {
		return message.Message{}
	}
	device := deviceMap[serialNo]
	result, err := device.device.RunCommand(cmd, args)
	if err != nil {
		return message.Message{}
	}
	var out message.Message
	out.SetString("result", result)
	return out
}

var getDeviceListHandle message.MethodHandle = func(m message.Message) message.Message {
	var deviceList []message.Message
	deviceList = make([]message.Message, 0)
	for _, dev := range deviceMap {
		var device message.Message
		device.SetString("serialNo", dev.serialNo)
		device.SetString("model", dev.model)
		device.SetString("product", dev.product)
		device.SetString("version", dev.version)
		device.SetString("buildId", dev.buildId)
		device.SetString("buildType", dev.buildType)
		device.SetString("sku", dev.sku)
		device.SetString("sdk", dev.sdk)
		device.SetString("security", dev.security)
		device.SetString("api", dev.api)
		device.SetString("config", dev.config)
		device.SetString("camera", dev.camera)
		device.SetString("scanner", dev.scanner)
		device.SetString("wwan", dev.wwan)
		device.SetBool("busy", dev.busy)
		deviceList = append(deviceList, device)
	}
	var result message.Message
	result.SetMessageArray("deviceList", deviceList)
	return result
}

var getTestResultHandle message.MethodHandle = func(m message.Message) message.Message {
	serialNo := m.GetString("serialNo", "")
	var output message.Message
	if _, exist := deviceMap[serialNo]; !exist {
		return message.Message{}
	}
	testId := m.GetInt("testId", -1)
	if testId < 0 {
		return message.Message{}
	}
	testResult := testResultMap[testId]
	output.SetInt("testId", testId)
	output.SetString("serialNo", serialNo)
	output.SetStringArray("results", testResult.result)
	output.SetString("logcat", testResult.logcatFileName)

	return output
}

var getRuntimeStateHandle message.MethodHandle = func(m message.Message) message.Message {
	serialNo := m.GetString("serialNo", "")
	var output message.Message
	if _, exist := deviceMap[serialNo]; !exist {
		output.SetString("serialNo", serialNo)
		output.SetInt("state", 2)
		return output
	}
	device := deviceMap[serialNo]
	output.SetString("serialNo", serialNo)
	if device.busy {
		output.SetInt("state", 1)
	} else {
		output.SetInt("state", 0)
	}
	duration := time.Now().Unix() - device.startTime
	output.SetInt("duration", int(duration))
	if device.busy {
		testcaseName := device.testcaseName
		output.SetString("testcaseName", testcaseName)
	}
	return output
}

var getTestRunnerStateHandle message.MethodHandle = func(m message.Message) message.Message {
	id := m.GetInt("testId", -1)
	var output message.Message
	if id < 0 {
		return message.Message{}
	}
	runner := testRunnerMap[id]
	output.SetInt("id", runner.Id)
	output.SetString("serialNo", runner.SerialNo)
	output.SetStringArray("testcases", runner.Testcases)
	output.SetStringArray("completed", runner.Completed)
	return output
}

var clientConnectedHandle message.ClientConnectedHandle = func() {
	log.Println("Client Connected")
}

var clientDisconnectHandle message.ClientDisconnectedHandle = func() {
	log.Println("Client Disconnected")
}

func uploadTestResult(result *TestResult) {
	if remoteService == nil {
		log.Println("remote service is nil")
		return
	}
	var msg message.Message
	msg.SetInt("id", result.id)
	msg.SetString("serialNo", result.serialNo)
	msg.SetStringArray("result", result.result)
	msg.SetString("logcat", result.logcatFileName)
	msg.SetString("buildId", result.buildId)
	remoteService.CallMethod("uploadResult", msg)
}

func uploadConsoleLog(device *Device) {
	if remoteService == nil {
		log.Println("remote service is nil")
		return
	}
	var msg message.Message
	msg.SetString("serialNo", device.serialNo)
	msg.SetString("console", device.console)
	remoteService.CallMethod("uploadConsole", msg)
}

func uploadDeviceProp(device *Device) {
	device.getDeviceProp()
	log.Printf("New device %s - %s - %s", device.serialNo, device.model, device.product)
	if remoteService != nil {
		var param message.Message
		param.SetBool("state", true)
		param.SetString("host", "Android-"+uid)
		param.SetString("serialNo", device.serialNo)
		param.SetString("model", device.model)
		param.SetString("product", device.product)
		param.SetString("version", device.version)
		param.SetString("buildId", device.buildId)
		param.SetString("buildType", device.buildType)
		param.SetString("sku", device.sku)
		param.SetString("sdk", device.sdk)
		param.SetString("security", device.security)
		param.SetString("api", device.api)
		param.SetString("config", device.config)
		param.SetString("camera", device.camera)
		param.SetString("scanner", device.scanner)
		param.SetString("wwan", device.wwan)
		remoteService.CallMethod("deviceStateChange", param)
	}
}

func getFormatTime() string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d",
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second())
}

func Start() {
	uid = uuid.New()[:8]
	message.RegisterBusListener(message.MessageBusListener{ConnectHandle: clientConnectedHandle, DisconnectHandle: clientDisconnectHandle})
	androidService = message.CreateLocalService("Android-" + uid)
	androidService.RegisterMethod("shellCmd", shellCmdHandle)
	androidService.RegisterMethod("runTestcase", runTestcaseHandle)
	androidService.RegisterMethod("getDeviceList", getDeviceListHandle)
	androidService.RegisterMethod("getRuntimeState", getRuntimeStateHandle)
	androidService.RegisterMethod("getTestRunnerState", getTestRunnerStateHandle)
	message.AddService(androidService)
	remoteService = message.GetService("server")
}
