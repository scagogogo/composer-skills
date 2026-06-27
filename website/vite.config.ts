import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  base: '/composer-skills/',
  plugins: [react()],
  build: {
    outDir: 'dist',
    chunkSizeWarningLimit: 1000,
    rollupOptions: {
      output: {
        manualChunks: {
          'antd': ['antd', '@ant-design/icons'],
          'charts': ['@ant-design/charts'],
          'highlighter': ['react-syntax-highlighter'],
          'motion': ['framer-motion'],
          'react': ['react', 'react-dom'],
        },
      },
    },
  },
})
