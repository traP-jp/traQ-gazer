import { defineConfig } from 'vite'
import path from 'path'
import vue from '@vitejs/plugin-vue'
import dns from 'dns'
dns.setDefaultResultOrder('ipv4first')

const srcPath = path.resolve(import.meta.dirname, 'src').replace(/\\/g, '/')

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:4010',
        rewrite: (path) => (path.startsWith('/api') ? path.slice('/api'.length) : path),
        changeOrigin: true
      }
    }
  },
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: `
        @import "${srcPath}/styles/color.scss";
        @import "${srcPath}/styles/mixin.scss";
        `
      }
    }
  },
  plugins: [vue()]
})
