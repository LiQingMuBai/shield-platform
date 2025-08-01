package ushield

import (
	
	"github.com/ushield/aurora-admin/server/global"
    "github.com/ushield/aurora-admin/server/model/common/response"
    "github.com/ushield/aurora-admin/server/model/ushield"
    ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UserBundleEnergyOrdersApi struct {}



// CreateUserBundleEnergyOrders 创建userBundleEnergyOrders表
// @Tags UserBundleEnergyOrders
// @Summary 创建userBundleEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserBundleEnergyOrders true "创建userBundleEnergyOrders表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userBundleEnergyOrders/createUserBundleEnergyOrders [post]
func (userBundleEnergyOrdersApi *UserBundleEnergyOrdersApi) CreateUserBundleEnergyOrders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var userBundleEnergyOrders ushield.UserBundleEnergyOrders
	err := c.ShouldBindJSON(&userBundleEnergyOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userBundleEnergyOrdersService.CreateUserBundleEnergyOrders(ctx,&userBundleEnergyOrders)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteUserBundleEnergyOrders 删除userBundleEnergyOrders表
// @Tags UserBundleEnergyOrders
// @Summary 删除userBundleEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserBundleEnergyOrders true "删除userBundleEnergyOrders表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userBundleEnergyOrders/deleteUserBundleEnergyOrders [delete]
func (userBundleEnergyOrdersApi *UserBundleEnergyOrdersApi) DeleteUserBundleEnergyOrders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := userBundleEnergyOrdersService.DeleteUserBundleEnergyOrders(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserBundleEnergyOrdersByIds 批量删除userBundleEnergyOrders表
// @Tags UserBundleEnergyOrders
// @Summary 批量删除userBundleEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userBundleEnergyOrders/deleteUserBundleEnergyOrdersByIds [delete]
func (userBundleEnergyOrdersApi *UserBundleEnergyOrdersApi) DeleteUserBundleEnergyOrdersByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userBundleEnergyOrdersService.DeleteUserBundleEnergyOrdersByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserBundleEnergyOrders 更新userBundleEnergyOrders表
// @Tags UserBundleEnergyOrders
// @Summary 更新userBundleEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserBundleEnergyOrders true "更新userBundleEnergyOrders表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userBundleEnergyOrders/updateUserBundleEnergyOrders [put]
func (userBundleEnergyOrdersApi *UserBundleEnergyOrdersApi) UpdateUserBundleEnergyOrders(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var userBundleEnergyOrders ushield.UserBundleEnergyOrders
	err := c.ShouldBindJSON(&userBundleEnergyOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userBundleEnergyOrdersService.UpdateUserBundleEnergyOrders(ctx,userBundleEnergyOrders)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserBundleEnergyOrders 用id查询userBundleEnergyOrders表
// @Tags UserBundleEnergyOrders
// @Summary 用id查询userBundleEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询userBundleEnergyOrders表"
// @Success 200 {object} response.Response{data=ushield.UserBundleEnergyOrders,msg=string} "查询成功"
// @Router /userBundleEnergyOrders/findUserBundleEnergyOrders [get]
func (userBundleEnergyOrdersApi *UserBundleEnergyOrdersApi) FindUserBundleEnergyOrders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reuserBundleEnergyOrders, err := userBundleEnergyOrdersService.GetUserBundleEnergyOrders(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reuserBundleEnergyOrders, c)
}
// GetUserBundleEnergyOrdersList 分页获取userBundleEnergyOrders表列表
// @Tags UserBundleEnergyOrders
// @Summary 分页获取userBundleEnergyOrders表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserBundleEnergyOrdersSearch true "分页获取userBundleEnergyOrders表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userBundleEnergyOrders/getUserBundleEnergyOrdersList [get]
func (userBundleEnergyOrdersApi *UserBundleEnergyOrdersApi) GetUserBundleEnergyOrdersList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo ushieldReq.UserBundleEnergyOrdersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userBundleEnergyOrdersService.GetUserBundleEnergyOrdersInfoList(ctx,pageInfo)
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

// GetUserBundleEnergyOrdersPublic 不需要鉴权的userBundleEnergyOrders表接口
// @Tags UserBundleEnergyOrders
// @Summary 不需要鉴权的userBundleEnergyOrders表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userBundleEnergyOrders/getUserBundleEnergyOrdersPublic [get]
func (userBundleEnergyOrdersApi *UserBundleEnergyOrdersApi) GetUserBundleEnergyOrdersPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    userBundleEnergyOrdersService.GetUserBundleEnergyOrdersPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的userBundleEnergyOrders表接口信息",
    }, "获取成功", c)
}
