import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import dns from 'dns'
dns.setDefaultResultOrder('ipv4first')

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
  plugins: [vue()]
})
