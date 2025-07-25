package ushield

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/ushield"
    ushieldReq "github.com/flipped-aurora/gin-vue-admin/server/model/ushield/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UserPackageSubscriptionsApi struct {}



// CreateUserPackageSubscriptions 创建userPackageSubscriptions表
// @Tags UserPackageSubscriptions
// @Summary 创建userPackageSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserPackageSubscriptions true "创建userPackageSubscriptions表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userPackageSubscriptions/createUserPackageSubscriptions [post]
func (userPackageSubscriptionsApi *UserPackageSubscriptionsApi) CreateUserPackageSubscriptions(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var userPackageSubscriptions ushield.UserPackageSubscriptions
	err := c.ShouldBindJSON(&userPackageSubscriptions)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userPackageSubscriptionsService.CreateUserPackageSubscriptions(ctx,&userPackageSubscriptions)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteUserPackageSubscriptions 删除userPackageSubscriptions表
// @Tags UserPackageSubscriptions
// @Summary 删除userPackageSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserPackageSubscriptions true "删除userPackageSubscriptions表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userPackageSubscriptions/deleteUserPackageSubscriptions [delete]
func (userPackageSubscriptionsApi *UserPackageSubscriptionsApi) DeleteUserPackageSubscriptions(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := userPackageSubscriptionsService.DeleteUserPackageSubscriptions(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserPackageSubscriptionsByIds 批量删除userPackageSubscriptions表
// @Tags UserPackageSubscriptions
// @Summary 批量删除userPackageSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userPackageSubscriptions/deleteUserPackageSubscriptionsByIds [delete]
func (userPackageSubscriptionsApi *UserPackageSubscriptionsApi) DeleteUserPackageSubscriptionsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userPackageSubscriptionsService.DeleteUserPackageSubscriptionsByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserPackageSubscriptions 更新userPackageSubscriptions表
// @Tags UserPackageSubscriptions
// @Summary 更新userPackageSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserPackageSubscriptions true "更新userPackageSubscriptions表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userPackageSubscriptions/updateUserPackageSubscriptions [put]
func (userPackageSubscriptionsApi *UserPackageSubscriptionsApi) UpdateUserPackageSubscriptions(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var userPackageSubscriptions ushield.UserPackageSubscriptions
	err := c.ShouldBindJSON(&userPackageSubscriptions)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userPackageSubscriptionsService.UpdateUserPackageSubscriptions(ctx,userPackageSubscriptions)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserPackageSubscriptions 用id查询userPackageSubscriptions表
// @Tags UserPackageSubscriptions
// @Summary 用id查询userPackageSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询userPackageSubscriptions表"
// @Success 200 {object} response.Response{data=ushield.UserPackageSubscriptions,msg=string} "查询成功"
// @Router /userPackageSubscriptions/findUserPackageSubscriptions [get]
func (userPackageSubscriptionsApi *UserPackageSubscriptionsApi) FindUserPackageSubscriptions(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reuserPackageSubscriptions, err := userPackageSubscriptionsService.GetUserPackageSubscriptions(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reuserPackageSubscriptions, c)
}
// GetUserPackageSubscriptionsList 分页获取userPackageSubscriptions表列表
// @Tags UserPackageSubscriptions
// @Summary 分页获取userPackageSubscriptions表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserPackageSubscriptionsSearch true "分页获取userPackageSubscriptions表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userPackageSubscriptions/getUserPackageSubscriptionsList [get]
func (userPackageSubscriptionsApi *UserPackageSubscriptionsApi) GetUserPackageSubscriptionsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo ushieldReq.UserPackageSubscriptionsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userPackageSubscriptionsService.GetUserPackageSubscriptionsInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetUserPackageSubscriptionsPublic 不需要鉴权的userPackageSubscriptions表接口
// @Tags UserPackageSubscriptions
// @Summary 不需要鉴权的userPackageSubscriptions表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userPackageSubscriptions/getUserPackageSubscriptionsPublic [get]
func (userPackageSubscriptionsApi *UserPackageSubscriptionsApi) GetUserPackageSubscriptionsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    userPackageSubscriptionsService.GetUserPackageSubscriptionsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的userPackageSubscriptions表接口信息",
    }, "获取成功", c)
}
