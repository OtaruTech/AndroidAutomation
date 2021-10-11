package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ReportRouter struct {
}

// InitReportRouter 初始化 Report 路由信息
func (s *ReportRouter) InitReportRouter(Router *gin.RouterGroup) {
	reportRouter := Router.Group("report").Use(middleware.OperationRecord())
	reportRouterWithoutRecord := Router.Group("report")
	var reportApi = v1.ApiGroupApp.AutoCodeApiGroup.ReportApi
	{
		reportRouter.POST("createReport", reportApi.CreateReport)   // 新建Report
		reportRouter.DELETE("deleteReport", reportApi.DeleteReport) // 删除Report
		reportRouter.DELETE("deleteReportByIds", reportApi.DeleteReportByIds) // 批量删除Report
		reportRouter.PUT("updateReport", reportApi.UpdateReport)    // 更新Report
	}
	{
		reportRouterWithoutRecord.GET("findReport", reportApi.FindReport)        // 根据ID获取Report
		reportRouterWithoutRecord.GET("getReportList", reportApi.GetReportList)  // 获取Report列表
	}
}
