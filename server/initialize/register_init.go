package initialize

import (
	_ "github.com/ushield/aurora-admin/server/source/example"
	_ "github.com/ushield/aurora-admin/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
