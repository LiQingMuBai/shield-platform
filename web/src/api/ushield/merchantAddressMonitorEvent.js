import service from '@/utils/request'
// @Tags MerchantAddressMonitorEvent
// @Summary 创建merchantAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAddressMonitorEvent true "创建merchantAddressMonitorEvent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /merchantAddressMonitorEvent/createMerchantAddressMonitorEvent [post]
export const createMerchantAddressMonitorEvent = (data) => {
  return service({
    url: '/merchantAddressMonitorEvent/createMerchantAddressMonitorEvent',
    method: 'post',
    data
  })
}

// @Tags MerchantAddressMonitorEvent
// @Summary 删除merchantAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAddressMonitorEvent true "删除merchantAddressMonitorEvent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /merchantAddressMonitorEvent/deleteMerchantAddressMonitorEvent [delete]
export const deleteMerchantAddressMonitorEvent = (params) => {
  return service({
    url: '/merchantAddressMonitorEvent/deleteMerchantAddressMonitorEvent',
    method: 'delete',
    params
  })
}

// @Tags MerchantAddressMonitorEvent
// @Summary 批量删除merchantAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除merchantAddressMonitorEvent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /merchantAddressMonitorEvent/deleteMerchantAddressMonitorEvent [delete]
export const deleteMerchantAddressMonitorEventByIds = (params) => {
  return service({
    url: '/merchantAddressMonitorEvent/deleteMerchantAddressMonitorEventByIds',
    method: 'delete',
    params
  })
}

// @Tags MerchantAddressMonitorEvent
// @Summary 更新merchantAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAddressMonitorEvent true "更新merchantAddressMonitorEvent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /merchantAddressMonitorEvent/updateMerchantAddressMonitorEvent [put]
export const updateMerchantAddressMonitorEvent = (data) => {
  return service({
    url: '/merchantAddressMonitorEvent/updateMerchantAddressMonitorEvent',
    method: 'put',
    data
  })
}

// @Tags MerchantAddressMonitorEvent
// @Summary 用id查询merchantAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MerchantAddressMonitorEvent true "用id查询merchantAddressMonitorEvent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /merchantAddressMonitorEvent/findMerchantAddressMonitorEvent [get]
export const findMerchantAddressMonitorEvent = (params) => {
  return service({
    url: '/merchantAddressMonitorEvent/findMerchantAddressMonitorEvent',
    method: 'get',
    params
  })
}

// @Tags MerchantAddressMonitorEvent
// @Summary 分页获取merchantAddressMonitorEvent表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取merchantAddressMonitorEvent表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /merchantAddressMonitorEvent/getMerchantAddressMonitorEventList [get]
export const getMerchantAddressMonitorEventList = (params) => {
  return service({
    url: '/merchantAddressMonitorEvent/getMerchantAddressMonitorEventList',
    method: 'get',
    params
  })
}

// @Tags MerchantAddressMonitorEvent
// @Summary 不需要鉴权的merchantAddressMonitorEvent表接口
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.MerchantAddressMonitorEventSearch true "分页获取merchantAddressMonitorEvent表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /merchantAddressMonitorEvent/getMerchantAddressMonitorEventPublic [get]
export const getMerchantAddressMonitorEventPublic = () => {
  return service({
    url: '/merchantAddressMonitorEvent/getMerchantAddressMonitorEventPublic',
    method: 'get',
  })
}
