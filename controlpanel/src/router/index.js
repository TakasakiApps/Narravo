// import Login from ''
import { createRouter, createWebHashHistory } from 'vue-router';
import login from '../components/Login.vue'
import Sysviem from '../components/Sysviem.vue'
import Adminpw from '../components/Adminpw.vue'
import AddBook from '../components/AddBook.vue'
import AddDesc from '../components/AddDesc.vue'
// import RolaCard from '../components/layouts/RoleCard.vue'
import { ElMessage } from 'element-plus'
import VueCookies from 'vue-cookies'
import Editpas from '../components/Editpas.vue'
const cookie = VueCookies


const routes  =[
  {
    path: '/',
    name: 'login',
    component: login,
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: Sysviem,

  },
  {
    path:'/Admain',
    name:'admin',
    component:Adminpw
  },
  {
    path:'/Addbook',
    name:'AddBook',
    component:AddBook
  },
  {
    path:'/AddDesc',
    name:'AddDesc',
    component:AddDesc
  },
  {
    path:'/Editpas',
    name:'Editpas',
    component:Editpas
  }
  // {
  //   path:'/RolaCard',
  //   name:"RolaCard",
  //   component:RolaCard
  // }

]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  // const token = localStorage.token; // 假设存储在 localStorage 中
  const token = cookie.get('token')
  console.log(token)
  if (to.name !== 'login' && !token) {
    // 如果目标页面不是登录页且没有 token，则阻止跳转并跳转到登录页面
    next({ name: 'login' });
    // 或者使用原生 alert 进行提示
    ElMessage({
      type:'warning',
      message:'请登录'
    })

  } else {
    // 有 token 或目标页面为登录页，则允许跳转
    next();
  }
});
export default router;