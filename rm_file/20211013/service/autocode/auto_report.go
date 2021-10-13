package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type ReportService struct {
}

// CreateReport 创建Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (reportService *ReportService) CreateReport(report autocode.Report) (err error) {
	err = global.GVA_DB.Create(&report).Error
	return err
}

// DeleteReport 删除Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (reportService *ReportService)DeleteReport(report autocode.Report) (err error) {
	err = global.GVA_DB.Delete(&report).Error
	return err
}

// DeleteReportByIds 批量删除Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (reportService *ReportService)DeleteReportByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Report{},"id in ?",ids.Ids).Error
	return err
}

// UpdateReport 更新Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (reportService *ReportService)UpdateReport(report autocode.Report) (err error) {
	err = global.GVA_DB.Save(&report).Error
	return err
}

// GetReport 根据id获取Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (reportService *ReportService)GetReport(id uint) (err error, report autocode.Report) {
	err = global.GVA_DB.Where("id = ?", id).First(&report).Error
	return
}

// GetReportInfoList 分页获取Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (reportService *ReportService)GetReportInfoList(info autoCodeReq.ReportSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Report{})
    var reports []autocode.Report
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
    if info.Result != "" {
        db = db.Where("`result` LIKE ?","%"+ info.Result+"%")
    }
    if info.Logcat != "" {
        db = db.Where("`logcat` LIKE ?","%"+ info.Logcat+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&reports).Error
	return err, reports, total
}
