<template>
  <div>
    <div>
      <el-input class="searchinput" v-model="inputvalue" :suffix-icon="Search" @blur="searchbook(1, inputvalue)">
      </el-input>
    </div>
    <ul v-infinite-scroll="load" class="infinite-list">
      <BookLiem v-for="i in Nowsearch" :key="i.id" :name="i.name" :bookid="i.id" :mtime="i.author.UpdatedAt"
        :act="i.author.name" :coverid="i.coverId" />
      <el-card v-if="!Nowsearch.length || Nowsearch.length == null" class="tag">
        <img src="../assets/nodata.67f7a1c9.png" alt="">
      </el-card>
    </ul>
  </div>
</template>

<script setup>
import BookLiem from '../components/BookLiem.vue'
import { ref } from 'vue'
import { client } from '../http/client'
import { Search } from '@element-plus/icons-vue'
let Searchinput = ref(null)
//当前数据量
const count = ref(1)
//当前收缩列表
const Nowsearch = ref([])
//当前搜索目标
const oldsearch = ref(null)
//实际输入框
const inputvalue = ref(null)
//
const load = () => {
  //改变当前数据量
  count.value += 1
  searchbook(count.value, oldsearch)
}
// 初始化搜索
function searchinit(newdata) {
  //将页数清空
  count.value = 1;
  Nowsearch.value = [];
  oldsearch.value = Searchinput.value;
  // searchbook(count.value, oldsearch)
}
function searchbook(pg, event) {
  client.get('/api/v1/novel/search', {
    params: {
      token: localStorage.token,
      page: pg,
      count: '5',
      keyword: event
    }
  }).then(
    res => {
      console.log(res.data.list)
      if (res.data.list !== null) {
        Nowsearch.value = res.data.list
      }
      console.log(Nowsearch.value, '2123123')
    }
  ).catch(
    err => {
      console.log(err)
      setTimeout(() => {
        ElNotification.closeAll()
        ElNotification.error({ title: '操作失败!', message: '请检查是否登录' })
      }, 100)
    }
  )
}
</script>

<style lang="scss" scoped>
.infinite-list {
  padding: 0;
}

.searchinput {
  position: relative;
  width: 70vw;
  height: 7vh;
  left: 50%;
  transform: translate(-50%);
}

.tag {
  position: absolute;
  top: 50%;
  /*将子元素的顶部边缘置于父元素的中间位置*/
  left: 55%;
  /*将子元素的左侧边缘置于父元素的中间位置*/
  transform: translate(-50%, -50%);
  /*通过负的50%偏移，使子元素居中*/
}
</style>