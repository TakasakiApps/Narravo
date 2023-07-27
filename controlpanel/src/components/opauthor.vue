<template>
  <el-select v-model="value" filterable placeholder="输入作者名后选择" remote :remote-method="remoteMethod" :loading="loading">
    <el-option v-for="(item) in authordata" :key="item.avatarId" :label="item.name" :value="item.name"
      @click="useAut(item)" />
  </el-select>
</template>

<script lang="ts" setup>
import { onMounted, reactive, ref } from 'vue'
import axios from 'axios'
import { ElMessageBox } from 'element-plus'
import VueCookies from 'vue-cookies'
const cookie = VueCookies

interface AuthorDataItem {
  avatarId: string;
  description: string;
  name: string;
  id: string
}

const emit = defineEmits(['ChildEvent'])
onMounted(() => {
  // console.log('authordata')
})



const dialogVisible = ref(false)

const handleClose = (done: () => void) => {
  ElMessageBox.confirm('Are you sure to close this dialog?')
    .then(() => {
      done()
    })
    .catch(() => {
      // catch error
    })
}

const useAut = (e: any) => {
  console.log(e)
  emit('ChildEvent', e)
}
const authordata = reactive<AuthorDataItem[]>([])
const value = ref('')
//加载状态
const loading = ref(false)
const author = reactive({
  token: '',
  name: '',
  description: ''
})
const remoteMethod = (query: String) => {
  if (query) {
    loading.value = true
    axios.get('/api/v1/author/search', {
      params: {
        token: cookie.get('token'),
        page: 1,
        count: 8,
        keyword: query,
      }
    }).then(
      (e) => {
        loading.value = false
        const mydata = e.data.list;
        mydata.forEach((item: AuthorDataItem) => {
          authordata.push({
            avatarId: item.avatarId,
            description: item.description,
            name: item.name,
            id: item.id
          });
        });
        console.log(authordata)
      }
    )
  }
}
// const options = [
//   {
//     value: 'Option1',
//     label: 'Option1',
//   },
//   {
//     value: 'Option2',
//     label: 'Option2',
//   },
//   {
//     value: 'Option3',
//     label: 'Option3',
//   },
//   {
//     value: 'Option4',
//     label: 'Option4',
//   },
//   {
//     value: 'Option5',
//     label: 'Option5',
//   },

// ]
</script>
<style>
.button {
  margin-left: 50px;
  width: 80px;
}
</style>