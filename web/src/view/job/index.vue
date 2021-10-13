<template>
  <div id="job_list">
    <el-table
      ref="multipleTable"
      style="width: 100%"
      tooltip-effect="dark"
      :data="jobList"
      row-key="ID"
    >
      <el-table-column align="left" label="名称" prop="name" min-width="40%" />
      <el-table-column align="left" label="时间点" min-width="15%">
        <template #default="scope">
          <span>每天{{ scope.row.hour }}点
          </span>
        </template>
      </el-table-column>
      <el-table-column align="left" label="产品" prop="product" min-width="20%" />
      <el-table-column align="left" label="用户" prop="owner" min-width="20%" />
      <el-table-column align="left" label="打开" min-width="5%">
        <template #default="scope">
          <el-switch
            v-model="scope.row.enable"
            type="text"
            icon="el-icon-edit"
            size="small"
            @change="changeSwitch(scope.row)"
          >{{ scope.row.enable }}</el-switch>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>

import {
  getJobList
} from '@/api/autoJob'
import {
  jobChanged
} from '@/api/android'
export default {
  name: 'Job',
  data() {
    return {
      getJobListApi: getJobList,
      jobList: [],
    }
  },
  async created() {
    const jobListData = await this.getJobListApi({ page: 1, size: 100 })
    this.jobList = jobListData.data.list
    console.log(this.jobList)
  },
  methods: {
    async changeSwitch(row) {
      const req = {
        id: row.ID,
        enable: row.enable
      }
      await jobChanged(req)
    }
  },
}
</script>

<style>
</style>
