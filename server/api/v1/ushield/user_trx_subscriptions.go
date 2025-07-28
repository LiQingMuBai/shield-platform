package ushield

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/common/response"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
	"go.uber.org/zap"
)

type UserTrxSubscriptionsApi struct{}

// CreateUserTrxSubscriptions 创建userTrxSubscriptions表
// @Tags UserTrxSubscriptions
// @Summary 创建userTrxSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserTrxSubscriptions true "创建userTrxSubscriptions表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userTrxSubscriptions/createUserTrxSubscriptions [post]
func (userTrxSubscriptionsApi *UserTrxSubscriptionsApi) CreateUserTrxSubscriptions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var userTrxSubscriptions ushield.UserTrxSubscriptions
	err := c.ShouldBindJSON(&userTrxSubscriptions)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userTrxSubscriptionsService.CreateUserTrxSubscriptions(ctx, &userTrxSubscriptions)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteUserTrxSubscriptions 删除userTrxSubscriptions表
// @Tags UserTrxSubscriptions
// @Summary 删除userTrxSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserTrxSubscriptions true "删除userTrxSubscriptions表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userTrxSubscriptions/deleteUserTrxSubscriptions [delete]
func (userTrxSubscriptionsApi *UserTrxSubscriptionsApi) DeleteUserTrxSubscriptions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := userTrxSubscriptionsService.DeleteUserTrxSubscriptions(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserTrxSubscriptionsByIds 批量删除userTrxSubscriptions表
// @Tags UserTrxSubscriptions
// @Summary 批量删除userTrxSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userTrxSubscriptions/deleteUserTrxSubscriptionsByIds [delete]
func (userTrxSubscriptionsApi *UserTrxSubscriptionsApi) DeleteUserTrxSubscriptionsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userTrxSubscriptionsService.DeleteUserTrxSubscriptionsByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserTrxSubscriptions 更新userTrxSubscriptions表
// @Tags UserTrxSubscriptions
// @Summary 更新userTrxSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserTrxSubscriptions true "更新userTrxSubscriptions表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userTrxSubscriptions/updateUserTrxSubscriptions [put]
func (userTrxSubscriptionsApi *UserTrxSubscriptionsApi) UpdateUserTrxSubscriptions(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var userTrxSubscriptions ushield.UserTrxSubscriptions
	err := c.ShouldBindJSON(&userTrxSubscriptions)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userTrxSubscriptionsService.UpdateUserTrxSubscriptions(ctx, userTrxSubscriptions)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserTrxSubscriptions 用id查询userTrxSubscriptions表
// @Tags UserTrxSubscriptions
// @Summary 用id查询userTrxSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询userTrxSubscriptions表"
// @Success 200 {object} response.Response{data=ushield.UserTrxSubscriptions,msg=string} "查询成功"
// @Router /userTrxSubscriptions/findUserTrxSubscriptions [get]
func (userTrxSubscriptionsApi *UserTrxSubscriptionsApi) FindUserTrxSubscriptions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	reuserTrxSubscriptions, err := userTrxSubscriptionsService.GetUserTrxSubscriptions(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reuserTrxSubscriptions, c)
}

// GetUserTrxSubscriptionsList 分页获取userTrxSubscriptions表列表
// @Tags UserTrxSubscriptions
// @Summary 分页获取userTrxSubscriptions表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserTrxSubscriptionsSearch true "分页获取userTrxSubscriptions表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userTrxSubscriptions/getUserTrxSubscriptionsList [get]
func (userTrxSubscriptionsApi *UserTrxSubscriptionsApi) GetUserTrxSubscriptionsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo ushieldReq.UserTrxSubscriptionsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userTrxSubscriptionsService.GetUserTrxSubscriptionsInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetUserTrxSubscriptionsPublic 不需要鉴权的userTrxSubscriptions表接口
// @Tags UserTrxSubscriptions
// @Summary 不需要鉴权的userTrxSubscriptions表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userTrxSubscriptions/getUserTrxSubscriptionsPublic [get]
func (userTrxSubscriptionsApi *UserTrxSubscriptionsApi) GetUserTrxSubscriptionsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	userTrxSubscriptionsService.GetUserTrxSubscriptionsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的userTrxSubscriptions表接口信息",
	}, "获取成功", c)
}
