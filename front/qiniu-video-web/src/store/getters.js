import variables from '@/styles/variables.module.scss'

const getters = {
  cssVar: state => variables,

  videoData: state => state.video.data,
  total: state => state.video.total,
  base: state => state.video.base
}
export default getters
