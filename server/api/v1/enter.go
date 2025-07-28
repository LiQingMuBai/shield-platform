package v1

import (
	"github.com/ushield/aurora-admin/server/api/v1/example"
	"github.com/ushield/aurora-admin/server/api/v1/system"
	"github.com/ushield/aurora-admin/server/api/v1/ushield"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	UshieldApiGroup ushield.ApiGroup
}
