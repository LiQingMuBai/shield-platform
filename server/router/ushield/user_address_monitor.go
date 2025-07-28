package ushield

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/middleware"
)

type UserAddressMonitorRouter struct{}

// InitUserAddressMonitorRouter 初始化 userAddressMonitor表 路由信息
func (s *UserAddressMonitorRouter) InitUserAddressMonitorRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userAddressMonitorRouter := Router.Group("userAddressMonitor").Use(middleware.OperationRecord())
	userAddressMonitorRouterWithoutRecord := Router.Group("userAddressMonitor")
	userAddressMonitorRouterWithoutAuth := PublicRouter.Group("userAddressMonitor")
	{
		userAddressMonitorRouter.POST("createUserAddressMonitor", userAddressMonitorApi.CreateUserAddressMonitor)             // 新建userAddressMonitor表
		userAddressMonitorRouter.DELETE("deleteUserAddressMonitor", userAddressMonitorApi.DeleteUserAddressMonitor)           // 删除userAddressMonitor表
		userAddressMonitorRouter.DELETE("deleteUserAddressMonitorByIds", userAddressMonitorApi.DeleteUserAddressMonitorByIds) // 批量删除userAddressMonitor表
		userAddressMonitorRouter.PUT("updateUserAddressMonitor", userAddressMonitorApi.UpdateUserAddressMonitor)              // 更新userAddressMonitor表
	}
	{
		userAddressMonitorRouterWithoutRecord.GET("findUserAddressMonitor", userAddressMonitorApi.FindUserAddressMonitor)       // 根据ID获取userAddressMonitor表
		userAddressMonitorRouterWithoutRecord.GET("getUserAddressMonitorList", userAddressMonitorApi.GetUserAddressMonitorList) // 获取userAddressMonitor表列表
	}
	{
		userAddressMonitorRouterWithoutAuth.GET("getUserAddressMonitorPublic", userAddressMonitorApi.GetUserAddressMonitorPublic) // userAddressMonitor表开放接口
	}
}
