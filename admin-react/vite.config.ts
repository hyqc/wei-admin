import { defineConfig, loadEnv } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path';
import { createMockPlugin } from './src/mock/plugin';

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '');
  const useMock = env.VITE_USE_MOCK === 'true';

  return {
    plugins: [react(), createMockPlugin(useMock)],
    resolve: {
      alias: {
        '@': path.resolve(__dirname, 'src'),
      },
    },
    server: {
      host: '0.0.0.0',
      port: 8001,
      open: false,
      proxy: {
        // 未启用 mock 时，将 /api 代理到真实后端（后端路由无 /api 前缀，转发时需移除）
        '/api': {
          target: 'http://127.0.0.1:3000',
          changeOrigin: true,
          rewrite: (p: string) => p.replace(/^\/api/, ''),
        },
        // 本地存储返回 /upload 相对路径，代理到后端静态目录展示（后端路由无 /api 前缀，不重写）
        '/upload': {
          target: 'http://127.0.0.1:3000',
          changeOrigin: true,
        },
      },
    },
    build: {
      rollupOptions: {
        output: {
          // 大依赖单独分包，避免单 chunk 过大
          manualChunks: {
            react: ['react', 'react-dom', 'react-router', 'react-router-dom'],
            antd: ['antd', '@ant-design/icons'],
          },
        },
      },
    },
  };
});
