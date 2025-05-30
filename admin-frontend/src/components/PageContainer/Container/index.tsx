import ForbiddenPage from '@/pages/403';
import NoFoundPage from '@/pages/404';
import { PageContainer } from '@ant-design/pro-layout';
import React from 'react';
import { useModel } from 'umi';

export type ContentType = {
  wrapperStyle?: React.CSSProperties;
  cardStyle?: React.CSSProperties;
  children?: React.ReactNode;
};

const Content: React.FC<ContentType> = (props: any) => {
  const { wrapperStyle } = props;
  const { initialState } = useModel('@@initialState');
  const menuMap = initialState?.menuData || {};
  const pathname = path(location.pathname);
  let canAccessLocalMenu = true
  let notFound = false
  if(menuMap[pathname]!==undefined){
    canAccessLocalMenu = menuMap[pathname]?.access;
    notFound = false
  }else{
    notFound = true
  }
  const initCardStyle: any = () => {
    let wrapperStyless = {
      padding: '0 0 1.4rem',
    };
    if (wrapperStyle && Object.keys(wrapperStyle).length) {
      wrapperStyless = { ...wrapperStyless, ...wrapperStyle };
    }
    return wrapperStyless;
  };

  const wrapperStyless = initCardStyle();

  function path(path: string) {
    return path.length > 1 && path.substring(path.length - 1) === '/'
      ? path.substring(0, path.length - 1)
      : path;
  }
  return (
    canAccessLocalMenu ? (
      <PageContainer style={wrapperStyless}>{props?.children}</PageContainer>
    ) : (
      <ForbiddenPage />
    )
  )
};

export default Content;
