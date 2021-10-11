import service from '@/utils/request'

// @Tags Testset
// @Summary 创建Testset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Testset true "创建Testset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testset/createTestset [post]
export const createTestset = (data) => {
  return service({
    url: '/testset/createTestset',
    method: 'post',
    data
  })
}

// @Tags Testset
// @Summary 删除Testset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Testset true "删除Testset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testset/deleteTestset [delete]
export const deleteTestset = (data) => {
  return service({
    url: '/testset/deleteTestset',
    method: 'delete',
    data
  })
}

// @Tags Testset
// @Summary 删除Testset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Testset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testset/deleteTestset [delete]
export const deleteTestsetByIds = (data) => {
  return service({
    url: '/testset/deleteTestsetByIds',
    method: 'delete',
    data
  })
}

// @Tags Testset
// @Summary 更新Testset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Testset true "更新Testset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testset/updateTestset [put]
export const updateTestset = (data) => {
  return service({
    url: '/testset/updateTestset',
    method: 'put',
    data
  })
}

// @Tags Testset
// @Summary 用id查询Testset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Testset true "用id查询Testset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testset/findTestset [get]
export const findTestset = (params) => {
  return service({
    url: '/testset/findTestset',
    method: 'get',
    params
  })
}

// @Tags Testset
// @Summary 分页获取Testset列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Testset列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testset/getTestsetList [get]
export const getTestsetList = (params) => {
  return service({
    url: '/testset/getTestsetList',
    method: 'get',
    params
  })
}
