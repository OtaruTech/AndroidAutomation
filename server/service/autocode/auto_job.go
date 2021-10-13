package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type JobService struct {
}

// CreateJob 创建Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService) CreateJob(job autocode.Job) (err error) {
	err = global.GVA_DB.Create(&job).Error
	return err
}

// DeleteJob 删除Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService)DeleteJob(job autocode.Job) (err error) {
	err = global.GVA_DB.Delete(&job).Error
	return err
}

// DeleteJobByIds 批量删除Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService)DeleteJobByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Job{},"id in ?",ids.Ids).Error
	return err
}

// UpdateJob 更新Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService)UpdateJob(job autocode.Job) (err error) {
	err = global.GVA_DB.Save(&job).Error
	return err
}

// GetJob 根据id获取Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService)GetJob(id uint) (err error, job autocode.Job) {
	err = global.GVA_DB.Where("id = ?", id).First(&job).Error
	return
}

// GetJobInfoList 分页获取Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService)GetJobInfoList(info autoCodeReq.JobSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Job{})
    var jobs []autocode.Job
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` = ?",info.Name)
    }
    if info.OtaPath != "" {
        db = db.Where("`otaPath` = ?",info.OtaPath)
    }
    if info.OtaFormat != "" {
        db = db.Where("`otaFormat` = ?",info.OtaFormat)
    }
    if info.Hour != nil {
        db = db.Where("`hour` = ?",info.Hour)
    }
    if info.Product != "" {
        db = db.Where("`product` = ?",info.Product)
    }
    if info.Testcases != "" {
        db = db.Where("`testcases` LIKE ?","%"+ info.Testcases+"%")
    }
    if info.Owner != "" {
        db = db.Where("`owner` = ?",info.Owner)
    }
    if info.OtaUrl != "" {
        db = db.Where("`otaUrl` = ?",info.OtaUrl)
    }
    if info.Enable != nil {
        db = db.Where("`enable` = ?",info.Enable)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&jobs).Error
	return err, jobs, total
}
