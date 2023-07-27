<template>
  <div class="main">
    <!--使用ref属性给图片元素命名为imageRef-->
    <img ref="imageRef" :src="imageSrc" alt="image">
    <button @click="cropImage">裁剪图片</button>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, watch } from "vue";
import Cropper from 'cropperjs';
import "cropperjs/dist/cropper.css";
import router from "~/router";

const props = defineProps({
  //图片地址
  imageSrc: {
    type: String,
    required: true
  },
  //aspectRatio:置裁剪框为固定的宽高比
  aspectRatio: {
    type: Number,
    default: 1,
  },
  //viewMode: 视图控制
  viewMode: {
    type: Number,
    default: 1,
  },
  //autoCropArea: 设置裁剪区域占图片的大小 值为 0-1 默认 0.8 表示 80%的区域
  autoCropArea: {
    type: Number,
    default: 1,
  },
  //是否可以改变尺寸
  cropBoxResizable: {
    type: Boolean,
    default: false,
  }
})
//绑定图片的dom对象
const imageRef = ref<HTMLInputElement | null>(null)
let cropper: any = null;
//使用Cropper构造函数创建裁剪器实例，并将图片元素和一些裁剪选项传入
onMounted(() => {
  console.log('打开了')
  if (imageRef.value) {
    cropper = new Cropper(imageRef.value, {
      aspectRatio: props.aspectRatio,
      viewMode: props.viewMode,
      autoCropArea: props.autoCropArea,
    });
  } else {
    console.log('出现了意de')
  }
});
//定义方法
const emit = defineEmits(['updateImageSrc'])
//在cropImage函数中，获取裁剪后的图片数据URL，并使用emit方法触发updateImageSrc事件，
// 将裁剪后的图片数据URL传递给父组件。
const cropImage = () => {
  const canvas = cropper.getCroppedCanvas();
  const croppedImage = canvas.toDataURL();
  console.log(canvas, croppedImage);

  canvas.toBlob((e: any) => {
    const reader = new FileReader
    reader.readAsArrayBuffer(e);
    reader.onload = () => {
      console.log('BufferBufferBufferBuffer')
      emit('updateImageSrc', croppedImage, reader.result);
      console.log(reader.result)
    }


  })
  // console.log('xxxx', xxx)

}
//销毁
onBeforeUnmount(() => {
  cropper.destroy()
})
watch(props, () => {
  console.log('11')
  cropper.replace(props.imageSrc)
})
</script>

<style scoped>
img {
  width: 400px;
  height: 200px;
}

.main {}
</style>