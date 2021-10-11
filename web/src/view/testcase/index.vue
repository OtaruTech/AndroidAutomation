<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="模块">
          <el-input v-model="searchInfo.component" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="标签">
          <el-input v-model="searchInfo.tags" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="超时时间">
          <el-input v-model="searchInfo.timeout" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="el-icon-refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          size="mini"
          type="primary"
          icon="el-icon-plus"
          :disabled="multipleSelection.length == 0"
          @click="createTestsetDialog"
        >创建测试集 </el-button>
        <el-button
          size="mini"
          type="danger"
          icon="el-icon-download"
          :disabled="multipleSelection.length == 0"
          @click="runTestcase"
        >运行测试用例</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" min-width="10%" />
        <el-table-column align="left" label="名称" prop="name" min-width="50%" />
        <el-table-column align="left" label="模块" prop="component" min-width="15%" />
        <el-table-column align="left" label="标签" prop="tags" min-width="15%" />
        <el-table-column align="left" label="超时时间" prop="timeout" min-width="10%" />
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
    <el-dialog v-model="createTestsetDialogVisible" :before-close="closeCreateTestsetDialog" title="创建测试集">
      <el-form :model="testsetFormData" label-position="right" label-width="80px">
        <el-form-item label="名称:">
          <el-input v-model="testsetFormData.name" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeCreateTestsetDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterCreateTestsetDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog v-model="runTestcaseDialogVisible" :before-close="closeRunTestcaseDialog" title="运行测试用例">
      <el-form :model="runTestcaseFormData" label-position="right" label-width="80px">
        <el-form-item label="名称:">
          <el-input v-model="runTestcaseFormData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="版本:">
          <el-input v-model="runTestcaseFormData.otaUrl" clearable placeholder="可选参数" />
        </el-form-item>
        <el-form-item label="设备:">
          <el-input v-model="runTestcaseFormData.serialNo" clearable placeholder="可选参数" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeRunTestcaseDialog">取消</el-button>
          <el-button size="small" type="primary" @click="confirmRunTestcaseDialog">确定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>

</template>

<script>
import {
  getTestcaseList
} from '@/api/autoTestcase'
import {
  getDeviceList
} from '@/api/autoDevice'
import {
  runTestcase,
  getRuntimeState
} from '@/api/android'
import {
  createTestRunner,
} from '@/api/autoTestRunner'
import {
  createTestset,
} from '@/api/auto_testset'
import { getUserInfo } from '@/api/user'
import infoList from '@/mixins/infoList'
export default {
  name: 'Testcase',
  mixins: [infoList],
  data() {
    return {
      listApi: getTestcaseList,
      deviceListApi: getDeviceList,
      createTestsetApi: createTestset,
      deviceList: [],
      createTestsetDialogVisible: false,
      multipleSelection: [],
      testsetFormData: {
        name: '',
      },
      runTestcaseDialogVisible: false,
      runTestcaseFormData: {
        name: '',
        otaUrl: '',
        serialNo: '',
      },
      userInfo: '',
    }
  },
  async created() {
    await this.getTableData()
    var deviceData = await this.deviceListApi()
    this.deviceList = deviceData.data.list
    this.userInfo = await getUserInfo()
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
    closeCreateTestsetDialog() {
      this.createTestsetDialogVisible = false
      this.testsetFormData = {
        name: '',
      }
    },
    async enterCreateTestsetDialog() {
      if (this.testsetFormData.name === '') {
        this.$message({
          type: 'warning',
          message: '请输入测试集名称'
        })
        return
      }
      var testset = {
        name: this.testsetFormData.name,
        testcases: ''
      }
      var testcases = []
      this.multipleSelection.forEach(item => {
        testcases.push(item.name)
      })
      testset.testcases = testcases.join(',')
      console.log(testset)
      const res = await createTestset(testset)
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建测试集: ' + testset.name
        })
      } else {
        this.$message({
          type: 'error',
          message: '创建失败'
        })
      }
      this.createTestsetDialogVisible = false
    },
    createTestsetDialog() {
      this.createTestsetDialogVisible = true
    },
    showRunTestcaseDialog() {
      this.runTestcaseDialogVisible = true
    },
    closeRunTestcaseDialog() {
      this.runTestcaseDialogVisible = false
    },
    async confirmRunTestcaseDialog() {
      if (this.runTestcaseFormData.name === '') {
        this.$message({
          type: 'warning',
          message: '请输入任务名称'
        })
        return
      }
      var serialNo = ''
      if (this.runTestcaseFormData.serialNo === '') {
        for (var i = 0; i < this.deviceList.length; i++) {
          var param = {
            serialNo: this.deviceList[i].serialNo,
          }
          const rsp = await getRuntimeState(param)
          if (rsp.code === 0) {
            console.log('serialNo: ' + param.serialNo + ', state: ' + rsp.data.state)
            if (rsp.data.state === 0) {
              serialNo = param.serialNo
              break
            }
          }
        }
        if (serialNo === '') {
          this.$message({
            type: 'warning',
            message: '找不到空闲设备'
          })
          return
        }
      } else {
        serialNo = this.runTestcaseFormData.serialNo
      }
      // run testcase
      var runParam = {
        testcases: [],
        serialNo: serialNo,
        timeout: 0,
        otaUrl: this.runTestcaseFormData.otaUrl
      }
      this.multipleSelection.forEach(item => {
        runParam.timeout += item.timeout
        runParam.testcases.push(item.name)
      })
      var response = await runTestcase(runParam)
      console.log(response)
      if (response.code === 0) {
        const loading = this.$loading({
          lock: true,
          text: '正在运行测试用例，请稍候',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.7)'
        })
        this.timer = setTimeout(() => { // 设置延迟执行
          this.$message({
            type: 'success',
            message: '运行测试用例'
          })
          loading.close()
        }, 1000)
      } else {
        this.$message({
          type: 'error',
          message: '运行测试用例失败'
        })
        return
      }
      // based on the test id to store the test runner
      var testRunner = {
        testId: response.data.testId,
        testcases: runParam.testcases.join(','),
        name: this.runTestcaseFormData.name,
        owner: this.userInfo.data.userInfo.userName,
        serialNo: serialNo
      }
      const res = await createTestRunner(testRunner)
      console.log('createTestRunner: ' + res)
      this.runTestcaseDialogVisible = false
      this.toDeviceDetail(serialNo)
    },
    runTestcase() {
      this.showRunTestcaseDialog()
    },
    toDeviceDetail(serialNo) {
      this.$router.push({ name: 'DeviceDetail', query: { serialNo: serialNo }})
    }
  },
}
</script>

<style>
</style>
