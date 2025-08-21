package ushield

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/common/response"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
	"github.com/ushield/aurora-admin/server/utils"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

type MerchantAddressMonitorEventApi struct{}

// FindMerchantAddressMonitorEvent 用id查询merchantAddressMonitorEvent表
// @Tags MerchantAddressMonitorEvent
// @Summary 用id查询merchantAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询merchantAddressMonitorEvent表"
// @Success 200 {object} response.Response{data=ushield.MerchantAddressMonitorEvent,msg=string} "查询成功"
// @Router /merchantAddressMonitorEvent/findMerchantAddressMonitorEvent [get]
func (merchantAddressMonitorEventApi *MerchantAddressMonitorEventApi) FindMerchantAddressMonitorEventByAddressAndUser(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var l ushieldReq.MerchantAddressMonitorEventReq
	err := c.ShouldBindJSON(&l)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return

	}
	log.Printf("Address : %v", l.Address)
	userID := utils.GetUserID(c)
	remerchantAddressMonitorEvent, err := merchantAddressMonitorEventService.GetMerchantAddressMonitorEventByAddressAndUser(ctx, l.Address, userID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	status := remerchantAddressMonitorEvent.Status
	code := 20000
	message := "20000"

	if status == 1 {
		//正常
		code = 20000
		message = "正常"
	} else if status == 0 {
		//暂停服务
		code = 20001
		message = "服务已经暂停该预警功能"
	} else if status == 2 {
		//将被拉入黑名单
		code = 20002
		message = "即将被拉入黑名单"
	} else if status == 4 {
		//已经是黑名单
		code = 20004
		message = "已经被拉入黑名单"
	} else {
		code = 20005
		message = "无反馈"
	}

	response.Result(code, nil, message, c)
}

// CreateMerchantAddressMonitorEvent 创建merchantAddressMonitorEvent表
// @Tags MerchantAddressMonitorEvent
// @Summary 创建merchantAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.MerchantAddressMonitorEvent true "创建merchantAddressMonitorEvent表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /merchantAddressMonitorEvent/createMerchantAddressMonitorEvent [post]
func (merchantAddressMonitorEventApi *MerchantAddressMonitorEventApi) CreateMerchantAddressMonitorEvent(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	fmt.Println("进入厂商模式")

	var merchantAddressMonitorEvent ushield.MerchantAddressMonitorEvent
	merchantAddressMonitorEvent.UserId = utils.GetUserID(c)

	log.Printf("userID : %v", merchantAddressMonitorEvent.UserId)
	var l ushieldReq.MerchantAddressMonitorEventReq
	err := c.ShouldBindJSON(&l)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return

	}
	log.Printf("Address : %v", l.Address)
	log.Printf("Callback : %v", l.Callback)

	//valid, message := utils.IsValidCryptoAddress(merchantAddressMonitorEvent.Address)

	//if !valid {
	//	//global.GVA_LOG.Error("非法地址!", zap.Error(message))
	//	response.FailWithMessage("非法地址:"+message, c)
	//	return
	//}

	//isUrl := utils.IsValidURL(merchantAddressMonitorEvent.Callback)
	//if !isUrl {
	//	//global.GVA_LOG.Error("非法地址!", zap.Error(message))
	//	response.FailWithMessage("非法回调地址:"+merchantAddressMonitorEvent.Callback, c)
	//	return
	//}

	//if merchantAddressMonitorEvent.Address

	merchantAddressMonitorEvent.Address = l.Address
	merchantAddressMonitorEvent.Callback = l.Callback
	merchantAddressMonitorEvent.CreatedAt = time.Now()
	merchantAddressMonitorEvent.UpdatedAt = time.Now()
	merchantAddressMonitorEvent.Days = 0
	merchantAddressMonitorEvent.Status = 1
	//merchantAddressMonitorEvent.Network = message

	err = merchantAddressMonitorEventService.CreateMerchantAddressMonitorEvent(ctx, &merchantAddressMonitorEvent)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMerchantAddressMonitorEvent 删除merchantAddressMonitorEvent表
// @Tags MerchantAddressMonitorEvent
// @Summary 删除merchantAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.MerchantAddressMonitorEvent true "删除merchantAddressMonitorEvent表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /merchantAddressMonitorEvent/deleteMerchantAddressMonitorEvent [delete]
func (merchantAddressMonitorEventApi *MerchantAddressMonitorEventApi) InvokeMerchantAddressMonitorEvent(c *gin.Context) {
	// 创建业务用Context
	//ctx := c.Request.Context()
	//
	id := c.Query("id")

	event, _ := merchantAddressMonitorEventService.GetMerchantAddressMonitorEvent(context.Background(), id)

	callback(event.Callback, event.Address, event.Id)
	fmt.Printf(event.Callback)

	////err := merchantAddressMonitorEventService.DeleteMerchantAddressMonitorEvent(ctx, id)
	//if err != nil {
	//	global.GVA_LOG.Error("删除失败!", zap.Error(err))
	//	response.FailWithMessage("删除失败:"+err.Error(), c)
	//	return
	//}

	response.OkWithMessage("回调成功", c)
}

