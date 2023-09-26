import { createRouter, createWebHistory } from 'vue-router'
import DataTable from '../components/DataTable.vue'
import UsageGuide from '../components/UsageGuide.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: UsageGuide
    },
    {
      path: '/loop',
      name: 'loop',
      component: DataTable
    },
    {
      path: '/serialization',
      name: 'serialization',
      component: DataTable
    },
    {
      path: '/file',
      name: 'file',
      component: DataTable
    },
    {
      path: '/db-query',
      name: 'db-query',
      component: DataTable
    },
    {
      path: '/db-update',
      name: 'db-update',
      component: DataTable
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
