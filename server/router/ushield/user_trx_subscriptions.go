package ushield

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/middleware"
)

type UserTrxSubscriptionsRouter struct{}

// InitUserTrxSubscriptionsRouter 初始化 userTrxSubscriptions表 路由信息
func (s *UserTrxSubscriptionsRouter) InitUserTrxSubscriptionsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userTrxSubscriptionsRouter := Router.Group("userTrxSubscriptions").Use(middleware.OperationRecord())
	userTrxSubscriptionsRouterWithoutRecord := Router.Group("userTrxSubscriptions")
	userTrxSubscriptionsRouterWithoutAuth := PublicRouter.Group("userTrxSubscriptions")
	{
		userTrxSubscriptionsRouter.POST("createUserTrxSubscriptions", userTrxSubscriptionsApi.CreateUserTrxSubscriptions)             // 新建userTrxSubscriptions表
		userTrxSubscriptionsRouter.DELETE("deleteUserTrxSubscriptions", userTrxSubscriptionsApi.DeleteUserTrxSubscriptions)           // 删除userTrxSubscriptions表
		userTrxSubscriptionsRouter.DELETE("deleteUserTrxSubscriptionsByIds", userTrxSubscriptionsApi.DeleteUserTrxSubscriptionsByIds) // 批量删除userTrxSubscriptions表
		userTrxSubscriptionsRouter.PUT("updateUserTrxSubscriptions", userTrxSubscriptionsApi.UpdateUserTrxSubscriptions)              // 更新userTrxSubscriptions表
	}
	{
		userTrxSubscriptionsRouterWithoutRecord.GET("findUserTrxSubscriptions", userTrxSubscriptionsApi.FindUserTrxSubscriptions)       // 根据ID获取userTrxSubscriptions表
		userTrxSubscriptionsRouterWithoutRecord.GET("getUserTrxSubscriptionsList", userTrxSubscriptionsApi.GetUserTrxSubscriptionsList) // 获取userTrxSubscriptions表列表
	}
	{
		userTrxSubscriptionsRouterWithoutAuth.GET("getUserTrxSubscriptionsPublic", userTrxSubscriptionsApi.GetUserTrxSubscriptionsPublic) // userTrxSubscriptions表开放接口
	}
}
