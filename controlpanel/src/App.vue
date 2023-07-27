<template>
  <el-config-provider namespace="ep">
    <BaseHeader />
    <div class="flex main-container">
      <BaseSide />
      <div w="full" py="4" class="back">
        <!-- <Logos my="4" /> -->
        <router-view></router-view>
      </div>
    </div>
  </el-config-provider>
</template>

<style>
#app {
  text-align: center;
  color: var(--ep-text-color-primary);
}

.main-container {
  height: calc(100vh - var(--ep-menu-item-height) - 3px);
}

.back {
  background: linear-gradient(150deg, rgb(254, 245, 230) 0%, rgb(230, 244, 254) 62%, rgb(255, 255, 255) 90%, rgb(255, 255, 255) 100%);
}
</style>

<script setup>

// import Cookies from 'universal-cookie';
// import axios from 'axios';
// import { onMounted } from 'vue'
// import { ElMessage } from 'element-plus';
// import { useRouter } from 'vue-router'
// const router = useRouter()
// const cookies = new Cookies
// onMounted(() => {
//   const token = cookies.get('token')
//   if (token) {
//     axios.post('/api/auth/renew', {
//       token: token
//     }).then(
//       e => {
//         cookies.set('token', e.data.token)
//         router.push({
//           name: 'dashboard'
//         })
//       }
//     ).catch(
//       () => {
//         ElMessage({
//           message: '登录失效',
//           type: 'warning',
//         })
//         cookies.remove('token')
//       }
//     )
//   } else {

//   }
// })
</script>