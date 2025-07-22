import service from '@/utils/request'
// @Tags UserUsdtDeposits
// @Summary 创建userUsdtDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserUsdtDeposits true "创建userUsdtDeposits表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userUsdtDeposits/createUserUsdtDeposits [post]
export const createUserUsdtDeposits = (data) => {
  return service({
    url: '/userUsdtDeposits/createUserUsdtDeposits',
    method: 'post',
    data
  })
}

// @Tags UserUsdtDeposits
// @Summary 删除userUsdtDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserUsdtDeposits true "删除userUsdtDeposits表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userUsdtDeposits/deleteUserUsdtDeposits [delete]
export const deleteUserUsdtDeposits = (params) => {
  return service({
    url: '/userUsdtDeposits/deleteUserUsdtDeposits',
    method: 'delete',
    params
  })
}

// @Tags UserUsdtDeposits
// @Summary 批量删除userUsdtDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userUsdtDeposits表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userUsdtDeposits/deleteUserUsdtDeposits [delete]
export const deleteUserUsdtDepositsByIds = (params) => {
  return service({
    url: '/userUsdtDeposits/deleteUserUsdtDepositsByIds',
    method: 'delete',
    params
  })
}

// @Tags UserUsdtDeposits
// @Summary 更新userUsdtDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserUsdtDeposits true "更新userUsdtDeposits表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userUsdtDeposits/updateUserUsdtDeposits [put]
export const updateUserUsdtDeposits = (data) => {
  return service({
    url: '/userUsdtDeposits/updateUserUsdtDeposits',
    method: 'put',
    data
  })
}

// @Tags UserUsdtDeposits
// @Summary 用id查询userUsdtDeposits表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserUsdtDeposits true "用id查询userUsdtDeposits表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userUsdtDeposits/findUserUsdtDeposits [get]
export const findUserUsdtDeposits = (params) => {
  return service({
    url: '/userUsdtDeposits/findUserUsdtDeposits',
    method: 'get',
    params
  })
}

// @Tags UserUsdtDeposits
// @Summary 分页获取userUsdtDeposits表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userUsdtDeposits表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userUsdtDeposits/getUserUsdtDepositsList [get]
export const getUserUsdtDepositsList = (params) => {
  return service({
    url: '/userUsdtDeposits/getUserUsdtDepositsList',
    method: 'get',
    params
  })
}

// @Tags UserUsdtDeposits
// @Summary 不需要鉴权的userUsdtDeposits表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.UserUsdtDepositsSearch true "分页获取userUsdtDeposits表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userUsdtDeposits/getUserUsdtDepositsPublic [get]
export const getUserUsdtDepositsPublic = () => {
  return service({
    url: '/userUsdtDeposits/getUserUsdtDepositsPublic',
    method: 'get',
  })
}
