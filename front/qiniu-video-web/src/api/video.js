import request from '@/utils/request'

/**
 * 获取视频列表
 */
export const getVideoList = (type, param) => {
  return request({
    url: `/api/video/v1/list/${type}`,
    params: param,
    method: 'GET'
  })
}

/**
 * 上传视频
 */
export const uploadVideo = data => {
  return request({
    url: '/api/video/v1/upload',
    method: 'POST',
    data
  })
}
