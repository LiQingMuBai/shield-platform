package response

import "github.com/ushield/aurora-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
