<template>
  <div>
    <div class="gva-form-box">
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
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {
  createDevice,
  updateDevice,
  findDevice
} from '@/api/autoDevice' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Device',
  mixins: [infoList],
  data() {
    return {
      type: '',
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
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findDevice({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.redevice
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
  },
  methods: {
    async save() {
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
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>
