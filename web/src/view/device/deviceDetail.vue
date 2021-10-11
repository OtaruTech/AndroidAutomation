<template>
  <div>
    <div class="deviceInfo">
      <span>Model: {{ device.model }}</span>
      <span style="margin-left:20px;">Product: {{ device.product }}</span>
      <span style="margin-left:20px;">Serial: {{ device.serialNo }}</span>
      <span style="margin-left:20px;">Build ID: {{ device.buildId }}</span>
      <span style="margin-left:20px;">Build Type: {{ device.buildType }}</span>
      <span style="margin-left:20px;">Android {{ device.version }}</span>
    </div>
    <div class="deviceInfoDetail">
      <span>First API Level: {{ device.api }}</span>
      <span style="margin-left:20px;">Configuration Number: {{ device.config }}</span>
      <span style="margin-left:20px;">Camera: {{ device.camera }}</span>
      <span style="margin-left:20px;">Scanner: {{ device.scanner }}</span>
      <span style="margin-left:20px;">Security Level: {{ device.security }}</span>
      <span style="margin-left:20px;">WWAN: {{ device.wwan }}</span>
    </div>
    <div class="console" style="white-space: pre-wrap;">
      <p class="consoleLog">{{ console }}</p>
      <i class="el-icon-loading" style="margin-top:8px" :style="{display:loadingDisplay}" />
    </div>
    <!-- <i :class="el-icon-loading" /> -->
  </div>
</template>

<script>
import {
  getDeviceList,
} from '@/api/autoDevice'
import {
  getDeviceConsoleList,
} from '@/api/autoDeviceConsole'
import {
  getRuntimeState
} from '@/api/android'
export default {
  name: 'DeviceDetail',
  data() {
    return {
      getDeviceListApi: getDeviceList,
      getDeviceConsoleListApi: getDeviceConsoleList,
      device: {
        serialNo: ''
      },
      console,
      timer: null,
      loadingDisplay: ''
    }
  },
  async created() {
    const serialNo = this.$route.query.serialNo
    console.log('Device Detail: ' + serialNo)
    const deviceListData = await this.getDeviceListApi({ page: 1, size: 1000 })
    const deviceList = deviceListData.data.list

    for (let i = 0; i < deviceList.length; i++) {
      if (deviceList[i].serialNo === serialNo) {
        this.device = deviceList[i]
        break
      }
    }
    console.log(this.device)
    if (this.device.serialNo === '') {
      this.$message({
        type: 'error',
        message: '查找不到指定设备'
      })
    }
    const deviceConsoleListData = await this.getDeviceConsoleListApi({ page: 1, size: 1000 })
    const deviceConsoleList = deviceConsoleListData.data.list
    for (let i = 0; i < deviceConsoleList.length; i++) {
      if (deviceConsoleList[i].serialNo === serialNo) {
        this.console = deviceConsoleList[i].console
        break
      }
    }
    var param = {
      serialNo: this.device.serialNo
    }
    const rsp = await getRuntimeState(param)
    if (rsp.code === 0) {
      if (rsp.data.state === 1) {
        this.loadingDisplay = ''
      } else {
        this.loadingDisplay = 'none'
      }
    }
    this.timer = setInterval(async() => {
      const deviceConsoleListData = await this.getDeviceConsoleListApi({ page: 1, size: 1000 })
      const deviceConsoleList = deviceConsoleListData.data.list
      for (let i = 0; i < deviceConsoleList.length; i++) {
        if (deviceConsoleList[i].serialNo === serialNo) {
          this.console = deviceConsoleList[i].console
          break
        }
      }
      const rsp = await getRuntimeState(param)
      if (rsp.code === 0) {
        if (rsp.data.state === 1) {
          this.loadingDisplay = ''
        } else {
          this.loadingDisplay = 'none'
        }
      }
    }, 1000 * 10)
  },
  beforeUnmount() {
    clearInterval(this.timer)
    this.timer = null
  },
  methods: {
  }
}
</script>

<style scoped>

.deviceInfo {
  background-color: #fff;
  padding: 16px;
  font-size: 16px;
  font-weight: bolder;
}

.deviceInfoDetail {
  background-color: #fff;
  padding: 16px;
  font-size: 16px;
  font-weight: bolder;
}

.console {
  margin-top: 16px;
  background-color: #fff;
  padding: 16px;
}

.consoleLog {
  font-size: 15px;
}
</style>
