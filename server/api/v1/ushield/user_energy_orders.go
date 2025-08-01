package ushield

import (
	
	"github.com/ushield/aurora-admin/server/global"
    "github.com/ushield/aurora-admin/server/model/common/response"
    "github.com/ushield/aurora-admin/server/model/ushield"
    ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UserEnergyOrdersApi struct {}



// CreateUserEnergyOrders 创建userEnergyOrders表
// @Tags UserEnergyOrders
// @Summary 创建userEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserEnergyOrders true "创建userEnergyOrders表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userEnergyOrders/createUserEnergyOrders [post]
func (userEnergyOrdersApi *UserEnergyOrdersApi) CreateUserEnergyOrders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var userEnergyOrders ushield.UserEnergyOrders
	err := c.ShouldBindJSON(&userEnergyOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userEnergyOrdersService.CreateUserEnergyOrders(ctx,&userEnergyOrders)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteUserEnergyOrders 删除userEnergyOrders表
// @Tags UserEnergyOrders
// @Summary 删除userEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserEnergyOrders true "删除userEnergyOrders表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userEnergyOrders/deleteUserEnergyOrders [delete]
func (userEnergyOrdersApi *UserEnergyOrdersApi) DeleteUserEnergyOrders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := userEnergyOrdersService.DeleteUserEnergyOrders(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserEnergyOrdersByIds 批量删除userEnergyOrders表
// @Tags UserEnergyOrders
// @Summary 批量删除userEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userEnergyOrders/deleteUserEnergyOrdersByIds [delete]
func (userEnergyOrdersApi *UserEnergyOrdersApi) DeleteUserEnergyOrdersByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userEnergyOrdersService.DeleteUserEnergyOrdersByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserEnergyOrders 更新userEnergyOrders表
// @Tags UserEnergyOrders
// @Summary 更新userEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserEnergyOrders true "更新userEnergyOrders表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userEnergyOrders/updateUserEnergyOrders [put]
func (userEnergyOrdersApi *UserEnergyOrdersApi) UpdateUserEnergyOrders(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var userEnergyOrders ushield.UserEnergyOrders
	err := c.ShouldBindJSON(&userEnergyOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userEnergyOrdersService.UpdateUserEnergyOrders(ctx,userEnergyOrders)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserEnergyOrders 用id查询userEnergyOrders表
// @Tags UserEnergyOrders
// @Summary 用id查询userEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询userEnergyOrders表"
// @Success 200 {object} response.Response{data=ushield.UserEnergyOrders,msg=string} "查询成功"
// @Router /userEnergyOrders/findUserEnergyOrders [get]
func (userEnergyOrdersApi *UserEnergyOrdersApi) FindUserEnergyOrders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reuserEnergyOrders, err := userEnergyOrdersService.GetUserEnergyOrders(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reuserEnergyOrders, c)
}
// GetUserEnergyOrdersList 分页获取userEnergyOrders表列表
// @Tags UserEnergyOrders
// @Summary 分页获取userEnergyOrders表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserEnergyOrdersSearch true "分页获取userEnergyOrders表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userEnergyOrders/getUserEnergyOrdersList [get]
func (userEnergyOrdersApi *UserEnergyOrdersApi) GetUserEnergyOrdersList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo ushieldReq.UserEnergyOrdersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userEnergyOrdersService.GetUserEnergyOrdersInfoList(ctx,pageInfo)
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

// GetUserEnergyOrdersPublic 不需要鉴权的userEnergyOrders表接口
// @Tags UserEnergyOrders
// @Summary 不需要鉴权的userEnergyOrders表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userEnergyOrders/getUserEnergyOrdersPublic [get]
func (userEnergyOrdersApi *UserEnergyOrdersApi) GetUserEnergyOrdersPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    userEnergyOrdersService.GetUserEnergyOrdersPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的userEnergyOrders表接口信息",
    }, "获取成功", c)
}
