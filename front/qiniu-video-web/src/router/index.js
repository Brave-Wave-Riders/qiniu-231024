import { createRouter, createWebHashHistory } from 'vue-router'
import layout from '@/layout/index'

const routes = [
  {
    path: '/',
    component: () => layout,
    redirect: '/discover',
    children: [
      {
        path: '/discover',
        name: 'discover',
        component: () =>
          import('@/views/discover/index'),
        meta: {
          title: '首页',
          icon: 'el-icon-house'
        }
      },
      {
        path: '/404',
        name: '404',
        component: () =>
          import(/* webpackChunkName: "error-page" */ '@/views/error-page/404')
      },
      {
        path: '/401',
        name: '401',
        component: () =>
          import(/* webpackChunkName: "error-page" */ '@/views/error-page/401')
      }
    ]
  },
  {
    path: '/video/:id',
    component: () =>
      import('@/views/video-detail/index'),
    meta: {
      title: 'videoDetail'
    }
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
