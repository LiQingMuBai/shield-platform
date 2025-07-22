package ushield

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserTrxPlaceholdersRouter struct {}

// InitUserTrxPlaceholdersRouter 初始化 userTrxPlaceholders表 路由信息
func (s *UserTrxPlaceholdersRouter) InitUserTrxPlaceholdersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	userTrxPlaceholdersRouter := Router.Group("userTrxPlaceholders").Use(middleware.OperationRecord())
	userTrxPlaceholdersRouterWithoutRecord := Router.Group("userTrxPlaceholders")
	userTrxPlaceholdersRouterWithoutAuth := PublicRouter.Group("userTrxPlaceholders")
	{
		userTrxPlaceholdersRouter.POST("createUserTrxPlaceholders", userTrxPlaceholdersApi.CreateUserTrxPlaceholders)   // 新建userTrxPlaceholders表
		userTrxPlaceholdersRouter.DELETE("deleteUserTrxPlaceholders", userTrxPlaceholdersApi.DeleteUserTrxPlaceholders) // 删除userTrxPlaceholders表
		userTrxPlaceholdersRouter.DELETE("deleteUserTrxPlaceholdersByIds", userTrxPlaceholdersApi.DeleteUserTrxPlaceholdersByIds) // 批量删除userTrxPlaceholders表
		userTrxPlaceholdersRouter.PUT("updateUserTrxPlaceholders", userTrxPlaceholdersApi.UpdateUserTrxPlaceholders)    // 更新userTrxPlaceholders表
	}
	{
		userTrxPlaceholdersRouterWithoutRecord.GET("findUserTrxPlaceholders", userTrxPlaceholdersApi.FindUserTrxPlaceholders)        // 根据ID获取userTrxPlaceholders表
		userTrxPlaceholdersRouterWithoutRecord.GET("getUserTrxPlaceholdersList", userTrxPlaceholdersApi.GetUserTrxPlaceholdersList)  // 获取userTrxPlaceholders表列表
	}
	{
	    userTrxPlaceholdersRouterWithoutAuth.GET("getUserTrxPlaceholdersPublic", userTrxPlaceholdersApi.GetUserTrxPlaceholdersPublic)  // userTrxPlaceholders表开放接口
	}
}
