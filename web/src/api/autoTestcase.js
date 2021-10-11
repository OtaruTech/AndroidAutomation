import service from '@/utils/request'

// @Tags Testcase
// @Summary 创建Testcase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Testcase true "创建Testcase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testcase/createTestcase [post]
export const createTestcase = (data) => {
  return service({
    url: '/testcase/createTestcase',
    method: 'post',
    data
  })
}

// @Tags Testcase
// @Summary 删除Testcase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Testcase true "删除Testcase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testcase/deleteTestcase [delete]
export const deleteTestcase = (data) => {
  return service({
    url: '/testcase/deleteTestcase',
    method: 'delete',
    data
  })
}

// @Tags Testcase
// @Summary 删除Testcase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Testcase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testcase/deleteTestcase [delete]
export const deleteTestcaseByIds = (data) => {
  return service({
    url: '/testcase/deleteTestcaseByIds',
    method: 'delete',
    data
  })
}

// @Tags Testcase
// @Summary 更新Testcase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Testcase true "更新Testcase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testcase/updateTestcase [put]
export const updateTestcase = (data) => {
  return service({
    url: '/testcase/updateTestcase',
    method: 'put',
    data
  })
}

// @Tags Testcase
// @Summary 用id查询Testcase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Testcase true "用id查询Testcase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testcase/findTestcase [get]
export const findTestcase = (params) => {
  return service({
    url: '/testcase/findTestcase',
    method: 'get',
    params
  })
}

// @Tags Testcase
// @Summary 分页获取Testcase列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Testcase列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testcase/getTestcaseList [get]
export const getTestcaseList = (params) => {
  return service({
    url: '/testcase/getTestcaseList',
    method: 'get',
    params
  })
}
