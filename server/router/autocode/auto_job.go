package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type JobRouter struct {
}

// InitJobRouter 初始化 Job 路由信息
func (s *JobRouter) InitJobRouter(Router *gin.RouterGroup) {
	jobRouter := Router.Group("job").Use(middleware.OperationRecord())
	jobRouterWithoutRecord := Router.Group("job")
	var jobApi = v1.ApiGroupApp.AutoCodeApiGroup.JobApi
	{
		jobRouter.POST("createJob", jobApi.CreateJob)   // 新建Job
		jobRouter.DELETE("deleteJob", jobApi.DeleteJob) // 删除Job
		jobRouter.DELETE("deleteJobByIds", jobApi.DeleteJobByIds) // 批量删除Job
		jobRouter.PUT("updateJob", jobApi.UpdateJob)    // 更新Job
	}
	{
		jobRouterWithoutRecord.GET("findJob", jobApi.FindJob)        // 根据ID获取Job
		jobRouterWithoutRecord.GET("getJobList", jobApi.GetJobList)  // 获取Job列表
	}
}
