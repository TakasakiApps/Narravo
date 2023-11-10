<template>
  <div class="header">
    <!-- <div class="titles">
      <p style="width: 150px; margin-left: 10px;">筛选:</p>
      <el-cascader v-model="cascader" :options="options" :show-all-levels="false" :props="props" collapse-tags-tooltip
        collapse-tags clearable />
    </div> -->
    <el-button round type="danger" style="position: relative; left: 80vw;" size="large" @click="date()">导入</el-button>
  </div>
  <div style="height:640px">
    <el-scrollbar height="600px">
      <div v-masonry transition-duration="0.3s" item-selector=".item-card" fit-width="true" gutter="20" class="body">
        <div v-masonry-tile class="item-card">
          <el-card :body-style="{ padding: '0px' }" class="book-card" v-contextmenu:contextmenu shadow="hover"
            v-for="(liem, index) in dabk" :id="index" @click="openbook(liem)" @contextmenu="handleContextMenu(liem)">
            <img
              :src="liem.jpg ? liem.jpg : 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png'"
              class="image" />
            <div style="padding: 5px">
              <div class="title">
                <h1 class="title-txt">{{ removeExtension(liem.name) }}</h1>
                <el-icon size="25">
                  <svg t="1698674595063" class="icon" viewBox="0 0 1024 1024" version="1.1"
                    xmlns="http://www.w3.org/2000/svg" p-id="4957" width="32" height="32" v-if="!liem.local">
                    <path
                      d="M909.989 343.281c-21.76-51.446-52.904-97.643-92.568-137.307-39.663-39.664-85.86-70.808-137.306-92.568-53.28-22.536-109.858-33.962-168.164-33.962S397.067 90.87 343.787 113.406c-51.446 21.76-97.643 52.904-137.307 92.568-39.664 39.664-70.808 85.86-92.568 137.307-22.536 53.28-33.962 109.858-33.962 168.164s11.426 114.884 33.962 168.164c21.76 51.445 52.904 97.643 92.568 137.306 39.664 39.664 85.86 70.809 137.307 92.568 53.28 22.535 109.858 33.962 168.164 33.962s114.884-11.427 168.164-33.962c51.445-21.76 97.643-52.904 137.306-92.568 39.664-39.663 70.809-85.86 92.568-137.306 22.535-53.28 33.962-109.858 33.962-168.164s-11.427-114.884-33.962-168.164zM543.951 376.03c17.57-0.243 35.095-0.683 52.536-1.368a1990.762 1990.762 0 0 0 91.131-5.678c10.157 34.378 16.643 71.314 19.389 110.46H543.951V376.03z m0-64.069V163.41c17.617 12.176 39.372 29.845 61.074 54.197 23.42 26.28 43.109 56.138 58.784 89.019-41.937 3.147-82.117 4.783-119.858 5.335z m-64-144.763v144.705a2038.572 2038.572 0 0 1-103.41-4.309 1929.6 1929.6 0 0 1-11.183-0.786c15.195-31.915 34.175-60.991 56.676-86.711 20.284-23.185 40.722-40.49 57.917-52.899zM370.842 371.354a2124.301 2124.301 0 0 0 109.109 4.606v103.484H322.25c2.742-39.079 9.208-75.96 19.335-110.285 9.453 0.776 19.21 1.512 29.257 2.195z m-112.731 108.09H145.342c3.916-45.757 16.208-89.778 36.139-130.276 20.257 3.578 52.836 8.759 95.407 13.641-9.902 36.699-16.182 75.682-18.777 116.635z m-0.859 64c1.975 39.144 9.712 78.242 23.101 116.799-44.093 4.976-77.797 10.319-98.621 13.992-20.08-40.64-32.458-84.841-36.39-130.792h111.91z m64.098 0h158.601v104.048c-36.37 0.602-72.818 2.13-109.109 4.601-8.41 0.573-16.612 1.182-24.609 1.821-14.322-36.523-22.659-73.543-24.883-110.47z m158.601 168.103v131.712c-16.974-15.877-36.574-35.962-56.038-59.735-18.064-22.064-33.892-44.698-47.379-67.668a2038.38 2038.38 0 0 1 103.417-4.309z m64 136.665V711.487c34.387 0.503 70.782 1.903 108.68 4.527-13.467 22.915-29.263 45.496-47.286 67.509-21.547 26.32-43.262 48.119-61.394 64.689z m52.536-199.426a2063.187 2063.187 0 0 0-52.536-1.384V543.444h163.958c-2.227 36.975-10.582 74.045-24.939 110.615a1994.972 1994.972 0 0 0-86.483-5.273z m175.521-105.342H878.56c-3.928 45.898-16.283 90.051-36.322 130.652a1822.84 1822.84 0 0 0-93.476-13.439c13.478-38.691 21.264-77.929 23.246-117.213z m-0.859-64c-2.605-41.111-8.924-80.238-18.891-117.061a1821.491 1821.491 0 0 0 90.233-13.075c19.89 40.46 32.158 84.432 36.07 130.136H771.149z m1.017-228.215a373.188 373.188 0 0 1 34.046 39.155 1772.853 1772.853 0 0 1-75.117 9.98c-19.529-46.785-45.825-88.91-78.291-125.338a423.863 423.863 0 0 0-5.386-5.927c46.185 18.263 88.573 45.955 124.748 82.13zM384.81 165.923a420.788 420.788 0 0 0-8.355 9.102c-32.552 36.525-58.904 78.775-78.449 125.71-32.357-3.484-59.48-7.244-80.226-10.469a373.386 373.386 0 0 1 33.956-39.037c38.338-38.338 83.654-67.152 133.074-85.306zM251.736 771.659a373.444 373.444 0 0 1-33.584-38.548c22.526-3.498 52.529-7.614 88.626-11.328 18.229 35.542 41.31 70.356 68.867 103.808a678.068 678.068 0 0 0 35.727 39.995c-59.757-16.875-114.524-48.814-159.636-93.927z m368.249 91.745a677.032 677.032 0 0 0 33.629-37.814c27.469-33.344 50.487-68.043 68.689-103.466a1780.315 1780.315 0 0 1 83.528 10.88 373.34 373.34 0 0 1-33.665 38.654c-43.23 43.232-95.324 74.373-152.181 91.746z"
                      fill="" p-id="4958"></path>
                  </svg>
                  <svg t="1698674687432" class="icon" viewBox="0 0 1024 1024" version="1.1" v-else
                    xmlns="http://www.w3.org/2000/svg" p-id="9635" width="32" height="32">
                    <path
                      d="M496 256.1l-64-112L128 144.1c-57-3-64 32-64 48l0 688 896 0 0-624L496 256.1zM112 192.1l288 0 37.3 64L112 256.1 112 192.1zM912 832.1 112 832.1l0-528 800 0L912 832.1z"
                      p-id="9636"></path>
                  </svg>
                </el-icon>
              </div>

              <div class="bottom">
                <time class="time"></time>
                <p text class="author" @click="date">{{ liem.author !== 0 ? liem.author : '群星' }}</p>
              </div>
            </div>
          </el-card>
        </div>
      </div>
      <!-- <Story /> -->
    </el-scrollbar>
  </div>



  <v-contextmenu ref="contextmenu">
    <v-contextmenu-item @click="goinfo('/info', ref)"><span>书籍详情</span></v-contextmenu-item>
    <v-contextmenu-item @click="removebook(ref)"><span>移出书架</span></v-contextmenu-item>
  </v-contextmenu>
