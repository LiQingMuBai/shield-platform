package ushield

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TgUsersRouter struct {}

// InitTgUsersRouter 初始化 tgUsers表 路由信息
func (s *TgUsersRouter) InitTgUsersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	tgUsersRouter := Router.Group("tgUsers").Use(middleware.OperationRecord())
	tgUsersRouterWithoutRecord := Router.Group("tgUsers")
	tgUsersRouterWithoutAuth := PublicRouter.Group("tgUsers")
	{
		tgUsersRouter.POST("createTgUsers", tgUsersApi.CreateTgUsers)   // 新建tgUsers表
		tgUsersRouter.DELETE("deleteTgUsers", tgUsersApi.DeleteTgUsers) // 删除tgUsers表
		tgUsersRouter.DELETE("deleteTgUsersByIds", tgUsersApi.DeleteTgUsersByIds) // 批量删除tgUsers表
		tgUsersRouter.PUT("updateTgUsers", tgUsersApi.UpdateTgUsers)    // 更新tgUsers表
	}
	{
		tgUsersRouterWithoutRecord.GET("findTgUsers", tgUsersApi.FindTgUsers)        // 根据ID获取tgUsers表
		tgUsersRouterWithoutRecord.GET("getTgUsersList", tgUsersApi.GetTgUsersList)  // 获取tgUsers表列表
	}
	{
	    tgUsersRouterWithoutAuth.GET("getTgUsersPublic", tgUsersApi.GetTgUsersPublic)  // tgUsers表开放接口
	}
}
