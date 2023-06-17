<template>
  <div class="group">
    <div class="button" @click="main">
      <img src="../assets/24gl-minimization.svg">
    </div>
    <div  class="button" @click="max" >
      <img src="../assets/24gl-minimize2.svg" v-if="win.max">
      <img src="../assets/24gl-maxmize.svg" v-else>
    </div>
    <div  class="button cross" @click="cross">
      <img src="../assets/24gl-cross.svg">
    </div>
    <van-config-provider :theme ="win.theme"></van-config-provider>
  </div>
</template>
<script setup lang="ts">
import { ipcRenderer } from 'electron'
import { reactive,ref } from 'vue';
import mitt from '../samples/EventBus'

const win = reactive({
  // 用于改变全局样式以及监视最大最小化按钮
   max:false,
   theme :"dark"
})
function cross (){
  ipcRenderer.send('window-close')
}
function main(){
  ipcRenderer.send('window-main')
}
function max(){
  ipcRenderer.send('window-max')
}
ipcRenderer.on('mainWin-max', (_, status) => {
  win.max=status;
  console.log('1')
})
const changeType = () =>{
  if(win.theme === 'dark'){
    win.theme = 'light'
    // console.log('lig')
  }else{
    win.theme = 'dark'
    // console.log('dark')
  }
}

mitt.on('change',changeType)
</script>

<style scoped>
.group{
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  -webkit-app-region: no-drag;
}
.button{
  display: inline-flex;
  width: 40px;
  height: 100%;
  justify-content: center;
  align-items: center;
  
}
.button-mouseover {
  background-color: #d0d0d0;
}
.button:hover{
  background-color: #d0d0d0;
}
.cross:hover{
  background-color: #e71123;
}
</style>