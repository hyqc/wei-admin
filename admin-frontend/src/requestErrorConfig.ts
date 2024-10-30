import type { RequestOptions } from '@@/plugin-request/request';
import type { RequestConfig } from '@umijs/max';
import { message, notification } from 'antd';
import { GetLoginToken, IsLongPage } from './utils/common';
import { SUCCESS } from './services/apis/code';

// 错误处理方案： 错误类型
enum ErrorShowType {
  SILENT = 0,
  WARN_MESSAGE = 1,
  ERROR_MESSAGE = 2,
  NOTIFICATION = 3,
  REDIRECT = 9,
}
// 与后端约定的响应数据格式
interface ResponseStructure {
  code: number;
  message: string;
  data?: any;
  type?: ErrorShowType,
}

const isDev = process.env.NODE_ENV === 'development';

/**
 * @name 错误处理
 * pro 自带的错误处理， 可以在这里做自己的改动
 * @doc https://umijs.org/docs/max/request#配置
 */
export const errorConfig: RequestConfig = {
  // 错误处理： umi@3 的错误处理方案。
  errorConfig: {
    // 错误抛出
    errorThrower: (res) => {
      const { code, data, message, type } = res as unknown as ResponseStructure;
      if (code !== SUCCESS ) {
        const error: any = new Error(message);
        error.name = 'BizError';
        error.info = { code, message, type, data };
        throw error; // 抛出自制的错误
      }
    },
    // // 错误接收及处理
    // errorHandler: (error: any, opts: any) => {
    //   console.log('=======', error, opts)
    //   if (opts?.skipErrorHandler) throw error;
    //   // 我们的 errorThrower 抛出的错误。
    //   if (error.name === 'BizError') {
    //     const errorInfo: ResponseStructure | undefined = error.info;
    //     if (errorInfo) {
    //       switch (errorInfo.type) {
    //         case ErrorShowType.SILENT:
    //           // do nothing
    //           break;
    //         case ErrorShowType.WARN_MESSAGE:
    //           message.warning(errorInfo.message);
    //           break;
    //         case ErrorShowType.ERROR_MESSAGE:
    //           message.error(errorInfo.message);
    //           break;
    //         case ErrorShowType.NOTIFICATION:
    //           notification.open({
    //             description: errorInfo.message,
    //             message: errorInfo.message,
    //           });
    //           break;
    //         case ErrorShowType.REDIRECT:
    //           // TODO: redirect
    //           break;
    //         default:
    //           message.error(errorInfo.message);
    //       }
    //     }
    //   } else if (error.response) {
    //     // Axios 的错误
    //     // 请求成功发出且服务器也响应了状态码，但状态代码超出了 2xx 的范围
    //     message.error(`Response status:${error.response.status}`);
    //   } else if (error.request) {
    //     // 请求已经成功发起，但没有收到响应
    //     // \`error.request\` 在浏览器中是 XMLHttpRequest 的实例，
    //     // 而在node.js中是 http.ClientRequest 的实例
    //     message.error('None response! Please retry.');
    //   } else {
    //     // 发送请求时出了点问题
    //     message.error('Request error, please retry.');
    //   }
    // },
  },

  // 请求拦截器
  requestInterceptors: [
    (config: RequestOptions) => {
      // 拦截请求配置，进行个性化处理。
      const url = config?.url;
      const realyUrl = `${BaseAPI}${url}`;
      console.log('请求拦截器：', BaseAPI, url, realyUrl);
      let headers = config.headers
      if (!IsLongPage()) {
        const tokenInfo = GetLoginToken();
        const token = tokenInfo !== undefined ? tokenInfo.token : '';
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        headers = { ...config.headers, 'Authorization': 'Bearer ' + token }

      }

      return {
        ...config,
        headers,
        url: realyUrl
      };
    },
  ],

  // 响应拦截器
  responseInterceptors: [
    (response) => {
      // 拦截响应数据，进行个性化处理
      console.log('返回数据：', response.data);
      const data: any = response.data;
      return new Promise((resolve, reject) => {
        message.destroy();
        if (response.status !== 200) {
          const msg: string = data.msg;
          message.error(msg, MessageDuritain);
          // eslint-disable-next-line no-promise-executor-return
          return reject(msg);
        }else if (data.code !== SUCCESS) {
          message.error(data.msg, MessageDuritain);
          // eslint-disable-next-line no-promise-executor-return
          return reject(data.msg);
        }
        // eslint-disable-next-line no-promise-executor-return
        return resolve(response);
      });
    },
  ],
};
