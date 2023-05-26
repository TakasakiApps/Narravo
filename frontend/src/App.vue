<template>
  <div>
    <div class="tags">
      <van-tabs v-model:active="active" class="Systga"  gutter="50">
        <van-tab title="首页" to="/"  title-class="vanTab" title-style="width: 12.5rem;">
        </van-tab>
        <van-tab title="创造" to="/CreateVue" title-class="vanTab">
        </van-tab>
        <van-tab title="书架" to="/BookShelfVue" title-class="vanTab">
        </van-tab>
        <van-tab title="设置" to="/SysHomeVue" title-class="vanTab">
        </van-tab>
      </van-tabs>
      <router-view />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Tab as vanTab,Tabs as vanTabs, Col as vanCol, Row as vanRow } from "vant"
import {ref} from 'vue';
import  NewToken from './samples/node-api'

const active = ref(0)

</script>

<style scoped>
.Systga{
  display: flex;
  justify-content: center;
}
.van-tab--active{
  width: 200px;
}
</style>