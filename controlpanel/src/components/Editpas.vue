<template>
  <div>
    <el-card style="width: 55vw; margin: auto;">
      <el-form label-width="100px" :model="formLabelAlign" style="max-width: 500px; margin: auto;">
        <el-form-item label="旧密码" class="item">
          <el-input v-model="formLabelAlign.oldpasswd" />
        </el-form-item>
        <el-form-item label="新密码" class="item">
          <el-input v-model="formLabelAlign.newpasswd" type="password" />
        </el-form-item>
        <el-form-item label="确认密码" class="item">
          <el-input v-model="formLabelAlign.affpasswd" type="password" />
        </el-form-item>
        <el-form-item class="item">
          <el-button type="primary" @click="changewd">
            确认修改
          </el-button>
          <el-button>
            取消
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'
const formLabelAlign = reactive({
  oldpasswd: '',
  newpasswd: '',
  affpasswd: '',
})
const changewd = () => {
  if (formLabelAlign.newpasswd !== '' && formLabelAlign.newpasswd == formLabelAlign.affpasswd) {
    axios.post('/api/v1/auth/reset/password', {
      username: 'admin',
      password: formLabelAlign.oldpasswd,
      newPassword: formLabelAlign.newpasswd,
    }).then(() => {
      ElMessage({
        message: '重设密码成功,请从小登录',
        type: 'success',
      })
    }).catch(() => {
      ElMessage({
        message: '重设密码失败',
        type: 'error'
      })
    })
  } else {
    ElMessage({
      message: '密码不能为空,两次密码需相同',
      type: 'warning',
    })
  }

}
</script>

<style scoped>
.item {
  margin-bottom: 35px;
}
</style>