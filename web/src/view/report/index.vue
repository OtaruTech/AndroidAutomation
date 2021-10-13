<template>
  <div id="report_list">
    <report-item
      v-for="(result, index) in reportList"
      :key="index"
      :test-result="result"
      :show-log="consoleLogShow"
    />
    <el-dialog v-model="consoleLogDialogVisible" title="控制台日志">
      <!-- <p>{{ theConsoleLog }}</p> -->
      <p
        v-for="(line, index) in logLines"
        :key="index"
      >{{ line }}
      </p>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="consoleLogDialogDismiss">关闭</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  getReportList
} from '@/api/autoReport'
import infoList from '@/mixins/infoList'
import reportItem from './reportItem.vue'

export default {
  name: 'Report',
  components: {
    reportItem
  },
  mixins: [infoList],
  data() {
    return {
      logLines: [],
      consoleLogDialogVisible: false,
      listApi: getReportList,
      reportList: [],
      spanArr: [],
      deviceList: []
    }
  },
  async created() {
    const reportListData = await this.listApi({ page: 1, size: 1000 })
    this.reportList = reportListData.data.list
    this.reportList = this.reportList.reverse()
  },
  methods: {
    consoleLogShow(consoleLog) {
      this.consoleLogDialogVisible = true
      this.logLines = consoleLog.split('\n')
    },
    consoleLogDialogDismiss() {
      this.consoleLogDialogVisible = false
    }
  },
}
</script>

<style>
</style>
