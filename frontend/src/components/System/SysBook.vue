<template>
  <div>
    <div class="Sys-tag">
      外观设置
    </div>
    <van-switch v-model="type" @click="checked()" active-color="#7a7374" inactive-color="#dcdee0" class="Sys-main"  size="22px"/>
  </div>
</template>

<script setup lang="ts">
import { ipcRenderer } from 'electron';
import { ref } from 'vue';
import mitt from '../../samples/EventBus';

let type = ref(true)
function checked() {
  if (type.value === true) {
    mitt.emit('change')
  } else {
    mitt.emit('change')
  }

  return !type.value
}

</script>

<style scoped>
.Sys-main::after{
  contain: '切换主题颜色';
  color: #000;
}
</style>