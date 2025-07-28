package ushield

import (
	"github.com/ushield/aurora-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysAnnouncementsInfoRouter struct {}

// InitSysAnnouncementsInfoRouter 初始化 sysAnnouncementsInfo表 路由信息
func (s *SysAnnouncementsInfoRouter) InitSysAnnouncementsInfoRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	sysAnnouncementsInfoRouter := Router.Group("sysAnnouncementsInfo").Use(middleware.OperationRecord())
	sysAnnouncementsInfoRouterWithoutRecord := Router.Group("sysAnnouncementsInfo")
	sysAnnouncementsInfoRouterWithoutAuth := PublicRouter.Group("sysAnnouncementsInfo")
	{
		sysAnnouncementsInfoRouter.POST("createSysAnnouncementsInfo", sysAnnouncementsInfoApi.CreateSysAnnouncementsInfo)   // 新建sysAnnouncementsInfo表
		sysAnnouncementsInfoRouter.DELETE("deleteSysAnnouncementsInfo", sysAnnouncementsInfoApi.DeleteSysAnnouncementsInfo) // 删除sysAnnouncementsInfo表
		sysAnnouncementsInfoRouter.DELETE("deleteSysAnnouncementsInfoByIds", sysAnnouncementsInfoApi.DeleteSysAnnouncementsInfoByIds) // 批量删除sysAnnouncementsInfo表
		sysAnnouncementsInfoRouter.PUT("updateSysAnnouncementsInfo", sysAnnouncementsInfoApi.UpdateSysAnnouncementsInfo)    // 更新sysAnnouncementsInfo表
	}
	{
		sysAnnouncementsInfoRouterWithoutRecord.GET("findSysAnnouncementsInfo", sysAnnouncementsInfoApi.FindSysAnnouncementsInfo)        // 根据ID获取sysAnnouncementsInfo表
		sysAnnouncementsInfoRouterWithoutRecord.GET("getSysAnnouncementsInfoList", sysAnnouncementsInfoApi.GetSysAnnouncementsInfoList)  // 获取sysAnnouncementsInfo表列表
	}
	{
	    sysAnnouncementsInfoRouterWithoutAuth.GET("getSysAnnouncementsInfoPublic", sysAnnouncementsInfoApi.GetSysAnnouncementsInfoPublic)  // sysAnnouncementsInfo表开放接口
	}
}
