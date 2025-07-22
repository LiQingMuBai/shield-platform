import service from '@/utils/request'
// @Tags UserPackageSubscriptions
// @Summary 创建userPackageSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPackageSubscriptions true "创建userPackageSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userPackageSubscriptions/createUserPackageSubscriptions [post]
export const createUserPackageSubscriptions = (data) => {
  return service({
    url: '/userPackageSubscriptions/createUserPackageSubscriptions',
    method: 'post',
    data
  })
}

// @Tags UserPackageSubscriptions
// @Summary 删除userPackageSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPackageSubscriptions true "删除userPackageSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userPackageSubscriptions/deleteUserPackageSubscriptions [delete]
export const deleteUserPackageSubscriptions = (params) => {
  return service({
    url: '/userPackageSubscriptions/deleteUserPackageSubscriptions',
    method: 'delete',
    params
  })
}

// @Tags UserPackageSubscriptions
// @Summary 批量删除userPackageSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userPackageSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userPackageSubscriptions/deleteUserPackageSubscriptions [delete]
export const deleteUserPackageSubscriptionsByIds = (params) => {
  return service({
    url: '/userPackageSubscriptions/deleteUserPackageSubscriptionsByIds',
    method: 'delete',
    params
  })
}

// @Tags UserPackageSubscriptions
// @Summary 更新userPackageSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPackageSubscriptions true "更新userPackageSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userPackageSubscriptions/updateUserPackageSubscriptions [put]
export const updateUserPackageSubscriptions = (data) => {
  return service({
    url: '/userPackageSubscriptions/updateUserPackageSubscriptions',
    method: 'put',
    data
  })
}

// @Tags UserPackageSubscriptions
// @Summary 用id查询userPackageSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserPackageSubscriptions true "用id查询userPackageSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userPackageSubscriptions/findUserPackageSubscriptions [get]
export const findUserPackageSubscriptions = (params) => {
  return service({
    url: '/userPackageSubscriptions/findUserPackageSubscriptions',
    method: 'get',
    params
  })
}

// @Tags UserPackageSubscriptions
// @Summary 分页获取userPackageSubscriptions表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userPackageSubscriptions表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userPackageSubscriptions/getUserPackageSubscriptionsList [get]
export const getUserPackageSubscriptionsList = (params) => {
  return service({
    url: '/userPackageSubscriptions/getUserPackageSubscriptionsList',
    method: 'get',
    params
  })
}

// @Tags UserPackageSubscriptions
// @Summary 不需要鉴权的userPackageSubscriptions表接口
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserPackageSubscriptionsSearch true "分页获取userPackageSubscriptions表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userPackageSubscriptions/getUserPackageSubscriptionsPublic [get]
export const getUserPackageSubscriptionsPublic = () => {
  return service({
    url: '/userPackageSubscriptions/getUserPackageSubscriptionsPublic',
    method: 'get',
  })
}
