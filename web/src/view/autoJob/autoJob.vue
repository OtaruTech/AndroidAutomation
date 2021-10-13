<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="版本路劲">
          <el-input v-model="searchInfo.otaPath" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="版本格式">
          <el-input v-model="searchInfo.otaFormat" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="时间点">
          <el-input v-model="searchInfo.hour" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="产品">
          <el-input v-model="searchInfo.product" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="测试用例">
          <el-input v-model="searchInfo.testcases" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="用户">
          <el-input v-model="searchInfo.owner" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="版本链接">
          <el-input v-model="searchInfo.otaUrl" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="使能" prop="enable">
          <el-select v-model="searchInfo.enable" clearable placeholder="请选择">
            <el-option
              key="true"
              label="是"
              value="true"
            />
            <el-option
              key="false"
              label="否"
              value="false"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="el-icon-refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="el-icon-delete" size="mini" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="版本路劲" prop="otaPath" width="120" />
        <el-table-column align="left" label="版本格式" prop="otaFormat" width="120" />
        <el-table-column align="left" label="文件格式" prop="fileFormat" width="120" />
        <el-table-column align="left" label="时间点" prop="hour" width="120" />
        <el-table-column align="left" label="产品" prop="product" width="120" />
        <el-table-column align="left" label="测试用例" prop="testcases" width="120" />
        <el-table-column align="left" label="用户" prop="owner" width="120" />
        <el-table-column align="left" label="版本链接" prop="otaUrl" width="120" />
        <el-table-column align="left" label="使能" prop="enable" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.enable) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="text" icon="el-icon-edit" size="small" class="table-button" @click="updateJob(scope.row)">变更</el-button>
            <el-button type="text" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="版本路劲:">
          <el-input v-model="formData.otaPath" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="版本格式:">
          <el-input v-model="formData.otaFormat" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件格式:">
          <el-input v-model="formData.fileFormat" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="时间点:">
          <el-input v-model.number="formData.hour" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="产品:">
          <el-input v-model="formData.product" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="测试用例:">
          <el-input v-model="formData.testcases" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户:">
          <el-input v-model="formData.owner" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="版本链接:">
          <el-input v-model="formData.otaUrl" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="使能:">
          <el-switch v-model="formData.enable" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createJob,
  deleteJob,
  deleteJobByIds,
  updateJob,
  findJob,
  getJobList
} from '@/api/autoJob' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Job',
  mixins: [infoList],
  data() {
    return {
      listApi: getJobList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        name: '',
        otaPath: '',
        otaFormat: '',
        fileFormat: '',
        hour: 0,
        product: '',
        testcases: '',
        owner: '',
        otaUrl: '',
        enable: false,
      }
    }
  },
  async created() {
    await this.getTableData()
  },
  methods: {
    onReset() {
      this.searchInfo = {}
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.enable === '') {
        this.searchInfo.enable = null
      }
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteJob(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteJobByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateJob(row) {
      const res = await findJob({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.rejob
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        name: '',
        otaPath: '',
        otaFormat: '',
        fileFormat: '',
        hour: 0,
        product: '',
        testcases: '',
        owner: '',
        otaUrl: '',
        enable: false,
      }
    },
    async deleteJob(row) {
      const res = await deleteJob({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createJob(this.formData)
          break
        case 'update':
          res = await updateJob(this.formData)
          break
        default:
          res = await createJob(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
}
</script>

<style>
</style>
