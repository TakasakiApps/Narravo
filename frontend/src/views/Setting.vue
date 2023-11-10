<template>
    <div class="container">
        <el-upload action="" :before-upload="beforeAvatarUpload" :http-request="upLoad" :show-file-list="false">
            <el-avatar :src="avatar" fit="cover">
            </el-avatar>
        </el-upload>
        <div class="info">

        </div>
    </div>
    <el-dialog title="头像裁剪" v-model="isCaper">
        <VueCropper ref="cropper" :img="options">

        </VueCropper>
    </el-dialog>
</template>
    
<script setup lang='ts'>
import { ref, onMounted } from 'vue'
import defaultImage from '../assets/默认头像.png'
import { client } from '../http/client'
//引入裁剪插件
import 'vue-cropper/dist/index.css'
import { VueCropper } from "vue-cropper";
//头像图片
const avatar = ref(defaultImage)
//加载头像
onMounted(() => {
    if (localStorage.imgUrl) {
        avatar.value = localStorage.imgUrl
    }
})
//获取头像图片所需参数

//上传头像图片
//上传头像所要的参数
const date = {
    token: localStorage.token,
    type: 'avatar'
}
//裁剪图片参数
const isCaper = ref(false)
let options: any
let outType: any
//上传图片函数
const upLoad = (param) => {
    let formData = new FormData()
    formData.append('file', param.file)

    console.log(formData.get('file'));
    if (date.token) {
        client.post('/api/v1/assets/upload/image', formData, {
            params: {
                token: date.token,
                type: date.type
            }
        }, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        }).then(res => {
            if (res.status == 200) {
                ElMessage.success('上传头像成功')
                localStorage.imgId = res.data.id


                //上传完后加载图片
                client.get('/api/v1/assets/fetch/image', {
                    params: {
                        token: date.token,
                        type: date.type,
                        id: localStorage.imgId
                    }
                }).then(res => {
                    localStorage.imgUrl = res.request.responseURL
                    isCaper.value = true
                    options = res.request.responseURL

                    avatar.value = res.request.responseURL


                }).catch(err => {
                    console.log(err);

                })
            }
            else {
                ElMessage.error('头像上传失败')
            }
        }).catch(err => {
            console.log(err);

        })
    } else {
        ElNotification.error({ message: '登录状态已过期！请重新登录' })
    }
    console.log('触发上传事件');

}

//1.检查图片格式和大小符不符合要求 
const beforeAvatarUpload: UploadProps['beforeUpload'] = (param) => {

    if (param.size / 1024 / 1024 > 2) {
        ElNotification.error({ title: '头像太大了', message: '请上传小于2MB的头像' })
        return false
    }
    if (param.type == 'image/jpeg' || param.type == 'image/png' || param.type == 'image/gif') {

        return true
    }
    ElNotification.error({ title: '上传头像不符合格式', message: '请上传jpeg,png或gif格式的头像' })
    return false
}
//裁剪头像

</script>
    
<style scoped lang="scss">
.el-avatar {
    width: 200px;
    height: 200px;
}
</style>