// 自动生成模板Job
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Job 结构体
// 如果含有time.Time 请自行import time包
type Job struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(128);"`
      OtaPath  string `json:"otaPath" form:"otaPath" gorm:"column:otaPath;comment:;type:varchar(512);"`
      OtaFormat  string `json:"otaFormat" form:"otaFormat" gorm:"column:otaFormat;comment:;type:varchar(64);"`
      FileFormat  string `json:"fileFormat" form:"fileFormat" gorm:"column:fileFormat;comment:;type:varchar(64);"`
      Hour  *int `json:"hour" form:"hour" gorm:"column:hour;comment:;type:smallint"`
      Product  string `json:"product" form:"product" gorm:"column:product;comment:;type:varchar(64);"`
      Testcases  string `json:"testcases" form:"testcases" gorm:"column:testcases;comment:;type:varchar(4096);"`
      Owner  string `json:"owner" form:"owner" gorm:"column:owner;comment:;type:varchar(64);"`
      OtaUrl  string `json:"otaUrl" form:"otaUrl" gorm:"column:otaUrl;comment:;type:varchar(512);"`
      Enable  *bool `json:"enable" form:"enable" gorm:"column:enable;comment:;type:tinyint"`
}


// TableName Job 表名
func (Job) TableName() string {
  return "job"
}

