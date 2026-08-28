import type { Plugin } from 'vite';
import type { Connect } from 'vite';
import { mockEntries } from './index';

/** 模拟网络延迟（毫秒） */
const LATENCY = 300;

export interface MockRequest {
  body: Record<string, any>;
  query: URLSearchParams;
  headers: any;
}

export interface MockEntry {
  url: string;
  method: 'GET' | 'POST' | 'PUT' | 'DELETE';
  response: (req: MockRequest) => any;
}

/**
 * 自定义 mock 插件：在开发服务器中拦截 /api 请求并返回模拟数据
 */
export function createMockPlugin(enabled: boolean): Plugin {
  return {
    name: 'admin-vue3x-mock',
    apply: 'serve',
    configureServer(server) {
      if (!enabled) return;
      const handler: Connect.NextHandleFunction = (req, res, next) => {
        const url = (req.url || '').split('?')[0];
        const method = (req.method || 'GET').toUpperCase();
        const entry = mockEntries.find(
          (e) => e.url === url && e.method.toUpperCase() === method,
        );
        if (!entry) {
          next();
          return;
        }
        let rawBody = '';
        req.on('data', (chunk: { toString(encoding?: string): string }) => {
          rawBody += chunk.toString('utf-8');
        });
        req.on('end', () => {
          const query = new URLSearchParams((req.url || '').split('?')[1] || '');
          let body: Record<string, any> = {};
          try {
            body = rawBody ? JSON.parse(rawBody) : {};
          } catch {
            body = {};
          }
          let data: any;
          try {
            data = entry.response({ body, query, headers: req.headers });
          } catch (err) {
            data = { code: 500, msg: (err as Error).message || 'mock 处理出错', data: null };
          }
          setTimeout(() => {
            res.statusCode = 200;
            res.setHeader('Content-Type', 'application/json; charset=utf-8');
            res.end(JSON.stringify(data));
          }, LATENCY);
        });
      };
      server.middlewares.use(handler);
    },
  };
}
