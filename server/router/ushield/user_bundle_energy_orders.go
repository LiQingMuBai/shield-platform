package ushield

import (
	"github.com/ushield/aurora-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserBundleEnergyOrdersRouter struct {}

// InitUserBundleEnergyOrdersRouter 初始化 userBundleEnergyOrders表 路由信息
func (s *UserBundleEnergyOrdersRouter) InitUserBundleEnergyOrdersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	userBundleEnergyOrdersRouter := Router.Group("userBundleEnergyOrders").Use(middleware.OperationRecord())
	userBundleEnergyOrdersRouterWithoutRecord := Router.Group("userBundleEnergyOrders")
	userBundleEnergyOrdersRouterWithoutAuth := PublicRouter.Group("userBundleEnergyOrders")
	{
		userBundleEnergyOrdersRouter.POST("createUserBundleEnergyOrders", userBundleEnergyOrdersApi.CreateUserBundleEnergyOrders)   // 新建userBundleEnergyOrders表
		userBundleEnergyOrdersRouter.DELETE("deleteUserBundleEnergyOrders", userBundleEnergyOrdersApi.DeleteUserBundleEnergyOrders) // 删除userBundleEnergyOrders表
		userBundleEnergyOrdersRouter.DELETE("deleteUserBundleEnergyOrdersByIds", userBundleEnergyOrdersApi.DeleteUserBundleEnergyOrdersByIds) // 批量删除userBundleEnergyOrders表
		userBundleEnergyOrdersRouter.PUT("updateUserBundleEnergyOrders", userBundleEnergyOrdersApi.UpdateUserBundleEnergyOrders)    // 更新userBundleEnergyOrders表
	}
	{
		userBundleEnergyOrdersRouterWithoutRecord.GET("findUserBundleEnergyOrders", userBundleEnergyOrdersApi.FindUserBundleEnergyOrders)        // 根据ID获取userBundleEnergyOrders表
		userBundleEnergyOrdersRouterWithoutRecord.GET("getUserBundleEnergyOrdersList", userBundleEnergyOrdersApi.GetUserBundleEnergyOrdersList)  // 获取userBundleEnergyOrders表列表
	}
	{
	    userBundleEnergyOrdersRouterWithoutAuth.GET("getUserBundleEnergyOrdersPublic", userBundleEnergyOrdersApi.GetUserBundleEnergyOrdersPublic)  // userBundleEnergyOrders表开放接口
	}
}
