import service from '@/utils/request'
// @Tags SysAnnouncementsInfo
// @Summary 创建sysAnnouncementsInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysAnnouncementsInfo true "创建sysAnnouncementsInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysAnnouncementsInfo/createSysAnnouncementsInfo [post]
export const createSysAnnouncementsInfo = (data) => {
  return service({
    url: '/sysAnnouncementsInfo/createSysAnnouncementsInfo',
    method: 'post',
    data
  })
}

// @Tags SysAnnouncementsInfo
// @Summary 删除sysAnnouncementsInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysAnnouncementsInfo true "删除sysAnnouncementsInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysAnnouncementsInfo/deleteSysAnnouncementsInfo [delete]
export const deleteSysAnnouncementsInfo = (params) => {
  return service({
    url: '/sysAnnouncementsInfo/deleteSysAnnouncementsInfo',
    method: 'delete',
    params
  })
}

// @Tags SysAnnouncementsInfo
// @Summary 批量删除sysAnnouncementsInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysAnnouncementsInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysAnnouncementsInfo/deleteSysAnnouncementsInfo [delete]
export const deleteSysAnnouncementsInfoByIds = (params) => {
  return service({
    url: '/sysAnnouncementsInfo/deleteSysAnnouncementsInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags SysAnnouncementsInfo
// @Summary 更新sysAnnouncementsInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysAnnouncementsInfo true "更新sysAnnouncementsInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysAnnouncementsInfo/updateSysAnnouncementsInfo [put]
export const updateSysAnnouncementsInfo = (data) => {
  return service({
    url: '/sysAnnouncementsInfo/updateSysAnnouncementsInfo',
    method: 'put',
    data
  })
}

// @Tags SysAnnouncementsInfo
// @Summary 用id查询sysAnnouncementsInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SysAnnouncementsInfo true "用id查询sysAnnouncementsInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysAnnouncementsInfo/findSysAnnouncementsInfo [get]
export const findSysAnnouncementsInfo = (params) => {
  return service({
    url: '/sysAnnouncementsInfo/findSysAnnouncementsInfo',
    method: 'get',
    params
  })
}

// @Tags SysAnnouncementsInfo
// @Summary 分页获取sysAnnouncementsInfo表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysAnnouncementsInfo表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysAnnouncementsInfo/getSysAnnouncementsInfoList [get]
export const getSysAnnouncementsInfoList = (params) => {
  return service({
    url: '/sysAnnouncementsInfo/getSysAnnouncementsInfoList',
    method: 'get',
    params
  })
}

// @Tags SysAnnouncementsInfo
// @Summary 不需要鉴权的sysAnnouncementsInfo表接口
// @Accept application/json
// @Produce application/json
// @Param data query ushieldReq.SysAnnouncementsInfoSearch true "分页获取sysAnnouncementsInfo表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysAnnouncementsInfo/getSysAnnouncementsInfoPublic [get]
export const getSysAnnouncementsInfoPublic = () => {
  return service({
    url: '/sysAnnouncementsInfo/getSysAnnouncementsInfoPublic',
    method: 'get',
  })
}
