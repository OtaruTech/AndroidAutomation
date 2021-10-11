<template>
  <div id="device_list">
    <device-item
      v-for="(deviceInfo, index) in deviceList"
      :key="index"
      :device-info="deviceInfo"
      :run-testcase="runTestcase"
      :index="index"
    />
    <el-dialog v-model="runTestcaseDialogVisible" :before-close="runTestcaseCloseDialog" title="运行测试用例">
      <div>
        <el-checkbox v-model="otaUpgradeChecked" /><span class="ota_upgrade_text">系统版本升级</span>
        <el-input v-model="otaUrl" placeholder="输入系统版本号" :disabled="!otaUpgradeChecked" />
      </div>
      <div class="testcase_type">
        <span class="testcase_type_text">测试用例分类</span>
        <el-radio-group
          v-for="item in testcaseTypeTab"
          :key="item.id"
          v-model="testcaseType"
          class="testcase_type_radio_group"
          @change="testcaseToggleTab(item.id)"
        >
          <div class="radio_button">
            <el-radio-button
              :label="item.label"
              class="testcase_type_radio"
            />
          </div>
        </el-radio-group>
        <el-select v-model="testcaseTypeValue" placeholder="请选择" class="testcase_type_select">
          <el-option
            v-for="testcase in testcaseOptions"
            :key="testcase.value"
            :label="testcase.label"
            :value="testcase.value"
          />
        </el-select>
        <div class="named">
          <span class="testcase_named">任务命名</span>
          <br>
          <el-input v-model="testRunnerName" class="testcase_named_input" placeholder="给测试任务起个名字" />
        </div>
      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="runTestcaseCloseDialog">取消</el-button>
          <el-button size="small" type="primary" @click="runTestcaseEnterDialog">确定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>

</template>

