package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type DeviceConsoleService struct {
}

// CreateDeviceConsole 创建DeviceConsole记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceConsoleService *DeviceConsoleService) CreateDeviceConsole(deviceConsole autocode.DeviceConsole) (err error) {
	err = global.GVA_DB.Create(&deviceConsole).Error
	return err
}

// DeleteDeviceConsole 删除DeviceConsole记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceConsoleService *DeviceConsoleService)DeleteDeviceConsole(deviceConsole autocode.DeviceConsole) (err error) {
	err = global.GVA_DB.Delete(&deviceConsole).Error
	return err
}

// DeleteDeviceConsoleByIds 批量删除DeviceConsole记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceConsoleService *DeviceConsoleService)DeleteDeviceConsoleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.DeviceConsole{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDeviceConsole 更新DeviceConsole记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceConsoleService *DeviceConsoleService)UpdateDeviceConsole(deviceConsole autocode.DeviceConsole) (err error) {
	err = global.GVA_DB.Save(&deviceConsole).Error
	return err
}

// GetDeviceConsole 根据id获取DeviceConsole记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceConsoleService *DeviceConsoleService)GetDeviceConsole(id uint) (err error, deviceConsole autocode.DeviceConsole) {
	err = global.GVA_DB.Where("id = ?", id).First(&deviceConsole).Error
	return
}

// GetDeviceConsoleInfoList 分页获取DeviceConsole记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceConsoleService *DeviceConsoleService)GetDeviceConsoleInfoList(info autoCodeReq.DeviceConsoleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.DeviceConsole{})
    var deviceConsoles []autocode.DeviceConsole
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.SerialNo != "" {
        db = db.Where("`serialNo` = ?",info.SerialNo)
    }
    if info.Console != "" {
        db = db.Where("`console` LIKE ?","%"+ info.Console+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&deviceConsoles).Error
	return err, deviceConsoles, total
}
