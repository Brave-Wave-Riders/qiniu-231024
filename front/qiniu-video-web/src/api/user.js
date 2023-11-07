import request from '@/utils/request'

/**
 * 登录
 */
export const login = data => {
  return request({
    url: '/api/user/v1/login',
    method: 'POST',
    data
  })
}
/**
 * 获取用户信息
 */
export const getUserInfo = () => {
  return request({
    url: '/api/user/v1/profile'
  })
}

/**
 * 登录
 */
export const register = data => {
  return request({
    url: '/api/user/v1/register',
    method: 'POST',
    data
  })
}
