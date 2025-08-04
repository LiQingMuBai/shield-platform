package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/router"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		systemRouter := router.RouterGroupApp.System
		systemRouter.InitUserUsdtDepositsRouter(privateGroup, publicGroup)
	}
	{
		ushieldRouter := router.RouterGroupApp.Ushield
		ushieldRouter.InitUserTrxSubscriptionsRouter(privateGroup, publicGroup)
		ushieldRouter.InitUserUsdtSubscriptionsRouter(privateGroup, publicGroup)
		ushieldRouter.InitUserPackageSubscriptionsRouter(privateGroup, publicGroup)
		ushieldRouter.InitTgUsersRouter(privateGroup, publicGroup)
		ushieldRouter.InitUserOperationBundlesRouter(privateGroup, publicGroup)
		ushieldRouter.InitUserAddressMonitorEventRouter(privateGroup, publicGroup)
		ushieldRouter.InitUserAddressMonitorRouter(privateGroup, publicGroup)
		ushieldRouter.InitUserTrxPlaceholdersRouter(privateGroup, publicGroup)
		ushieldRouter.InitUserUsdtPlaceholdersRouter(privateGroup, publicGroup)
		ushieldRouter.InitUserTrxDepositsRouter(privateGroup, publicGroup)
		ushieldRouter.InitUserUsdtDepositsRouter(privateGroup, publicGroup)
		ushieldRouter.InitSysAnnouncementsInfoRouter(privateGroup, publicGroup)
		ushieldRouter.InitUserEnergyOrdersRouter(privateGroup, publicGroup)
		ushieldRouter.InitUserBundleEnergyOrdersRouter(privateGroup, publicGroup) // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
		ushieldRouter.InitMerchantAddressMonitorEventRouter(privateGroup, publicGroup)
	}
}
