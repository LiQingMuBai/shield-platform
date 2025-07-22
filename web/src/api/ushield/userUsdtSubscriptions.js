import service from '@/utils/request'
// @Tags UserUsdtSubscriptions
// @Summary 创建userUsdtSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserUsdtSubscriptions true "创建userUsdtSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userUsdtSubscriptions/createUserUsdtSubscriptions [post]
export const createUserUsdtSubscriptions = (data) => {
  return service({
    url: '/userUsdtSubscriptions/createUserUsdtSubscriptions',
    method: 'post',
    data
  })
}

// @Tags UserUsdtSubscriptions
// @Summary 删除userUsdtSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserUsdtSubscriptions true "删除userUsdtSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userUsdtSubscriptions/deleteUserUsdtSubscriptions [delete]
export const deleteUserUsdtSubscriptions = (params) => {
  return service({
    url: '/userUsdtSubscriptions/deleteUserUsdtSubscriptions',
    method: 'delete',
    params
  })
}

// @Tags UserUsdtSubscriptions
// @Summary 批量删除userUsdtSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userUsdtSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userUsdtSubscriptions/deleteUserUsdtSubscriptions [delete]
export const deleteUserUsdtSubscriptionsByIds = (params) => {
  return service({
    url: '/userUsdtSubscriptions/deleteUserUsdtSubscriptionsByIds',
    method: 'delete',
    params
  })
}

// @Tags UserUsdtSubscriptions
// @Summary 更新userUsdtSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserUsdtSubscriptions true "更新userUsdtSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userUsdtSubscriptions/updateUserUsdtSubscriptions [put]
export const updateUserUsdtSubscriptions = (data) => {
  return service({
    url: '/userUsdtSubscriptions/updateUserUsdtSubscriptions',
    method: 'put',
    data
  })
}

// @Tags UserUsdtSubscriptions
// @Summary 用id查询userUsdtSubscriptions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserUsdtSubscriptions true "用id查询userUsdtSubscriptions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userUsdtSubscriptions/findUserUsdtSubscriptions [get]
export const findUserUsdtSubscriptions = (params) => {
  return service({
    url: '/userUsdtSubscriptions/findUserUsdtSubscriptions',
    method: 'get',
    params
  })
}

// @Tags UserUsdtSubscriptions
// @Summary 分页获取userUsdtSubscriptions表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userUsdtSubscriptions表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userUsdtSubscriptions/getUserUsdtSubscriptionsList [get]
export const getUserUsdtSubscriptionsList = (params) => {
  return service({
    url: '/userUsdtSubscriptions/getUserUsdtSubscriptionsList',
    method: 'get',
    params
  })
}

// @Tags UserUsdtSubscriptions
// @Summary 不需要鉴权的userUsdtSubscriptions表接口
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserUsdtSubscriptionsSearch true "分页获取userUsdtSubscriptions表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userUsdtSubscriptions/getUserUsdtSubscriptionsPublic [get]
export const getUserUsdtSubscriptionsPublic = () => {
  return service({
    url: '/userUsdtSubscriptions/getUserUsdtSubscriptionsPublic',
    method: 'get',
  })
}
