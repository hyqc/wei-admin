import { Switch } from 'antd';

interface RowEnabledButtonProps {
  isEnabled?: boolean;
  isEnabledButtonDisabled?: boolean;
  onChange?: (checked: boolean) => void;
}

/** 列表行的启用/禁用开关 */
export default function RowEnabledButton({
  isEnabled,
  isEnabledButtonDisabled,
  onChange,
}: RowEnabledButtonProps) {
  return (
    <Switch
      checked={isEnabled}
      disabled={isEnabledButtonDisabled}
      checkedChildren="启用"
      unCheckedChildren="禁用"
      onChange={onChange}
    />
  );
}
