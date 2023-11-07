import variables from '@/styles/variables.module.scss'

const getters = {
  cssVar: state => variables,

  videoData: state => state.video.data,
  total: state => state.video.total,
  base: state => state.video.base,

  token: state => state.user.token
}
export default getters
