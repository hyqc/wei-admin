import { PermissionsType } from '@/services/apis/types';
import { useModel } from 'umi';

export type AuthorizationType = {
  name?: string;
  children?: React.ReactNode;
  forbidden?: React.ReactNode;
};

const Authorization: React.FC<AuthorizationType> = (props) => {
  const { name, children, forbidden } = props;
  const { initialState } = useModel('@@initialState');
  const permissions: PermissionsType|undefined = initialState?.permissions;
  if (permissions !== undefined && name !== undefined && permissions[name]) {
    return <>{children}</>;
  }
  if(forbidden===undefined){
    return <></>
  }
  return <>{forbidden}</>;
};

export default Authorization;
