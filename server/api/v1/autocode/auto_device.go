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

type DeviceApi struct {
}

var deviceService = service.ServiceGroupApp.AutoCodeServiceGroup.DeviceService


// CreateDevice 创建Device
// @Tags Device
// @Summary 创建Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Device true "创建Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /device/createDevice [post]
func (deviceApi *DeviceApi) CreateDevice(c *gin.Context) {
	var device autocode.Device
	_ = c.ShouldBindJSON(&device)
	if err := deviceService.CreateDevice(device); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDevice 删除Device
// @Tags Device
// @Summary 删除Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Device true "删除Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /device/deleteDevice [delete]
func (deviceApi *DeviceApi) DeleteDevice(c *gin.Context) {
	var device autocode.Device
	_ = c.ShouldBindJSON(&device)
	if err := deviceService.DeleteDevice(device); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDeviceByIds 批量删除Device
// @Tags Device
// @Summary 批量删除Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /device/deleteDeviceByIds [delete]
func (deviceApi *DeviceApi) DeleteDeviceByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := deviceService.DeleteDeviceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDevice 更新Device
// @Tags Device
// @Summary 更新Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Device true "更新Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /device/updateDevice [put]
func (deviceApi *DeviceApi) UpdateDevice(c *gin.Context) {
	var device autocode.Device
	_ = c.ShouldBindJSON(&device)
	if err := deviceService.UpdateDevice(device); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDevice 用id查询Device
// @Tags Device
// @Summary 用id查询Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Device true "用id查询Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /device/findDevice [get]
func (deviceApi *DeviceApi) FindDevice(c *gin.Context) {
	var device autocode.Device
	_ = c.ShouldBindQuery(&device)
	if err, redevice := deviceService.GetDevice(device.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redevice": redevice}, c)
	}
}

// GetDeviceList 分页获取Device列表
// @Tags Device
// @Summary 分页获取Device列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.DeviceSearch true "分页获取Device列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /device/getDeviceList [get]
func (deviceApi *DeviceApi) GetDeviceList(c *gin.Context) {
	var pageInfo autocodeReq.DeviceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := deviceService.GetDeviceInfoList(pageInfo); err != nil {
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
