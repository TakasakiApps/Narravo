<template>
    <div style="overflow-y:hidden; margin: auto; ">
        <div v-masonry transition-duration="0.3s" item-selector=".item-card" fit-width="true" gutter="20" class="body">
            <div v-masonry-tile class="item-card" v-for="item in 12">
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
                            <el-button text class="button">作者</el-button>
                        </div>
                    </div>
                </el-card>
            </div>
        </div>
    </div>
    <v-contextmenu ref="contextmenu" >
    <v-contextmenu-item><span @click="getInfo('/info')">书籍详情</span></v-contextmenu-item>
    <v-contextmenu-item><span>加入书架</span></v-contextmenu-item>
    
  </v-contextmenu>
</template>
    
<script setup lang='ts'>
import { Service } from '@element-plus/icons-vue'
import router from '../router'

const getInfo = (path) => {
router.push({path})
}
</script>
    
<style scoped>
.time {
    font-size: 12px;
    color: #888686;
}

.bottom {
    margin-top: 13px;
    line-height: 12px;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.button {
    padding: 0;
    min-height: auto;
}

.image {
    width: 100%;
    height: 202px;
    display: block;
}

.body {
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 20px auto;
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

.time {
    margin-top: -20px;
}
</style>