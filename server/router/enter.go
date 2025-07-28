package router

import (
	"github.com/ushield/aurora-admin/server/router/example"
	"github.com/ushield/aurora-admin/server/router/system"
	"github.com/ushield/aurora-admin/server/router/ushield"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Ushield ushield.RouterGroup
}
