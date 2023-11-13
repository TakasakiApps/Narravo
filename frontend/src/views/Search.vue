<template>
    <!-- <div class="header"> -->
    <!-- <div class="saixvan">
            <div class="titles">
                <p style="  ">分类:</p>
                <el-select v-model="value">
                    <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </div>
            <div class="biaoqian">
                <p style=" margin-left: 10px;">筛选:</p>
                <el-select v-model="values">
                    <el-option v-for="item in type" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </div>
        </div> -->


    <!-- <div class="search-box">
            <input type="text" class="search-txt" placeholder="想搜啥？" @blur="rmFouce"
                :class="[searchInput == true ? 'searchInput' : '']">
            <a class="search-btn" @click="toggleInput">
                <svg fill="currentColor" t="1688636529593" class="icon" viewBox="0 0 1024 1024" version="1.1"
                    xmlns="http://www.w3.org/2000/svg" p-id="38228" width="32" height="32">
                    <path
                        d="M192 480a256 256 0 1 1 512 0 256 256 0 0 1-512 0m631.776 362.496l-143.2-143.168A318.464 318.464 0 0 0 768 480c0-176.736-143.264-320-320-320S128 303.264 128 480s143.264 320 320 320a318.016 318.016 0 0 0 184.16-58.592l146.336 146.368c12.512 12.48 32.768 12.48 45.28 0 12.48-12.512 12.48-32.768 0-45.28"
                        p-id="38229"></path>
                </svg>
            </a>
        </div> -->
    <!-- </div> -->
    <ul class="box" v-infinite-scroll="load" infinite-scroll-immediate="false" style="overflow: auto; ">
        <el-scrollbar height="600px">
            <div class="box-item">
                <Story v-for="(item, index) in bookliem.list " :key="index" :item="item" />
            </div>
        </el-scrollbar>
    </ul>
</template>
    
<script setup lang='ts'>
import Story from '../components/Story.vue'
import { client } from '../http/client'
import { ref, onMounted } from 'vue'
import { ipcRenderer } from 'electron'
onMounted(() => {
    getAllBook(1)

})
let bookliem = ref([])
function getAllBook(index: number) {
    client.get('/api/v1/novel/getAllInfo',
        {
            params: {
                token: localStorage.token,
                page: index,
                count: 10
            }
        }).then(res => {
            console.log(res.request)
            bookliem.value = res.data
        }).catch(err => {
            console.log(err);
            if (err.response.status == '401') {
                setTimeout(() => {
                    ElNotification.closeAll()
                    ElNotification.error({ title: '操作失败!', message: '请检查是否登录' })
                }, 100)
            } else {
                ElNotification.error({ title: '操作失败!', message: '请检查网络' })
            }
        })

}
function a() {
    console.log('11')
}
// 无线滚动条（未使用）
const count = ref(1)
const load = () => {
    getAllBook(count)
    count.value++
}
//选择器值
const options = [
    {
        value: 'Option1',
        label: '玄幻',
    },
    {
        value: 'Option2',
        label: '修仙',
    },

]
const value = ref('')
const values = ref('')
const type = [
    {
        values: '有声',
        label: '有声'
    },
    {
        values: '全部',
        label: '全部'
    }
]


//点击搜索按钮，文本框变化
const searchInput = ref(false)
const toggleInput = () => {
    searchInput.value = !searchInput.value


}
const rmFouce = () => {
    searchInput.value = !searchInput.value
}

</script>
    
<style scoped lang="scss">
.header {
    height: 100px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .saixvan {
        width: 350px;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    .titles {
        width: 150px;
        height: 51px;
        background-color: #eeecec;

        border-radius: 5%;
        display: flex;
        align-items: center;
        justify-content: center;

        p {
            font-size: 18px;
            font-family: '更纱黑体 UI SC';
        }

        .el-select {
            width: 76px;
        }
    }

    .biaoqian {
        width: 150px;
        height: 51px;
        background-color: #eeecec;

        border-radius: 5%;
        display: flex;
        align-items: center;
        justify-content: center;

        p {
            font-size: 18px;
            font-family: '更纱黑体 UI SC';
        }

        .el-select {
            width: 76px;
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
    position: relative;
    right: 1px;
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

html.dark .search-box {
    background-color: #767676;

    .search-btn {
        background: #767676;
        color: #e84118 !important;
    }
}

/* 鼠标移入文本框变化 */
.searchInput {
    width: 200px !important;
    padding: 0 6px;
}

/* 鼠标移入搜索按钮变化 */
.search-box:hover .search-btn {
    background: #fff;
}
</style>