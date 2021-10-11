package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TestRunnerApi struct {
}

var testRunnerService = service.ServiceGroupApp.AutoCodeServiceGroup.TestRunnerService


// CreateTestRunner 创建TestRunner
// @Tags TestRunner
// @Summary 创建TestRunner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TestRunner true "创建TestRunner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testRunner/createTestRunner [post]
func (testRunnerApi *TestRunnerApi) CreateTestRunner(c *gin.Context) {
	var testRunner autocode.TestRunner
	_ = c.ShouldBindJSON(&testRunner)
	if err := testRunnerService.CreateTestRunner(testRunner); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestRunner 删除TestRunner
// @Tags TestRunner
// @Summary 删除TestRunner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TestRunner true "删除TestRunner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testRunner/deleteTestRunner [delete]
func (testRunnerApi *TestRunnerApi) DeleteTestRunner(c *gin.Context) {
	var testRunner autocode.TestRunner
	_ = c.ShouldBindJSON(&testRunner)
	if err := testRunnerService.DeleteTestRunner(testRunner); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestRunnerByIds 批量删除TestRunner
// @Tags TestRunner
// @Summary 批量删除TestRunner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestRunner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /testRunner/deleteTestRunnerByIds [delete]
func (testRunnerApi *TestRunnerApi) DeleteTestRunnerByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := testRunnerService.DeleteTestRunnerByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestRunner 更新TestRunner
// @Tags TestRunner
// @Summary 更新TestRunner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TestRunner true "更新TestRunner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testRunner/updateTestRunner [put]
func (testRunnerApi *TestRunnerApi) UpdateTestRunner(c *gin.Context) {
	var testRunner autocode.TestRunner
	_ = c.ShouldBindJSON(&testRunner)
	if err := testRunnerService.UpdateTestRunner(testRunner); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestRunner 用id查询TestRunner
// @Tags TestRunner
// @Summary 用id查询TestRunner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.TestRunner true "用id查询TestRunner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testRunner/findTestRunner [get]
func (testRunnerApi *TestRunnerApi) FindTestRunner(c *gin.Context) {
	var testRunner autocode.TestRunner
	_ = c.ShouldBindQuery(&testRunner)
	if err, retestRunner := testRunnerService.GetTestRunner(testRunner.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestRunner": retestRunner}, c)
	}
}

// GetTestRunnerList 分页获取TestRunner列表
// @Tags TestRunner
// @Summary 分页获取TestRunner列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.TestRunnerSearch true "分页获取TestRunner列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testRunner/getTestRunnerList [get]
func (testRunnerApi *TestRunnerApi) GetTestRunnerList(c *gin.Context) {
	var pageInfo autocodeReq.TestRunnerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := testRunnerService.GetTestRunnerInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
