import { createApp } from 'vue'
import './style.css';
import App from './App.vue'
import router from './router'
//axios全局配置文件
import '../network/axios.config'
// 右键插件
import contextmenu from "v-contextmenu"
import "v-contextmenu/dist/themes/default.css"
//瀑布插件
import {VueMasonryPlugin} from 'vue-masonry';
import 'element-plus/theme-chalk/dark/css-vars.css'

const app = createApp(App);

app.use(router)
.use(contextmenu)
.use(VueMasonryPlugin)
.mount('#app');
// .$nextTick(() => postMessage({ payload: 'removeLoading' }, '*'))
