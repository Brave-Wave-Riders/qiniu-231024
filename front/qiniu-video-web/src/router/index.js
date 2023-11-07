import { createRouter, createWebHashHistory } from 'vue-router'
import layout from '@/layout/index'

const routes = [
  {
    path: '/',
    component: () => layout,
    redirect: '/discover/0',
    children: [
      {
        path: '/discover/:id',
        name: 'discover',
        component: () =>
          import('@/views/discover/index'),
        meta: {
          title: '首页',
          icon: 'el-icon-house'
        }
      },
      {
        path: '/video/:id',
        component: () =>
          import('@/views/video-detail/index'),
        meta: {
          title: 'videoDetail'
        }
      },
      {
        path: '/video/upload',
        component: () =>
          import('@/views/video-upload/index'),
        meta: {
          title: 'videoUpload'
        }
      },
      // 个人中心
      {
        path: '/profile',
        name: 'profile',
        redirect: '/discover/0',
        component: () =>
          import('@/views/profile/index'),
        meta: {
          title: '我的',
          icon: 'el-icon-user'
        }
      },
      {
        path: '/knowledge',
        name: 'knowledge',
        redirect: '/discover/1',
        meta: {
          title: '知识',
          icon: 'el-icon-reading'
        }
      },
      {
        path: '/game',
        name: 'game',
        redirect: '/discover/2',
        meta: {
          title: '游戏',
          icon: 'el-icon-mobile-phone'
        }
      },
      {
        path: '/enjoy',
        name: 'enjoy',
        redirect: '/discover/3',
        meta: {
          title: '娱乐',
          icon: 'el-icon-star-off'
        }
      },
      {
        path: '/music',
        name: 'music',
        redirect: '/discover/4',
        meta: {
          title: '音乐',
          icon: 'el-icon-headset'
        }
      },
      {
        path: '/food',
        name: 'food',
        redirect: '/discover/5',
        meta: {
          title: '美食',
          icon: 'el-icon-tableware'
        }
      },
      {
        path: '/sport',
        name: 'sport',
        redirect: '/discover/6',
        meta: {
          title: '体育',
          icon: 'el-icon-football'
        }
      },
      {
        path: '/fade',
        name: 'fade',
        redirect: '/discover/7',
        meta: {
          title: '时尚',
          icon: 'el-icon-s-goods'
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
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
