import { createStore } from 'vuex'
import getters from './getters'
import video from './modules/video'
import user from './modules/user'

export default createStore({
  getters,
  modules: {
    video,
    user
  }
})
