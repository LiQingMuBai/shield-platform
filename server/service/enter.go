package service

import (
	"github.com/ushield/aurora-admin/server/service/example"
	"github.com/ushield/aurora-admin/server/service/system"
	"github.com/ushield/aurora-admin/server/service/ushield"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	UshieldServiceGroup ushield.ServiceGroup
}
