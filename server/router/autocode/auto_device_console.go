package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceConsoleRouter struct {
}

// InitDeviceConsoleRouter 初始化 DeviceConsole 路由信息
func (s *DeviceConsoleRouter) InitDeviceConsoleRouter(Router *gin.RouterGroup) {
	deviceConsoleRouter := Router.Group("deviceConsole").Use(middleware.OperationRecord())
	deviceConsoleRouterWithoutRecord := Router.Group("deviceConsole")
	var deviceConsoleApi = v1.ApiGroupApp.AutoCodeApiGroup.DeviceConsoleApi
	{
		deviceConsoleRouter.POST("createDeviceConsole", deviceConsoleApi.CreateDeviceConsole)   // 新建DeviceConsole
		deviceConsoleRouter.DELETE("deleteDeviceConsole", deviceConsoleApi.DeleteDeviceConsole) // 删除DeviceConsole
		deviceConsoleRouter.DELETE("deleteDeviceConsoleByIds", deviceConsoleApi.DeleteDeviceConsoleByIds) // 批量删除DeviceConsole
		deviceConsoleRouter.PUT("updateDeviceConsole", deviceConsoleApi.UpdateDeviceConsole)    // 更新DeviceConsole
	}
	{
		deviceConsoleRouterWithoutRecord.GET("findDeviceConsole", deviceConsoleApi.FindDeviceConsole)        // 根据ID获取DeviceConsole
		deviceConsoleRouterWithoutRecord.GET("getDeviceConsoleList", deviceConsoleApi.GetDeviceConsoleList)  // 获取DeviceConsole列表
	}
}
