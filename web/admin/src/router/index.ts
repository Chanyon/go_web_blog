import {createRouter,createWebHistory} from 'vue-router';

const meta = {
  title: 'Blog后台管理系统',
}
const routes = [
  { 
    path: '/',
    name: 'admin',   
    meta:meta,
    component: ()=>import("../views/Admin.vue"),
    children: [
      {
        path: '/index',
        name: 'index',
        component: ()=>import("../components/admin/Index.vue"),
        meta: {
          title: 'Blog后台管理页面'
        }
      },
      {
        path: '/catelist',
        name: 'category',
        component: ()=>import("../components/category/CategoryList.vue"),
        meta: {
          title: 'Blog后台管理页面'
        }
      },
      {
        path: '/addart',
        name: 'addArticle',
        component: ()=>import("../components/article/AddArticle.vue"),
        meta: {
          title: 'Blog后台管理页面-新增文章'
        }
      },
      {
        path: '/addart/:id',
        name: 'editArticle',
        component: ()=>import("../components/article/AddArticle.vue"),
        meta: {
          title: 'Blog后台管理页面-文章编辑'
        },
        props: true
      },
      {
        path: '/artlist',
        name: 'ArticleList',
        component: ()=>import("../components/article/ArticleList.vue"),
        meta: {
          title: 'Blog后台管理页面'
        }
      },
      {
        path: '/profile',
        name: 'profile',
        component: ()=>import("../components/user/Profile.vue"),
        meta: {
          title: 'Blog后台管理页面'
        }
      },
      {
        path: '/userlist',
        name: 'userList',
        component: ()=>import("../components/user/UserList.vue"),
        meta: {
          title: 'Blog后台管理页面'
        }
      },
      {
        path: '/commentlist',
        name: 'commentList',
        component: ()=>import("../components/comment/commentList.vue"),
        meta: {
          title: 'Blog后台管理页面'
        }
      },
    ],
  },
  {
    path: '/login',
    name: 'Login',
    component: ()=>import("../views/Login.vue"),
    meta: {
      title: 'Blog后台登录页面'
    }
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to,from,next) =>{
  if(to.meta.title){
    document.title = to.meta.title as string;
  }
  const userToken = sessionStorage.getItem('token')
  if(!userToken){
    if(to.path === '/login'){
      next();
    }else{
      next('/login');
    }
  }else{
    next();
  }
}) 

export default router