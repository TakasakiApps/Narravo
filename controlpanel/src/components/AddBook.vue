<template>
  <div>
    <el-card class="box-card">
      新增书籍
      <el-form :model="book" label-width="120px">
        <el-form-item label="书名" class="fromIt">
          <el-input v-model="book.name" class="fromVa" placeholder="请输入书名" />
        </el-form-item>
        <el-form-item label="选择作者" class="fromIt">
          <opauthor :authordata="author" class="fromVa" @childEvent="ChildEvent" />
        </el-form-item>
        <el-form-item label="简介" class="fromIt">
          <el-input v-model="book.description" autosize type="textarea" resize="none" placeholder="请输入简介"
            class="fromVa" />
        </el-form-item>
        <el-form-item>
          <div class="fromIt">作品头像</div>
          <!-- <imageUp></imageUp> -->
          <div class="cropper fromVa">
            <!-- <el-upload :before-upload="bfUpload" :on-change="uploadChange" :show-file-list="false" accept=".jpg,.png"
              :limit="1">
              <el-button type="primary">Click to upload</el-button></el-upload> -->
            <el-button @click="upnowimg" class="inputbtn" type="primary">点击上传图片</el-button>
            <input type="file" ref="uploadRef" style="display: none;" @change="openUpload" accept="image/*">
            <!-- 对话框 -->

            <div v-if="true">
              <!-- <p>预览图片</p> -->
              <!-- <button @click="dialogVisible = true">111</button> -->
              <!-- <img :src="imageNew" alt="" class="croimg"> -->
            </div>
          </div>
        </el-form-item>
        <el-form-item class="butns">
          <el-button type="primary" @click="onSubmit">创建</el-button>
          <el-button @click="backrouter">取消</el-button>
        </el-form-item>
      </el-form>
      <el-dialog v-model="dialogVisible" title="裁剪图片" width="50%">
        <template #footer>
          <CropperJs :imageSrc="imageSrc" @updateImageSrc="updateImageSrc" />
        </template>
      </el-dialog>
      <div class="sampleCard">
        <div>
          <el-card :body-style="{ padding: '0px' }">
            <img :src="imageNew" alt="" v-if="imageNew !== ''" ref="newcropimg" class="croimg image">
            <!-- <img :src="img1" class="image" /> -->
            <div style="height: 69px;">
              <span class="sampName">{{ book.name }}
              </span>
              <div class="newchapter">最新章节</div>
              <div class="bottom">
                <time class="sampAut">{{ book.author }}</time>
              </div>
            </div>
          </el-card>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { reactive, onMounted, ref } from 'vue';
import VueCookies from 'vue-cookies'
const cookie = VueCookies
import CropperJs from './layouts/CropperJs.vue';
import { ElMessageBox, ElMessage } from 'element-plus'
import type { UploadProps, UploadFile, UploadFiles } from 'element-plus'
// import { client as axios } from '../http/client';
import { client as axios } from '/src/http/client'
import opauthor from './opauthor.vue';
import imageUp from './ImageUp.vue';
//
const limgs = ref()
// import img1 fr
const img1 = ref('1.jpg')
onMounted(() => {
  book.token = cookie.get('token')

})
const newcropimg = ref<HTMLInputElement | null>(null)
const author = ref([])
const dialogVisible = ref(false)
// const cookie = new Cookies
const book = reactive({
  token: '',
  name: '',
  description: '',
  coverId: 'e1ellezv5oyp5aynlwfeqw0367kt6r0h',
  authorId: '',
  author: ''
})
const upimg = reactive({
  token: '',
  url: '',
  name: '',
  size: '',
  type: 'image/png',
})
const imageSrc = ref('')
//定义一个imageNew变量来接收裁剪之后的图片
const imageNew = ref('1.jpg')
//点击裁剪按钮
const updateImageSrc = (updateImageSrc: any, croppedBlob: any) => {
  // uploadRef.value?.files[0]
  imageNew.value = updateImageSrc
  bloArray.value = croppedBlob
}
//上传图片
//获取实例
const uploadRef = ref<HTMLInputElement | null>(null);
const upnowimg = () => {
  uploadRef.value?.click()
  // console.log(uploadRef.value)
}
const openUpload = (e: Event) => {
  const inputElement = e.target as HTMLInputElement;
  if (inputElement.files && inputElement.files.length > 0) {
    const uploadedFile = inputElement.files[0];
    if (uploadedFile.type.startsWith('image/')) {
      // 处理上传的图片文件
      // console.log('上传的图片:', uploadedFile);
      if (uploadedFile.size < 1024 * 1024 * 3) {
        displayImage(uploadedFile);
        dialogVisible.value = true;

      }
      // limgs.value = uploadedFile
      console.log(uploadedFile)

    } else {
      // 非图片文件错误处理
      console.error('只能上传图片文件');
      // 清空文件选择框
      inputElement.value = '';
    }
  }
}
//
// const uploadChange = (file: File, uploadFiles: UploadFile) => {
//   const reader = new FileReader();
//   // reader.readAsDataURL(url)
//   console.log('111111')
//   reader.onload = () => {
//     imageSrc.value = reader.result as string;
//   };
//   reader.readAsDataURL(file);
// }
//阻止上传
// const bfUpload = (rawFile: UploadFile) => {
//   return false
// }
//处理文件
const bloArray = ref()
const displayImage = (file: File) => {
  const reader = new FileReader();
  reader.onload = () => {
    imageSrc.value = reader.result as string;
  };
  reader.readAsDataURL(file);

  // console.log('ppppppppppppp', reader.readAsDataURL(file))
};
//传递
const ChildEvent = (e) => {
  book.author = e.name
  book.authorId = e.id
}
// const addavatar = () => {
//   axios.post('/api/author/add', {
//     token: book.token,
//     name: '唐家六少',
//     description: "string",
//     avatarId: 'ou01eeudolo3nk7p6z83011n3glgdymi'
//   }).then(
//     () => {
//       console.log('成功')
//     }
//   )
// }

