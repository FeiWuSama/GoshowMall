import { createRouter, createWebHistory } from 'vue-router'
// @ts-ignore
import HomeView from '@/pages/HomeView.vue'
// @ts-ignore
import AboutView from '@/pages/AboutView.vue'
// @ts-ignore
import LarkAuth from '@/pages/user/LarkAuth.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView,
    },
    {
      path: '/user/lark/auth',
      name: 'lark-auth',
      component: LarkAuth,
    },
  ],
})

export default router
