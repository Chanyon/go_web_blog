import { createRouter, createWebHistory } from 'vue-router';

const ArticleList = () => import('../components/ArticleList.vue')
const Details = () => import('../components/Details.vue');

const routes = [
  {
    path:'/',
    name: 'articleList',
    component: ArticleList,
  },
  {
    path: "/search/:title",
    name: "search",
    component: () => import("../components/Search.vue"),
    props: true
  },
  {
    path: "/article/category/:cid",
    name: "article",
    component: ArticleList,
    props: true,
  },
  {
    path: "/article/details/:id",
    name: "details",
    component: Details,
    props: true,
  }
];
const router = createRouter({
  history: createWebHistory(),
  routes,
  strict: true,
});

export default router;