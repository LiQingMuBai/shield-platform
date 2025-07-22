package ushield

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserPackageSubscriptionsRouter struct {}

// InitUserPackageSubscriptionsRouter 初始化 userPackageSubscriptions表 路由信息
func (s *UserPackageSubscriptionsRouter) InitUserPackageSubscriptionsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	userPackageSubscriptionsRouter := Router.Group("userPackageSubscriptions").Use(middleware.OperationRecord())
	userPackageSubscriptionsRouterWithoutRecord := Router.Group("userPackageSubscriptions")
	userPackageSubscriptionsRouterWithoutAuth := PublicRouter.Group("userPackageSubscriptions")
	{
		userPackageSubscriptionsRouter.POST("createUserPackageSubscriptions", userPackageSubscriptionsApi.CreateUserPackageSubscriptions)   // 新建userPackageSubscriptions表
		userPackageSubscriptionsRouter.DELETE("deleteUserPackageSubscriptions", userPackageSubscriptionsApi.DeleteUserPackageSubscriptions) // 删除userPackageSubscriptions表
		userPackageSubscriptionsRouter.DELETE("deleteUserPackageSubscriptionsByIds", userPackageSubscriptionsApi.DeleteUserPackageSubscriptionsByIds) // 批量删除userPackageSubscriptions表
		userPackageSubscriptionsRouter.PUT("updateUserPackageSubscriptions", userPackageSubscriptionsApi.UpdateUserPackageSubscriptions)    // 更新userPackageSubscriptions表
	}
	{
		userPackageSubscriptionsRouterWithoutRecord.GET("findUserPackageSubscriptions", userPackageSubscriptionsApi.FindUserPackageSubscriptions)        // 根据ID获取userPackageSubscriptions表
		userPackageSubscriptionsRouterWithoutRecord.GET("getUserPackageSubscriptionsList", userPackageSubscriptionsApi.GetUserPackageSubscriptionsList)  // 获取userPackageSubscriptions表列表
	}
	{
	    userPackageSubscriptionsRouterWithoutAuth.GET("getUserPackageSubscriptionsPublic", userPackageSubscriptionsApi.GetUserPackageSubscriptionsPublic)  // userPackageSubscriptions表开放接口
	}
}
