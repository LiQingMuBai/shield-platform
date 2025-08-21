package request

import (
	"github.com/ushield/aurora-admin/server/model/common/request"
)

type MerchantAddressMonitorEventSearch struct {
	request.PageInfo
}
type MerchantAddressMonitorEventReq struct {
	Callback string `json:"callback_url"`
	Address  string `json:"address"`
}
