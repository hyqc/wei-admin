import { Footer, Question, SelectLang, AvatarDropdown, AvatarName } from '@/components';
import { LinkOutlined } from '@ant-design/icons';
import type { Settings as LayoutSettings, MenuDataItem } from '@ant-design/pro-components';
import { SettingDrawer } from '@ant-design/pro-components';
import type { RunTimeLayoutConfig } from '@umijs/max';
import { history, Link } from '@umijs/max';
import defaultSettings from '../config/defaultSettings';
import { errorConfig } from './requestErrorConfig';
import React from 'react';
import { currentAdminInfo } from './services/apis/admin/account';
import { GetLoginToken, HandleMenusToMap, HandleRemoteMenuIntoLocal, IsLogin, IsLongPage, Logout, MenusMapType } from './utils/common';
import { AdminInfo, RespAccountInfoData, RespAccountPermissionData } from './proto/admin_ts/admin_account';
const isDev = process.env.NODE_ENV === 'development';

/**
 * @see  https://umijs.org/zh-CN/plugins/plugin-initial-state
 * */
export async function getInitialState(): Promise<{
  settings?: Partial<LayoutSettings>;
  currentUser?: AdminInfo;
  permissions?: Map<string, string>;
  menuData?: MenusMapType;
  fetchUserInfo?: () => Promise<RespAccountInfoData | undefined>;
}> {
  //定义获取账号详情的请求方法
  const fetchUserInfo = async () => {
    try {
      const tokenInfo = GetLoginToken();
      if (tokenInfo?.token) {
        const res = await currentAdminInfo(tokenInfo?.remember);
        return res.data;
      }
      return undefined;
    } catch (error) {
      console.log('=========', error)
      //history.push(LoginPath);
    }
    return undefined;
  };
  // 如果不是登录页面，执行
  if (!IsLongPage()) {
    const currentUserData: RespAccountInfoData = await fetchUserInfo();
    const currentUser = currentUserData?.data;
    const permissions = { ...currentUser.permissions };
    return {
      fetchUserInfo,
      currentUser,
      permissions,
      settings: defaultSettings as Partial<LayoutSettings>,
    };
  }
  return {
    fetchUserInfo,
    settings: defaultSettings as Partial<LayoutSettings>,
  };
}

// ProLayout 支持的api https://procomponents.ant.design/components/layout
export const layout: RunTimeLayoutConfig = ({ initialState, setInitialState }) => {
  return {
    actionsRender: () => [<Question key="doc" />, <SelectLang key="SelectLang" />],
    avatarProps: {
      src: initialState?.currentUser?.avatar,
      title: <AvatarName />,
      render: (_, avatarChildren) => {
        return <AvatarDropdown>{avatarChildren}</AvatarDropdown>;
      },
    },
    waterMarkProps: {
      content: initialState?.currentUser?.name,
    },
    footerRender: () => <Footer />,
    onPageChange: () => {
      const { location } = history;
      // 如果没有登录，重定向到 login
      // 如果没有登录，重定向到 login
      if (location.pathname !== LoginPath && !IsLogin(initialState)) {
        return Logout();
      }
    },
    bgLayoutImgList: [
      {
        src: 'https://mdn.alipayobjects.com/yuyan_qk0oxh/afts/img/D2LWSqNny4sAAAAAAAAAAAAAFl94AQBr',
        left: 85,
        bottom: 100,
        height: '303px',
      },
      {
        src: 'https://mdn.alipayobjects.com/yuyan_qk0oxh/afts/img/C2TWRpJpiC0AAAAAAAAAAAAAFl94AQBr',
        bottom: -68,
        right: -45,
        height: '303px',
      },
      {
        src: 'https://mdn.alipayobjects.com/yuyan_qk0oxh/afts/img/F6vSTbj8KpYAAAAAAAAAAAAAFl94AQBr',
        bottom: 0,
        left: 0,
        width: '331px',
      },
    ],
    links: isDev
      ? [
        <Link key="openapi" to="/umi/plugin/openapi" target="_blank">
          <LinkOutlined />
          <span>OpenAPI 文档</span>
        </Link>,
      ]
      : [],
    menuItemRender: (menuItemProps, defaultDom) => {
      // if (menuItemProps.isUrl) {
      //   return defaultDom;
      // }
      // 支持二级菜单显示icon
      const styleSpan = { display: 'inline-block', marginRight: '2px' };
      return (
        <Link to={menuItemProps.path}>
          <span style={styleSpan}>
            {menuItemProps.pro_layout_parentKeys &&
              menuItemProps.pro_layout_parentKeys.length > 0 &&
              menuItemProps.icon}
          </span>
          <span style={styleSpan}> {defaultDom}</span>
        </Link>
      );
    },
    menuHeaderRender: undefined,
    menu: {
      locale: true,
      defaultOpenAll: true,
      request: (_params: any, defaultMenuData: MenuDataItem[]) => {
        const menuData = initialState?.currentUser?.menus;
        const tmpMenuList: MenuDataItem[] = HandleRemoteMenuIntoLocal(
          [],
          defaultMenuData,
          menuData,
          'children',
        );
        const menuList: MenuDataItem[] = HandleRemoteMenuIntoLocal(
          [],
          tmpMenuList,
          menuData,
          'routes',
        );
        setInitialState({
          ...initialState,
          menuData: HandleMenusToMap({}, menuList, 'children'),
        });
        return new Promise((resolve) => {
          resolve(menuList);
        });
      },
    },
    // 自定义 403 页面
    // unAccessible: <div>unAccessible</div>,
    // 增加一个 loading 的状态
    childrenRender: (children) => {
      // if (initialState?.loading) return <PageLoading />;
      return (
        <>
          {children}
          {isDev && (
            <SettingDrawer
              disableUrlParams
              enableDarkTheme
              settings={initialState?.settings}
              onSettingChange={(settings) => {
                setInitialState((preInitialState) => ({
                  ...preInitialState,
                  settings,
                }));
              }}
            />
          )}
        </>
      );
    },
    ...initialState?.settings,
  };
};

/**
 * @name request 配置，可以配置错误处理
 * 它基于 axios 和 ahooks 的 useRequest 提供了一套统一的网络请求和错误处理方案。
 * @doc https://umijs.org/docs/max/request#配置
 */
export const request = {
  ...errorConfig,
};
