package system

import (
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/common/response"
	"github.com/ushield/aurora-admin/server/model/system"
	systemReq "github.com/ushield/aurora-admin/server/model/system/request"
	"go.uber.org/zap"
)

type UserUsdtDepositsApi struct{}

// CreateUserUsdtDeposits 创建userUsdtDeposits表
// @Tags UserUsdtDeposits
// @Summary 创建userUsdtDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.UserUsdtDeposits true "创建userUsdtDeposits表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userUsdtDeposits/createUserUsdtDeposits [post]
func (userUsdtDepositsApi *UserUsdtDepositsApi) CreateUserUsdtDeposits(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var userUsdtDeposits system.UserUsdtDeposits
	err := c.ShouldBindJSON(&userUsdtDeposits)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userUsdtDepositsService.CreateUserUsdtDeposits(ctx, &userUsdtDeposits)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteUserUsdtDeposits 删除userUsdtDeposits表
// @Tags UserUsdtDeposits
// @Summary 删除userUsdtDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.UserUsdtDeposits true "删除userUsdtDeposits表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userUsdtDeposits/deleteUserUsdtDeposits [delete]
func (userUsdtDepositsApi *UserUsdtDepositsApi) DeleteUserUsdtDeposits(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := userUsdtDepositsService.DeleteUserUsdtDeposits(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserUsdtDepositsByIds 批量删除userUsdtDeposits表
// @Tags UserUsdtDeposits
// @Summary 批量删除userUsdtDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userUsdtDeposits/deleteUserUsdtDepositsByIds [delete]
func (userUsdtDepositsApi *UserUsdtDepositsApi) DeleteUserUsdtDepositsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userUsdtDepositsService.DeleteUserUsdtDepositsByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserUsdtDeposits 更新userUsdtDeposits表
// @Tags UserUsdtDeposits
// @Summary 更新userUsdtDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.UserUsdtDeposits true "更新userUsdtDeposits表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userUsdtDeposits/updateUserUsdtDeposits [put]
func (userUsdtDepositsApi *UserUsdtDepositsApi) UpdateUserUsdtDeposits(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var userUsdtDeposits system.UserUsdtDeposits
	err := c.ShouldBindJSON(&userUsdtDeposits)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userUsdtDepositsService.UpdateUserUsdtDeposits(ctx, userUsdtDeposits)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserUsdtDeposits 用id查询userUsdtDeposits表
// @Tags UserUsdtDeposits
// @Summary 用id查询userUsdtDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询userUsdtDeposits表"
// @Success 200 {object} response.Response{data=system.UserUsdtDeposits,msg=string} "查询成功"
// @Router /userUsdtDeposits/findUserUsdtDeposits [get]
func (userUsdtDepositsApi *UserUsdtDepositsApi) FindUserUsdtDeposits(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	reuserUsdtDeposits, err := userUsdtDepositsService.GetUserUsdtDeposits(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reuserUsdtDeposits, c)
}

// GetUserUsdtDepositsList 分页获取userUsdtDeposits表列表
// @Tags UserUsdtDeposits
// @Summary 分页获取userUsdtDeposits表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.UserUsdtDepositsSearch true "分页获取userUsdtDeposits表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userUsdtDeposits/getUserUsdtDepositsList [get]
func (userUsdtDepositsApi *UserUsdtDepositsApi) GetUserUsdtDepositsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo systemReq.UserUsdtDepositsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userUsdtDepositsService.GetUserUsdtDepositsInfoList(ctx, pageInfo)
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

// GetUserUsdtDepositsPublic 不需要鉴权的userUsdtDeposits表接口
// @Tags UserUsdtDeposits
// @Summary 不需要鉴权的userUsdtDeposits表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userUsdtDeposits/getUserUsdtDepositsPublic [get]
func (userUsdtDepositsApi *UserUsdtDepositsApi) GetUserUsdtDepositsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	userUsdtDepositsService.GetUserUsdtDepositsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的userUsdtDeposits表接口信息",
	}, "获取成功", c)
}
