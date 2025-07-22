import service from '@/utils/request'
// @Tags UserAddressMonitorEvent
// @Summary 创建userAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserAddressMonitorEvent true "创建userAddressMonitorEvent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userAddressMonitorEvent/createUserAddressMonitorEvent [post]
export const createUserAddressMonitorEvent = (data) => {
  return service({
    url: '/userAddressMonitorEvent/createUserAddressMonitorEvent',
    method: 'post',
    data
  })
}

// @Tags UserAddressMonitorEvent
// @Summary 删除userAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserAddressMonitorEvent true "删除userAddressMonitorEvent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAddressMonitorEvent/deleteUserAddressMonitorEvent [delete]
export const deleteUserAddressMonitorEvent = (params) => {
  return service({
    url: '/userAddressMonitorEvent/deleteUserAddressMonitorEvent',
    method: 'delete',
    params
  })
}

// @Tags UserAddressMonitorEvent
// @Summary 批量删除userAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userAddressMonitorEvent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAddressMonitorEvent/deleteUserAddressMonitorEvent [delete]
export const deleteUserAddressMonitorEventByIds = (params) => {
  return service({
    url: '/userAddressMonitorEvent/deleteUserAddressMonitorEventByIds',
    method: 'delete',
    params
  })
}

// @Tags UserAddressMonitorEvent
// @Summary 更新userAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserAddressMonitorEvent true "更新userAddressMonitorEvent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAddressMonitorEvent/updateUserAddressMonitorEvent [put]
export const updateUserAddressMonitorEvent = (data) => {
  return service({
    url: '/userAddressMonitorEvent/updateUserAddressMonitorEvent',
    method: 'put',
    data
  })
}

// @Tags UserAddressMonitorEvent
// @Summary 用id查询userAddressMonitorEvent表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserAddressMonitorEvent true "用id查询userAddressMonitorEvent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAddressMonitorEvent/findUserAddressMonitorEvent [get]
export const findUserAddressMonitorEvent = (params) => {
  return service({
    url: '/userAddressMonitorEvent/findUserAddressMonitorEvent',
    method: 'get',
    params
  })
}

// @Tags UserAddressMonitorEvent
// @Summary 分页获取userAddressMonitorEvent表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userAddressMonitorEvent表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAddressMonitorEvent/getUserAddressMonitorEventList [get]
export const getUserAddressMonitorEventList = (params) => {
  return service({
    url: '/userAddressMonitorEvent/getUserAddressMonitorEventList',
    method: 'get',
    params
  })
}

// @Tags UserAddressMonitorEvent
// @Summary 不需要鉴权的userAddressMonitorEvent表接口
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserAddressMonitorEventSearch true "分页获取userAddressMonitorEvent表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userAddressMonitorEvent/getUserAddressMonitorEventPublic [get]
export const getUserAddressMonitorEventPublic = () => {
  return service({
    url: '/userAddressMonitorEvent/getUserAddressMonitorEventPublic',
    method: 'get',
  })
}
