// 自动生成模板DeviceConsole
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// DeviceConsole 结构体
// 如果含有time.Time 请自行import time包
type DeviceConsole struct {
      global.GVA_MODEL
      SerialNo  string `json:"serialNo" form:"serialNo" gorm:"column:serialNo;comment:;type:varchar(64);"`
      Console  string `json:"console" form:"console" gorm:"column:console;comment:;type:text(65536);"`
}


// TableName DeviceConsole 表名
func (DeviceConsole) TableName() string {
  return "deviceConsole"
}

