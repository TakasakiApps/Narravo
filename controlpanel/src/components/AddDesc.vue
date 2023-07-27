<template>
  <div>
    <el-card class="box-card">
      <el-table :data="Desc" style="width: 100%" table-layout="auto">
        <el-table-column prop="name" label="名字" width="180" />
        <el-table-column prop="description" label="简介" />
        <el-table-column label="操作">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleClick">更新</el-button>
            <el-button link type="primary" size="small" @click="delClick(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <el-dialog width="30%" v-model="dialogshow">
      <autcar></autcar>
    </el-dialog>
  </div>
</template>
<script lang="ts" setup>
import axios from "axios";
import { onMounted, reactive, ref } from "vue";
import VueCookies from 'vue-cookies'
import autcar from "./layouts/authcar.vue";

const dialogshow = ref(false)
const cookie = VueCookies
const Desc = reactive([])
onMounted(() => {
  axios.get('/api/v1/author/getAllInfo', {
    params: {
      token: cookie.get('token'),
      page: 1,
      count: 20,
    }
  }).then(
    (e) => {
      const mydata = e.data.list
      console.log(e.data)
      //暂存数据
      mydata.forEach(element => {
        //赋值
        Desc.push(element)
        console.log(Desc.values)
      });
    }
  )
})
const handleClick = () => {
  dialogshow.value = true
}
const delClick = (e: any) => {
  console.log(e.id)
  if (e.id !== '') {
    const authorId = e.id
    axios.get(`/api/v1/author/delete/${authorId}`, {
      params: {
        token: cookie.get('token')
      }
    }).then(
      () => {
        console.log('删除成功')
      }
    ).catch(
      err => {
        console.log(err.message)
      }
    )
  }
}
</script>

<style lang="scss" scoped>
.box-card {
  width: 70%;
  margin: auto;
}
</style>