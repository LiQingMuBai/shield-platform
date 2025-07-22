package ushield

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/ushield"
    ushieldReq "github.com/flipped-aurora/gin-vue-admin/server/model/ushield/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TgUsersApi struct {}



// CreateTgUsers 创建tgUsers表
// @Tags TgUsers
// @Summary 创建tgUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.TgUsers true "创建tgUsers表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /tgUsers/createTgUsers [post]
func (tgUsersApi *TgUsersApi) CreateTgUsers(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var tgUsers ushield.TgUsers
	err := c.ShouldBindJSON(&tgUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = tgUsersService.CreateTgUsers(ctx,&tgUsers)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteTgUsers 删除tgUsers表
// @Tags TgUsers
// @Summary 删除tgUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.TgUsers true "删除tgUsers表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /tgUsers/deleteTgUsers [delete]
func (tgUsersApi *TgUsersApi) DeleteTgUsers(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := tgUsersService.DeleteTgUsers(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTgUsersByIds 批量删除tgUsers表
// @Tags TgUsers
// @Summary 批量删除tgUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /tgUsers/deleteTgUsersByIds [delete]
func (tgUsersApi *TgUsersApi) DeleteTgUsersByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := tgUsersService.DeleteTgUsersByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTgUsers 更新tgUsers表
// @Tags TgUsers
// @Summary 更新tgUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.TgUsers true "更新tgUsers表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /tgUsers/updateTgUsers [put]
func (tgUsersApi *TgUsersApi) UpdateTgUsers(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var tgUsers ushield.TgUsers
	err := c.ShouldBindJSON(&tgUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = tgUsersService.UpdateTgUsers(ctx,tgUsers)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTgUsers 用id查询tgUsers表
// @Tags TgUsers
// @Summary 用id查询tgUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询tgUsers表"
// @Success 200 {object} response.Response{data=ushield.TgUsers,msg=string} "查询成功"
// @Router /tgUsers/findTgUsers [get]
func (tgUsersApi *TgUsersApi) FindTgUsers(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	retgUsers, err := tgUsersService.GetTgUsers(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(retgUsers, c)
}
// GetTgUsersList 分页获取tgUsers表列表
// @Tags TgUsers
// @Summary 分页获取tgUsers表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.TgUsersSearch true "分页获取tgUsers表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tgUsers/getTgUsersList [get]
func (tgUsersApi *TgUsersApi) GetTgUsersList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo ushieldReq.TgUsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := tgUsersService.GetTgUsersInfoList(ctx,pageInfo)
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

// GetTgUsersPublic 不需要鉴权的tgUsers表接口
// @Tags TgUsers
// @Summary 不需要鉴权的tgUsers表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tgUsers/getTgUsersPublic [get]
func (tgUsersApi *TgUsersApi) GetTgUsersPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    tgUsersService.GetTgUsersPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的tgUsers表接口信息",
    }, "获取成功", c)
}
