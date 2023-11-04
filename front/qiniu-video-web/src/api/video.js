import request from '@/utils/request'

/**
 * 获取视频列表
 */
export const getVideoList = data => {
  return request({
    url: '/api/video/list',
    method: 'POST',
    data
  })
}
