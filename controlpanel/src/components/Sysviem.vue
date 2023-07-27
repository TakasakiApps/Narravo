<template>
  <div>
    <el-card class="box-card">
      <div class="mt-4 search-main">
        <el-input v-model="SearchInput" placeholder="在此搜索" class="input-with-select">
          <template #prepend>
            <el-select v-model="select" placeholder="Select" style="width: 100px">
              <el-option label="书本名" value="1" />
            </el-select>
          </template>
          <template #append>
            <div>
              <el-button :icon="Search" size="small" />
            </div>
          </template>
        </el-input>
      </div>
      <el-table :data="filterTableData">
        <el-table-column label="书名" prop="BookName" />
        <el-table-column label="作者" prop="name" />
        <el-table-column label="作者" prop="name" />、
        <el-table-column label="tag" prop="name" />
        <el-table-column align="right">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)">Edit</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">Delete</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination background layout="prev, pager, next" :total="55" class="page" />
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { Search } from '@element-plus/icons-vue'
import readRoledir from '../nodee/node-api'
interface User {
  BookName: string
  name: string
  address: string
}

const search = ref('')
// 搜索
const SearchInput = ref('')
const select = ref('书本名')
const filterTableData = computed(() =>
  tableData.filter(
    (data) =>
      !search.value ||
      data.name.toLowerCase().includes(search.value.toLowerCase())
  )
)
const handleEdit = (index: number, row: User) => {
  console.log(index, row)
}
const handleDelete = (index: number, row: User) => {
  console.log(index, row)
}

const tableData: User[] = [
  {
    BookName: 'xxxx',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    BookName: '2016-05-02',
    name: 'John',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    BookName: '2016-05-04',
    name: 'Morgan',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    BookName: '2016-05-01',
    name: 'Jessy',
    address: 'No. 189, Grove St, Los Angeles',
  },
]
</script>

<style scoped>
.box-card {
  /* width: 100%; */
  /* height: 100px; */
}

.page {
  /* float: left; */
  display: flex;
  justify-content: center;
}

.search-main {
  width: 28rem;
  position: relative;
  right: 0px;
  margin-bottom: 1rem;
}
</style>