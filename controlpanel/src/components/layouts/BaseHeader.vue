<script lang="ts" setup>
import { toggleDark } from "~/composables";
import { useRouter } from 'vue-router'
const router = useRouter()
function Allrouter() {
  router.push({
    name: 'dashboard'
  })
}
function AddBook() {
  router.push({
    name: 'AddBook'
  })
}

</script>

<template>
  <el-menu class="el-menu-demo" mode="horizontal">
    <el-menu-item index="1">novel后台管理</el-menu-item>
    <el-sub-menu index="2">
      <template #title>书籍管理</template>
      <el-menu-item @click="Allrouter">查看全部书籍</el-menu-item>
      <el-menu-item @click="AddBook">添加书籍</el-menu-item>
      <el-menu-item>管理书籍</el-menu-item>
    </el-sub-menu>
    <el-menu-item index="3" disabled>用户管理</el-menu-item>
    <!-- <el-menu-item index="4">Orders</el-menu-item> -->
    <el-menu-item h="full" @click="toggleDark()" class="toggleDark">
      <button class="border-none w-full bg-transparent cursor-pointer" style="height: var(--ep-menu-item-height)">
        <i inline-flex i="dark:ep-moon ep-sunny" />
      </button>
    </el-menu-item>
  </el-menu>
</template>



<style lang="scss" scoped>
.toggleDark {
  position: absolute;
  right: 0px;
}
</style>



