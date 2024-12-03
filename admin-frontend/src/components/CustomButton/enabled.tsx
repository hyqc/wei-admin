import { Button } from "antd";

type rowEnabledType = {
  isEnabled?: boolean,
  disabled?: boolean,
  children?: React.ReactNode
  trueText?: string,
  falseText?: string
  danger?: boolean
}

//列表禁用/启用按钮
export const RowEnabledButton: React.FC<rowEnabledType> = (props) => {
  const { isEnabled, disabled, trueText, falseText,danger } = props;
  const tText = trueText ?? '启用中'
  const fText = falseText ?? '已禁用'
  const dangerVal = danger ?? !!!isEnabled
  return (
    <Button {...props} type='primary' disabled={!!disabled} danger={dangerVal} size='small'>{!!isEnabled ? tText : fText}</Button>
  );
}
