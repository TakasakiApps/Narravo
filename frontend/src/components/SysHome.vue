<template>
  <div>
    <!-- 切换主题 -->
      <van-cell>
        <!-- 使用 title 插槽来自定义标题 -->
        <template #title>
            <span >
              主题颜色
            </span>
        </template>
        <template #label>
              <span>
                切换主题
              </span>
          </template>
        <!-- 使用插槽来自定义右侧图标 -->
        <template #value>
          <van-switch v-model="type" @click="checked()" active-color="#1c1c1e" inactive-color="#dcdee0"  />
        </template>
      </van-cell>
      <van-cell>
        
      </van-cell>
  </div>
</template>

<script lang="ts" setup>
import { ipcRenderer } from 'electron';
import { ref } from 'vue';
import mitt from '../samples/EventBus';

let type = ref(true)
function checked(){
  if(type.value===true){
    mitt.emit('change')

  }else{
    mitt.emit('change')


  }

  return !type.value
}

</script>

<style scoped>
.cell-lift{
  float: left;
}
</style>