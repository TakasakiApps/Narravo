<template>
  <el-menu default-active="2" class="el-menu-vertical-demo" :collapse="isCollapse" @open="handleOpen"
    @close="handleClose">
    <el-menu-item index="1">
      <el-icon @click="User">
        <icoUser />
      </el-icon>
      <template #title>作者管理</template>
    </el-menu-item>
    <el-menu-item index="2" disabled>
      <el-icon>
        <document />
      </el-icon>
      <template #title>在线管理</template>
    </el-menu-item>

    <!-- <el-menu-item index="4">
      <el-icon><setting /></el-icon> -->
    <el-sub-menu index="3">
      <template #title>
        <el-icon>
          <setting />
        </el-icon>
        <span>Navigator One</span>
      </template>
      <el-menu-item-group title="Group Two">
        <el-menu-item @click="clearCookies">账号登出</el-menu-item>
        <el-menu-item @click="changewd">账号管理</el-menu-item>
      </el-menu-item-group>
    </el-sub-menu>
    <!-- <el-menu-item  @click="toggleDark()">
      <button
        class="border-none w-full bg-transparent cursor-pointer"
        style="height: var(--ep-menu-item-height)"
      >
        <i inline-flex i="dark:ep-moon ep-sunny" />
      </button>
    </el-menu-item> -->
    <!-- <el-menu-item @click="toggleDark()" class="toggleDark">
      <el-icon>
        <document />
      </el-icon>
      <template #title>Navigator Four</template>
    </el-menu-item> -->
  </el-menu>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import {
  Location,
  Document,
  Menu as IconMenu,
  Setting,
  User as icoUser
} from "@element-plus/icons-vue";
import VueCookies from 'vue-cookies'
const cookie: any = VueCookies
import { useRouter } from 'vue-router'

const router = useRouter()

const isCollapse = ref(true);
const handleOpen = (key: string, keyPath: string[]) => {
  // console.log(key, keyPath);
};
const handleClose = (key: string, keyPath: string[]) => {
  // console.log(key, keyPath);
};
const clearCookies = () => {
  cookie.remove('token')
  router.push({
    path: '/'
  })
}
const User = () => {
  router.push({
    path: '/AddDesc'
  })
}
const changewd = () => {
  router.push({
    path: '/Editpas'
  })
}
</script>
<style></style>