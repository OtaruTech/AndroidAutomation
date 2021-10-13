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

type ReportApi struct {
}

var reportService = service.ServiceGroupApp.AutoCodeServiceGroup.ReportService


// CreateReport 创建Report
// @Tags Report
// @Summary 创建Report
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Report true "创建Report"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /report/createReport [post]
func (reportApi *ReportApi) CreateReport(c *gin.Context) {
	var report autocode.Report
	_ = c.ShouldBindJSON(&report)
	if err := reportService.CreateReport(report); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteReport 删除Report
// @Tags Report
// @Summary 删除Report
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Report true "删除Report"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /report/deleteReport [delete]
func (reportApi *ReportApi) DeleteReport(c *gin.Context) {
	var report autocode.Report
	_ = c.ShouldBindJSON(&report)
	if err := reportService.DeleteReport(report); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteReportByIds 批量删除Report
// @Tags Report
// @Summary 批量删除Report
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Report"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /report/deleteReportByIds [delete]
func (reportApi *ReportApi) DeleteReportByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := reportService.DeleteReportByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateReport 更新Report
// @Tags Report
// @Summary 更新Report
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Report true "更新Report"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /report/updateReport [put]
func (reportApi *ReportApi) UpdateReport(c *gin.Context) {
	var report autocode.Report
	_ = c.ShouldBindJSON(&report)
	if err := reportService.UpdateReport(report); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindReport 用id查询Report
// @Tags Report
// @Summary 用id查询Report
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Report true "用id查询Report"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /report/findReport [get]
func (reportApi *ReportApi) FindReport(c *gin.Context) {
	var report autocode.Report
	_ = c.ShouldBindQuery(&report)
	if err, rereport := reportService.GetReport(report.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rereport": rereport}, c)
	}
}

// GetReportList 分页获取Report列表
// @Tags Report
// @Summary 分页获取Report列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ReportSearch true "分页获取Report列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /report/getReportList [get]
func (reportApi *ReportApi) GetReportList(c *gin.Context) {
	var pageInfo autocodeReq.ReportSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := reportService.GetReportInfoList(pageInfo); err != nil {
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
