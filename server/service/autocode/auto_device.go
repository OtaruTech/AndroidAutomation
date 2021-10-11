package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type DeviceService struct {
}

// CreateDevice 创建Device记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceService *DeviceService) CreateDevice(device autocode.Device) (err error) {
	err = global.GVA_DB.Create(&device).Error
	return err
}

// DeleteDevice 删除Device记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceService *DeviceService)DeleteDevice(device autocode.Device) (err error) {
	err = global.GVA_DB.Delete(&device).Error
	return err
}

// DeleteDeviceByIds 批量删除Device记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceService *DeviceService)DeleteDeviceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Device{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDevice 更新Device记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceService *DeviceService)UpdateDevice(device autocode.Device) (err error) {
	err = global.GVA_DB.Save(&device).Error
	return err
}

// GetDevice 根据id获取Device记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceService *DeviceService)GetDevice(id uint) (err error, device autocode.Device) {
	err = global.GVA_DB.Where("id = ?", id).First(&device).Error
	return
}

// GetDeviceInfoList 分页获取Device记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceService *DeviceService)GetDeviceInfoList(info autoCodeReq.DeviceSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Device{})
    var devices []autocode.Device
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.SerialNo != "" {
        db = db.Where("`serialNo` = ?",info.SerialNo)
    }
    if info.Model != "" {
        db = db.Where("`model` = ?",info.Model)
    }
    if info.Product != "" {
        db = db.Where("`product` = ?",info.Product)
    }
    if info.Version != "" {
        db = db.Where("`version` = ?",info.Version)
    }
    if info.BuildId != "" {
        db = db.Where("`buildId` = ?",info.BuildId)
    }
    if info.BuildType != "" {
        db = db.Where("`buildType` = ?",info.BuildType)
    }
    if info.Sku != "" {
        db = db.Where("`sku` = ?",info.Sku)
    }
    if info.Sdk != "" {
        db = db.Where("`sdk` = ?",info.Sdk)
    }
    if info.Security != "" {
        db = db.Where("`security` = ?",info.Security)
    }
    if info.Api != "" {
        db = db.Where("`api` = ?",info.Api)
    }
    if info.Config != "" {
        db = db.Where("`config` = ?",info.Config)
    }
    if info.Camera != "" {
        db = db.Where("`camera` = ?",info.Camera)
    }
    if info.Scanner != "" {
        db = db.Where("`scanner` = ?",info.Scanner)
    }
    if info.Wwan != "" {
        db = db.Where("`wwan` = ?",info.Wwan)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&devices).Error
	return err, devices, total
}
