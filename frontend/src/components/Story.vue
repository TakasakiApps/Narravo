<template>
        <div style="overflow-y:auto; margin: auto; ">
        <el-row >
            <el-col :span="6" v-for="item in 6" >
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
            </el-col>
        </el-row>
    </div>
</template>
    
<script setup lang='ts'>
    
</script>
    
<style scoped>
    .time {
    font-size: 12px;
    color: #999;
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

.el-card {
    width: 176px;
    border-radius: 10px;
}

.title {
    /* border: 1px solid black; */
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