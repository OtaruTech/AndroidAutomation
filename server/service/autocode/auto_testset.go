package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type TestsetService struct {
}

// CreateTestset 创建Testset记录
// Author [piexlmax](https://github.com/piexlmax)
func (testsetService *TestsetService) CreateTestset(testset autocode.Testset) (err error) {
	err = global.GVA_DB.Create(&testset).Error
	return err
}

// DeleteTestset 删除Testset记录
// Author [piexlmax](https://github.com/piexlmax)
func (testsetService *TestsetService)DeleteTestset(testset autocode.Testset) (err error) {
	err = global.GVA_DB.Delete(&testset).Error
	return err
}

// DeleteTestsetByIds 批量删除Testset记录
// Author [piexlmax](https://github.com/piexlmax)
func (testsetService *TestsetService)DeleteTestsetByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Testset{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTestset 更新Testset记录
// Author [piexlmax](https://github.com/piexlmax)
func (testsetService *TestsetService)UpdateTestset(testset autocode.Testset) (err error) {
	err = global.GVA_DB.Save(&testset).Error
	return err
}

// GetTestset 根据id获取Testset记录
// Author [piexlmax](https://github.com/piexlmax)
func (testsetService *TestsetService)GetTestset(id uint) (err error, testset autocode.Testset) {
	err = global.GVA_DB.Where("id = ?", id).First(&testset).Error
	return
}

// GetTestsetInfoList 分页获取Testset记录
// Author [piexlmax](https://github.com/piexlmax)
func (testsetService *TestsetService)GetTestsetInfoList(info autoCodeReq.TestsetSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Testset{})
    var testsets []autocode.Testset
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Testcases != "" {
        db = db.Where("`testcases` LIKE ?","%"+ info.Testcases+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&testsets).Error
	return err, testsets, total
}
