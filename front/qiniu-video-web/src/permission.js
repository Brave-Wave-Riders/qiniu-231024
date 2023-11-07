import router from './router'

/**
 * 路由前置守卫
 */
router.beforeEach(async (to, from, next) => {
  // if (to.name === from.name && to.params.type !== from.params.type) {
  //   next({ name: 'empty', query: { toPath: to.fullPath } })
  // } else {
  //   next()
  // }
})
