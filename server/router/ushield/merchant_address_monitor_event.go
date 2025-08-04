package ushield

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/middleware"
)

type MerchantAddressMonitorEventRouter struct{}

// InitMerchantAddressMonitorEventRouter 初始化 merchantAddressMonitorEvent表 路由信息
func (s *MerchantAddressMonitorEventRouter) InitMerchantAddressMonitorEventRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	merchantAddressMonitorEventRouter := Router.Group("merchantAddressMonitorEvent").Use(middleware.OperationRecord())
	merchantAddressMonitorEventRouterWithoutRecord := Router.Group("merchantAddressMonitorEvent")
	merchantAddressMonitorEventRouterWithoutAuth := PublicRouter.Group("merchantAddressMonitorEvent")
	{
		merchantAddressMonitorEventRouter.POST("createMerchantAddressMonitorEvent", merchantAddressMonitorEventApi.CreateMerchantAddressMonitorEvent)             // 新建merchantAddressMonitorEvent表
		merchantAddressMonitorEventRouter.POST("order", merchantAddressMonitorEventApi.CreateMerchantAddressMonitorEvent)                                         // 新建merchantAddressMonitorEvent表
		merchantAddressMonitorEventRouter.DELETE("deleteMerchantAddressMonitorEvent", merchantAddressMonitorEventApi.DeleteMerchantAddressMonitorEvent)           // 删除merchantAddressMonitorEvent表
		merchantAddressMonitorEventRouter.DELETE("deleteMerchantAddressMonitorEventByIds", merchantAddressMonitorEventApi.DeleteMerchantAddressMonitorEventByIds) // 批量删除merchantAddressMonitorEvent表
		merchantAddressMonitorEventRouter.PUT("updateMerchantAddressMonitorEvent", merchantAddressMonitorEventApi.UpdateMerchantAddressMonitorEvent)              // 更新merchantAddressMonitorEvent表
	}
	{
		merchantAddressMonitorEventRouterWithoutRecord.GET("findMerchantAddressMonitorEvent", merchantAddressMonitorEventApi.FindMerchantAddressMonitorEvent)       // 根据ID获取merchantAddressMonitorEvent表
		merchantAddressMonitorEventRouterWithoutRecord.GET("getMerchantAddressMonitorEventList", merchantAddressMonitorEventApi.GetMerchantAddressMonitorEventList) // 获取merchantAddressMonitorEvent表列表
	}
	{
		merchantAddressMonitorEventRouterWithoutAuth.GET("getMerchantAddressMonitorEventPublic", merchantAddressMonitorEventApi.GetMerchantAddressMonitorEventPublic) // merchantAddressMonitorEvent表开放接口
	}
}
