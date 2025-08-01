package ushield

import (
	"github.com/ushield/aurora-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserEnergyOrdersRouter struct {}

// InitUserEnergyOrdersRouter 初始化 userEnergyOrders表 路由信息
func (s *UserEnergyOrdersRouter) InitUserEnergyOrdersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	userEnergyOrdersRouter := Router.Group("userEnergyOrders").Use(middleware.OperationRecord())
	userEnergyOrdersRouterWithoutRecord := Router.Group("userEnergyOrders")
	userEnergyOrdersRouterWithoutAuth := PublicRouter.Group("userEnergyOrders")
	{
		userEnergyOrdersRouter.POST("createUserEnergyOrders", userEnergyOrdersApi.CreateUserEnergyOrders)   // 新建userEnergyOrders表
		userEnergyOrdersRouter.DELETE("deleteUserEnergyOrders", userEnergyOrdersApi.DeleteUserEnergyOrders) // 删除userEnergyOrders表
		userEnergyOrdersRouter.DELETE("deleteUserEnergyOrdersByIds", userEnergyOrdersApi.DeleteUserEnergyOrdersByIds) // 批量删除userEnergyOrders表
		userEnergyOrdersRouter.PUT("updateUserEnergyOrders", userEnergyOrdersApi.UpdateUserEnergyOrders)    // 更新userEnergyOrders表
	}
	{
		userEnergyOrdersRouterWithoutRecord.GET("findUserEnergyOrders", userEnergyOrdersApi.FindUserEnergyOrders)        // 根据ID获取userEnergyOrders表
		userEnergyOrdersRouterWithoutRecord.GET("getUserEnergyOrdersList", userEnergyOrdersApi.GetUserEnergyOrdersList)  // 获取userEnergyOrders表列表
	}
	{
	    userEnergyOrdersRouterWithoutAuth.GET("getUserEnergyOrdersPublic", userEnergyOrdersApi.GetUserEnergyOrdersPublic)  // userEnergyOrders表开放接口
	}
}
