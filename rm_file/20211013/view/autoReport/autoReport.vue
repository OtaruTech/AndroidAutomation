<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="测试ID">
          <el-input v-model="searchInfo.testId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="测试用例">
          <el-input v-model="searchInfo.testcases" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="结果">
          <el-input v-model="searchInfo.result" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="日志文件">
          <el-input v-model="searchInfo.logcat" placeholder="搜索条件" />
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
        <el-table-column align="left" label="测试ID" prop="testId" width="120" />
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="测试用例" prop="testcases" width="120" />
        <el-table-column align="left" label="结果" prop="result" width="120" />
        <el-table-column align="left" label="日志文件" prop="logcat" width="120" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="text" icon="el-icon-edit" size="small" class="table-button" @click="updateReport(scope.row)">变更</el-button>
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
        <el-form-item label="测试ID:">
          <el-input v-model.number="formData.testId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="测试用例:">
          <el-input v-model="formData.testcases" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结果:">
          <el-input v-model="formData.result" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="日志文件:">
          <el-input v-model="formData.logcat" clearable placeholder="请输入" />
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
  createReport,
  deleteReport,
  deleteReportByIds,
  updateReport,
  findReport,
  getReportList
} from '@/api/autoReport' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Report',
  mixins: [infoList],
  data() {
    return {
      listApi: getReportList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        testId: 0,
        name: '',
        testcases: '',
        result: '',
        logcat: '',
      }
    }
  },
  async created() {
    await this.getTableData()
    this.tableData.forEach(data => {
      data.result = data.result.substr(0, 16)
    })
  },
  methods: {
    onReset() {
      this.searchInfo = {}
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
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
        this.deleteReport(row)
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
      const res = await deleteReportByIds({ ids })
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
    async updateReport(row) {
      const res = await findReport({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.rereport
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        testId: 0,
        name: '',
        testcases: '',
        result: '',
        logcat: '',
      }
    },
    async deleteReport(row) {
      const res = await deleteReport({ ID: row.ID })
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
          res = await createReport(this.formData)
          break
        case 'update':
          res = await updateReport(this.formData)
          break
        default:
          res = await createReport(this.formData)
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
