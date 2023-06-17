<template>
  <div>
    <div class="Sys-tag">
      账户设置
    </div>
    <div class="Sys-main">
      <van-button type="default" size="small" v-if="true" @click="userlogin">登录</van-button>
      <van-button type="default" size="small" v-else >登出</van-button>
      <van-overlay :show="show" @click="show = false">
        <template #default>
          <div class="card">
            <div class="block">
              <input type="text" v-model="user.name">
            </div>
          </div>
        </template>
      </van-overlay>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { showDialog,Dialog } from 'vant';
const show = ref(false);
const user = ref({
  name:'',
  password:''
})
function userlogin(){
  show.value=true
}
function dialogClose(){
  show.value=false
  return true
}
</script>

<style lang="scss" scoped>
.card{
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}
.block{
    width: 120px;
    height: 120px;
    border: 1px solid white;
}
</style>