</template>
    
<script setup lang='ts'>
import { Service } from '@element-plus/icons-vue'
import router from '../router'
import { ref, onMounted, computed } from 'vue'
import { ipcRenderer } from 'electron'
import { messageConfig } from 'element-plus';
import { dataType } from 'element-plus/es/components/table-v2/src/common.mjs';

// import data from '../../electron/conslo'
// import { nodeinit, cutFile } from '../node-api/cutFile.js'
// const { nodeinit, cutFile } = require('../node-api/cutFile')
//选择器值
const options = [
  {
    value: 'year',
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
    value: 'local',
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
  // {
  //   value: '',
  //   label: '类型',
  //   children: [
  //     {
  //       value: '',
  //       label: '玄幻'
  //     },
  //     {
  //       value: '',
  //       label: '修仙'
  //     },
  //     {
  //       value: '',
  //       label: '武侠'
  //     },
  //     {
  //       value: '',
  //       label: '都市'
  //     },
  //   ]
  // }
]
//设置选择器多选
const props = {
  multiple: true,
}
//选择器当前选中
const cascader = ref(null)
const cardValue = ref(null)
onMounted(() => {
  // 获取最新的
  date()
  // console.log(dabk.value)
})
// let contextmenu = ref(null)
//转换字符串
function convertToString(value: unknown): string {
  if (typeof value == 'string') {
    return value; // 如果参数已经是字符串类型，则直接返回原本的字符串
  } else {
    return String(value); // 将其他类型的参数转换为字符串并返回
  }
}
//右键书籍详情
const goinfo = (path: any, ref: any) => {
  let ispath = '/info'
  console.log(cardValue.value)

  // console.log(cardValue.value.local)
  router.push({
    // path: ispath,
    name: "info",
    params: {
      name: cardValue.value.name,
      local: cardValue.value.local,
      author: convertToString(cardValue.value.author),
      size: cardValue.value.size,
      mtime: cardValue.modifiedTime,
      coverId: cardValue.value.coverId,
      noverid: cardValue.value.id
    }
  })
  console.log(ref)
}
ipcRenderer.on('GetBookliem', (event, message: Array) => {
  // console.log('书籍列表:', message.fileNames);
  lockbook.value = message.fileNames
  // console.log(dabk.value)
});
ipcRenderer.on('GetwebBook', (event, message: Array) => {
  // console.log(message)
  webbook.value = message
  console.log(webbook)
})
// 过滤后缀
function removeExtension(filename) {
  return filename.replace('.txt', '');
}
function handleContextMenu(event: any) {
  // console.log('event', event)
  cardValue.value = event
}
function removebook(ref) {
  if (cardValue.value.local == 0) {
    // 删除网络书本
    ipcRenderer.send('removewebBoke', cardValue.value.id)
    date()
  } else {
    //删除本地文件
    ipcRenderer.send('removelocalBook', cardValue.value.name)
    setTimeout(() => {
      date()
    }, 150)
  }
}
ipcRenderer.on('okdel', (name) => {
  // console.log(name)
})
ipcRenderer.on('okremove', (e, data) => {
  setTimeout(() => {
    ElNotification.closeAll()
    ElNotification.success({ title: '操作成功!', message: '操作成功' })
  }, 100)
})
ipcRenderer.on('noremove', () => {
  setTimeout(() => {
    ElNotification.closeAll()
    ElNotification.error({ title: '添加失败!', message: '出现了错误' })
  }, 100)
})
//本地数据
let lockbook = ref([])
//网络数据
let webbook = ref([])
// 合并后
let dabk = computed(() => {
  return [...lockbook.value, ...webbook.value]
})
// 获取本地
function copy() {
  dialog.showOpenDialog({
    title: '选择文件',
    properties: ['openFile'],
    filters: [{ name: 'All Files', extensions: ['*'] }],
  }, (filePaths) => {
    if (filePaths && filePaths.length > 0) {
      console.log(`Selected file: ${filePaths[0]}`)
    } else {
      console.log('No file selected')
    }
  })
}
function date() {
  //获取最新
  // dabk.value = data()
  ipcRenderer.send('Upbookdata', '1');
  ipcRenderer.send('GetwebBook', 'web');
  // console.log(webbook.value)
  setTimeout(() => {
    ElNotification.closeAll()
    console.log('111')
    ElNotification.success({ title: '添加成功!', message: '更新书架' })
  }, 100)
  // console.log('1212', dabk.value)
}
// 打开书本详情
function openbook(e: any) {
  console.log(e)
  // local为true为本地,直接跳转阅读界面
  if (e.local) {
    console.log("本地书");
    router.push({
      // path: ispath,
      name: "info",
      params: {
        name: e.name,
        local: e.local,
        size: e.size,
        author: convertToString(e.author),
        mtime: e.modifiedTime,
      }
    })
  } else {
    console.log('不在')
    router.push({
      // path: ispath,
      name: "info",
      params: {
        name: e.name,
        local: e.local,
        author: convertToString(e.author),
        size: e.size,
        mtime: e.modifiedTime,
        coverId: e.coverId,
        noverid: e.id
      }
    })
  }
}
</script>
    
<style >
.v-contextmenu {
  /* height: 200px; */
  width: 100px;
  animation: tanchu 0.2s;
}

/* 菜单淡入淡出动画 */
@keyframes tanchu {
  from {
    width: 0;
  }

  to {
    width: 100px;
  }
}
</style>
<style scoped lang="scss">
.header {
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: space-between;

  .titles {
    width: 200px;
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

.time {
  font-size: 12px;
  color: #888686;
}

.book-card {
  display: inline-block;
  height: 283px;
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
}

.image {
  width: 100%;
  height: 180px;
  display: block;
}

.body {
  display: flex;
  justify-content: center;
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
  width: 88vw;
}

.title {
  font-size: large;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 34px;

  white-space: nowrap;
  text-overflow: ellipsis;

  // width: 100px;
  .title-txt {
    width: 9rem;
    overflow: hidden;
  }
}



.time {
  margin-top: -20px;
}

html.dark .titles {

  animation: asideBackground 1.5s;
  background-color: rgb(67, 63, 63);
}
</style>