import service from '@/utils/request'
// @Tags UserTrxDeposits
// @Summary 创建userTrxDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserTrxDeposits true "创建userTrxDeposits表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userTrxDeposits/createUserTrxDeposits [post]
export const createUserTrxDeposits = (data) => {
  return service({
    url: '/userTrxDeposits/createUserTrxDeposits',
    method: 'post',
    data
  })
}

// @Tags UserTrxDeposits
// @Summary 删除userTrxDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserTrxDeposits true "删除userTrxDeposits表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userTrxDeposits/deleteUserTrxDeposits [delete]
export const deleteUserTrxDeposits = (params) => {
  return service({
    url: '/userTrxDeposits/deleteUserTrxDeposits',
    method: 'delete',
    params
  })
}

// @Tags UserTrxDeposits
// @Summary 批量删除userTrxDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userTrxDeposits表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userTrxDeposits/deleteUserTrxDeposits [delete]
export const deleteUserTrxDepositsByIds = (params) => {
  return service({
    url: '/userTrxDeposits/deleteUserTrxDepositsByIds',
    method: 'delete',
    params
  })
}

// @Tags UserTrxDeposits
// @Summary 更新userTrxDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserTrxDeposits true "更新userTrxDeposits表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userTrxDeposits/updateUserTrxDeposits [put]
export const updateUserTrxDeposits = (data) => {
  return service({
    url: '/userTrxDeposits/updateUserTrxDeposits',
    method: 'put',
    data
  })
}

// @Tags UserTrxDeposits
// @Summary 用id查询userTrxDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserTrxDeposits true "用id查询userTrxDeposits表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userTrxDeposits/findUserTrxDeposits [get]
export const findUserTrxDeposits = (params) => {
  return service({
    url: '/userTrxDeposits/findUserTrxDeposits',
    method: 'get',
    params
  })
}

// @Tags UserTrxDeposits
// @Summary 分页获取userTrxDeposits表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userTrxDeposits表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userTrxDeposits/getUserTrxDepositsList [get]
export const getUserTrxDepositsList = (params) => {
  return service({
    url: '/userTrxDeposits/getUserTrxDepositsList',
    method: 'get',
    params
  })
}

// @Tags UserTrxDeposits
// @Summary 不需要鉴权的userTrxDeposits表接口
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserTrxDepositsSearch true "分页获取userTrxDeposits表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userTrxDeposits/getUserTrxDepositsPublic [get]
export const getUserTrxDepositsPublic = () => {
  return service({
    url: '/userTrxDeposits/getUserTrxDepositsPublic',
    method: 'get',
  })
}
