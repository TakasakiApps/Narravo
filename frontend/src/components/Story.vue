<template>
    <div style="overflow-y:hidden;  ">
        <div v-masonry transition-duration="0.3s" item-selector=".item-card" fit-width="true" gutter="20" class="body">
            <div v-masonry-tile class="item-card"  >
                <el-card :body-style="{ padding: '0px' }" v-contextmenu:contextmenu shadow="hover">
                    <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                        class="image" />
                    <div style="padding: 5px">
                        <div class="title">
                            <h1>标题</h1>
                            <el-icon size="25">
                                <Service />
                            </el-icon>
                        </div>

                        <div class="bottom">
                            <time class="time">最新章节</time>
                            <p text class="author">作者</p>
                        </div>
                    </div>
                </el-card>
            </div>
        </div>
    </div>
    <v-contextmenu ref="contextmenu">
        <v-contextmenu-item><span @click="getInfo('/info')">书籍详情</span></v-contextmenu-item>
        <v-contextmenu-item><span>加入书架</span></v-contextmenu-item>
    </v-contextmenu>
</template>
    
<script setup lang='ts'>
import { Service } from '@element-plus/icons-vue'
import { inject } from 'vue'
import router from '../router'
const globalVars = inject('globalVars')
const getInfo = (path: any) => {
    
    globalVars.isInfo = true;
    console.log('值：',globalVars.isInfo);
    router.push(path)
}

</script>
    
<style scoped>
.time {
    font-size: 13px;
    color: #888686;
    font-family: '更纱黑体 SC';
    margin-top: -40px;
}

.bottom {
    margin-top: 13px;
    line-height: 12px;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.author {
    padding: 0;
    min-height: auto;
    font-family: 'Roboto';
    font-size: 12px;
}

.image {
    width: 100%;
    height: 202px;
    display: block;
}

.body {
    display: flex;
    
    align-items: center;
    margin: 20px 20px;
}

.el-card {
    width: 176px;
    border-radius: 10px;
    margin-left: 10px;
    margin-right: 10px;

}

div.item-card {
    margin: 10px auto;
}

.title {

    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 34px;

    .el-icon {
        float: right;
    }

}

h1 {
    font-size: 20px;
    font-family: '更纱黑体 SC';

    color: #333333;
}
</style>