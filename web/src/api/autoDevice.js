import service from '@/utils/request'

// @Tags Device
// @Summary 创建Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Device true "创建Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /device/createDevice [post]
export const createDevice = (data) => {
  return service({
    url: '/device/createDevice',
    method: 'post',
    data
  })
}

// @Tags Device
// @Summary 删除Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Device true "删除Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /device/deleteDevice [delete]
export const deleteDevice = (data) => {
  return service({
    url: '/device/deleteDevice',
    method: 'delete',
    data
  })
}

// @Tags Device
// @Summary 删除Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /device/deleteDevice [delete]
export const deleteDeviceByIds = (data) => {
  return service({
    url: '/device/deleteDeviceByIds',
    method: 'delete',
    data
  })
}

// @Tags Device
// @Summary 更新Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Device true "更新Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /device/updateDevice [put]
export const updateDevice = (data) => {
  return service({
    url: '/device/updateDevice',
    method: 'put',
    data
  })
}

// @Tags Device
// @Summary 用id查询Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Device true "用id查询Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /device/findDevice [get]
export const findDevice = (params) => {
  return service({
    url: '/device/findDevice',
    method: 'get',
    params
  })
}

// @Tags Device
// @Summary 分页获取Device列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Device列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /device/getDeviceList [get]
export const getDeviceList = (params) => {
  return service({
    url: '/device/getDeviceList',
    method: 'get',
    params
  })
}
