import { createStore } from 'vuex'
import getters from './getters'
import video from './modules/video'

export default createStore({
  getters,
  modules: {
    video
  }
})
