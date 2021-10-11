// 自动生成模板TestRunner
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TestRunner 结构体
// 如果含有time.Time 请自行import time包
type TestRunner struct {
      global.GVA_MODEL
      TestId  *int `json:"testId" form:"testId" gorm:"column:testId;comment:;type:int"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(128);"`
      Testcases  string `json:"testcases" form:"testcases" gorm:"column:testcases;comment:;type:varchar(1024);"`
      Owner  string `json:"owner" form:"owner" gorm:"column:owner;comment:;type:varchar(128);"`
      SerialNo  string `json:"serialNo" form:"serialNo" gorm:"column:serialNo;comment:;type:varchar(64);"`
}


// TableName TestRunner 表名
func (TestRunner) TableName() string {
  return "testRunner"
}

