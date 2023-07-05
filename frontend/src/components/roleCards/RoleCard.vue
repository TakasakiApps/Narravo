<template>
  <div>
    <el-button @click="dialogVisible = true">
      新建角色卡
    </el-button>
    <el-dialog v-model="dialogVisible" title="角色信息" width="30%" :style="{ background: rgb, transition: transition }"
      class="card-main">
      <el-form label-position="left">
        <el-form-item>
          <template #label>
            <div class="max-txt">
              角色名:
            </div>
          </template>
          <el-input v-model="Role.name" />
        </el-form-item>
        <el-form-item>
          <template #label>
            <div class="max-txt">
              角色模型:
            </div>
          </template>
          <el-select v-model="Reader" class="m-2" placeholder="Select">
            <el-option v-for="item in Readerliem" :key="item.label" :label="item.name" :value="item.label" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <template #label>
            <div class="max-txt">
              代表颜色:
            </div>
          </template>
          <el-color-picker v-model="rgb" show-alpha />
        </el-form-item>
        <el-form-item>
          <template #label>
            <div class="max-txt">
              角色语速:
            </div>
          </template>
          <el-slider v-model="Role.speed" :max="10" />
        </el-form-item>
        <el-form-item>
          <template #label>
            <div class="max-txt">
              角色语调:
            </div>
          </template>
          <el-slider v-model="Role.tone" :max="10" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">创建</el-button>
          <el-button @click="dialogVisible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, computed, watch, onMounted } from 'vue';
import { ElMessageBox } from 'element-plus'
import readRoledir from '~/nodee/node-api';
onMounted(() => {
  //获取角色信息
})
const dialogVisible = ref(false)
const runoo = /rgba\((\d+),\s*(\d+),\s*(\d+),\s*([0-9]*\.?[0-9]+)\)/
const Txtcolor = ref('black')
//角色颜色值
const rgb = ref("rgb(255,255,255,1)");
const transition = ' background-color 2s linear'
const Role = reactive({
  name: '',
  speed: 10,
  tone: 10
});
const depth = ref(0)
//监视属性改变计算深浅度
watch(rgb, () => {
  console.log(runoo.exec(rgb.value))
  depth.value = 1
  let liem = runoo.exec(rgb.value)
  if (liem) {
    // console.log(Number(liem[0]))
    for (let x = 1; x < liem.length - 1; x++) {
      let i = Number(liem[x])
      // console.log(i, x)
      switch (x) {
        case x = 1:
          depth.value += i * 0.2126
          break
        case x = 2:
          depth.value += i * 0.7152
          break
        case x = 3:
          depth.value += i * 0.0722
          break
      }
      if (depth.value / 255 < 0.5) {
        Txtcolor.value = 'white'
      } else {
        Txtcolor.value = 'black'
      }
    }
  }

  // if (liem) {
  //   for (let i of liem) {
  //     if()
  //   }
  // } else {

  // }
})
const Readerliem = [
  {
    name: '小帅',
    label: '小'
  }
]
const Reader = ref('')
//角色卡数据
const RolaCard = reactive({
  name: Role.name,
  speed: Role.speed,
  tone: Role.tone,
  Reader: Reader,
  rgb,
})
// const handleClose = (done: () => void) => {
//   ElMessageBox.confirm('Are you sure to close this dialog?')
//     .then(() => {
//       done()
//     })
//     .catch(() => {
//       // catch error
//     })
// }
const onSubmit = () => {
  //
  readRoledir(undefined, RolaCard)

}
</script>

<style lang="scss" >
// .card-main {
//   transition: background-color 2s linear;
// }
.color-picker {
  width: 100px;

}

.ep-color-picker__trigger {
  width: 200px;
}

.ep-form-item__content {
  justify-content: center;
}

.max-txt {
  color: v-bind(Txtcolor);
  transition: background-color 2s linear;
}
</style>
