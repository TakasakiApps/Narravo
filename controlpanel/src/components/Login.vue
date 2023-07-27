<template>
  <div>
    <el-card class="box-card">
      <div class="tag">
        管理员登录
      </div>
      <div class="login-main">
        <el-row :gutter="20">
          <div class="login-liem inline-flex">
            <span class="ml-3 w-35 text-gray-600 inline-flex items-center">账号：</span>
            <el-input v-model="user.act" placeholder="请输入账号" />
          </div>
        </el-row>
        <div class="login-liem">
          <el-row :gutter="20">
            <div class="login-liem inline-flex">
              <span class="ml-3 w-35 text-gray-600 inline-flex items-center">密码：</span>
              <el-input v-model="user.passwd" placeholder="请输入密码" type="password" />
            </div>
          </el-row>
        </div>
        <el-button class="login-button" @click="login">登录</el-button>
      </div>
    </el-card>
  </div>
</template>

<style scoped lang="scss">
.tag {
  margin-bottom: 1.5rem;
}

.tag::after {
  content: "";
  display: flex;
  width: 50px;
  border-bottom: 3px solid rgb(157, 146, 95);
  position: absolute;
  left: 47%;
  top: 18px;
  transform: translateX(-25%);
}

.box-card {
  /* margin-left: 50%; */
  width: 480px;
  height: 300px;
  /* text-align: center; */
  /* justify-self: center; */
  margin: auto;
  margin-top: 10%;
  background-color: rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(10px)
}

.login-main {
  // border: 1px solid rgba(90, 86, 86, 0.384);
  border-radius: 5px;
  margin-left: 30px;
}

.login-liem {
  margin-top: 1rem;
}

.login-button {
  margin-top: 40px;

}
</style>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus';
// import { client as axios } from '../http/client';
import VueCookies from 'vue-cookies'
const cookie = VueCookies
// import axios from 'axios';
import { client as axios } from '../http/client'
const router = useRouter()
const ps = ''
const user = ref({
  act: '',
  passwd: ''
})


const login = () => {
  if (user.value.act == 'admin') {
    axios.post('/api/v1/auth/login', {
      username: user.value.act,
      password: user.value.passwd
    }).then(
      (e) => {
        cookie.set('token', e.data.token)
        router.push({
          name: 'dashboard'
        })
      }
    ).catch(
      () => {
        ElMessage({
          message: '登录失败',
          type: 'warning',
        })
      }
    )
  } else {
    ElMessage({
      message: '只允许管理员登录',
      type: 'warning',
    })
    // console.log(user.value.act)
  }
  // router.push({
  //   name: 'dashboard'
  // })
}
onMounted(() => {
  const token = cookie.get('token')
  if (token) {
    axios.post('/api/v1/auth/renew', {
      data: token
    }).then(
      e => {
        cookie.set('token', e.data.token)
        router.push({
          name: 'dashboard'
        })
      }
    ).catch(
      () => {
        ElMessage({
          message: '登录失效',
          type: 'warning',
        })
        // cookie.remove('token')
      }
    )
  } else {

  }
})
</script>