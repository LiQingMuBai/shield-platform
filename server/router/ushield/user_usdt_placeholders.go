package ushield

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/middleware"
)

type UserUsdtPlaceholdersRouter struct{}

// InitUserUsdtPlaceholdersRouter 初始化 userUsdtPlaceholders表 路由信息
func (s *UserUsdtPlaceholdersRouter) InitUserUsdtPlaceholdersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userUsdtPlaceholdersRouter := Router.Group("userUsdtPlaceholders").Use(middleware.OperationRecord())
	userUsdtPlaceholdersRouterWithoutRecord := Router.Group("userUsdtPlaceholders")
	userUsdtPlaceholdersRouterWithoutAuth := PublicRouter.Group("userUsdtPlaceholders")
	{
		userUsdtPlaceholdersRouter.POST("createUserUsdtPlaceholders", userUsdtPlaceholdersApi.CreateUserUsdtPlaceholders)             // 新建userUsdtPlaceholders表
		userUsdtPlaceholdersRouter.DELETE("deleteUserUsdtPlaceholders", userUsdtPlaceholdersApi.DeleteUserUsdtPlaceholders)           // 删除userUsdtPlaceholders表
		userUsdtPlaceholdersRouter.DELETE("deleteUserUsdtPlaceholdersByIds", userUsdtPlaceholdersApi.DeleteUserUsdtPlaceholdersByIds) // 批量删除userUsdtPlaceholders表
		userUsdtPlaceholdersRouter.PUT("updateUserUsdtPlaceholders", userUsdtPlaceholdersApi.UpdateUserUsdtPlaceholders)              // 更新userUsdtPlaceholders表
	}
	{
		userUsdtPlaceholdersRouterWithoutRecord.GET("findUserUsdtPlaceholders", userUsdtPlaceholdersApi.FindUserUsdtPlaceholders)       // 根据ID获取userUsdtPlaceholders表
		userUsdtPlaceholdersRouterWithoutRecord.GET("getUserUsdtPlaceholdersList", userUsdtPlaceholdersApi.GetUserUsdtPlaceholdersList) // 获取userUsdtPlaceholders表列表
	}
	{
		userUsdtPlaceholdersRouterWithoutAuth.GET("getUserUsdtPlaceholdersPublic", userUsdtPlaceholdersApi.GetUserUsdtPlaceholdersPublic) // userUsdtPlaceholders表开放接口
	}
}
