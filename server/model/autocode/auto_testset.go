// 自动生成模板Testset
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Testset 结构体
// 如果含有time.Time 请自行import time包
type Testset struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(128);"`
      Testcases  string `json:"testcases" form:"testcases" gorm:"column:testcases;comment:;type:text(4096);"`
}


// TableName Testset 表名
func (Testset) TableName() string {
  return "testset"
}

