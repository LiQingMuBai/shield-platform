package ushield

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserUsdtSubscriptionsRouter struct {}

// InitUserUsdtSubscriptionsRouter 初始化 userUsdtSubscriptions表 路由信息
func (s *UserUsdtSubscriptionsRouter) InitUserUsdtSubscriptionsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	userUsdtSubscriptionsRouter := Router.Group("userUsdtSubscriptions").Use(middleware.OperationRecord())
	userUsdtSubscriptionsRouterWithoutRecord := Router.Group("userUsdtSubscriptions")
	userUsdtSubscriptionsRouterWithoutAuth := PublicRouter.Group("userUsdtSubscriptions")
	{
		userUsdtSubscriptionsRouter.POST("createUserUsdtSubscriptions", userUsdtSubscriptionsApi.CreateUserUsdtSubscriptions)   // 新建userUsdtSubscriptions表
		userUsdtSubscriptionsRouter.DELETE("deleteUserUsdtSubscriptions", userUsdtSubscriptionsApi.DeleteUserUsdtSubscriptions) // 删除userUsdtSubscriptions表
		userUsdtSubscriptionsRouter.DELETE("deleteUserUsdtSubscriptionsByIds", userUsdtSubscriptionsApi.DeleteUserUsdtSubscriptionsByIds) // 批量删除userUsdtSubscriptions表
		userUsdtSubscriptionsRouter.PUT("updateUserUsdtSubscriptions", userUsdtSubscriptionsApi.UpdateUserUsdtSubscriptions)    // 更新userUsdtSubscriptions表
	}
	{
		userUsdtSubscriptionsRouterWithoutRecord.GET("findUserUsdtSubscriptions", userUsdtSubscriptionsApi.FindUserUsdtSubscriptions)        // 根据ID获取userUsdtSubscriptions表
		userUsdtSubscriptionsRouterWithoutRecord.GET("getUserUsdtSubscriptionsList", userUsdtSubscriptionsApi.GetUserUsdtSubscriptionsList)  // 获取userUsdtSubscriptions表列表
	}
	{
	    userUsdtSubscriptionsRouterWithoutAuth.GET("getUserUsdtSubscriptionsPublic", userUsdtSubscriptionsApi.GetUserUsdtSubscriptionsPublic)  // userUsdtSubscriptions表开放接口
	}
}