<script>
import {
  getDeviceList
} from '@/api/autoDevice'
import {
  getTestcaseList
} from '@/api/autoTestcase'
import {
  runTestcase,
  getRuntimeState
} from '@/api/android'
import {
  createTestRunner,
} from '@/api/autoTestRunner'
import {
  getTestsetList
} from '@/api/auto_testset'
import infoList from '@/mixins/infoList'
import deviceItem from './deviceItem.vue'
import { getUserInfo } from '@/api/user'
export default {
  name: 'Device',
  components: {
    deviceItem
  },
  mixins: [infoList],
  data() {
    return {
      selectDeviceIndex: -1,
      listApi: getDeviceList,
      getTestsetListApi: getTestsetList,
      deviceList: [],
      testsetList: [],
      testcaseListApi: getTestcaseList,
      testcaseDataTable: [],
      runTestcaseDialogVisible: false,
      testcaseTypeTab: [
        { id: 0, label: '预定义' },
        { id: 1, label: '模块' },
        { id: 2, label: '标签' },
      ],
      testcaseType: '模块',
      testcaseTypeId: 1,
      testcaseOptions: [],
      testcaseTypeValue: '',
      testRunnerName: '',
      otaUrl: '',
      otaUpgradeChecked: false,
      userInfo: '',
      timer: null
    }
  },
  async created() {
    await this.getTableData()
    var response = await this.testcaseListApi({ page: 1, size: 1000 })
    this.testcaseDataTable = response.data.list
    this.testcaseOptions.length = 0
    var testsetData = await this.getTestsetListApi({ page: 1, size: 100 })
    this.testsetList = testsetData.data.list
    this.testcaseDataTable.forEach(tc => {
      var exist = false
      this.testcaseOptions.forEach(option => {
        if (option.label === tc.component) {
          exist = true
        }
      })
      if (exist !== true) {
        this.testcaseOptions.push({
          value: tc.component,
          label: tc.component
        })
      }
    })
    this.updateRuntimeState()
    this.userInfo = await getUserInfo()
    // update runtime state
    this.timer = setInterval(() => {
      this.updateRuntimeState()
    }, 1000 * 30)
  },
  beforeUnmount() {
    clearInterval(this.timer)
    this.timer = null
  },
  methods: {
    runTestcaseCloseDialog() {
      this.runTestcaseDialogVisible = false
    },
    async updateRuntimeState() {
      this.deviceList.length = 0
      for (let i = 0; i < this.tableData.length; i++) {
        const devInfo = {
          model: this.tableData[i].model,
          serialNo: this.tableData[i].serialNo,
          buildId: this.tableData[i].buildId,
          state: 2,
          duration: 0,
          testcaseName: ''
        }
        var param = {
          serialNo: this.tableData[i].serialNo,
        }
        const rsp = await getRuntimeState(param)
        if (rsp.code === 0) {
          devInfo.state = rsp.data.state
          devInfo.testcaseName = rsp.data.testcaseName
          devInfo.duration = rsp.data.duration
        }
        this.deviceList.push(devInfo)
      }
    },
    async runTestcaseEnterDialog() {
      if (this.otaUpgradeChecked && this.otaUrl === '') {
        this.$message({
          type: 'warning',
          message: '请输入OTA下载链接'
        })
        return
      }
      if (this.testcaseTypeValue === '') {
        this.$message({
          type: 'warning',
          message: '请选择正确的测试用例'
        })
      } else {
        var serialNo = this.tableData[this.selectDeviceIndex].serialNo
        var param = {
          testcases: [],
          serialNo: serialNo,
          timeout: 0,
          otaUrl: ''
        }
        if (this.otaUpgradeChecked) {
          param.otaUrl = this.otaUrl
        } else {
          param.otaUrl = ''
        }
        var timeout = 0
        if (this.testcaseTypeId === 1) { // component
          this.testcaseDataTable.forEach(tc => {
            if (this.testcaseTypeValue === tc.component) {
              param.testcases.push(tc.name)
              timeout += tc.timeout
            }
          })
        } else if (this.testcaseTypeId === 2) { // tags
          this.testcaseDataTable.forEach(tc => {
            if (this.testcaseTypeValue === tc.tags) {
              param.testcases.push(tc.name)
              timeout += tc.timeout
            }
          })
        } else if (this.testcaseTypeId === 0) {
          for (var i = 0; i < this.testsetList.length; i++) {
            if (this.testcaseTypeValue === this.testsetList[i].name) {
              param.testcases = this.testsetList[i].testcases.split(',')
              break
            }
          }
        }
        param.timeout = timeout
        var response = await runTestcase(param)
        // console.log(response)
        // based on the test id to store the test runner
        var runnerName = ''
        // console.log('test runner name: ' + this.testRunnerName)
        if (this.testRunnerName !== '') {
          runnerName = this.testRunnerName
        } else {
          runnerName = 'dev-' + serialNo + '-' + response.data.testId
        }
        var testRunner = {
          testId: response.data.testId,
          testcases: param.testcases.join(','),
          name: runnerName,
          owner: this.userInfo.data.userInfo.userName,
          serialNo: serialNo
        }
        await createTestRunner(testRunner)
        // console.log('createTestRunner: ' + res)
        this.runTestcaseDialogVisible = false
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
            this.updateRuntimeState()
            loading.close()
            this.toDeviceDetail(serialNo)
          }, 1000)
        } else {
          this.$message({
            type: 'error',
            message: '运行测试用例失败'
          })
        }
      }
      // run testcase
      // var response = runTestcase({ testcases: ['camera', 'scanner'], serialNo: '123456' })
      // console.log(response)
    },
    runTestcase(id) {
      // console.log('Run Testcase device: ' + id)
      this.runTestcaseDialogVisible = true
      this.selectDeviceIndex = id
      this.testRunnerName = ''
    },
    testcaseToggleTab(id) {
      this.testcaseTypeValue = ''
      this.testcaseTypeId = id
      this.testcaseOptions.length = 0
      // testset
      if (id === 0) {
        this.testsetList.forEach(set => {
          this.testcaseOptions.push({
            value: set.name,
            label: set.name
          })
        })
        return
      }
      this.testcaseDataTable.forEach(tc => {
        var exist = false
        this.testcaseOptions.forEach(option => {
          if (id === 1) {
            if (option.label === tc.component) {
              exist = true
            }
          } else if (id === 2) {
            if (option.label === tc.tags) {
              exist = true
            }
          }
        })
        if (exist !== true) {
          if (id === 1) {
            this.testcaseOptions.push({
              value: tc.component,
              label: tc.component
            })
          } else if (id === 2) {
            this.testcaseOptions.push({
              value: tc.tags,
              label: tc.tags
            })
          }
        }
      })
    },
    toDeviceDetail(serialNo) {
      this.$router.push({ name: 'DeviceDetail', query: { serialNo: serialNo }})
    }
  },
}
</script>

<style scoped>
.ota_upgrade_text {
  font-size: 15px;
  font-weight: bold;
  margin-left: 6px;
}

.testcase_type {
  margin-top: 10px;
}

.testcase_type_text {
  font-size: 15px;
  font-weight: bold;
}

.testcase_type_radio_group {
  margin-top: 10px;
}

.testcase_type_radio {
  display: inline;
}

.testcase_type_select {
  margin-top: 10px;
}

.radio_button {
  display: inline;
}

.named {
    margin-top: 10px;
}

.testcase_named {
  font-size: 15px;
  font-weight: bold;
}

.named_input {
    margin-top: 10px;
}
</style>
