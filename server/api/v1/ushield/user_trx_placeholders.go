package ushield

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/ushield"
    ushieldReq "github.com/flipped-aurora/gin-vue-admin/server/model/ushield/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UserTrxPlaceholdersApi struct {}



// CreateUserTrxPlaceholders 创建userTrxPlaceholders表
// @Tags UserTrxPlaceholders
// @Summary 创建userTrxPlaceholders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserTrxPlaceholders true "创建userTrxPlaceholders表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userTrxPlaceholders/createUserTrxPlaceholders [post]
func (userTrxPlaceholdersApi *UserTrxPlaceholdersApi) CreateUserTrxPlaceholders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var userTrxPlaceholders ushield.UserTrxPlaceholders
	err := c.ShouldBindJSON(&userTrxPlaceholders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userTrxPlaceholdersService.CreateUserTrxPlaceholders(ctx,&userTrxPlaceholders)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteUserTrxPlaceholders 删除userTrxPlaceholders表
// @Tags UserTrxPlaceholders
// @Summary 删除userTrxPlaceholders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserTrxPlaceholders true "删除userTrxPlaceholders表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userTrxPlaceholders/deleteUserTrxPlaceholders [delete]
func (userTrxPlaceholdersApi *UserTrxPlaceholdersApi) DeleteUserTrxPlaceholders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := userTrxPlaceholdersService.DeleteUserTrxPlaceholders(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserTrxPlaceholdersByIds 批量删除userTrxPlaceholders表
// @Tags UserTrxPlaceholders
// @Summary 批量删除userTrxPlaceholders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userTrxPlaceholders/deleteUserTrxPlaceholdersByIds [delete]
func (userTrxPlaceholdersApi *UserTrxPlaceholdersApi) DeleteUserTrxPlaceholdersByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userTrxPlaceholdersService.DeleteUserTrxPlaceholdersByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserTrxPlaceholders 更新userTrxPlaceholders表
// @Tags UserTrxPlaceholders
// @Summary 更新userTrxPlaceholders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.UserTrxPlaceholders true "更新userTrxPlaceholders表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userTrxPlaceholders/updateUserTrxPlaceholders [put]
func (userTrxPlaceholdersApi *UserTrxPlaceholdersApi) UpdateUserTrxPlaceholders(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var userTrxPlaceholders ushield.UserTrxPlaceholders
	err := c.ShouldBindJSON(&userTrxPlaceholders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userTrxPlaceholdersService.UpdateUserTrxPlaceholders(ctx,userTrxPlaceholders)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserTrxPlaceholders 用id查询userTrxPlaceholders表
// @Tags UserTrxPlaceholders
// @Summary 用id查询userTrxPlaceholders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询userTrxPlaceholders表"
// @Success 200 {object} response.Response{data=ushield.UserTrxPlaceholders,msg=string} "查询成功"
// @Router /userTrxPlaceholders/findUserTrxPlaceholders [get]
func (userTrxPlaceholdersApi *UserTrxPlaceholdersApi) FindUserTrxPlaceholders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reuserTrxPlaceholders, err := userTrxPlaceholdersService.GetUserTrxPlaceholders(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reuserTrxPlaceholders, c)
}
// GetUserTrxPlaceholdersList 分页获取userTrxPlaceholders表列表
// @Tags UserTrxPlaceholders
// @Summary 分页获取userTrxPlaceholders表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserTrxPlaceholdersSearch true "分页获取userTrxPlaceholders表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userTrxPlaceholders/getUserTrxPlaceholdersList [get]
func (userTrxPlaceholdersApi *UserTrxPlaceholdersApi) GetUserTrxPlaceholdersList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo ushieldReq.UserTrxPlaceholdersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userTrxPlaceholdersService.GetUserTrxPlaceholdersInfoList(ctx,pageInfo)
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

// GetUserTrxPlaceholdersPublic 不需要鉴权的userTrxPlaceholders表接口
// @Tags UserTrxPlaceholders
// @Summary 不需要鉴权的userTrxPlaceholders表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userTrxPlaceholders/getUserTrxPlaceholdersPublic [get]
func (userTrxPlaceholdersApi *UserTrxPlaceholdersApi) GetUserTrxPlaceholdersPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    userTrxPlaceholdersService.GetUserTrxPlaceholdersPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的userTrxPlaceholders表接口信息",
    }, "获取成功", c)
}
