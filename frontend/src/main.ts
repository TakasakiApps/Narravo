import { createApp } from 'vue'
import "./style.css"
import App from './App.vue'
import './samples/node-api'
import 'vant/lib/index.css'
import router from "./Router/index"
createApp(App)
  .use(router)
  .mount('#app')
  .$nextTick(() => {
    postMessage({ payload: 'removeLoading' }, '*')
  })
