package initialize

import (
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/system"
	"github.com/ushield/aurora-admin/server/model/ushield"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(system.UserUsdtDeposits{}, ushield.UserTrxSubscriptions{}, ushield.UserUsdtSubscriptions{}, ushield.UserPackageSubscriptions{}, ushield.TgUsers{}, ushield.UserOperationBundles{}, ushield.UserAddressMonitorEvent{}, ushield.UserAddressMonitor{}, ushield.UserTrxPlaceholders{}, ushield.UserTrxDeposits{}, ushield.UserUsdtDeposits{}, ushield.SysAnnouncementsInfo{}, ushield.UserEnergyOrders{}, ushield.UserBundleEnergyOrders{}, ushield.MerchantAddressMonitorEvent{})
	if err != nil {
		return err
	}
	return nil
}
