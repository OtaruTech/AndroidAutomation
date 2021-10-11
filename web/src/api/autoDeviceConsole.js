import service from '@/utils/request'

// @Tags DeviceConsole
// @Summary 创建DeviceConsole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceConsole true "创建DeviceConsole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceConsole/createDeviceConsole [post]
export const createDeviceConsole = (data) => {
  return service({
    url: '/deviceConsole/createDeviceConsole',
    method: 'post',
    data
  })
}

// @Tags DeviceConsole
// @Summary 删除DeviceConsole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceConsole true "删除DeviceConsole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceConsole/deleteDeviceConsole [delete]
export const deleteDeviceConsole = (data) => {
  return service({
    url: '/deviceConsole/deleteDeviceConsole',
    method: 'delete',
    data
  })
}

// @Tags DeviceConsole
// @Summary 删除DeviceConsole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DeviceConsole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceConsole/deleteDeviceConsole [delete]
export const deleteDeviceConsoleByIds = (data) => {
  return service({
    url: '/deviceConsole/deleteDeviceConsoleByIds',
    method: 'delete',
    data
  })
}

// @Tags DeviceConsole
// @Summary 更新DeviceConsole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceConsole true "更新DeviceConsole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deviceConsole/updateDeviceConsole [put]
export const updateDeviceConsole = (data) => {
  return service({
    url: '/deviceConsole/updateDeviceConsole',
    method: 'put',
    data
  })
}

// @Tags DeviceConsole
// @Summary 用id查询DeviceConsole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DeviceConsole true "用id查询DeviceConsole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deviceConsole/findDeviceConsole [get]
export const findDeviceConsole = (params) => {
  return service({
    url: '/deviceConsole/findDeviceConsole',
    method: 'get',
    params
  })
}

// @Tags DeviceConsole
// @Summary 分页获取DeviceConsole列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DeviceConsole列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceConsole/getDeviceConsoleList [get]
export const getDeviceConsoleList = (params) => {
  return service({
    url: '/deviceConsole/getDeviceConsoleList',
    method: 'get',
    params
  })
}
