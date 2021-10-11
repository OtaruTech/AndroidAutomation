package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type TestcaseService struct {
}

// CreateTestcase 创建Testcase记录
// Author [piexlmax](https://github.com/piexlmax)
func (testcaseService *TestcaseService) CreateTestcase(testcase autocode.Testcase) (err error) {
	err = global.GVA_DB.Create(&testcase).Error
	return err
}

// DeleteTestcase 删除Testcase记录
// Author [piexlmax](https://github.com/piexlmax)
func (testcaseService *TestcaseService)DeleteTestcase(testcase autocode.Testcase) (err error) {
	err = global.GVA_DB.Delete(&testcase).Error
	return err
}

// DeleteTestcaseByIds 批量删除Testcase记录
// Author [piexlmax](https://github.com/piexlmax)
func (testcaseService *TestcaseService)DeleteTestcaseByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Testcase{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTestcase 更新Testcase记录
// Author [piexlmax](https://github.com/piexlmax)
func (testcaseService *TestcaseService)UpdateTestcase(testcase autocode.Testcase) (err error) {
	err = global.GVA_DB.Save(&testcase).Error
	return err
}

// GetTestcase 根据id获取Testcase记录
// Author [piexlmax](https://github.com/piexlmax)
func (testcaseService *TestcaseService)GetTestcase(id uint) (err error, testcase autocode.Testcase) {
	err = global.GVA_DB.Where("id = ?", id).First(&testcase).Error
	return
}

// GetTestcaseInfoList 分页获取Testcase记录
// Author [piexlmax](https://github.com/piexlmax)
func (testcaseService *TestcaseService)GetTestcaseInfoList(info autoCodeReq.TestcaseSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Testcase{})
    var testcases []autocode.Testcase
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` = ?",info.Name)
    }
    if info.Component != "" {
        db = db.Where("`component` = ?",info.Component)
    }
    if info.Tags != "" {
        db = db.Where("`tags` LIKE ?","%"+ info.Tags+"%")
    }
    if info.Timeout != nil {
        db = db.Where("`timeout` > ?",info.Timeout)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&testcases).Error
	return err, testcases, total
}
