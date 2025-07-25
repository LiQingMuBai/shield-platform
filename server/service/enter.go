package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ushield"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	UshieldServiceGroup ushield.ServiceGroup
}