func callback(url string, _address string, status int64) {

	code := 20000
	message := "20000"

	if status == 1 {
		//正常
		code = 20000
		message = "正常"
	} else if status == 0 {
		//暂停服务
		code = 20001
		message = "服务已经暂停该预警功能"
	} else if status == 2 {
		//将被拉入黑名单
		code = 20002
		message = "即将被拉入黑名单"
	} else if status == 4 {
		//已经是黑名单
		code = 20004
		message = "已经被拉入黑名单"
	}

	data := map[string]interface{}{
		"code":    code,     // 或直接用 chat_id 如 "123456789"=
		"address": _address, // 或直接用 chat_id 如 "123456789"=
		"message": message,
	}
	// 转换为 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON  parse error...:", err)
		return
	}

	// 发送 POST 请求

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("发送消息失败:", err)
		return
	}
	defer resp.Body.Close()

	// 打印响应结果
	//fmt.Println("消息发送状态:", resp.Status)
}

func (merchantAddressMonitorEventApi *MerchantAddressMonitorEventApi) DeleteMerchantAddressMonitorEvent(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := merchantAddressMonitorEventService.DeleteMerchantAddressMonitorEvent(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMerchantAddressMonitorEventByIds 批量删除merchantAddressMonitorEvent表
// @Tags MerchantAddressMonitorEvent
// @Summary 批量删除merchantAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /merchantAddressMonitorEvent/deleteMerchantAddressMonitorEventByIds [delete]
func (merchantAddressMonitorEventApi *MerchantAddressMonitorEventApi) DeleteMerchantAddressMonitorEventByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := merchantAddressMonitorEventService.DeleteMerchantAddressMonitorEventByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMerchantAddressMonitorEvent 更新merchantAddressMonitorEvent表
// @Tags MerchantAddressMonitorEvent
// @Summary 更新merchantAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ushield.MerchantAddressMonitorEvent true "更新merchantAddressMonitorEvent表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /merchantAddressMonitorEvent/updateMerchantAddressMonitorEvent [put]
func (merchantAddressMonitorEventApi *MerchantAddressMonitorEventApi) UpdateMerchantAddressMonitorEvent(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var merchantAddressMonitorEvent ushield.MerchantAddressMonitorEvent
	err := c.ShouldBindJSON(&merchantAddressMonitorEvent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = merchantAddressMonitorEventService.UpdateMerchantAddressMonitorEvent(ctx, merchantAddressMonitorEvent)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMerchantAddressMonitorEvent 用id查询merchantAddressMonitorEvent表
// @Tags MerchantAddressMonitorEvent
// @Summary 用id查询merchantAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询merchantAddressMonitorEvent表"
// @Success 200 {object} response.Response{data=ushield.MerchantAddressMonitorEvent,msg=string} "查询成功"
// @Router /merchantAddressMonitorEvent/findMerchantAddressMonitorEvent [get]
func (merchantAddressMonitorEventApi *MerchantAddressMonitorEventApi) FindMerchantAddressMonitorEvent(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	remerchantAddressMonitorEvent, err := merchantAddressMonitorEventService.GetMerchantAddressMonitorEvent(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(remerchantAddressMonitorEvent, c)
}

// GetMerchantAddressMonitorEventList 分页获取merchantAddressMonitorEvent表列表
// @Tags MerchantAddressMonitorEvent
// @Summary 分页获取merchantAddressMonitorEvent表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.MerchantAddressMonitorEventSearch true "分页获取merchantAddressMonitorEvent表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /merchantAddressMonitorEvent/getMerchantAddressMonitorEventList [get]
func (merchantAddressMonitorEventApi *MerchantAddressMonitorEventApi) GetMerchantAddressMonitorEventList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo ushieldReq.MerchantAddressMonitorEventSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := merchantAddressMonitorEventService.GetMerchantAddressMonitorEventInfoList(ctx, pageInfo)
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

// GetMerchantAddressMonitorEventPublic 不需要鉴权的merchantAddressMonitorEvent表接口
// @Tags MerchantAddressMonitorEvent
// @Summary 不需要鉴权的merchantAddressMonitorEvent表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /merchantAddressMonitorEvent/getMerchantAddressMonitorEventPublic [get]
func (merchantAddressMonitorEventApi *MerchantAddressMonitorEventApi) GetMerchantAddressMonitorEventPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	merchantAddressMonitorEventService.GetMerchantAddressMonitorEventPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的merchantAddressMonitorEvent表接口信息",
	}, "获取成功", c)
}
