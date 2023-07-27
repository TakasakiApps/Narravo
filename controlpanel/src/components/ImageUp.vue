<template>
  <div>
    <el-upload class="avatar-uploader" action="" :show-file-list="false" :on-success="handleAvatarSuccess"
      :before-upload="beforeAvatarUpload"><img v-if="imageUrl" :src="imageUrl" class="avatar" />
      <el-icon v-else class="avatar-uploader-icon">
        <Plus />
      </el-icon>
    </el-upload>
    <el-dialog v-model="showCropperDialog" title="裁剪图片" width="60%">
      <vue-cropper ref="cropper" v-model="croppedImage" :options="cropperOptions"
        @initialized="onCropperInitialized"></vue-cropper>

      <div class="dialog-footer">
        <el-button @click="handleCropAndUpload">确定</el-button>
        <el-button @click="showCropperDialog = false">取消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { ElButton, ElDialog } from 'element-plus';
import VueCropper from 'vue-cropper';

const showCropperDialog = ref(false);
const croppedImage = ref('');
const cropperRef = ref();
const imageUrl = ref()
const cropperOptions = {
  guides: true,
  viewMode: 1,
  dragMode: 'crop',
  background: true,
  rotatable: true,
  aspectRatio: 1,
};

const onCropperInitialized = (event) => {
  const cropperInstance = event.detail.cropper;
  cropperRef.value = cropperInstance;
};

const handleCropAndUpload = () => {
  const cropper = cropperRef.value;

  if (!cropper) {
    return;
  }

  const canvasData = cropper.getCroppedCanvas();

  if (canvasData) {
    const base64Image = canvasData.toDataURL('image/jpeg');
    // 执行上传逻辑，使用base64Image
  }

  showCropperDialog.value = false;
};
</script>

<style scoped>
.el-button {
  margin-bottom: 20px;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
}
</style>
