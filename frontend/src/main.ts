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
//图标库
import { library } from '@fortawesome/fontawesome-svg-core'

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

/* import specific icons */
import { fas } from '@fortawesome/free-solid-svg-icons'

/* add icons to the library */
library.add(fas)

const app = createApp(App);
app.use(router)
.component('font-awesome-icon', FontAwesomeIcon)
.use(contextmenu)
.use(VueMasonryPlugin)

.mount('#app').$nextTick(() => postMessage({ payload: 'removeLoading' }, '*'))
