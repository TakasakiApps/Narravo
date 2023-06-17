import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import SysHomeVue from '../components/SysHome.vue';
import Commonend from '../components/Commend.vue';
import BookShelfVue from '../components/BookShelf.vue';
import CreateVue from '../components/Create.vue';
import SearchLiemVue from '../components/SearchLiem.vue';
const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Commonend',
    component: Commonend,
  },
  {
    path:'/SysHomeVue',
    name:'SysHomeVue', 
    component:SysHomeVue
  },
  {
    path:'/BookShelfVue',
    name:'BookShelfVue',
    component:BookShelfVue
  },
  {
    path:'/CreateVue',
    name:'CreateVue',
    component:CreateVue
  },
  {
    path:'/SearchLiem/:search',
    name:'SearchLiem',
    component:SearchLiemVue
  }
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;