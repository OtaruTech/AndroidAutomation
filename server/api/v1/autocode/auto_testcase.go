package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TestcaseApi struct {
}

var testcaseService = service.ServiceGroupApp.AutoCodeServiceGroup.TestcaseService


// CreateTestcase 创建Testcase
// @Tags Testcase
// @Summary 创建Testcase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Testcase true "创建Testcase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testcase/createTestcase [post]
func (testcaseApi *TestcaseApi) CreateTestcase(c *gin.Context) {
	var testcase autocode.Testcase
	_ = c.ShouldBindJSON(&testcase)
	if err := testcaseService.CreateTestcase(testcase); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestcase 删除Testcase
// @Tags Testcase
// @Summary 删除Testcase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Testcase true "删除Testcase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testcase/deleteTestcase [delete]
func (testcaseApi *TestcaseApi) DeleteTestcase(c *gin.Context) {
	var testcase autocode.Testcase
	_ = c.ShouldBindJSON(&testcase)
	if err := testcaseService.DeleteTestcase(testcase); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestcaseByIds 批量删除Testcase
// @Tags Testcase
// @Summary 批量删除Testcase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Testcase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /testcase/deleteTestcaseByIds [delete]
func (testcaseApi *TestcaseApi) DeleteTestcaseByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := testcaseService.DeleteTestcaseByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestcase 更新Testcase
// @Tags Testcase
// @Summary 更新Testcase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Testcase true "更新Testcase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testcase/updateTestcase [put]
func (testcaseApi *TestcaseApi) UpdateTestcase(c *gin.Context) {
	var testcase autocode.Testcase
	_ = c.ShouldBindJSON(&testcase)
	if err := testcaseService.UpdateTestcase(testcase); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestcase 用id查询Testcase
// @Tags Testcase
// @Summary 用id查询Testcase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Testcase true "用id查询Testcase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testcase/findTestcase [get]
func (testcaseApi *TestcaseApi) FindTestcase(c *gin.Context) {
	var testcase autocode.Testcase
	_ = c.ShouldBindQuery(&testcase)
	if err, retestcase := testcaseService.GetTestcase(testcase.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestcase": retestcase}, c)
	}
}

// GetTestcaseList 分页获取Testcase列表
// @Tags Testcase
// @Summary 分页获取Testcase列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.TestcaseSearch true "分页获取Testcase列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testcase/getTestcaseList [get]
func (testcaseApi *TestcaseApi) GetTestcaseList(c *gin.Context) {
	var pageInfo autocodeReq.TestcaseSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := testcaseService.GetTestcaseInfoList(pageInfo); err != nil {
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
