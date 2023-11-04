const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true
})
module.exports = {
  devServer: {
    // 配置反向代理
    proxy: {
      // 当地址中有/api的时候会触发代理机制
      '/api': {
        // 要代理的服务器地址  这里不用写 api
        target: 'http://localhost:8888/',
        changeOrigin: true // 是否跨域
      },
      '/hls': {
        // 要代理的服务器地址  这里不用写 api
        target: 'http://localhost',
        changeOrigin: true, // 是否跨域
        pathRewrite: {
          '/hls': '/'
        }
      }
    }
  },
}