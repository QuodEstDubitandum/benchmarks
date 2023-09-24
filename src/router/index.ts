import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/loop',
      name: 'loop',
      component: HomeView
    },
    {
      path: '/serialization',
      name: 'serialization',
      component: HomeView
    },
    {
      path: '/file',
      name: 'file',
      component: HomeView
    },
    {
      path: '/db-query',
      name: 'db-query',
      component: HomeView
    },
    {
      path: '/db-update',
      name: 'db-update',
      component: HomeView
    }
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue')
    // }
  ]
})

export default router
