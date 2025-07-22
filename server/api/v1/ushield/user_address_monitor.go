package ushield

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/ushield"
    ushieldReq "github.com/flipped-aurora/gin-vue-admin/server/model/ushield/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UserAddressMonitorApi struct {}



// CreateUserAddressMonitor 创建userAddressMonitor表
// @Tags UserAddressMonitor
// @Summary 创建userAddressMonitor表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserAddressMonitor true "创建userAddressMonitor表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userAddressMonitor/createUserAddressMonitor [post]
func (userAddressMonitorApi *UserAddressMonitorApi) CreateUserAddressMonitor(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var userAddressMonitor ushield.UserAddressMonitor
	err := c.ShouldBindJSON(&userAddressMonitor)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userAddressMonitorService.CreateUserAddressMonitor(ctx,&userAddressMonitor)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteUserAddressMonitor 删除userAddressMonitor表
// @Tags UserAddressMonitor
// @Summary 删除userAddressMonitor表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserAddressMonitor true "删除userAddressMonitor表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userAddressMonitor/deleteUserAddressMonitor [delete]
func (userAddressMonitorApi *UserAddressMonitorApi) DeleteUserAddressMonitor(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := userAddressMonitorService.DeleteUserAddressMonitor(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserAddressMonitorByIds 批量删除userAddressMonitor表
// @Tags UserAddressMonitor
// @Summary 批量删除userAddressMonitor表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userAddressMonitor/deleteUserAddressMonitorByIds [delete]
func (userAddressMonitorApi *UserAddressMonitorApi) DeleteUserAddressMonitorByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userAddressMonitorService.DeleteUserAddressMonitorByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserAddressMonitor 更新userAddressMonitor表
// @Tags UserAddressMonitor
// @Summary 更新userAddressMonitor表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserAddressMonitor true "更新userAddressMonitor表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userAddressMonitor/updateUserAddressMonitor [put]
func (userAddressMonitorApi *UserAddressMonitorApi) UpdateUserAddressMonitor(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var userAddressMonitor ushield.UserAddressMonitor
	err := c.ShouldBindJSON(&userAddressMonitor)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userAddressMonitorService.UpdateUserAddressMonitor(ctx,userAddressMonitor)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserAddressMonitor 用id查询userAddressMonitor表
// @Tags UserAddressMonitor
// @Summary 用id查询userAddressMonitor表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询userAddressMonitor表"
// @Success 200 {object} response.Response{data=ushield.UserAddressMonitor,msg=string} "查询成功"
// @Router /userAddressMonitor/findUserAddressMonitor [get]
func (userAddressMonitorApi *UserAddressMonitorApi) FindUserAddressMonitor(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reuserAddressMonitor, err := userAddressMonitorService.GetUserAddressMonitor(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reuserAddressMonitor, c)
}
// GetUserAddressMonitorList 分页获取userAddressMonitor表列表
// @Tags UserAddressMonitor
// @Summary 分页获取userAddressMonitor表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserAddressMonitorSearch true "分页获取userAddressMonitor表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userAddressMonitor/getUserAddressMonitorList [get]
func (userAddressMonitorApi *UserAddressMonitorApi) GetUserAddressMonitorList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo ushieldReq.UserAddressMonitorSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userAddressMonitorService.GetUserAddressMonitorInfoList(ctx,pageInfo)
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

// GetUserAddressMonitorPublic 不需要鉴权的userAddressMonitor表接口
// @Tags UserAddressMonitor
// @Summary 不需要鉴权的userAddressMonitor表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userAddressMonitor/getUserAddressMonitorPublic [get]
func (userAddressMonitorApi *UserAddressMonitorApi) GetUserAddressMonitorPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    userAddressMonitorService.GetUserAddressMonitorPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的userAddressMonitor表接口信息",
    }, "获取成功", c)
}
