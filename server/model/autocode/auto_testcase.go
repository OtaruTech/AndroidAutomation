// 自动生成模板Testcase
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Testcase 结构体
// 如果含有time.Time 请自行import time包
type Testcase struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(128);"`
      Component  string `json:"component" form:"component" gorm:"column:component;comment:;type:varchar(64);"`
      Tags  string `json:"tags" form:"tags" gorm:"column:tags;comment:;type:varchar(128);"`
      Timeout  *int `json:"timeout" form:"timeout" gorm:"column:timeout;comment:;type:int"`
}


// TableName Testcase 表名
func (Testcase) TableName() string {
  return "testcase"
}

