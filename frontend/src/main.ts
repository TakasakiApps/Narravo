import { createApp } from 'vue'
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


const app = createApp(App);
app.use(router)
.use(contextmenu)
.use(VueMasonryPlugin)
.mount('#app').$nextTick(() => postMessage({ payload: 'removeLoading' }, '*'))
