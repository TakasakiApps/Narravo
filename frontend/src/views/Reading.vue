<template>
        <div class="container">
                <h1>章节名</h1>
                <div class="sidenav">
                        <el-aside :class="[right==false ? 'open':'' ]">
                                <div class="body">
                                        <h2>目录</h2>
                                        <el-scrollbar max-height="600px" always="true">
                                                <div class="chapter" v-for="item in optison">
                                                        <el-text truncated>{{ item.value }}</el-text>
                                                </div>
                                        </el-scrollbar>
                                </div>

                                <div class="bottom">
                                        <p style="font-family: '更纱黑体 SC'; font-size: 14px;">字体大小</p>
                                        <el-input-number style="width: 78px; border-radius: 25% ;  margin-left: 12px; display: flex; justify-content: center;"></el-input-number>
                                </div>
                        </el-aside>
                        <div class="right" @click="checkLight">

                                <svg t="1688729279262" class="icon" viewBox="0 0 1024 1024" version="1.1"
                                        xmlns="http://www.w3.org/2000/svg" p-id="2586" width="20" height="20" v-if="!right">
                                        <path d="M610.88 512L192 93.12 285.12 0l512 512-512 512L192 930.88z" fill="#262626"
                                                p-id="2587"></path>
                                </svg>
                                <svg t="1688729212223" class="icon" viewBox="0 0 1024 1024" version="1.1"
                                        xmlns="http://www.w3.org/2000/svg" p-id="1610" width="20" height="20" v-if="right">
                                        <path d="M378.24 512l418.88 418.88L704 1024 192 512l512-512 93.12 93.12z" fill="#262626"
                                                p-id="1611"></path>
                                </svg>

                        </div>
                </div>
        </div>
</template>
    
<script setup lang='ts'>
import { reactive, ref } from 'vue'
const right = ref(true)//true为右箭头，false为左箭头
const checkLight = () => {
right.value = !right.value
}
//章节名
const optison = reactive([
        { value: '第一章：绯红' },
        { value: '第二章：情况' },
        { value: '第三章：梅丽莎' },
        { value: '第四章：占卜' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
        { value: '第五章：“职业”和住房都是严肃的事情' },
])
</script>
    
<style scoped lang="scss">
.container {
        height: 100%;
        min-height: auto;
}

.sidenav {
        position: fixed;
        z-index: 1;
        transition: 0.5s;
        top: 0;
        right: 0;
        height: 100%;
}

.el-aside {
        height: 100%;
        float: right;
        // width: 286px;
        width: 0;
        background-color: #F2F2F2;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        transition: 0.5s;
        overflow: hidden;

        .body {
                display: flex;
                flex-direction: column;
                justify-content: center;
                align-items: center;

        }


        h2 {
                font-family: '更纱黑体 UI SC';
                font-size: 24px;
        }

        .el-text {
                font-family: '更纱黑体 UI SC';
                font-size: 16px;
        }

        .chapter {
                width: 232px;
                height: 38px;
                line-height:38px;
                border-radius: 5px;

        }
        .chapter:hover{
                background-color: #D7D7D7;
        }
        .bottom {
                display: flex;
                align-items: center;
                margin: 20px 10px;
                // 修改input Number样式
                
                :deep(.el-input-number .el-input__wrapper){
                        background-color: #D7D7D7;
                        border-radius: 100px ;
                        
                        padding-left: 10px;
                        padding-right: 10px;
                }
                :deep(.el-input-number__decrease),:deep(.el-input-number__increase){
                        width: 15px;
                        height: 30px;
                        background: none;
                        border: none;
                        border-radius: 25% ;
                        padding:0 5px
                }
               
                

        }


}

.right {
        background-color: #F2F2F2;
        width: 20px;
        height: 50px;
        border: none;
        position: relative;
        z-index: 1;
        // right: 30px;
        margin-top: 340px;
        display: flex;
        align-items: center;
        border-radius: 15px 0 0 15px;
}
.right:hover{
        background-color: #b7b7b7;
}
.open{
        width: 286px;
}
</style>