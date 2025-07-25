package ushield

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	UserTrxSubscriptionsRouter
	UserUsdtSubscriptionsRouter
	UserPackageSubscriptionsRouter
	TgUsersRouter
	UserOperationBundlesRouter
	UserAddressMonitorEventRouter
	UserAddressMonitorRouter
	UserTrxPlaceholdersRouter
	UserUsdtPlaceholdersRouter
	UserTrxDepositsRouter
	UserUsdtDepositsRouter
}

var (
	userTrxSubscriptionsApi     = api.ApiGroupApp.UshieldApiGroup.UserTrxSubscriptionsApi
	userUsdtSubscriptionsApi    = api.ApiGroupApp.UshieldApiGroup.UserUsdtSubscriptionsApi
	userPackageSubscriptionsApi = api.ApiGroupApp.UshieldApiGroup.UserPackageSubscriptionsApi
	tgUsersApi                  = api.ApiGroupApp.UshieldApiGroup.TgUsersApi
	userOperationBundlesApi     = api.ApiGroupApp.UshieldApiGroup.UserOperationBundlesApi
	userAddressMonitorEventApi  = api.ApiGroupApp.UshieldApiGroup.UserAddressMonitorEventApi
	userAddressMonitorApi       = api.ApiGroupApp.UshieldApiGroup.UserAddressMonitorApi
	userTrxPlaceholdersApi      = api.ApiGroupApp.UshieldApiGroup.UserTrxPlaceholdersApi
	userUsdtPlaceholdersApi     = api.ApiGroupApp.UshieldApiGroup.UserUsdtPlaceholdersApi
	userTrxDepositsApi          = api.ApiGroupApp.UshieldApiGroup.UserTrxDepositsApi
	userUsdtDepositsApi         = api.ApiGroupApp.UshieldApiGroup.UserUsdtDepositsApi
)
