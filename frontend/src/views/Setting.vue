<template>
    <!-- <div>主题颜色: <el-color-picker v-model="themeColor" @change="changeTheme"/></div> -->



    <!-- 表单 -->
    <!-- <el-form>
        <el-form-item>
            <el-avatar></el-avatar>
        </el-form-item>
        <el-form-item>
            <el-input></el-input>
        </el-form-item>

    </el-form> -->
    <div class="switch" @click="btnSwitch" :class="[isDark == true ? 'dark' : '']">
        <span class="sun"><font-awesome-icon :icon="['fas', 'sun']" /></span>
        <span class="moon"><font-awesome-icon :icon="['fas', 'moon']" /></span>

    </div>
</template>
    
<script setup lang='ts'>
import { ref, onBeforeMount } from 'vue'
import { ipcRenderer } from 'electron'
import { useColorMode } from "@vueuse/core";
// import { useElementPlusTheme } from "../hooks/useElementPlusTheme";

const mode = useColorMode({
    // 如果模式为auto也需要回显回auto
    emitAuto: true,
    // 默认模式先默认auto,后续通过Electorn拿到当前App主题
    initialValue: "auto",
});


//修改颜色主题
// const themeColor = ref("#cc312c");
// const { changeTheme } = useElementPlusTheme(themeColor.value);
//切换
const isDark = ref(false)

const btnSwitch = () => {
    isDark.value = !isDark.value
    if (isDark.value == true) {
        ipcRenderer.invoke("dark")
        console.log(isDark.value);

    } else {
        ipcRenderer.invoke('light')
        console.log('light');

    }


}
</script>
    
<style scoped>
.switch {
    border: 1px solid #c4c7ce;
    width: 60px;
    height: 30px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 6px;
    border-radius: 15px;
    cursor: pointer;
    position: fixed;
    top: 15px;
    right: 15px;
}

/* 滑块 */
.switch::before {
    content: '';
    width: 24px;
    height: 24px;
    border-radius: 50%;
    background-color:rgb(239, 122, 122); 
    position: absolute;

    left: 45px;
    transition: 0.2s ease-out;
    animation: switchDark 0.2s;

}

.switch.dark::before {
    left: 3px;

    animation: switchLight 0.2s;
}

/* 图标 */
.sun{
    font-size: 18px;
    color: rgb(198, 198, 198);
}
.moon {
    font-size: 18px;
    color: #151414;
}

@keyframes switchDark {
    from {
        left: 3px;
    }

    to {
        left: 45px;
    }
}

@keyframes switchLight {
    from {
        left: 45px;
    }

    to {
        left: 3px;
    }
}</style>