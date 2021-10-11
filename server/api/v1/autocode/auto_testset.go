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

type TestsetApi struct {
}

var testsetService = service.ServiceGroupApp.AutoCodeServiceGroup.TestsetService


// CreateTestset 创建Testset
// @Tags Testset
// @Summary 创建Testset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Testset true "创建Testset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testset/createTestset [post]
func (testsetApi *TestsetApi) CreateTestset(c *gin.Context) {
	var testset autocode.Testset
	_ = c.ShouldBindJSON(&testset)
	if err := testsetService.CreateTestset(testset); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestset 删除Testset
// @Tags Testset
// @Summary 删除Testset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Testset true "删除Testset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testset/deleteTestset [delete]
func (testsetApi *TestsetApi) DeleteTestset(c *gin.Context) {
	var testset autocode.Testset
	_ = c.ShouldBindJSON(&testset)
	if err := testsetService.DeleteTestset(testset); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestsetByIds 批量删除Testset
// @Tags Testset
// @Summary 批量删除Testset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Testset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /testset/deleteTestsetByIds [delete]
func (testsetApi *TestsetApi) DeleteTestsetByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := testsetService.DeleteTestsetByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestset 更新Testset
// @Tags Testset
// @Summary 更新Testset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Testset true "更新Testset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testset/updateTestset [put]
func (testsetApi *TestsetApi) UpdateTestset(c *gin.Context) {
	var testset autocode.Testset
	_ = c.ShouldBindJSON(&testset)
	if err := testsetService.UpdateTestset(testset); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestset 用id查询Testset
// @Tags Testset
// @Summary 用id查询Testset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Testset true "用id查询Testset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testset/findTestset [get]
func (testsetApi *TestsetApi) FindTestset(c *gin.Context) {
	var testset autocode.Testset
	_ = c.ShouldBindQuery(&testset)
	if err, retestset := testsetService.GetTestset(testset.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestset": retestset}, c)
	}
}

// GetTestsetList 分页获取Testset列表
// @Tags Testset
// @Summary 分页获取Testset列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.TestsetSearch true "分页获取Testset列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testset/getTestsetList [get]
func (testsetApi *TestsetApi) GetTestsetList(c *gin.Context) {
	var pageInfo autocodeReq.TestsetSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := testsetService.GetTestsetInfoList(pageInfo); err != nil {
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
