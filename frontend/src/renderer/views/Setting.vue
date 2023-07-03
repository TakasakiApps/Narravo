<template>
   

    
    <span @click.stop="toggleDark()">暗黑模式</span>
    <el-switch size="small" v-model="isDark" />
   

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
import { ref, onBeforeMount } from 'vue'
import { useDark, useToggle, useColorMode } from '@vueuse/core'
import { nativeTheme } from  'electron'

const isDark = useDark()
const toggleDark = () => {
    useToggle(isDark)
    
}
const theme = ref(false)
// const {ipcRenderer} = require('electron')
const system = ref('白天')
const mode = useColorMode({
    // 如果模式为auto也需要回显回auto
    emitAuto: true,
    // 默认模式先默认auto,后续通过Electorn拿到当前App主题
    initialValue: "auto",
});

</script>
    
<style></style>