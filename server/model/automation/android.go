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
