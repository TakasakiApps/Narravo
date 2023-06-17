<template>
  <div class="Swipe-view">
    <h2>
      今日推荐
    </h2>
      <van-space :size="20">
        <book-card  v-for=" itme in props" 
          :name="itme.name" 
          :url="itme.url"
          :from="itme.from" 
          :key="itme.id" />
      </van-space>
  </div>
</template>

<script lang="ts" setup>
const props = [
  {
    name: "a",
    url: "/src/assets/C1.png",
    from: '/a',
    id: '1'
  },
  {
    name: "b",
    url: "/src/assets/C1.png",
    from: '/b',
    id: '2'
  },
  {
    name: "b",
    url: "/src/assets/C1.png",
    from: '/b',
    id: '2'
  },
  {
    name: "b",
    url: "/src/assets/C1.png",
    from: '/b',
    id: '2'
  },
  {
    name: "b",
    url: "/src/assets/C1.png",
    from: '/b',
    id: '2'
  },
]

</script>

<style scoped>
h2{
  margin-top: 0rem;
  margin-bottom: 0.65rem;
}
.Swipe-view{
  margin: 0 30px;
}
</style>