import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  envDir: './src/env',
  resolve: {
    alias: {
      '@': resolve(__dirname, './src')
    }
  }
  // server: {
  //   proxy: {
  //     '/api': {
  //       target: 'http://127.0.0.1:5656',
  //       changeOrigin: true,
  //       rewrite: (path) => path.replace(/^\/api/, '')
  //     }
  //   }
  // }
})
// .env.dev
// .env.sit
// .env.prod
