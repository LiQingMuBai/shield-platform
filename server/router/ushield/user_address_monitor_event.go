package ushield

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/middleware"
)

type UserAddressMonitorEventRouter struct{}

// InitUserAddressMonitorEventRouter 初始化 userAddressMonitorEvent表 路由信息
func (s *UserAddressMonitorEventRouter) InitUserAddressMonitorEventRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userAddressMonitorEventRouter := Router.Group("userAddressMonitorEvent").Use(middleware.OperationRecord())
	userAddressMonitorEventRouterWithoutRecord := Router.Group("userAddressMonitorEvent")
	userAddressMonitorEventRouterWithoutAuth := PublicRouter.Group("userAddressMonitorEvent")
	{
		userAddressMonitorEventRouter.POST("createUserAddressMonitorEvent", userAddressMonitorEventApi.CreateUserAddressMonitorEvent)             // 新建userAddressMonitorEvent表
		userAddressMonitorEventRouter.DELETE("deleteUserAddressMonitorEvent", userAddressMonitorEventApi.DeleteUserAddressMonitorEvent)           // 删除userAddressMonitorEvent表
		userAddressMonitorEventRouter.DELETE("deleteUserAddressMonitorEventByIds", userAddressMonitorEventApi.DeleteUserAddressMonitorEventByIds) // 批量删除userAddressMonitorEvent表
		userAddressMonitorEventRouter.PUT("updateUserAddressMonitorEvent", userAddressMonitorEventApi.UpdateUserAddressMonitorEvent)              // 更新userAddressMonitorEvent表
	}
	{
		userAddressMonitorEventRouterWithoutRecord.GET("findUserAddressMonitorEvent", userAddressMonitorEventApi.FindUserAddressMonitorEvent)       // 根据ID获取userAddressMonitorEvent表
		userAddressMonitorEventRouterWithoutRecord.GET("getUserAddressMonitorEventList", userAddressMonitorEventApi.GetUserAddressMonitorEventList) // 获取userAddressMonitorEvent表列表
	}
	{
		userAddressMonitorEventRouterWithoutAuth.GET("getUserAddressMonitorEventPublic", userAddressMonitorEventApi.GetUserAddressMonitorEventPublic) // userAddressMonitorEvent表开放接口
	}
}
