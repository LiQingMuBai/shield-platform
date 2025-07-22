package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ushield"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(system.UserUsdtDeposits{}, ushield.UserTrxSubscriptions{}, ushield.UserUsdtSubscriptions{}, ushield.UserPackageSubscriptions{}, ushield.TgUsers{}, ushield.UserOperationBundles{}, ushield.UserAddressMonitorEvent{}, ushield.UserAddressMonitor{}, ushield.UserTrxPlaceholders{}, ushield.UserTrxDeposits{}, ushield.UserUsdtDeposits{})
	if err != nil {
		return err
	}
	return nil
}
