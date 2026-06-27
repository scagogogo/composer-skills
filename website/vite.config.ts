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
          'highlighter': ['react-syntax-highlighter'],
          'react': ['react', 'react-dom'],
        },
      },
    },
  },
})
