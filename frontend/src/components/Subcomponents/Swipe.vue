<template>
  <div>
    <van-swipe class="my-swipe" :autoplay="3000" indicator-color="white">
      <van-swipe-item v-for="index in 4" :key="index">
      <img>
      </van-swipe-item>
    </van-swipe>
  </div>
</template>

<script lang="ts" setup>
import{   Swipe as vanSwipe, SwipeItem as vanSwipeItem } from 'vant'
const mySwipe ={
  vanSwipe,
  vanSwipeItem
}
</script>

<style scoped>

</style>