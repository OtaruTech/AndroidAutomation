package automation

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AndroidApi
}

var androidService = service.ServiceGroupApp.AutomationServiceGroup.AndroidService
