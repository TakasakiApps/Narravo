<template>
  <div>
    <div class="tag">
      <h5>Narravo</h5>
      <div class="button">
        <electon-botton />
      </div>
    </div>
    <div class="tags">
      <van-tabs v-model:active="active" class="Systga"  gutter="50">
        <van-tab title="首页" to="/"  title-class="vanTab" title-style="width: 12.5rem;">
        </van-tab>
        <van-tab title="创造" to="/CreateVue" title-class="vanTab">
        </van-tab>
        <van-tab title="书架" to="/BookShelfVue" title-class="vanTab">
        </van-tab>
        <van-tab title="设置" to="/SysHomeVue" title-class="vanTab">
        </van-tab>
      </van-tabs>
        <van-cell-group inset class="Search">
        <van-field v-model="Search" placeholder="" left-icon="search" @keyup.enter="handleEnter" class="Search-box"/>
        </van-cell-group>
        </div>
      <div class="view">
        <router-view />
    </div>

  </div>
</template>

<script lang="ts" setup>
import {ref,onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router'
import  NewToken from './samples/node-api';
const router = useRouter();
const route = useRoute();
onMounted(() => {
  // console.log(`取得token.`)
})
const Search = ref('')

const active = ref(0)
function handleEnter(){
  const liem  = {
    1:{
      name:`${Search.value}`,
      id: `${Search.value}`,
      act:'孙悟空',
      Brief:'1111111111111'
    },
    2:{
      name:"斗罗大陆",
      id:'2',
      act:'唐山',
      Brief: '1111111111111'
    },
    3:{
      name:"异灵鼠cb",
      id:'3',
      act:"异灵鼠",
      Brief: '1111111111111'
    },
    4:{
      name: "异灵鼠cb",
      id: '3',
      act: "异灵鼠",
      Brief: '1111111111111'
    },
    5:{
      name: "异灵鼠cb",
      id: '3',
      act: "异灵鼠",
      Brief: '1111111111111'
    }
  }
  // 处理用户搜索内容
  const userSearch = Search.value
  // console.log(userSearch)
  // 从json转换成字符串用params传递给组件
  router.push({
    name:`SearchLiem`,
    params:{
      search:userSearch
    },
    query:{
      // search:userSearch,
      data: JSON.stringify(liem)
    }
  })
}
</script>


<style lang="scss" scoped>
.Systga{
  display: flex;
  justify-content: center;
  // background-color: $color-gree;
}
.tag{
  display: flex;
  /* align-items: center; 垂直居中 */
  justify-content: space-between; /* 左右对齐，中间留有空白间隔 */
  height: 36px;
  padding: 0 20px; 
  padding-right: 0px;
  z-index: 10;
  position:sticky;
  position: -webkit-sticky;    // 兼容 -webkit 内核的浏览器
  top: 0px;  
  -webkit-app-region: drag;
  -webkit-user-select: none;
  background-color: white ;
}

.van-tab--active{
  width: 100%;
}
h5{
  margin: .125rem .1875rem ;
  padding: 0;
  opacity: 0.85;
}
.tags{
  position: relative;
  top: 0;
}
.Search{
  // position: relative;
  // right: 10px;
  position: absolute;
  top: 0px;
  right: 6rem;
  display: flex;
  align-items: center;
}
.Search-box{
  height: 80%;
  background-color:#f2f2f4;

}

.view{
  overflow:hidden;
  text-align: center;
}
</style>