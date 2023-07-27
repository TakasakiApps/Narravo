<template>
  <div>
    <el-card class="box-card">
      修改密码
      <el-form :model="from" label-position="top">
        <el-form-item label="旧密码">
          <el-input v-model="from.oldpw" class="from-input" />
        </el-form-item>
        <el-form-item label="新密码">
          <el-input v-model="from.newps" class="from-input" type="password" />
        </el-form-item>
        <el-form-item label="确认密码">
          <el-input v-model="from.nowps" class="from-input" type="password" />
        </el-form-item>
      </el-form>
      <el-form-item class="from-button">
        <el-button type="primary" @click="onSubmit">确认</el-button>
        <el-button @click="Cancel">取消</el-button>
      </el-form-item>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router'
const router = useRouter()
const from = reactive({
  oldpw: '',
  newps: '',
  nowps: ''
})
const onSubmit = () => {
  if (from.nowps == from.newps) {
    console.log('发送成功')
  } else {
    ElMessage({
      message: '请确保两次密码相同',
      type: 'warning',
    })
  }
  console.log(from)
}
const Cancel = () => {
  router.go(-1)
  console.log('返回')
}
</script>

<style lang="scss" scoped>
.box-card {
  // display: flex;
  // justify-content: center;
  // /* 水平居中 */
  // align-items: center;
  /* 垂直居中 */
  position: absolute;
  width: 550px;
  height: 25rem;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.from-button {
  position: absolute;
  bottom: 10px;
  left: 50%;
  display: block;
  transform: translate(-50%, -50%);
}

.from-input {
  width: 500px;
}
</style>