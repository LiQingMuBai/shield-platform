package ushield

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/common/response"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
	"go.uber.org/zap"
)

type UserUsdtPlaceholdersApi struct{}

// CreateUserUsdtPlaceholders 创建userUsdtPlaceholders表
// @Tags UserUsdtPlaceholders
// @Summary 创建userUsdtPlaceholders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserUsdtPlaceholders true "创建userUsdtPlaceholders表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userUsdtPlaceholders/createUserUsdtPlaceholders [post]
func (userUsdtPlaceholdersApi *UserUsdtPlaceholdersApi) CreateUserUsdtPlaceholders(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var userUsdtPlaceholders ushield.UserUsdtPlaceholders
	err := c.ShouldBindJSON(&userUsdtPlaceholders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userUsdtPlaceholdersService.CreateUserUsdtPlaceholders(ctx, &userUsdtPlaceholders)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteUserUsdtPlaceholders 删除userUsdtPlaceholders表
// @Tags UserUsdtPlaceholders
// @Summary 删除userUsdtPlaceholders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserUsdtPlaceholders true "删除userUsdtPlaceholders表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userUsdtPlaceholders/deleteUserUsdtPlaceholders [delete]
func (userUsdtPlaceholdersApi *UserUsdtPlaceholdersApi) DeleteUserUsdtPlaceholders(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := userUsdtPlaceholdersService.DeleteUserUsdtPlaceholders(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserUsdtPlaceholdersByIds 批量删除userUsdtPlaceholders表
// @Tags UserUsdtPlaceholders
// @Summary 批量删除userUsdtPlaceholders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userUsdtPlaceholders/deleteUserUsdtPlaceholdersByIds [delete]
func (userUsdtPlaceholdersApi *UserUsdtPlaceholdersApi) DeleteUserUsdtPlaceholdersByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userUsdtPlaceholdersService.DeleteUserUsdtPlaceholdersByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserUsdtPlaceholders 更新userUsdtPlaceholders表
// @Tags UserUsdtPlaceholders
// @Summary 更新userUsdtPlaceholders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserUsdtPlaceholders true "更新userUsdtPlaceholders表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userUsdtPlaceholders/updateUserUsdtPlaceholders [put]
func (userUsdtPlaceholdersApi *UserUsdtPlaceholdersApi) UpdateUserUsdtPlaceholders(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var userUsdtPlaceholders ushield.UserUsdtPlaceholders
	err := c.ShouldBindJSON(&userUsdtPlaceholders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userUsdtPlaceholdersService.UpdateUserUsdtPlaceholders(ctx, userUsdtPlaceholders)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserUsdtPlaceholders 用id查询userUsdtPlaceholders表
// @Tags UserUsdtPlaceholders
// @Summary 用id查询userUsdtPlaceholders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询userUsdtPlaceholders表"
// @Success 200 {object} response.Response{data=ushield.UserUsdtPlaceholders,msg=string} "查询成功"
// @Router /userUsdtPlaceholders/findUserUsdtPlaceholders [get]
func (userUsdtPlaceholdersApi *UserUsdtPlaceholdersApi) FindUserUsdtPlaceholders(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	reuserUsdtPlaceholders, err := userUsdtPlaceholdersService.GetUserUsdtPlaceholders(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reuserUsdtPlaceholders, c)
}

// GetUserUsdtPlaceholdersList 分页获取userUsdtPlaceholders表列表
// @Tags UserUsdtPlaceholders
// @Summary 分页获取userUsdtPlaceholders表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserUsdtPlaceholdersSearch true "分页获取userUsdtPlaceholders表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userUsdtPlaceholders/getUserUsdtPlaceholdersList [get]
func (userUsdtPlaceholdersApi *UserUsdtPlaceholdersApi) GetUserUsdtPlaceholdersList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo ushieldReq.UserUsdtPlaceholdersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userUsdtPlaceholdersService.GetUserUsdtPlaceholdersInfoList(ctx, pageInfo)
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

// GetUserUsdtPlaceholdersPublic 不需要鉴权的userUsdtPlaceholders表接口
// @Tags UserUsdtPlaceholders
// @Summary 不需要鉴权的userUsdtPlaceholders表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userUsdtPlaceholders/getUserUsdtPlaceholdersPublic [get]
func (userUsdtPlaceholdersApi *UserUsdtPlaceholdersApi) GetUserUsdtPlaceholdersPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	userUsdtPlaceholdersService.GetUserUsdtPlaceholdersPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的userUsdtPlaceholders表接口信息",
	}, "获取成功", c)
}