// //上传打开对话框
// const handleChangeUpload = (uploadFile: UploadFile, uploadFiles: UploadFiles) => {
//   //拿到图片参数
//   console.log(uploadFile)
//   //url地址
//   upimg.url = document.getElementsByClassName('ep-upload__input')[0].value;
//   upimg.name = uploadFile.name;
//   console.log(upimg.url)
//   // upimg.type = uploadFile.t;
//   // upimg.size = uploadFile.size
//   dialogVisible.value = true
//   return true
// }
// 对话框
// const handleClose = (done: () => void) => {
//   ElMessageBox.confirm('Are you sure to close this dialog?')
//     .then(() => {
//       dialogVisible.value = false
//       done()
//     })
//     .catch(() => {
//       // catch error
//     })
// }
//定义上传文件大小
// const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
//   // dialogVisible.value = true
//   console.log('rawFile')
//   return true
//   // if (rawFile.type !== 'image/jpeg') {
//   //   ElMessage.error('Avatar picture must be JPG format!')
//   //   return false
//   // } else if (rawFile.size / 1024 / 1024 > 2) {
//   //   ElMessage.error('Avatar picture size can not exceed 2MB!')
//   //   return false
//   // }
//   // return true
// }
//上传到后端
const onSubmit = () => {
  let data = {
    token: cookie.get('token'),
    name: book.name,
    description: book.description,
    coverId: '',
    authorId: book.authorId
  }
  console.log('+++++++++++++++++++++', bloArray.value)
  const mycroimg = newcropimg.value
  console.log('123123123131221213123', mycroimg)
  const params = new URLSearchParams();
  const forms = new FormData()
  forms.append('file', limgs.value)
  console.log(forms)
  console.log(limgs.value)
  axios({
    method: 'post',
    url: '/api/v1/assets/upload/image',
    data: {
      forms
    },
    params: {
      token: data.token,
      type: 'cover'
    },
    headers: {
      'Content-Type': 'multipart/form-data',
    }
  }).then(
    (e) => {
      console.log(e)
    }
  ).catch(
    (err) => {
      console.log(err)
    }
  )
}
</script>

<style lang="scss" scoped>
.box-card {
  width: 80%;
  left: 20%;

  margin: auto;
}

.fromVa {
  width: 40vw;
  margin-bottom: 20px;
  // position: relative;
  // top: 10px;
  // left: 5px;
}

.fromIt {
  text-align: center;
  position: relative;
  left: calc(50% - 40vw);
}

.butns {
  // position: relative;
  // left: calc(50% - 20vw);
  margin-top: 10vh;
  margin-bottom: 10vh;

}

.image {
  height: 202px;
}

.sampleCard {
  position: absolute;
  width: 200px;
  right: 11vw;
  bottom: 40vh;
}

.sampName {
  font-weight: bold;
  position: absolute;
  left: 20px;
  bottom: 40px;
}

.sampAut {
  position: absolute;
  bottom: 10px;
  font-size: 8px;
  right: 20px;
  color: #ababab;
}

.newchapter {
  position: absolute;
  bottom: 20px;
  left: 20px;
  font-size: 5px;
}

.inputbtn {
  width: 15vw;
  margin-right: 200px;
}
</style>