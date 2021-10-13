// 自动生成模板Report
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Report 结构体
// 如果含有time.Time 请自行import time包
type Report struct {
      global.GVA_MODEL
      TestId  *int `json:"testId" form:"testId" gorm:"column:testId;comment:;type:int"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(128);"`
      Testcases  string `json:"testcases" form:"testcases" gorm:"column:testcases;comment:;type:varchar(1024);"`
      Result  string `json:"result" form:"result" gorm:"column:result;comment:;type:text(65536);"`
      Logcat  string `json:"logcat" form:"logcat" gorm:"column:logcat;comment:;type:varchar(256);"`
}


// TableName Report 表名
func (Report) TableName() string {
  return "report"
}

