import { createApp,reactive } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
// 右键插件
import contextmenu from "v-contextmenu"
import "v-contextmenu/dist/themes/default.css"
//瀑布插件
import {VueMasonryPlugin} from 'vue-masonry';
// 增加暗黑模式样式文件
import 'element-plus/theme-chalk/dark/css-vars.css'
//导入axios全局配置文件
import './http/client'


const globalVars = reactive({
    isInfo: false,
    isHidden:false,//进入阅读页面为true
    isRight:0//1:右箭头，2：左箭头
})


router.beforeResolve(to =>{
    // console.log(to.path);
    if(to.path == '/info' ){
        globalVars.isInfo = true
    }else if(to.path == '/read'){
        globalVars.isHidden = true
        globalVars.isInfo = true
    }
    else{
        globalVars.isInfo = false;
        globalVars.isHidden = false
       
    }
    
})
const app = createApp(App);
app.use(router)
.use(contextmenu)
.use(VueMasonryPlugin)
.provide('globalVars', globalVars)
.mount('#app').$nextTick(() => postMessage({ payload: 'removeLoading' }, '*'))
