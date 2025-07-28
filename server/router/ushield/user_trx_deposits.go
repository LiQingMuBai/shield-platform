package ushield

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/middleware"
)

type UserTrxDepositsRouter struct{}

// InitUserTrxDepositsRouter 初始化 userTrxDeposits表 路由信息
func (s *UserTrxDepositsRouter) InitUserTrxDepositsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userTrxDepositsRouter := Router.Group("userTrxDeposits").Use(middleware.OperationRecord())
	userTrxDepositsRouterWithoutRecord := Router.Group("userTrxDeposits")
	userTrxDepositsRouterWithoutAuth := PublicRouter.Group("userTrxDeposits")
	{
		userTrxDepositsRouter.POST("createUserTrxDeposits", userTrxDepositsApi.CreateUserTrxDeposits)             // 新建userTrxDeposits表
		userTrxDepositsRouter.DELETE("deleteUserTrxDeposits", userTrxDepositsApi.DeleteUserTrxDeposits)           // 删除userTrxDeposits表
		userTrxDepositsRouter.DELETE("deleteUserTrxDepositsByIds", userTrxDepositsApi.DeleteUserTrxDepositsByIds) // 批量删除userTrxDeposits表
		userTrxDepositsRouter.PUT("updateUserTrxDeposits", userTrxDepositsApi.UpdateUserTrxDeposits)              // 更新userTrxDeposits表
	}
	{
		userTrxDepositsRouterWithoutRecord.GET("findUserTrxDeposits", userTrxDepositsApi.FindUserTrxDeposits)       // 根据ID获取userTrxDeposits表
		userTrxDepositsRouterWithoutRecord.GET("getUserTrxDepositsList", userTrxDepositsApi.GetUserTrxDepositsList) // 获取userTrxDeposits表列表
	}
	{
		userTrxDepositsRouterWithoutAuth.GET("getUserTrxDepositsPublic", userTrxDepositsApi.GetUserTrxDepositsPublic) // userTrxDeposits表开放接口
	}
}
