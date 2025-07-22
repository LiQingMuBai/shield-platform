package ushield

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	UserTrxSubscriptionsApi
	UserUsdtSubscriptionsApi
	UserPackageSubscriptionsApi
	TgUsersApi
	UserOperationBundlesApi
	UserAddressMonitorEventApi
	UserAddressMonitorApi
	UserTrxPlaceholdersApi
	UserUsdtPlaceholdersApi
	UserTrxDepositsApi
	UserUsdtDepositsApi
}

var (
	userTrxSubscriptionsService     = service.ServiceGroupApp.UshieldServiceGroup.UserTrxSubscriptionsService
	userUsdtSubscriptionsService    = service.ServiceGroupApp.UshieldServiceGroup.UserUsdtSubscriptionsService
	userPackageSubscriptionsService = service.ServiceGroupApp.UshieldServiceGroup.UserPackageSubscriptionsService
	tgUsersService                  = service.ServiceGroupApp.UshieldServiceGroup.TgUsersService
	userOperationBundlesService     = service.ServiceGroupApp.UshieldServiceGroup.UserOperationBundlesService
	userAddressMonitorEventService  = service.ServiceGroupApp.UshieldServiceGroup.UserAddressMonitorEventService
	userAddressMonitorService       = service.ServiceGroupApp.UshieldServiceGroup.UserAddressMonitorService
	userTrxPlaceholdersService      = service.ServiceGroupApp.UshieldServiceGroup.UserTrxPlaceholdersService
	userUsdtPlaceholdersService     = service.ServiceGroupApp.UshieldServiceGroup.UserUsdtPlaceholdersService
	userTrxDepositsService          = service.ServiceGroupApp.UshieldServiceGroup.UserTrxDepositsService
	userUsdtDepositsService         = service.ServiceGroupApp.UshieldServiceGroup.UserUsdtDepositsService
)
