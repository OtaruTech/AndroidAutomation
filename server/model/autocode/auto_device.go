// 自动生成模板Device
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Device 结构体
// 如果含有time.Time 请自行import time包
type Device struct {
      global.GVA_MODEL
      SerialNo  string `json:"serialNo" form:"serialNo" gorm:"column:serialNo;comment:;type:varchar(32);"`
      Model  string `json:"model" form:"model" gorm:"column:model;comment:;type:varchar(32);"`
      Product  string `json:"product" form:"product" gorm:"column:product;comment:;type:varchar(32);"`
      Version  string `json:"version" form:"version" gorm:"column:version;comment:;type:varchar(8);"`
      BuildId  string `json:"buildId" form:"buildId" gorm:"column:buildId;comment:;type:varchar(32);"`
      BuildType  string `json:"buildType" form:"buildType" gorm:"column:buildType;comment:;type:varchar(12);"`
      Sku  string `json:"sku" form:"sku" gorm:"column:sku;comment:;type:varchar(32);"`
      Sdk  string `json:"sdk" form:"sdk" gorm:"column:sdk;comment:;type:varchar(32);"`
      Security  string `json:"security" form:"security" gorm:"column:security;comment:;type:varchar(32);"`
      Api  string `json:"api" form:"api" gorm:"column:api;comment:;type:varchar(32);"`
      Config  string `json:"config" form:"config" gorm:"column:config;comment:;type:varchar(32);"`
      Camera  string `json:"camera" form:"camera" gorm:"column:camera;comment:;type:varchar(12);"`
      Scanner  string `json:"scanner" form:"scanner" gorm:"column:scanner;comment:;type:varchar(12);"`
      Wwan  string `json:"wwan" form:"wwan" gorm:"column:wwan;comment:;type:varchar(12);"`
}

// TableName Device 表名
func (Device) TableName() string {
  return "Device"
}

