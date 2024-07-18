import { useModel } from 'umi';

export type AuthorizationType = {
  name?: string;
  children?: React.ReactNode;
  forbidden?: React.ReactNode;
};

const Authorization: React.FC<AuthorizationType> = (props) => {
  const { name, children, forbidden } = props;
  const { initialState } = useModel('@@initialState');
  const permissions = initialState?.permissions;
  if (permissions && name && permissions[name]) {
    return <>{children}</>;
  }
  return <>{forbidden}</>;
};

export default Authorization;
