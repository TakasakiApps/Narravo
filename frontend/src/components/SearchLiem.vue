<template>
  <div>
    <book-liem 
    v-for=" itme in nowdata "
    :key="itme.id"
    :name = "itme.name"
    :act = "itme.act"
    :Brief = "itme.Brief"
    />
  </div>
</template>

<script setup lang="ts">
import { onMounted,watch } from 'vue'
import { useRoute } from 'vue-router'
// const ins = getCurrentInstance() as ComponentInternalInstance | null
const  Router :any = useRoute()
let nowdata = JSON.parse(Router.query.data)
// 如果响应路由并不会销毁再调用，而是重复使用，使用生命周期钩子不会被调用
watch(
  ()=> Router.query.data,
  ()=>{
    // 强制更新dom
    const newdata = JSON.parse(Router.query.data)
    nowdata = newdata
  }
)

</script>

<style lang="scss" scoped>

</style>