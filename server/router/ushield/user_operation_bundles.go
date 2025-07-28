package ushield

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/middleware"
)

type UserOperationBundlesRouter struct{}

// InitUserOperationBundlesRouter 初始化 userOperationBundles表 路由信息
func (s *UserOperationBundlesRouter) InitUserOperationBundlesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userOperationBundlesRouter := Router.Group("userOperationBundles").Use(middleware.OperationRecord())
	userOperationBundlesRouterWithoutRecord := Router.Group("userOperationBundles")
	userOperationBundlesRouterWithoutAuth := PublicRouter.Group("userOperationBundles")
	{
		userOperationBundlesRouter.POST("createUserOperationBundles", userOperationBundlesApi.CreateUserOperationBundles)             // 新建userOperationBundles表
		userOperationBundlesRouter.DELETE("deleteUserOperationBundles", userOperationBundlesApi.DeleteUserOperationBundles)           // 删除userOperationBundles表
		userOperationBundlesRouter.DELETE("deleteUserOperationBundlesByIds", userOperationBundlesApi.DeleteUserOperationBundlesByIds) // 批量删除userOperationBundles表
		userOperationBundlesRouter.PUT("updateUserOperationBundles", userOperationBundlesApi.UpdateUserOperationBundles)              // 更新userOperationBundles表
	}
	{
		userOperationBundlesRouterWithoutRecord.GET("findUserOperationBundles", userOperationBundlesApi.FindUserOperationBundles)       // 根据ID获取userOperationBundles表
		userOperationBundlesRouterWithoutRecord.GET("getUserOperationBundlesList", userOperationBundlesApi.GetUserOperationBundlesList) // 获取userOperationBundles表列表
	}
	{
		userOperationBundlesRouterWithoutAuth.GET("getUserOperationBundlesPublic", userOperationBundlesApi.GetUserOperationBundlesPublic) // userOperationBundles表开放接口
	}
}
