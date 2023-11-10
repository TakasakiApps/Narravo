/*
    1.导入路由配置模块
    2.定义路由规则
    3.创建路由实例
    4.导出路由模块
*/
//1.导入模块只有一个时不加括号
import { createRouter, createWebHistory } from 'vue-router'

// 书架
import BookShelf from '../views/BookShelf.vue'
import Writer from '../views/Writer.vue'
// 设置
import Setting from '../views/Setting.vue'
// 搜索
import Search from '../views/Search.vue'
// import SearViem from '../views/Searviem.vue'
import SearviemVue from '../views/Searviem.vue'
// 书籍详情
import Bookinfo from '../views/Bookinfo.vue'
// import Addrole from '../views/Addrole.vue'
// 阅读
import Reading from '../views/Reading.vue'
import BookLiemVue from '../components/BookLiem.vue'
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
        path: '/setting',
        component: Setting
    },
    {
        path: '/search',
        component: Search
    },
    {
        path: '/info/:local/:name?/:author?/:coverId?/:mtime?/:noverid?',
        name:"info",
        component: Bookinfo,
        props: true,
        children: [

        ]
    },
    {
        path: '/read/:local/:noverid?/:name?',
        name: 'reading',
        props: true,
        component: Reading
    },
    // {
    //     path: '/role',
    //     component: Addrole
    // },
    {
        path:'/BookLiemVue',
        component:BookLiemVue
    },
    {
        path:'/SearviemVue',
        component:SearviemVue
    }
]

//3.
const router = createRouter({
    history: createWebHistory(),
    routes,
})
//路由守卫

//4.
export default router