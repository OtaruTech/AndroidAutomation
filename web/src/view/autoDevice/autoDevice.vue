<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="设备序列号">
          <el-input v-model="searchInfo.serialNo" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="设备型号">
          <el-input v-model="searchInfo.model" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="产品名称">
          <el-input v-model="searchInfo.product" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="版本">
          <el-input v-model="searchInfo.version" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="版本号">
          <el-input v-model="searchInfo.buildId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="版本类型">
          <el-input v-model="searchInfo.buildType" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="sku">
          <el-input v-model="searchInfo.sku" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="sdk">
          <el-input v-model="searchInfo.sdk" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="安全日期">
          <el-input v-model="searchInfo.security" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="API">
          <el-input v-model="searchInfo.api" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="配置信息">
          <el-input v-model="searchInfo.config" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="摄像头">
          <el-input v-model="searchInfo.camera" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="扫描引擎">
          <el-input v-model="searchInfo.scanner" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="移动网络">
          <el-input v-model="searchInfo.wwan" placeholder="搜索条件" />
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
        <el-table-column align="left" label="设备序列号" prop="serialNo" width="120" />
        <el-table-column align="left" label="设备型号" prop="model" width="120" />
        <el-table-column align="left" label="产品名称" prop="product" width="120" />
        <el-table-column align="left" label="版本" prop="version" width="120" />
        <el-table-column align="left" label="版本号" prop="buildId" width="120" />
        <el-table-column align="left" label="版本类型" prop="buildType" width="120" />
        <el-table-column align="left" label="sku" prop="sku" width="120" />
        <el-table-column align="left" label="sdk" prop="sdk" width="120" />
        <el-table-column align="left" label="安全日期" prop="security" width="120" />
        <el-table-column align="left" label="API" prop="api" width="120" />
        <el-table-column align="left" label="配置信息" prop="config" width="120" />
        <el-table-column align="left" label="摄像头" prop="camera" width="120" />
        <el-table-column align="left" label="扫描引擎" prop="scanner" width="120" />
        <el-table-column align="left" label="移动网络" prop="wwan" width="120" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="text" icon="el-icon-edit" size="small" class="table-button" @click="updateDevice(scope.row)">变更</el-button>
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
        <el-form-item label="设备序列号:">
          <el-input v-model="formData.serialNo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="设备型号:">
          <el-input v-model="formData.model" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="产品名称:">
          <el-input v-model="formData.product" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="版本:">
          <el-input v-model="formData.version" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="版本号:">
          <el-input v-model="formData.buildId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="版本类型:">
          <el-input v-model="formData.buildType" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="sku:">
          <el-input v-model="formData.sku" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="sdk:">
          <el-input v-model="formData.sdk" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="安全日期:">
          <el-input v-model="formData.security" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="API:">
          <el-input v-model="formData.api" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="配置信息:">
          <el-input v-model="formData.config" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="摄像头:">
          <el-input v-model="formData.camera" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="扫描引擎:">
          <el-input v-model="formData.scanner" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="移动网络:">
          <el-input v-model="formData.wwan" clearable placeholder="请输入" />
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
  createDevice,
  deleteDevice,
  deleteDeviceByIds,
  updateDevice,
  findDevice,
  getDeviceList
} from '@/api/autoDevice' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Device',
  mixins: [infoList],
  data() {
    return {
      listApi: getDeviceList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        serialNo: '',
        model: '',
        product: '',
        version: '',
        buildId: '',
        buildType: '',
        sku: '',
        sdk: '',
        security: '',
        api: '',
        config: '',
        camera: '',
        scanner: '',
        wwan: '',
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
        this.deleteDevice(row)
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
      const res = await deleteDeviceByIds({ ids })
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
    async updateDevice(row) {
      const res = await findDevice({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.redevice
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        serialNo: '',
        model: '',
        product: '',
        version: '',
        buildId: '',
        buildType: '',
        sku: '',
        sdk: '',
        security: '',
        api: '',
        config: '',
        camera: '',
        scanner: '',
        wwan: '',
      }
    },
    async deleteDevice(row) {
      const res = await deleteDevice({ ID: row.ID })
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
          res = await createDevice(this.formData)
          break
        case 'update':
          res = await updateDevice(this.formData)
          break
        default:
          res = await createDevice(this.formData)
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
