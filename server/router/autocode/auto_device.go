package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceRouter struct {
}

// InitDeviceRouter 初始化 Device 路由信息
func (s *DeviceRouter) InitDeviceRouter(Router *gin.RouterGroup) {
	deviceRouter := Router.Group("device").Use(middleware.OperationRecord())
	deviceRouterWithoutRecord := Router.Group("device")
	var deviceApi = v1.ApiGroupApp.AutoCodeApiGroup.DeviceApi
	{
		deviceRouter.POST("createDevice", deviceApi.CreateDevice)   // 新建Device
		deviceRouter.DELETE("deleteDevice", deviceApi.DeleteDevice) // 删除Device
		deviceRouter.DELETE("deleteDeviceByIds", deviceApi.DeleteDeviceByIds) // 批量删除Device
		deviceRouter.PUT("updateDevice", deviceApi.UpdateDevice)    // 更新Device
	}
	{
		deviceRouterWithoutRecord.GET("findDevice", deviceApi.FindDevice)        // 根据ID获取Device
		deviceRouterWithoutRecord.GET("getDeviceList", deviceApi.GetDeviceList)  // 获取Device列表
	}
}
