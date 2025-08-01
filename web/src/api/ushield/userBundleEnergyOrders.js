import service from '@/utils/request'
// @Tags UserBundleEnergyOrders
// @Summary 创建userBundleEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserBundleEnergyOrders true "创建userBundleEnergyOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userBundleEnergyOrders/createUserBundleEnergyOrders [post]
export const createUserBundleEnergyOrders = (data) => {
  return service({
    url: '/userBundleEnergyOrders/createUserBundleEnergyOrders',
    method: 'post',
    data
  })
}

// @Tags UserBundleEnergyOrders
// @Summary 删除userBundleEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserBundleEnergyOrders true "删除userBundleEnergyOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userBundleEnergyOrders/deleteUserBundleEnergyOrders [delete]
export const deleteUserBundleEnergyOrders = (params) => {
  return service({
    url: '/userBundleEnergyOrders/deleteUserBundleEnergyOrders',
    method: 'delete',
    params
  })
}

// @Tags UserBundleEnergyOrders
// @Summary 批量删除userBundleEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userBundleEnergyOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userBundleEnergyOrders/deleteUserBundleEnergyOrders [delete]
export const deleteUserBundleEnergyOrdersByIds = (params) => {
  return service({
    url: '/userBundleEnergyOrders/deleteUserBundleEnergyOrdersByIds',
    method: 'delete',
    params
  })
}

// @Tags UserBundleEnergyOrders
// @Summary 更新userBundleEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserBundleEnergyOrders true "更新userBundleEnergyOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userBundleEnergyOrders/updateUserBundleEnergyOrders [put]
export const updateUserBundleEnergyOrders = (data) => {
  return service({
    url: '/userBundleEnergyOrders/updateUserBundleEnergyOrders',
    method: 'put',
    data
  })
}

// @Tags UserBundleEnergyOrders
// @Summary 用id查询userBundleEnergyOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserBundleEnergyOrders true "用id查询userBundleEnergyOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userBundleEnergyOrders/findUserBundleEnergyOrders [get]
export const findUserBundleEnergyOrders = (params) => {
  return service({
    url: '/userBundleEnergyOrders/findUserBundleEnergyOrders',
    method: 'get',
    params
  })
}

// @Tags UserBundleEnergyOrders
// @Summary 分页获取userBundleEnergyOrders表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userBundleEnergyOrders表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userBundleEnergyOrders/getUserBundleEnergyOrdersList [get]
export const getUserBundleEnergyOrdersList = (params) => {
  return service({
    url: '/userBundleEnergyOrders/getUserBundleEnergyOrdersList',
    method: 'get',
    params
  })
}

// @Tags UserBundleEnergyOrders
// @Summary 不需要鉴权的userBundleEnergyOrders表接口
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserBundleEnergyOrdersSearch true "分页获取userBundleEnergyOrders表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userBundleEnergyOrders/getUserBundleEnergyOrdersPublic [get]
export const getUserBundleEnergyOrdersPublic = () => {
  return service({
    url: '/userBundleEnergyOrders/getUserBundleEnergyOrdersPublic',
    method: 'get',
  })
}
