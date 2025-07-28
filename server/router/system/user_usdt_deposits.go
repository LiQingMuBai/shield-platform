package system

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/middleware"
)

type UserUsdtDepositsRouter struct{}

// InitUserUsdtDepositsRouter 初始化 userUsdtDeposits表 路由信息
func (s *UserUsdtDepositsRouter) InitUserUsdtDepositsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userUsdtDepositsRouter := Router.Group("userUsdtDeposits").Use(middleware.OperationRecord())
	userUsdtDepositsRouterWithoutRecord := Router.Group("userUsdtDeposits")
	userUsdtDepositsRouterWithoutAuth := PublicRouter.Group("userUsdtDeposits")
	{
		userUsdtDepositsRouter.POST("createUserUsdtDeposits", userUsdtDepositsApi.CreateUserUsdtDeposits)             // 新建userUsdtDeposits表
		userUsdtDepositsRouter.DELETE("deleteUserUsdtDeposits", userUsdtDepositsApi.DeleteUserUsdtDeposits)           // 删除userUsdtDeposits表
		userUsdtDepositsRouter.DELETE("deleteUserUsdtDepositsByIds", userUsdtDepositsApi.DeleteUserUsdtDepositsByIds) // 批量删除userUsdtDeposits表
		userUsdtDepositsRouter.PUT("updateUserUsdtDeposits", userUsdtDepositsApi.UpdateUserUsdtDeposits)              // 更新userUsdtDeposits表
	}
	{
		userUsdtDepositsRouterWithoutRecord.GET("findUserUsdtDeposits", userUsdtDepositsApi.FindUserUsdtDeposits)       // 根据ID获取userUsdtDeposits表
		userUsdtDepositsRouterWithoutRecord.GET("getUserUsdtDepositsList", userUsdtDepositsApi.GetUserUsdtDepositsList) // 获取userUsdtDeposits表列表
	}
	{
		userUsdtDepositsRouterWithoutAuth.GET("getUserUsdtDepositsPublic", userUsdtDepositsApi.GetUserUsdtDepositsPublic) // userUsdtDeposits表开放接口
	}
}
