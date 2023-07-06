<template>
    <div class="header">
        <div class="titles">
            <p style="width: 100px; margin-left: 10px;">作品分类:</p>
            <el-cascader :options="options" :show-all-levels="false" :props="props" collapse-tags-tooltip collapse-tags
                clearable />
        </div>
        <div class="biaoqian">
            <div class="labelBtn">
                <button class="works" :class="[btnmenu == 0 ? 'btns' : '']" @click="btnmenuClick(0)">全部作品</button>
            </div>
           
            <div class="labelBtn">
                <button class="soundWorks" :class="[btnmenu == 2 ? 'btns' : '']" @click="btnmenuClick(2)">有声作品</button>
            </div>

        </div>
        
             <div class="search-box">
            <input type="text" class="search-txt" placeholder="想搜啥？">
            <a class="search-btn">
                <font-awesome-icon :icon="['fas', 'magnifying-glass']" />
            </a>
        </div>
        
       


    </div>
    <ul class="box" v-infinite-scroll="load" infinite-scroll-immediate="false" style="overflow: auto; ">
        <el-scrollbar height="600px">
            <div class="box-item">
                <div v-if="btnmenu == 0">
                    <Story />
                </div>
            </div>
        </el-scrollbar>
    </ul>
</template>
    
<script setup lang='ts'>
import Story from '../components/Story.vue'

import { ref } from 'vue'
//无线滚动条（未使用）
const count = ref(2)
const load = () => {
    console.log(111);
}
//选择器值
const options = [
    {
        value: 'guide',
        label: '完结状态',
        children: [
            {
                value: 'disciplines',
                label: '已完结',

            },
            {
                value: '11',
                label: '未完结'
            }
        ]
    },
    {
        value: '',
        label: '年份',
        children: [
            {
                value: '2023',
                label: '2023'
            },
            {
                value: '2022',
                label: '2022'
            },
            {
                value: '2021',
                label: '2021'
            },
            {
                value: '2020',
                label: '2020'
            },
            {
                value: '2019',
                label: '2019'
            },
        ]
    },
    {
        value: '',
        label: '状态',
        children: [
            {
                value: '',
                label: '本地'
            },
            {
                value: '',
                label: '有线'
            }
        ]
    },
    {
        value: '',
        label: '类型',
        children: [
            {
                value: '',
                label: '玄幻'
            },
            {
                value: '',
                label: '修仙'
            },
            {
                value: '',
                label: '武侠'
            },
            {
                value: '',
                label: '都市'
            },
        ]
    }
]
//设置多选
const props = {
    multiple: true,
}
//添加点击动画
const btnmenu = ref(0)
const btnmenuClick = (id) => {
    btnmenu.value = id
}
</script>
    
<style scoped lang="scss">
.header {
    height: 100px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .titles {
        width: 250px;
        height: 51px;
        background-color: #eeecec;

        border-radius: 5%;
        display: flex;
        align-items: center;
        justify-content: center;

        p {
            font-size: 16px;
        }
    }

    .biaoqian {
        width: 400px;
        // margin-right: 100px;
        
        display: flex;
        justify-content: space-around;
        .labelBtn{
            width: 150px;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        button {
            width: 130px;
            // background-color: rgb(239, 122, 122);
            background-color: #eeecec;
            
            height: 51px;
            border-radius: 5px;
            border: none;
            // margin-left: 10px;
        }
        button:hover{
            width: 150px;
        }
    }


    :deep(.el-input__wrapper) {
        background-color: transparent;
        box-shadow: none;
    }

    :deep(.el-input__wrapper:focus) {
        box-shadow: none;
    }

    :deep(.el-cascader__tags) {
        flex-wrap: nowrap !important;
    }
}

.box {
    margin-top: 2%;
}

.box-item {
    transition: box-shadow 0.3s ease-in-out;
}

.search-box {
    background: #eeecec;
    height: 40px;
    border-radius: 40px;
    padding: 8px;
    
    .search-txt {
        border: none;
        background: none;
        outline: none;
        float: left;
        padding: 0;
        color: rgb(46, 42, 42);
        font-size: 16px;
        line-height: 40px;
        width: 0;
        /* 动画过渡 */
        transition: 0.4s;
    }

    .search-btn {
        color: #e84118;
        float: right;
        width: 40px;
        height: 40px;
        border-radius: 50%;
        background: #eeecec;
        /* 弹性布局 水平垂直居中 */
        display: flex;
        justify-content: center;
        align-items: center;
        cursor: pointer;
        /* 动画过渡 */
        transition: 0.4s;
    }
}
html.dark .search-box{
    background-color: #767676;
    .search-btn{
        background: #767676;
        color: #e84118 !important;
    }
}
/* 鼠标移入文本框变化 */
.search-box:hover .search-txt {
    width: 200px;
    padding: 0 6px;
}

/* 鼠标移入搜索按钮变化 */
.search-box:hover .search-btn {
    background: #fff;
}
</style>