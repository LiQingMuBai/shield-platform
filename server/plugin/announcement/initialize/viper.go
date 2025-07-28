package initialize

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/plugin/announcement/plugin"
	"go.uber.org/zap"
)

func Viper() {
	err := global.GVA_VP.UnmarshalKey("announcement", &plugin.Config)
	if err != nil {
		err = errors.Wrap(err, "初始化配置文件失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
