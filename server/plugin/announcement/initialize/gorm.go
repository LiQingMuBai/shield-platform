package initialize

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/plugin/announcement/model"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.WithContext(ctx).AutoMigrate(
		new(model.Info),
	)
	if err != nil {
		err = errors.Wrap(err, "注册表失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
