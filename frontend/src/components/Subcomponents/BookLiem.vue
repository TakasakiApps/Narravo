<template>
  <div class="book-liem-main">
    <div class="book-img box-margin">
      <img src="../../assets/C1.png" alt="" class="book-img-main 
      shadow-effect 
      no-select 
      no-drag
      cursor-pointer"
      >
    </div>
    <div class="book-main box-margin">
      <h3 class="book-name">{{ props.name }}</h3>
      <div class="book-liem-act">
        <div class="book-liem-inln">
          {{ props.act }}
        </div>
        <div class="book-liem-inln">
          <van-tag plain type="primary">{{ props.act }}</van-tag>
        </div>
      </div>
      <div class="book-liem-Brief">
        {{ props.Brief }}
      </div>
      <van-button type="default" size="small">加入书架</van-button>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  id: Number,
  name: String,
  act: String,
  Brief: String
})
</script>

<style scoped>
.box-margin{
  margin-top: 1.25rem;
  margin-left: 3rem;
}
.book-liem-main{
  display: block;
  height: 200px;
}
.book-img{
  display: inline-flex;
  float: left;
  max-height: 200px;
}
.book-img-main{
  width: 200px;
}
.book-main{
  display: inline;
  float: left;
  text-align: left;
}
.book-liem-act{
  margin-top: 1rem;
}
.book-liem-inln{
  display: inline;
}
.book-liem-Brief{
  margin-top: 1rem;
}
</style>