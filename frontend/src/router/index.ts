/*
    1.导入路由配置模块
    2.定义路由规则
    3.创建路由实例
    4.导出路由模块
*/
//1.导入模块只有一个时不加括号
import { createRouter, createWebHistory } from 'vue-router'

import BookShelf from '../views/BookShelf.vue'
import Writer from '../views/Writer.vue'
import Setting from '../views/Setting.vue'
import Search from '../views/Search.vue'
import Bookinfo from '../views/Bookinfo.vue'
import Addrole from '../views/Addrole.vue'
//2.
const routes = [
    {
        path: '/',
        component: BookShelf,
    },
    {
        path: '/writer',
        component: Writer
    },
    {
        path:'/setting',
        component: Setting
    },
    {
        path:'/search',
        component: Search
    },
    {
        path:'/info',
        component:Bookinfo
    },
    {
        path:'/role',
        component:Addrole
    }
    

]

//3.
const router = createRouter({
    history: createWebHistory(),
    routes,
})

//4.
export default router