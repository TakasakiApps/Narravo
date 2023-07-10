<template>
  <div class="rolaCard">
    <div class="rolaCard-svg">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="115.648 92.224 792.8 835.1" width="176" height="176"
        v-if="props.gender" class="girlsvg">
        <path
          d="M613.312 626.496v-11.2h151.296V318.08c0-124.736-113.152-225.856-252.608-225.856S259.392 193.344 259.392 318.08v297.216h151.36v11.072h-0.256L115.648 820.16v107.136H908.48v-105.92l-295.168-194.88z m7.68 186.752c-72.192 69.824-109.504-36.416-109.504-36.416s-37.312 106.24-109.568 36.416c0 0-31.744-20.544-35.136-158.208l144.704 121.728L656.192 655.04c-3.392 137.6-35.2 158.208-35.2 158.208z"
          fill="currentColor" />
      </svg>
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0.00221157 1024 1024" width="176" height="176" v-else
        class="mansvg">
        <path
          d="M828.875 765.657c-191.159-16.86-202.516-102.922-202.516-102.922v-85.997c111.656-43.063 145.76-207.699 145.76-207.699 0-65.457-32.153-67.281-32.153-67.281V195.093c7.486-177.753-100.33-136.61-102.25-136.61-1.791 0-60.466-46.774-60.466-46.774-41.751-31.8-134.499 11.262-134.499 11.262l-3.839 0.064c-104.937 2.751-132.387 78.479-132.387 78.479l1.92 196.437c-54.804 0-47.35 76.719-47.35 76.719 0 84.142 147.616 200.148 147.616 200.148 9.47 110.472-49.174 166.46-49.174 166.46s-100.266 3.711-213.84 37.432c-2.72 0.863-5.344 1.6-7.967 2.463C30.617 817.421 0 920.855 0 920.855v101.033l463.579 1.888 68.465-1.92h491.732V943.25c1.92-115.815-194.901-177.593-194.901-177.593z m-244.65 129.22l-71.25 128.675-69.616-128.676-16.124-64.05 30.68-136.354h111.656l29.178 136.354-14.525 64.05z"
          fill="currentColor" />
      </svg>
    </div>
    <div class="rolaCard-main">
      <div class="useName">
        {{ name }}
      </div>
      <div class="readerName">
        {{ Reader }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, onMounted, computed } from 'vue';
// import Man from '@/assets/mman.svg';
// require('../../man.svg')
const props = defineProps<{
  name: String,
  gender: Number,
  Reader: String,
}>()
onMounted(() => {

})
const Genders = computed(() => {
  if (props.gender == 1) {
    return '#f06291e3'
  } else {
    return '#65B6F6'
  }
})
</script>

<style scoped lang="scss">
$girlColor: #F6A7C2;
$girlBackgroundColor : #f06298ae;
$manBackgroundColor: #65B6F6;
$manColor: #A7D5FA;

//整体卡片样式
.rolaCard {

  width: 178px;
  height: 271px;
  border-radius: 12px;
  background-color: v-bind(Genders);
}

.girlsvg {
  color: $girlColor;
}

.mansvg {
  color: $manColor;
}

//卡片头像样式
.rolaCard-svg {
  position: relative;
  top: 25px;
}

//下部主题样式
.rolaCard-main {
  position: relative;
  top: 21px;
  background-color: white;
  height: calc(271px - 200px);
  border: 1px solid rgba(0, 0, 0, 0.182);
  border-radius: 0% 0% 12px 12px;
}

.useName {
  margin: 8px;
  font-weight: bold;
}

.readerName {
  margin: 8px;
  color: #A2A2A2;
  font-size: 13px;
}
</style>