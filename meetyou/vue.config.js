const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  // 代理模式，解决前端直接请求Go接口跨域问题
  devServer: {
    proxy: {
      '/api': {
        target: 'http://go-app:8080', // Go 服务器地址
        changeOrigin: true,
        pathRewrite: {
          '^/api': '/api' // 把 /api 替换成 "/api" 字符串
        }
      }
    }
  }
})
