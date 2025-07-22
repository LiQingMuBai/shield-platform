import service from '@/utils/request'
// @Tags UserTrxSubscriptions
// @Summary 创建userTrxSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserTrxSubscriptions true "创建userTrxSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userTrxSubscriptions/createUserTrxSubscriptions [post]
export const createUserTrxSubscriptions = (data) => {
  return service({
    url: '/userTrxSubscriptions/createUserTrxSubscriptions',
    method: 'post',
    data
  })
}

// @Tags UserTrxSubscriptions
// @Summary 删除userTrxSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserTrxSubscriptions true "删除userTrxSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userTrxSubscriptions/deleteUserTrxSubscriptions [delete]
export const deleteUserTrxSubscriptions = (params) => {
  return service({
    url: '/userTrxSubscriptions/deleteUserTrxSubscriptions',
    method: 'delete',
    params
  })
}

// @Tags UserTrxSubscriptions
// @Summary 批量删除userTrxSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userTrxSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userTrxSubscriptions/deleteUserTrxSubscriptions [delete]
export const deleteUserTrxSubscriptionsByIds = (params) => {
  return service({
    url: '/userTrxSubscriptions/deleteUserTrxSubscriptionsByIds',
    method: 'delete',
    params
  })
}

// @Tags UserTrxSubscriptions
// @Summary 更新userTrxSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserTrxSubscriptions true "更新userTrxSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userTrxSubscriptions/updateUserTrxSubscriptions [put]
export const updateUserTrxSubscriptions = (data) => {
  return service({
    url: '/userTrxSubscriptions/updateUserTrxSubscriptions',
    method: 'put',
    data
  })
}

// @Tags UserTrxSubscriptions
// @Summary 用id查询userTrxSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserTrxSubscriptions true "用id查询userTrxSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userTrxSubscriptions/findUserTrxSubscriptions [get]
export const findUserTrxSubscriptions = (params) => {
  return service({
    url: '/userTrxSubscriptions/findUserTrxSubscriptions',
    method: 'get',
    params
  })
}

// @Tags UserTrxSubscriptions
// @Summary 分页获取userTrxSubscriptions表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userTrxSubscriptions表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userTrxSubscriptions/getUserTrxSubscriptionsList [get]
export const getUserTrxSubscriptionsList = (params) => {
  return service({
    url: '/userTrxSubscriptions/getUserTrxSubscriptionsList',
    method: 'get',
    params
  })
}

// @Tags UserTrxSubscriptions
// @Summary 不需要鉴权的userTrxSubscriptions表接口
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserTrxSubscriptionsSearch true "分页获取userTrxSubscriptions表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userTrxSubscriptions/getUserTrxSubscriptionsPublic [get]
export const getUserTrxSubscriptionsPublic = () => {
  return service({
    url: '/userTrxSubscriptions/getUserTrxSubscriptionsPublic',
    method: 'get',
  })
}
