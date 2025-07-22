import service from '@/utils/request'
// @Tags UserAddressMonitor
// @Summary 创建userAddressMonitor表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserAddressMonitor true "创建userAddressMonitor表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userAddressMonitor/createUserAddressMonitor [post]
export const createUserAddressMonitor = (data) => {
  return service({
    url: '/userAddressMonitor/createUserAddressMonitor',
    method: 'post',
    data
  })
}

// @Tags UserAddressMonitor
// @Summary 删除userAddressMonitor表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserAddressMonitor true "删除userAddressMonitor表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAddressMonitor/deleteUserAddressMonitor [delete]
export const deleteUserAddressMonitor = (params) => {
  return service({
    url: '/userAddressMonitor/deleteUserAddressMonitor',
    method: 'delete',
    params
  })
}

// @Tags UserAddressMonitor
// @Summary 批量删除userAddressMonitor表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userAddressMonitor表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAddressMonitor/deleteUserAddressMonitor [delete]
export const deleteUserAddressMonitorByIds = (params) => {
  return service({
    url: '/userAddressMonitor/deleteUserAddressMonitorByIds',
    method: 'delete',
    params
  })
}

// @Tags UserAddressMonitor
// @Summary 更新userAddressMonitor表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserAddressMonitor true "更新userAddressMonitor表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAddressMonitor/updateUserAddressMonitor [put]
export const updateUserAddressMonitor = (data) => {
  return service({
    url: '/userAddressMonitor/updateUserAddressMonitor',
    method: 'put',
    data
  })
}

// @Tags UserAddressMonitor
// @Summary 用id查询userAddressMonitor表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserAddressMonitor true "用id查询userAddressMonitor表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAddressMonitor/findUserAddressMonitor [get]
export const findUserAddressMonitor = (params) => {
  return service({
    url: '/userAddressMonitor/findUserAddressMonitor',
    method: 'get',
    params
  })
}

// @Tags UserAddressMonitor
// @Summary 分页获取userAddressMonitor表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userAddressMonitor表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAddressMonitor/getUserAddressMonitorList [get]
export const getUserAddressMonitorList = (params) => {
  return service({
    url: '/userAddressMonitor/getUserAddressMonitorList',
    method: 'get',
    params
  })
}

// @Tags UserAddressMonitor
// @Summary 不需要鉴权的userAddressMonitor表接口
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserAddressMonitorSearch true "分页获取userAddressMonitor表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userAddressMonitor/getUserAddressMonitorPublic [get]
export const getUserAddressMonitorPublic = () => {
  return service({
    url: '/userAddressMonitor/getUserAddressMonitorPublic',
    method: 'get',
  })
}
