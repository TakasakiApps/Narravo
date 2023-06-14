<template>
  <div>
    <h2 class="book-name">
      {{ bookxq.name }}
    </h2>
    <img :src=bookxq.img class="bookimg">
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
//根据名字来获取信息啥的
onMounted(() => {
  console.log(`应该获取了.`)

})
const bookxq = {
  img: '/public/C1.png',
  name:"天龙八部"
}


</script>

<style scoped>
.book-name{
  text-align: center;
}
.bookimg{
  max-height: 12.5rem;
  max-width: 12.5rem;
}
</style>