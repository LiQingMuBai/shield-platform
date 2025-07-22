import service from '@/utils/request'
// @Tags UserOperationBundles
// @Summary 创建userOperationBundles表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserOperationBundles true "创建userOperationBundles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userOperationBundles/createUserOperationBundles [post]
export const createUserOperationBundles = (data) => {
  return service({
    url: '/userOperationBundles/createUserOperationBundles',
    method: 'post',
    data
  })
}

// @Tags UserOperationBundles
// @Summary 删除userOperationBundles表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserOperationBundles true "删除userOperationBundles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userOperationBundles/deleteUserOperationBundles [delete]
export const deleteUserOperationBundles = (params) => {
  return service({
    url: '/userOperationBundles/deleteUserOperationBundles',
    method: 'delete',
    params
  })
}

// @Tags UserOperationBundles
// @Summary 批量删除userOperationBundles表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userOperationBundles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userOperationBundles/deleteUserOperationBundles [delete]
export const deleteUserOperationBundlesByIds = (params) => {
  return service({
    url: '/userOperationBundles/deleteUserOperationBundlesByIds',
    method: 'delete',
    params
  })
}

// @Tags UserOperationBundles
// @Summary 更新userOperationBundles表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserOperationBundles true "更新userOperationBundles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userOperationBundles/updateUserOperationBundles [put]
export const updateUserOperationBundles = (data) => {
  return service({
    url: '/userOperationBundles/updateUserOperationBundles',
    method: 'put',
    data
  })
}

// @Tags UserOperationBundles
// @Summary 用id查询userOperationBundles表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserOperationBundles true "用id查询userOperationBundles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userOperationBundles/findUserOperationBundles [get]
export const findUserOperationBundles = (params) => {
  return service({
    url: '/userOperationBundles/findUserOperationBundles',
    method: 'get',
    params
  })
}

// @Tags UserOperationBundles
// @Summary 分页获取userOperationBundles表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userOperationBundles表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userOperationBundles/getUserOperationBundlesList [get]
export const getUserOperationBundlesList = (params) => {
  return service({
    url: '/userOperationBundles/getUserOperationBundlesList',
    method: 'get',
    params
  })
}

// @Tags UserOperationBundles
// @Summary 不需要鉴权的userOperationBundles表接口
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.UserOperationBundlesSearch true "分页获取userOperationBundles表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userOperationBundles/getUserOperationBundlesPublic [get]
export const getUserOperationBundlesPublic = () => {
  return service({
    url: '/userOperationBundles/getUserOperationBundlesPublic',
    method: 'get',
  })
}
