<template>
  <div>
    <div class="gva-form-box">
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
          <el-switch v-model="formData.enable" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createJob,
  updateJob,
  findJob
} from '@/api/autoJob' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Job',
  mixins: [infoList],
  data() {
    return {
      type: '',
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
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findJob({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.rejob
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
