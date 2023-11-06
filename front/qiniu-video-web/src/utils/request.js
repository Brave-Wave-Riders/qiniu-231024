import axios from 'axios'
// import store from '@/store'
import { ElMessage } from 'element-plus'

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 5000
})

// // 请求拦截器
// service.interceptors.request.use(
//   config => {
//     // 在这个位置需要统一的去注入token
//     if (store.getters.token) {
//       // if (isCheckTimeout()) {
//       //   // 登出操作
//       //   store.dispatch('user/logout')
//       //   return Promise.reject(new Error('token 失效'))
//       // }
//       // 如果token存在 注入token
//       config.headers.Authorization = `Bearer ${store.getters.token}`
//     }
//     // 配置接口国际化
//     config.headers['Accept-Language'] = store.getters.language
//     return config // 必须返回配置
//   },
//   error => {
//     return Promise.reject(error)
//   }
// )

// 响应拦截器
service.interceptors.response.use(
  response => {
    const { code, errmsg, data } = response.data
    //   要根据success的成功与否决定下面的操作
    if (code === 0) {
      return data
    } else {
      // 业务错误
      ElMessage.error(errmsg) // 提示错误消息
      return Promise.reject(new Error(errmsg))
    }
  },
  error => {
    // 处理 token 超时问题
    if (
      error.response &&
      error.response.data &&
      error.response.data.code === 401
    ) {
      // token超时
      // store.dispatch('user/logout')
    }
    ElMessage.error(error.message) // 提示错误信息
    return Promise.reject(error)
  }
)

export default service
