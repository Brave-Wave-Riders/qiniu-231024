import { getItem, setItem } from '@/utils/storage'
const VIDEO = 'video_data'
const TOTAL = 'total'
const BASE = 'base'
export default {
  namespaced: true,
  state: () => ({
    data: getItem(VIDEO) || [
      { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
      { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
      { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' }],
    total: getItem(TOTAL) || 3,
    base: getItem(BASE) || 'http://s3hk0nelw.bkt.clouddn.com/'
  }),
  mutations: {
    setVideoData(state, data) {
      console.log('setVideoData: ', data)
      setItem(VIDEO, data)
      state.data = data
    },
    setVideoNum(state, total) {
      setItem(TOTAL, total)
      state.total = total
    },
    setBaseUrl(state, base) {
      console.log('setBaseUrl: ', base)
      setItem(BASE, base)
      state.base = base
    }
  }
}
