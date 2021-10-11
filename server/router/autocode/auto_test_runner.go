package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestRunnerRouter struct {
}

// InitTestRunnerRouter 初始化 TestRunner 路由信息
func (s *TestRunnerRouter) InitTestRunnerRouter(Router *gin.RouterGroup) {
	testRunnerRouter := Router.Group("testRunner").Use(middleware.OperationRecord())
	testRunnerRouterWithoutRecord := Router.Group("testRunner")
	var testRunnerApi = v1.ApiGroupApp.AutoCodeApiGroup.TestRunnerApi
	{
		testRunnerRouter.POST("createTestRunner", testRunnerApi.CreateTestRunner)   // 新建TestRunner
		testRunnerRouter.DELETE("deleteTestRunner", testRunnerApi.DeleteTestRunner) // 删除TestRunner
		testRunnerRouter.DELETE("deleteTestRunnerByIds", testRunnerApi.DeleteTestRunnerByIds) // 批量删除TestRunner
		testRunnerRouter.PUT("updateTestRunner", testRunnerApi.UpdateTestRunner)    // 更新TestRunner
	}
	{
		testRunnerRouterWithoutRecord.GET("findTestRunner", testRunnerApi.FindTestRunner)        // 根据ID获取TestRunner
		testRunnerRouterWithoutRecord.GET("getTestRunnerList", testRunnerApi.GetTestRunnerList)  // 获取TestRunner列表
	}
}
