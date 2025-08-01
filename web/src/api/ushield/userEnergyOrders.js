import service from '@/utils/request'
// @Tags UserEnergyOrders
// @Summary 创建userEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserEnergyOrders true "创建userEnergyOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userEnergyOrders/createUserEnergyOrders [post]
export const createUserEnergyOrders = (data) => {
  return service({
    url: '/userEnergyOrders/createUserEnergyOrders',
    method: 'post',
    data
  })
}

// @Tags UserEnergyOrders
// @Summary 删除userEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserEnergyOrders true "删除userEnergyOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userEnergyOrders/deleteUserEnergyOrders [delete]
export const deleteUserEnergyOrders = (params) => {
  return service({
    url: '/userEnergyOrders/deleteUserEnergyOrders',
    method: 'delete',
    params
  })
}

// @Tags UserEnergyOrders
// @Summary 批量删除userEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userEnergyOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userEnergyOrders/deleteUserEnergyOrders [delete]
export const deleteUserEnergyOrdersByIds = (params) => {
  return service({
    url: '/userEnergyOrders/deleteUserEnergyOrdersByIds',
    method: 'delete',
    params
  })
}

// @Tags UserEnergyOrders
// @Summary 更新userEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserEnergyOrders true "更新userEnergyOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userEnergyOrders/updateUserEnergyOrders [put]
export const updateUserEnergyOrders = (data) => {
  return service({
    url: '/userEnergyOrders/updateUserEnergyOrders',
    method: 'put',
    data
  })
}

// @Tags UserEnergyOrders
// @Summary 用id查询userEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserEnergyOrders true "用id查询userEnergyOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userEnergyOrders/findUserEnergyOrders [get]
export const findUserEnergyOrders = (params) => {
  return service({
    url: '/userEnergyOrders/findUserEnergyOrders',
    method: 'get',
    params
  })
}

// @Tags UserEnergyOrders
// @Summary 分页获取userEnergyOrders表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userEnergyOrders表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userEnergyOrders/getUserEnergyOrdersList [get]
export const getUserEnergyOrdersList = (params) => {
  return service({
    url: '/userEnergyOrders/getUserEnergyOrdersList',
    method: 'get',
    params
  })
}

// @Tags UserEnergyOrders
// @Summary 不需要鉴权的userEnergyOrders表接口
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserEnergyOrdersSearch true "分页获取userEnergyOrders表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userEnergyOrders/getUserEnergyOrdersPublic [get]
export const getUserEnergyOrdersPublic = () => {
  return service({
    url: '/userEnergyOrders/getUserEnergyOrdersPublic',
    method: 'get',
  })
}
