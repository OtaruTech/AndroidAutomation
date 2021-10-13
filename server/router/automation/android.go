package automation

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AndroidRouter struct {

}

func (s *AndroidRouter) InitAndroidRouter(Router *gin.RouterGroup) {
	androidRouter := Router.Group("android")
	var androidApi = v1.ApiGroupApp.AutomationApiGroup.AndroidApi
	{
		androidRouter.POST("getRuntimeState", androidApi.GetRuntimeState)
		androidRouter.POST("getTestRunnerState", androidApi.GetRuntimeState)
		androidRouter.POST("runTestcase", androidApi.RunTestcase)
		androidRouter.GET("downloadFile", androidApi.DownloadFile)
		androidRouter.POST("jobChanged", androidApi.JobChanged)
	}
}