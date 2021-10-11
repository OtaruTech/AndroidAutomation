package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type TestRunnerService struct {
}

// CreateTestRunner 创建TestRunner记录
// Author [piexlmax](https://github.com/piexlmax)
func (testRunnerService *TestRunnerService) CreateTestRunner(testRunner autocode.TestRunner) (err error) {
	err = global.GVA_DB.Create(&testRunner).Error
	return err
}

// DeleteTestRunner 删除TestRunner记录
// Author [piexlmax](https://github.com/piexlmax)
func (testRunnerService *TestRunnerService)DeleteTestRunner(testRunner autocode.TestRunner) (err error) {
	err = global.GVA_DB.Delete(&testRunner).Error
	return err
}

// DeleteTestRunnerByIds 批量删除TestRunner记录
// Author [piexlmax](https://github.com/piexlmax)
func (testRunnerService *TestRunnerService)DeleteTestRunnerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.TestRunner{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTestRunner 更新TestRunner记录
// Author [piexlmax](https://github.com/piexlmax)
func (testRunnerService *TestRunnerService)UpdateTestRunner(testRunner autocode.TestRunner) (err error) {
	err = global.GVA_DB.Save(&testRunner).Error
	return err
}

// GetTestRunner 根据id获取TestRunner记录
// Author [piexlmax](https://github.com/piexlmax)
func (testRunnerService *TestRunnerService)GetTestRunner(id uint) (err error, testRunner autocode.TestRunner) {
	err = global.GVA_DB.Where("id = ?", id).First(&testRunner).Error
	return
}

// GetTestRunnerInfoList 分页获取TestRunner记录
// Author [piexlmax](https://github.com/piexlmax)
func (testRunnerService *TestRunnerService)GetTestRunnerInfoList(info autoCodeReq.TestRunnerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.TestRunner{})
    var testRunners []autocode.TestRunner
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.TestId != nil {
        db = db.Where("`testId` = ?",info.TestId)
    }
    if info.Name != "" {
        db = db.Where("`name` = ?",info.Name)
    }
    if info.Testcases != "" {
        db = db.Where("`testcases` LIKE ?","%"+ info.Testcases+"%")
    }
    if info.Owner != "" {
        db = db.Where("`owner` = ?",info.Owner)
    }
    if info.SerialNo != "" {
        db = db.Where("`serialNo` = ?",info.SerialNo)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&testRunners).Error
	return err, testRunners, total
}
