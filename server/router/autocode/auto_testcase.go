package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestcaseRouter struct {
}

// InitTestcaseRouter 初始化 Testcase 路由信息
func (s *TestcaseRouter) InitTestcaseRouter(Router *gin.RouterGroup) {
	testcaseRouter := Router.Group("testcase").Use(middleware.OperationRecord())
	testcaseRouterWithoutRecord := Router.Group("testcase")
	var testcaseApi = v1.ApiGroupApp.AutoCodeApiGroup.TestcaseApi
	{
		testcaseRouter.POST("createTestcase", testcaseApi.CreateTestcase)   // 新建Testcase
		testcaseRouter.DELETE("deleteTestcase", testcaseApi.DeleteTestcase) // 删除Testcase
		testcaseRouter.DELETE("deleteTestcaseByIds", testcaseApi.DeleteTestcaseByIds) // 批量删除Testcase
		testcaseRouter.PUT("updateTestcase", testcaseApi.UpdateTestcase)    // 更新Testcase
	}
	{
		testcaseRouterWithoutRecord.GET("findTestcase", testcaseApi.FindTestcase)        // 根据ID获取Testcase
		testcaseRouterWithoutRecord.GET("getTestcaseList", testcaseApi.GetTestcaseList)  // 获取Testcase列表
	}
}
