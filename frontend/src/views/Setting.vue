<template>
 


   
<p>当前主题: <strong id="theme-source">系统默认</strong></p>

    <button id="toggle-dark-mode" @click="setNight">切换暗黑模式</button>
    <button id="reset-to-system" @click="setMorning">重置为系统默认</button>
    <!-- 表单 -->
    <el-form>
        <el-form-item>
            <el-avatar></el-avatar>
        </el-form-item>
        <el-form-item>
            <el-input></el-input>
        </el-form-item>

    </el-form>

</template>
    
<script setup lang='ts'>
import { ref,  } from 'vue'
import { useDark, useToggle,  } from '@vueuse/core'

  import { ipcRenderer } from 'electron';
const setNight = () =>{
   ipcRenderer.invoke('dark-mode:toggle')
  console.log('121212');
  
}
const setMorning = () =>{
  ipcRenderer.invoke('dark-mode:system')
}


</script>
    
<style></style>