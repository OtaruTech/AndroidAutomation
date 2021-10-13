<template>
  <div class="device" @click="onDeviceClick">
    <div class="device_left">
      <img src="@/assets/device/CT40.png" alt="" class="device_logo">
      <div class="device_info">
        <h5 class="device_title">{{ deviceInfo.model }}</h5>
        <h6 class="device_serial">Serial: {{ deviceInfo.serialNo }}</h6>
        <h6 class="device_build">Build: {{ deviceInfo.buildId.substr(0, 24) }}</h6>
      </div>
      <div class="device_state">
        <i :class="state.icon" :style="{ color: state.color }" style="margin-right: 4px" /><span>{{ state.status }}</span>
        <br>
        <span style="color: gray">{{ state.duration }}</span>
      </div>
    </div>
    <div class="device_right" @mouseenter="enterButton" @mouseleave="leaveButton">
      <el-button
        size="mini"
        type="primary"
        icon="el-icon-download"
        :disabled="deviceInfo.state != 0"
        @click="actionRunTestcase"
      >Run Testcase</el-button>
      <el-button
        size="mini"
        type="primary"
        icon="el-icon-top"
        :disabled="deviceInfo.state != 0"
      >OTA Upgrade</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DeviceItem',
  props: {
    index: Number,
    deviceInfo: Object,
    runTestcase: Function
  },
  data() {
    return {
      state: {
        icon: 'el-icon-error',
        color: '#c0c0c0',
        status: 'Offline',
        duration: '',
        onTheButton: false,
      }
    }
  },
  created() {
    // console.log(this.deviceInfo)
    if (this.deviceInfo.state === 1) {
      this.state.status = 'Running: ' + this.deviceInfo.testcaseName
      this.state.icon = 'el-icon-loading'
      this.state.color = '#000000'
      this.state.duration = 'Duration: ' + Math.round(this.deviceInfo.duration / 60) + 'min'
    } else if (this.deviceInfo.state === 0) {
      this.state.status = 'Available'
      this.state.icon = 'el-icon-success'
      this.state.color = '#00ff00'
    } else {
      this.state.status = 'Offline'
      this.state.icon = 'el-icon-error'
      this.state.color = '#c0c0c0'
    }
  },
  methods: {
    actionRunTestcase() {
      this.runTestcase(this.index)
    },
    onDeviceClick() {
      if (!this.onTheButton) {
        console.log('device click!')
        this.toDeviceDetail()
      }
    },
    enterButton() {
      this.onTheButton = true
    },
    leaveButton() {
      this.onTheButton = false
    },
    toDeviceDetail() {
      console.log(this.deviceInfo)
      this.$router.push({ name: 'DeviceDetail', query: { serialNo: this.deviceInfo.serialNo }})
    }
  }
}
</script>

<style scoped>
.device {
  margin: 20px;
  padding: 10px;
  height: 160px;
  width: 44%;
  display: inline-block;
  box-shadow: #ccc 0px 0px 10px;
  background-color: #fff;
  cursor: pointer;
}

.device_logo {
    height: 100px;
    width: 53px;
    float: left;
}

.device_info {
  margin-top: 8px;
  margin-left: 60px;
}

.device_title {
  font-weight: bold;
  font-size: 16px;
  color: black;
}

.device_serial {
  margin-top: 10px;
  color: gray;
}

.device_build {
  margin-top: 10px;
  color: gray;
}

.device_state {
  margin-top: 40px;
  margin-left: 10px;
  font-weight: bold;
}

.device_left {
  display: inline;
  float: left;
}

.device_right {
  margin-top: 10px;
  margin-right: 20px;
  float: right;
}

</style>
