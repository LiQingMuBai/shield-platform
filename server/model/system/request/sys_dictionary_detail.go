package request

import (
	"github.com/ushield/aurora-admin/server/model/common/request"
	"github.com/ushield/aurora-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
