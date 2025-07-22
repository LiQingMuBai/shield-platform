import service from '@/utils/request'
// @Tags TgUsers
// @Summary 创建tgUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TgUsers true "创建tgUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tgUsers/createTgUsers [post]
export const createTgUsers = (data) => {
  return service({
    url: '/tgUsers/createTgUsers',
    method: 'post',
    data
  })
}

// @Tags TgUsers
// @Summary 删除tgUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TgUsers true "删除tgUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tgUsers/deleteTgUsers [delete]
export const deleteTgUsers = (params) => {
  return service({
    url: '/tgUsers/deleteTgUsers',
    method: 'delete',
    params
  })
}

// @Tags TgUsers
// @Summary 批量删除tgUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除tgUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tgUsers/deleteTgUsers [delete]
export const deleteTgUsersByIds = (params) => {
  return service({
    url: '/tgUsers/deleteTgUsersByIds',
    method: 'delete',
    params
  })
}

// @Tags TgUsers
// @Summary 更新tgUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TgUsers true "更新tgUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tgUsers/updateTgUsers [put]
export const updateTgUsers = (data) => {
  return service({
    url: '/tgUsers/updateTgUsers',
    method: 'put',
    data
  })
}

// @Tags TgUsers
// @Summary 用id查询tgUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.TgUsers true "用id查询tgUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tgUsers/findTgUsers [get]
export const findTgUsers = (params) => {
  return service({
    url: '/tgUsers/findTgUsers',
    method: 'get',
    params
  })
}

// @Tags TgUsers
// @Summary 分页获取tgUsers表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取tgUsers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tgUsers/getTgUsersList [get]
export const getTgUsersList = (params) => {
  return service({
    url: '/tgUsers/getTgUsersList',
    method: 'get',
    params
  })
}

// @Tags TgUsers
// @Summary 不需要鉴权的tgUsers表接口
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.TgUsersSearch true "分页获取tgUsers表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tgUsers/getTgUsersPublic [get]
export const getTgUsersPublic = () => {
  return service({
    url: '/tgUsers/getTgUsersPublic',
    method: 'get',
  })
}
