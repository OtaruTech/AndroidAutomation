package automation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type RunTestcase struct {
	global.GVA_MODEL
	Testcases []string `json:"testcases"`
	SerialNo  string   `json:"serialNo"`
	Timeout   int      `json:"timeout"`
	OtaUrl    string   `json:"otaUrl"`
}

type DeviceRequest struct {
	global.GVA_MODEL
	SerialNo string `json:"serialNo"`
}

type TestId struct {
	global.GVA_MODEL
	TestId int `json:"testId"`
}

type FileInfo struct {
	global.GVA_MODEL
	FileName string `json:"fileName" form:"fileName"`
}

type LatestOtaRequest struct {
	OtaPath    string `json:"otaPath" form:"otaPath"`
	OtaFormat  string `json:"otaFormat" form:"otaFormat"`
	FileFormat string `json:"fileFormat" form:"fileFormat"`
}

type LatestOtaResponse struct {
	LatestVersion int    `json:"latestVersion" form:"latestVersion"`
	LatestOtaUrl  string `json:"latestOtaUrl" form:"latestOtaUrl"`
}

type JobChangedReq struct {
	Id     int  `json:"id" form:"id"`
	Enable bool `json:"enable" form:"enable"`
}
