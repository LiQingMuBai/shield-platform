package ushield

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/ushield"
    ushieldReq "github.com/flipped-aurora/gin-vue-admin/server/model/ushield/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UserAddressMonitorEventApi struct {}



// CreateUserAddressMonitorEvent 创建userAddressMonitorEvent表
// @Tags UserAddressMonitorEvent
// @Summary 创建userAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserAddressMonitorEvent true "创建userAddressMonitorEvent表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userAddressMonitorEvent/createUserAddressMonitorEvent [post]
func (userAddressMonitorEventApi *UserAddressMonitorEventApi) CreateUserAddressMonitorEvent(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var userAddressMonitorEvent ushield.UserAddressMonitorEvent
	err := c.ShouldBindJSON(&userAddressMonitorEvent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userAddressMonitorEventService.CreateUserAddressMonitorEvent(ctx,&userAddressMonitorEvent)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteUserAddressMonitorEvent 删除userAddressMonitorEvent表
// @Tags UserAddressMonitorEvent
// @Summary 删除userAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserAddressMonitorEvent true "删除userAddressMonitorEvent表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userAddressMonitorEvent/deleteUserAddressMonitorEvent [delete]
func (userAddressMonitorEventApi *UserAddressMonitorEventApi) DeleteUserAddressMonitorEvent(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := userAddressMonitorEventService.DeleteUserAddressMonitorEvent(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserAddressMonitorEventByIds 批量删除userAddressMonitorEvent表
// @Tags UserAddressMonitorEvent
// @Summary 批量删除userAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userAddressMonitorEvent/deleteUserAddressMonitorEventByIds [delete]
func (userAddressMonitorEventApi *UserAddressMonitorEventApi) DeleteUserAddressMonitorEventByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userAddressMonitorEventService.DeleteUserAddressMonitorEventByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserAddressMonitorEvent 更新userAddressMonitorEvent表
// @Tags UserAddressMonitorEvent
// @Summary 更新userAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserAddressMonitorEvent true "更新userAddressMonitorEvent表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userAddressMonitorEvent/updateUserAddressMonitorEvent [put]
func (userAddressMonitorEventApi *UserAddressMonitorEventApi) UpdateUserAddressMonitorEvent(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var userAddressMonitorEvent ushield.UserAddressMonitorEvent
	err := c.ShouldBindJSON(&userAddressMonitorEvent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userAddressMonitorEventService.UpdateUserAddressMonitorEvent(ctx,userAddressMonitorEvent)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserAddressMonitorEvent 用id查询userAddressMonitorEvent表
// @Tags UserAddressMonitorEvent
// @Summary 用id查询userAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询userAddressMonitorEvent表"
// @Success 200 {object} response.Response{data=ushield.UserAddressMonitorEvent,msg=string} "查询成功"
// @Router /userAddressMonitorEvent/findUserAddressMonitorEvent [get]
func (userAddressMonitorEventApi *UserAddressMonitorEventApi) FindUserAddressMonitorEvent(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reuserAddressMonitorEvent, err := userAddressMonitorEventService.GetUserAddressMonitorEvent(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reuserAddressMonitorEvent, c)
}
// GetUserAddressMonitorEventList 分页获取userAddressMonitorEvent表列表
// @Tags UserAddressMonitorEvent
// @Summary 分页获取userAddressMonitorEvent表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserAddressMonitorEventSearch true "分页获取userAddressMonitorEvent表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userAddressMonitorEvent/getUserAddressMonitorEventList [get]
func (userAddressMonitorEventApi *UserAddressMonitorEventApi) GetUserAddressMonitorEventList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo ushieldReq.UserAddressMonitorEventSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userAddressMonitorEventService.GetUserAddressMonitorEventInfoList(ctx,pageInfo)
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

// GetUserAddressMonitorEventPublic 不需要鉴权的userAddressMonitorEvent表接口
// @Tags UserAddressMonitorEvent
// @Summary 不需要鉴权的userAddressMonitorEvent表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userAddressMonitorEvent/getUserAddressMonitorEventPublic [get]
func (userAddressMonitorEventApi *UserAddressMonitorEventApi) GetUserAddressMonitorEventPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    userAddressMonitorEventService.GetUserAddressMonitorEventPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的userAddressMonitorEvent表接口信息",
    }, "获取成功", c)
}
