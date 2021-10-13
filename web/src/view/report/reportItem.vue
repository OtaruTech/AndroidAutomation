<template>
  <div class="div_table">
    <el-table
      ref="multipleTable"
      style="width: 100%"
      :header-cell-style="{fontWeight:900}"
      :data="testResults"
      tooltip-effect="dark"
      row-key="ID"
      :span-method="objectSpanMethod"
    >
      <el-table-column align="left" label="日期" min-width="14%">
        <template #default="scope">
          <span style="font-weight: bold;">{{ scope.row.serialNo }}</span>
          <br>
          <span>{{ formatDate(scope.row.date) }}</span>
        </template>
      </el-table-column>
      <el-table-column align="left" label="名称" prop="name" min-width="12%">
        <template #default="scope">
          <span style="font-weight: bold;">Build: {{ scope.row.buildId }}</span>
          <br>
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column align="left" label="用户" min-width="8%">
        <template #default="scope">
          <span style="font-weight: bold;">{{ scope.row.owner }}</span>
        </template></el-table-column>
      <el-table-column align="left" label="测试用例" prop="testcase" min-width="32%" />
      <el-table-column align="left" label="结果" min-width="7%">
        <template #default="scope">
          <span
            style="cursor: pointer;font-weight: bold;"
            :style="{color: (scope.row.result === 'Pass') ? 'green' : 'red'}"
            @click="resultOnClick(scope.row)"
          >{{ scope.row.result }}</span>
        </template>
      </el-table-column>
      <el-table-column align="left" label="用时/秒" prop="time" min-width="8%" />
      <el-table-column align="left" label="日志下载" min-width="16%">
        <template #default="scope">
          <span style="color: blue;cursor: pointer;" @click="downloadLogFile(scope.row.logcat)">{{ scope.row.logcat }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import infoList from '@/mixins/infoList'
import {
  getTestRunnerList
} from '@/api/autoTestRunner'
import {
  downloadFile,
} from '@/api/android'
export default {
  name: 'ReportItem',
  mixins: [infoList],
  props: {
    testResult: Object,
    showLog: Function
  },
  data() {
    return {
      border: true,
      spanArr: [],
      testResults: [],
      downloadFile: downloadFile
    }
  },
  async created() {
    const runnerData = await getTestRunnerList()
    var runnerList = runnerData.data.list
    var runnerName = ''
    var owner = ''
    var serialNo = ''
    for (let i = 0; i < runnerList.length; i++) {
      if (this.testResult.testId === runnerList[i].testId) {
        runnerName = runnerList[i].name
        owner = runnerList[i].owner
        serialNo = 'Serial: ' + runnerList[i].serialNo
        break
      }
    }

    const result = this.testResult.result
    const results = result.split('|||')
    results.forEach(ret => {
      var result = {}
      result.serialNo = serialNo
      result.date = this.testResult.CreatedAt
      result.logcat = this.testResult.logcat
      result.url = this.testResult.logcat
      result.resultDetail = ret
      result.testcase = this.getTestcaseFromResult(ret)
      result.result = this.getResultFromResult(ret)
      result.time = this.getTimeFromResult(ret)
      result.name = runnerName
      result.owner = owner
      result.buildId = this.testResult.buildId
      this.testResults.push(result)
    })
    this.getSpanArr(results)
  },
  methods: {
    async downloadLogFile(file) {
      await this.downloadFile(file)
    },
    resultOnClick(row) {
      this.showLog(row.resultDetail)
    },
    getSpanArr(data) {
      // data就是我们从后台拿到的数据
      for (var i = 0; i < data.length; i++) {
        if (i === 0) {
          this.spanArr.push(1)
          this.pos = 0
        } else {
          // 判断当前元素与上一个元素是否相同
          if (data[i].id === data[i - 1].id) {
            this.spanArr[this.pos] += 1
            this.spanArr.push(0)
          } else {
            this.spanArr.push(1)
            this.pos = i
          }
        }
      }
    },
    objectSpanMethod({ row, column, rowIndex, columnIndex }) {
      if (columnIndex === 0 || columnIndex === 1 || columnIndex === 2 || columnIndex === 6) {
        const _row = this.spanArr[rowIndex]
        const _col = _row > 0 ? 1 : 0
        return {
          rowspan: _row,
          colspan: _col
        }
      }
    },
    getTestcaseFromResult(result) {
      var lines = result.split('\n')
      if (lines.length > 7 && lines[0].length > 30 && lines[0].indexOf('INSTRUMENTATION_STATUS: class=') !== -1 &&
          lines[6].length > 29 && lines[6].indexOf('INSTRUMENTATION_STATUS: test=') !== -1) {
        const packageName = lines[0].substr(30)
        const className = lines[6].substr(29)
        return packageName + '#' + className
      }
      return '测试用例运行失败'
    },
    getResultFromResult(result) {
      if (result.indexOf('FAILURES!!!') !== -1) {
        return 'Fail'
      } else {
        if (result.indexOf('OK') !== -1) {
          return 'Pass'
        } else {
          return 'Unknown'
        }
      }
    },
    getTimeFromResult(result) {
      var lines = result.split('\n')
      for (let i = 0; i < lines.length; i++) {
        if (lines[i].indexOf('Time:') !== -1) {
          return lines[i].substr(6)
        }
      }
      return 0
    }
  },
}
</script>

<style>
.div_table {
  margin-bottom: 10px;
}
</style>
