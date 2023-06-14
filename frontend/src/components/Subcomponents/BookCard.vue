<template>
  <div class="BookCard">
    <a :to="from">
      <van-image width="10rem" height="10rem" radius="5" fit="cover" :src="url" alt="" />
      <h5>
        {{ name }}
      </h5>
      <div class="Name">
        {{ name }}
      </div>
      <!-- </van-image> -->
    </a>
  </div>
</template>

<script lang="ts" setup>
const props = defineProps({
  name:String,
  from:String,
  url:String
})
</script>

<style scoped>
.BookCard{

  /* 记得去掉 */
  border: 1px solid #ccc; /* 设置边框 */
}
*{
  margin: 0px;
}
.Name{
  opacity: 0.5;
  
}
</style>