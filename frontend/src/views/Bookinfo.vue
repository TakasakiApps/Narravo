<template>
    <div class="box">
        <div v-if="book.local" class="local-img">
            <img :src="bookcover" alt="" v-if="props.local == 0">
            <div class="local-title" v-else>
                <h3>{{ removeExtension(props.name) }}</h3>
            </div>
        </div>
        <img :src="img1" alt="" v-else>
        <div class="info">
            <h1>{{ removeExtension(props.name) }}</h1>
            <h2>{{ props.author !== '0' ? props.author : '群星' }}</h2>
            <!-- <h2 style="color: #7F7F7F;">{{ props.tpe }} {{ (props.size / 10000).toFixed(2) + '万字' }}</h2> -->
            <div style="display: flex;">
                <h3 style="margin-right: 32px;">最新章节：{{ props.local == 0 ? chapters[chapters.length - 1].name :
                    chapters[chapters.length - 1] }}</h3>
                <h3>更新时间:{{ formatDate(props.mtime) }}</h3>
            </div>
            <h3>书籍简介</h3>
            <p style="width: 450px; ">{{ props.resume }}</p>
            <div class="btn">
                <button style="width: 100px; height: 48px; margin-right: 20px;" @click="goRead('/read')">阅读</button>
                <button style="width: 142px; height: 48px; margin-right: 168px;" v-if="props.local == 0"
                    @click="addbookliem(props)">加入书架</button>
                <button class="no-hover" style="width: 142px; height: 48px; margin-right: 168px; background-color: #b8b8b8;"
                    v-else>已在书架</button>
                <!-- <button style="width: 120px; height: 48px;">去制作</button> -->
            </div>
        </div>
    </div>
    <router-view></router-view>
</template>
    
<script setup lang='ts'>
// import img1 from '../assets/诡秘.png'
import router from '../router'
import { ref, onMounted } from 'vue'
import { ipcRenderer } from 'electron'
import { client } from '../http/client'
onMounted(() => {
    console.log(props)
    // 网络小说获取封面
    if (props.local == 0) {
        client.get('/api/v1/assets/fetch/image', {
            params: {
                token: localStorage.token,
                type: "cover",
                id: props.coverId
            }
        }).then(
            res => {
                bookcover.value = res.request.responseURL
            }
        ).catch(
            err => {
                console.log(err)
            }
        )
        // 获取最新章节
        client.get(`/api/v1/novel/get/` + props.noverid + `/catalog`, {
            params: {
                token: localStorage.token
            }
        }).then(
            res => {
                chapters.value = res.data.chapters

            }
        ).catch(
            err => {
                console.log(err)
            }
        )
    } else if (props.local == '1') {
        console.log(props.name)
        ipcRenderer.send('testliem', props.name)
    }

})
ipcRenderer.on('sortedFiles', (e, data) => {
    // console.log(data)
    chapters.value = data
})
const props = defineProps({
    name: String,
    local: String,
    coverId: {
        type: String,
        default: "null"
    },
    author: {
        type: String,
        default: "群星"
    },
    mtime: {
        type: String,
        default: "2023-10-20"
    },
    resume: {
        type: String,
        default: "无"
    },
    size: {
        type: Number,
        default: 10000
    },
    tpe: {
        type: String,
        default: '日常'
    },
    noverid: {
        type: String,
        default: "null"
    }
})
let book = ref({
    local: 1
})
let bookcover = ref(null)
//章节数组
let chapters = ref([
    {
        name: '暂无'
    }
])
function addbookliem(e) {
    console.log(e)
    //定义存放数据
    let bookdata = {}
    bookdata = {
        name: e.name,
        id: e.noverid,
        coverId: e.coverId,
        author: e.author,
        jpg: bookcover.value,
        local: 0
    }
    console.log(bookdata)
    ipcRenderer.send('addbookliem', bookdata)
}
ipcRenderer.on('okadd', (e, data) => {
    setTimeout(() => {
        ElNotification.closeAll()
        ElNotification.success({ title: '添加成功!', message: '已添加至书架' })
    }, 100)
})
ipcRenderer.on('noaddliem', () => {
    setTimeout(() => {
        ElNotification.closeAll()
        console.log('2122')
        ElNotification.error({ title: '添加失败!', message: '出现了错误' })
    }, 100)
})
//去除后缀
function removeExtension(filename) {
    return filename.replace('.txt', '');
}
// 规则时间
function formatDate(dateString) {
    const originalDate = new Date(dateString);
    const year = originalDate.getFullYear();
    const month = originalDate.getMonth() + 1;
    const day = originalDate.getDate();

    const formattedDate = `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day}`;

    return formattedDate;
}
//获取数据并传输
const goRead = (path: any) => {
    // console.log(props.noverid)
    let data = props
    console.log(props, 'data')
    if (props.local == '0') {
        client.get(`/api/v1/novel/get/` + props.noverid + `/catalog`, {
            params: {
                token: localStorage.token,
            }
        }).then(
            res => {
                console.log(data.local, '121221121221')
                router.push({
                    name: 'reading',
                    params: {
                        name: data.name,
                        // local: data.local,
                        noverid: data.noverid
                    }
                })
            }
        ).catch(
            err => {
                console.log(err)
            }
        )

        console.log(props.local)
        // ipcRenderer.send('Upbookdata', props.name);
        // ipcRenderer.on('Getpassliem', (event, message: Array) => {
        //     console.log('书籍列表:', message);
        // });
        // 
        console.log(path, props);
    } else if (props.local == '1') {
        router.push({
            name: 'reading',
            params: {
                name: data.name,
                local: data.local,
                noverid: data.noverid
            }
        })
    }
}
</script>
    
<style scoped lang="scss">
.box {
    display: flex;
    align-items: center;
    justify-content: center;

    // margin: 170px auto;
    // border: 1px solid salmon;
    .local-img {
        width: 241px;
        height: 376px;
        margin-right: 105px;
        background-color: #EC808D;

        .local-title {
            background-color: wheat;
            width: 200px;
            height: 10vh;

            h3 {
                position: relative;
                font-size: x-large;
                left: 10px;
            }
        }

    }

    img {
        width: 241px;
        height: 376px;
        margin-right: 105px;
    }

    .info {
        h1 {
            font-family: 'Arial';
            font-size: 54px;
            color: #333333;
            margin-top: 0;
            margin-bottom: 24px;
        }

        h2 {
            font-family: 'Roboto';
            color: #333333;
            font-size: 20px;
            margin-top: 0;
            margin-bottom: 24px;
        }

        h3 {
            font-family: 'Arial';
            font-size: 20px;
            color: #7F7F7F;
            margin-top: 0;
            margin-bottom: 24px;
        }

        p {
            font-family: 'Roboto';
            font-size: 18px;
            color: #7F7F7F;
        }

    }

    .btn {
        button {
            background-color: #EC808D;
            border: none;
            border-radius: 5px;
            color: #F2F2F2;
            font-family: '更纱黑体 SC';
            font-size: 20px;
        }
    }
}

button {}
</style>