package ushield

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/common/response"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
	"go.uber.org/zap"
)

type UserOperationBundlesApi struct{}

// CreateUserOperationBundles 创建userOperationBundles表
// @Tags UserOperationBundles
// @Summary 创建userOperationBundles表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserOperationBundles true "创建userOperationBundles表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userOperationBundles/createUserOperationBundles [post]
func (userOperationBundlesApi *UserOperationBundlesApi) CreateUserOperationBundles(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var userOperationBundles ushield.UserOperationBundles
	err := c.ShouldBindJSON(&userOperationBundles)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userOperationBundlesService.CreateUserOperationBundles(ctx, &userOperationBundles)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteUserOperationBundles 删除userOperationBundles表
// @Tags UserOperationBundles
// @Summary 删除userOperationBundles表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserOperationBundles true "删除userOperationBundles表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userOperationBundles/deleteUserOperationBundles [delete]
func (userOperationBundlesApi *UserOperationBundlesApi) DeleteUserOperationBundles(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := userOperationBundlesService.DeleteUserOperationBundles(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserOperationBundlesByIds 批量删除userOperationBundles表
// @Tags UserOperationBundles
// @Summary 批量删除userOperationBundles表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userOperationBundles/deleteUserOperationBundlesByIds [delete]
func (userOperationBundlesApi *UserOperationBundlesApi) DeleteUserOperationBundlesByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userOperationBundlesService.DeleteUserOperationBundlesByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserOperationBundles 更新userOperationBundles表
// @Tags UserOperationBundles
// @Summary 更新userOperationBundles表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserOperationBundles true "更新userOperationBundles表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userOperationBundles/updateUserOperationBundles [put]
func (userOperationBundlesApi *UserOperationBundlesApi) UpdateUserOperationBundles(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var userOperationBundles ushield.UserOperationBundles
	err := c.ShouldBindJSON(&userOperationBundles)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userOperationBundlesService.UpdateUserOperationBundles(ctx, userOperationBundles)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserOperationBundles 用id查询userOperationBundles表
// @Tags UserOperationBundles
// @Summary 用id查询userOperationBundles表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询userOperationBundles表"
// @Success 200 {object} response.Response{data=ushield.UserOperationBundles,msg=string} "查询成功"
// @Router /userOperationBundles/findUserOperationBundles [get]
func (userOperationBundlesApi *UserOperationBundlesApi) FindUserOperationBundles(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	reuserOperationBundles, err := userOperationBundlesService.GetUserOperationBundles(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reuserOperationBundles, c)
}

// GetUserOperationBundlesList 分页获取userOperationBundles表列表
// @Tags UserOperationBundles
// @Summary 分页获取userOperationBundles表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserOperationBundlesSearch true "分页获取userOperationBundles表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userOperationBundles/getUserOperationBundlesList [get]
func (userOperationBundlesApi *UserOperationBundlesApi) GetUserOperationBundlesList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo ushieldReq.UserOperationBundlesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userOperationBundlesService.GetUserOperationBundlesInfoList(ctx, pageInfo)
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

// GetUserOperationBundlesPublic 不需要鉴权的userOperationBundles表接口
// @Tags UserOperationBundles
// @Summary 不需要鉴权的userOperationBundles表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userOperationBundles/getUserOperationBundlesPublic [get]
func (userOperationBundlesApi *UserOperationBundlesApi) GetUserOperationBundlesPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	userOperationBundlesService.GetUserOperationBundlesPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的userOperationBundles表接口信息",
	}, "获取成功", c)
}
