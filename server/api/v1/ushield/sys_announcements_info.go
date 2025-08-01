package ushield

import (
	
	"github.com/ushield/aurora-admin/server/global"
    "github.com/ushield/aurora-admin/server/model/common/response"
    "github.com/ushield/aurora-admin/server/model/ushield"
    ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SysAnnouncementsInfoApi struct {}



// CreateSysAnnouncementsInfo 创建sysAnnouncementsInfo表
// @Tags SysAnnouncementsInfo
// @Summary 创建sysAnnouncementsInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.SysAnnouncementsInfo true "创建sysAnnouncementsInfo表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sysAnnouncementsInfo/createSysAnnouncementsInfo [post]
func (sysAnnouncementsInfoApi *SysAnnouncementsInfoApi) CreateSysAnnouncementsInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var sysAnnouncementsInfo ushield.SysAnnouncementsInfo
	err := c.ShouldBindJSON(&sysAnnouncementsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysAnnouncementsInfoService.CreateSysAnnouncementsInfo(ctx,&sysAnnouncementsInfo)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSysAnnouncementsInfo 删除sysAnnouncementsInfo表
// @Tags SysAnnouncementsInfo
// @Summary 删除sysAnnouncementsInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.SysAnnouncementsInfo true "删除sysAnnouncementsInfo表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sysAnnouncementsInfo/deleteSysAnnouncementsInfo [delete]
func (sysAnnouncementsInfoApi *SysAnnouncementsInfoApi) DeleteSysAnnouncementsInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := sysAnnouncementsInfoService.DeleteSysAnnouncementsInfo(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysAnnouncementsInfoByIds 批量删除sysAnnouncementsInfo表
// @Tags SysAnnouncementsInfo
// @Summary 批量删除sysAnnouncementsInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysAnnouncementsInfo/deleteSysAnnouncementsInfoByIds [delete]
func (sysAnnouncementsInfoApi *SysAnnouncementsInfoApi) DeleteSysAnnouncementsInfoByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := sysAnnouncementsInfoService.DeleteSysAnnouncementsInfoByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSysAnnouncementsInfo 更新sysAnnouncementsInfo表
// @Tags SysAnnouncementsInfo
// @Summary 更新sysAnnouncementsInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.SysAnnouncementsInfo true "更新sysAnnouncementsInfo表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sysAnnouncementsInfo/updateSysAnnouncementsInfo [put]
func (sysAnnouncementsInfoApi *SysAnnouncementsInfoApi) UpdateSysAnnouncementsInfo(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var sysAnnouncementsInfo ushield.SysAnnouncementsInfo
	err := c.ShouldBindJSON(&sysAnnouncementsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysAnnouncementsInfoService.UpdateSysAnnouncementsInfo(ctx,sysAnnouncementsInfo)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysAnnouncementsInfo 用id查询sysAnnouncementsInfo表
// @Tags SysAnnouncementsInfo
// @Summary 用id查询sysAnnouncementsInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询sysAnnouncementsInfo表"
// @Success 200 {object} response.Response{data=ushield.SysAnnouncementsInfo,msg=string} "查询成功"
// @Router /sysAnnouncementsInfo/findSysAnnouncementsInfo [get]
func (sysAnnouncementsInfoApi *SysAnnouncementsInfoApi) FindSysAnnouncementsInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	resysAnnouncementsInfo, err := sysAnnouncementsInfoService.GetSysAnnouncementsInfo(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(resysAnnouncementsInfo, c)
}
// GetSysAnnouncementsInfoList 分页获取sysAnnouncementsInfo表列表
// @Tags SysAnnouncementsInfo
// @Summary 分页获取sysAnnouncementsInfo表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.SysAnnouncementsInfoSearch true "分页获取sysAnnouncementsInfo表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysAnnouncementsInfo/getSysAnnouncementsInfoList [get]
func (sysAnnouncementsInfoApi *SysAnnouncementsInfoApi) GetSysAnnouncementsInfoList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo ushieldReq.SysAnnouncementsInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysAnnouncementsInfoService.GetSysAnnouncementsInfoInfoList(ctx,pageInfo)
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

// GetSysAnnouncementsInfoPublic 不需要鉴权的sysAnnouncementsInfo表接口
// @Tags SysAnnouncementsInfo
// @Summary 不需要鉴权的sysAnnouncementsInfo表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysAnnouncementsInfo/getSysAnnouncementsInfoPublic [get]
func (sysAnnouncementsInfoApi *SysAnnouncementsInfoApi) GetSysAnnouncementsInfoPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    sysAnnouncementsInfoService.GetSysAnnouncementsInfoPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的sysAnnouncementsInfo表接口信息",
    }, "获取成功", c)
}
