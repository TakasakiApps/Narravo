<template>
  <div class="header">
    <div class="titles">
      <p style="width: 90px; margin-left: 10px;">筛选:</p>
      <el-cascader :options="options" :show-all-levels="false" :props="props" collapse-tags-tooltip collapse-tags
        clearable />
    </div>
    <el-button round type="danger" size="large">导入</el-button>
  </div>
  <div style="height:640px">
    <el-scrollbar height="600px">
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
      <!-- <Story /> -->
    </el-scrollbar>
  </div>



  <v-contextmenu ref="contextmenu">
    <v-contextmenu-item @click="goinfo('/info')"><span>书籍详情</span></v-contextmenu-item>
    <v-contextmenu-item><span>移出书架</span></v-contextmenu-item>

  </v-contextmenu>
</template>
    
<script setup lang='ts'>
import { Service } from '@element-plus/icons-vue'
import router from '../router'
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
const goinfo = (path)=> {
  router.push(path)
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
  height: 100px;
  display: flex;
  align-items: center;
  justify-content: space-between;

  .titles {
    width: 180px;
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

html.dark .titles {

  animation: asideBackground 1.5s;
  background-color: rgb(67, 63, 63);
}
</style>