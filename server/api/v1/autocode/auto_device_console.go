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

type DeviceConsoleApi struct {
}

var deviceConsoleService = service.ServiceGroupApp.AutoCodeServiceGroup.DeviceConsoleService


// CreateDeviceConsole 创建DeviceConsole
// @Tags DeviceConsole
// @Summary 创建DeviceConsole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.DeviceConsole true "创建DeviceConsole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceConsole/createDeviceConsole [post]
func (deviceConsoleApi *DeviceConsoleApi) CreateDeviceConsole(c *gin.Context) {
	var deviceConsole autocode.DeviceConsole
	_ = c.ShouldBindJSON(&deviceConsole)
	if err := deviceConsoleService.CreateDeviceConsole(deviceConsole); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDeviceConsole 删除DeviceConsole
// @Tags DeviceConsole
// @Summary 删除DeviceConsole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.DeviceConsole true "删除DeviceConsole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceConsole/deleteDeviceConsole [delete]
func (deviceConsoleApi *DeviceConsoleApi) DeleteDeviceConsole(c *gin.Context) {
	var deviceConsole autocode.DeviceConsole
	_ = c.ShouldBindJSON(&deviceConsole)
	if err := deviceConsoleService.DeleteDeviceConsole(deviceConsole); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDeviceConsoleByIds 批量删除DeviceConsole
// @Tags DeviceConsole
// @Summary 批量删除DeviceConsole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DeviceConsole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /deviceConsole/deleteDeviceConsoleByIds [delete]
func (deviceConsoleApi *DeviceConsoleApi) DeleteDeviceConsoleByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := deviceConsoleService.DeleteDeviceConsoleByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDeviceConsole 更新DeviceConsole
// @Tags DeviceConsole
// @Summary 更新DeviceConsole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.DeviceConsole true "更新DeviceConsole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deviceConsole/updateDeviceConsole [put]
func (deviceConsoleApi *DeviceConsoleApi) UpdateDeviceConsole(c *gin.Context) {
	var deviceConsole autocode.DeviceConsole
	_ = c.ShouldBindJSON(&deviceConsole)
	if err := deviceConsoleService.UpdateDeviceConsole(deviceConsole); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDeviceConsole 用id查询DeviceConsole
// @Tags DeviceConsole
// @Summary 用id查询DeviceConsole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.DeviceConsole true "用id查询DeviceConsole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deviceConsole/findDeviceConsole [get]
func (deviceConsoleApi *DeviceConsoleApi) FindDeviceConsole(c *gin.Context) {
	var deviceConsole autocode.DeviceConsole
	_ = c.ShouldBindQuery(&deviceConsole)
	if err, redeviceConsole := deviceConsoleService.GetDeviceConsole(deviceConsole.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redeviceConsole": redeviceConsole}, c)
	}
}

// GetDeviceConsoleList 分页获取DeviceConsole列表
// @Tags DeviceConsole
// @Summary 分页获取DeviceConsole列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.DeviceConsoleSearch true "分页获取DeviceConsole列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceConsole/getDeviceConsoleList [get]
func (deviceConsoleApi *DeviceConsoleApi) GetDeviceConsoleList(c *gin.Context) {
	var pageInfo autocodeReq.DeviceConsoleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := deviceConsoleService.GetDeviceConsoleInfoList(pageInfo); err != nil {
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
