package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestsetRouter struct {
}

// InitTestsetRouter 初始化 Testset 路由信息
func (s *TestsetRouter) InitTestsetRouter(Router *gin.RouterGroup) {
	testsetRouter := Router.Group("testset").Use(middleware.OperationRecord())
	testsetRouterWithoutRecord := Router.Group("testset")
	var testsetApi = v1.ApiGroupApp.AutoCodeApiGroup.TestsetApi
	{
		testsetRouter.POST("createTestset", testsetApi.CreateTestset)   // 新建Testset
		testsetRouter.DELETE("deleteTestset", testsetApi.DeleteTestset) // 删除Testset
		testsetRouter.DELETE("deleteTestsetByIds", testsetApi.DeleteTestsetByIds) // 批量删除Testset
		testsetRouter.PUT("updateTestset", testsetApi.UpdateTestset)    // 更新Testset
	}
	{
		testsetRouterWithoutRecord.GET("findTestset", testsetApi.FindTestset)        // 根据ID获取Testset
		testsetRouterWithoutRecord.GET("getTestsetList", testsetApi.GetTestsetList)  // 获取Testset列表
	}
}
