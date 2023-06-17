<template>
  <div>
    <SysBook  class="container no-select"/>
    <SysUser  class="container"/>
  </div>
</template>

<script lang="ts" setup>
import SysUser from './System/SysUser.vue';
import SysBook from './System/SysBook.vue';

</script>

<style>
.container {
  display: flex;
  justify-content: space-between; /* 左右对齐 */
  align-items: center; /* 垂直居中 */
  height: 3.125rem; 
}

.Sys-tag {
  display: inline-block;
  
}

.Sys-main {
  display: inline-block;
}
</style>