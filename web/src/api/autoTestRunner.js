import service from '@/utils/request'

// @Tags TestRunner
// @Summary 创建TestRunner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestRunner true "创建TestRunner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testRunner/createTestRunner [post]
export const createTestRunner = (data) => {
  return service({
    url: '/testRunner/createTestRunner',
    method: 'post',
    data
  })
}

// @Tags TestRunner
// @Summary 删除TestRunner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestRunner true "删除TestRunner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testRunner/deleteTestRunner [delete]
export const deleteTestRunner = (data) => {
  return service({
    url: '/testRunner/deleteTestRunner',
    method: 'delete',
    data
  })
}

// @Tags TestRunner
// @Summary 删除TestRunner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestRunner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testRunner/deleteTestRunner [delete]
export const deleteTestRunnerByIds = (data) => {
  return service({
    url: '/testRunner/deleteTestRunnerByIds',
    method: 'delete',
    data
  })
}

// @Tags TestRunner
// @Summary 更新TestRunner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestRunner true "更新TestRunner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testRunner/updateTestRunner [put]
export const updateTestRunner = (data) => {
  return service({
    url: '/testRunner/updateTestRunner',
    method: 'put',
    data
  })
}

// @Tags TestRunner
// @Summary 用id查询TestRunner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TestRunner true "用id查询TestRunner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testRunner/findTestRunner [get]
export const findTestRunner = (params) => {
  return service({
    url: '/testRunner/findTestRunner',
    method: 'get',
    params
  })
}

// @Tags TestRunner
// @Summary 分页获取TestRunner列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TestRunner列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testRunner/getTestRunnerList [get]
export const getTestRunnerList = (params) => {
  return service({
    url: '/testRunner/getTestRunnerList',
    method: 'get',
    params
  })
}
