import { Button } from "antd";

type rowEnabledType = {
  isEnabled?: boolean,
  disabled?: boolean,
  children?: React.ReactNode
}

//列表禁用/启用按钮
export const RowEnabledButton: React.FC<rowEnabledType> = (props) => {
  const { isEnabled, disabled, children } = props;

  return (
    <Button {...props} type='primary' disabled={!!disabled} danger={!!!isEnabled} size='small'>{!!isEnabled ? '生效中' : '已禁用'}</Button>
  );
}
