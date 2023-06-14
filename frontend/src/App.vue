<template>
  <div>
    <div class="tag">
      <h5>Narravo</h5>
      <div class="button">
        <electon-botton />
      </div>
    </div>
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
      <div class="view">
        <router-view />
      </div>
    </div>

  </div>
</template>

<script lang="ts" setup>

import {ref,onMounted } from 'vue';
import  NewToken from './samples/node-api'
onMounted(() => {
  console.log(`取得token.`)
})


const active = ref(0)

</script>


<style scoped>
.Systga{
  display: flex;
  justify-content: center;
}
.tag{
  display: flex;
  /* align-items: center; 垂直居中 */
  justify-content: space-between; /* 左右对齐，中间留有空白间隔 */
  height: 36px;
  padding: 0 20px; 
  padding-right: 0px;

}

.van-tab--active{
  width: 100%;
}
h5{
  margin: .125rem .1875rem ;
  padding: 0;
  opacity: 0.85;
}
.tag{
  -webkit-app-region: drag;
  -webkit-user-select: none;
}
button {
  /* 在设置禁用元素的拖动和选择功能的时候:hover不会触发*/
  /* -webkit-app-region: no-drag; */
}
.view{
  overflow:hidden;
  text-align: center;
}
</style>