package automation

import (
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/model/automation"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	auto "github.com/flipped-aurora/gin-vue-admin/server/rpc"
	"github.com/gin-gonic/gin"
)

type AndroidApi struct {
}

func (e *AndroidApi) GetRuntimeState(c *gin.Context) {
	var req automation.DeviceRequest
	_ = c.ShouldBindJSON(&req)
	ret, state := auto.AndroidGetRuntimeState(req.SerialNo)
	log.Println("automation: GetRuntimeState", ret, state)
	if(ret < 0) {
		// response.FailWithMessage("获取设备运行状态失败", c)
		response.Ok(c)
	} else {
		response.OkWithData(state, c)
	}
}

func (e *AndroidApi) GetTestRunnerState(c *gin.Context) {
	var req automation.TestId
	_ = c.ShouldBindJSON(&req)
	ret, state := auto.AndroidGetTestRunnerState(req.TestId)
	log.Println("automation: GetTestRunnerState", ret, state)
	if(ret < 0) {
		response.FailWithMessage("获取测试任务状态", c)
	} else {
		response.OkWithData(state, c)
	}
}

func (e *AndroidApi) RunTestcase(c *gin.Context) {
	var runTestcase automation.RunTestcase
	_ = c.ShouldBindJSON(&runTestcase)
	log.Println("automation:", runTestcase)
	ret := auto.AndroidRunTestcase(runTestcase.SerialNo, runTestcase.Testcases, runTestcase.Timeout, runTestcase.OtaUrl)
	if(ret < 0) {
		response.FailWithMessage("运行测试用例失败", c)
	} else {
		var testId automation.TestId
		testId.TestId = ret
		response.OkWithDetailed(testId, "开始运行测试用例", c)
	}
}
