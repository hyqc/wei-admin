import { Select } from 'antd';
import { antIconNames, antIcons } from '@/utils/icon';

interface IconSelectProps {
  value?: string;
  placeholder?: string;
  disabled?: boolean;
  onChange?: (value?: string) => void;
}

/** ant 图标下拉选择（支持按名称搜索） */
export default function IconSelect({ value, placeholder, disabled, onChange }: IconSelectProps) {
  return (
    <Select
      value={value}
      placeholder={placeholder}
      disabled={disabled}
      listHeight={320}
      showSearch
      allowClear
      filterOption={(input, option) =>
        String((option?.value as string) ?? '')
          .toLowerCase()
          .includes(input.toLowerCase())
      }
      onChange={onChange}
    >
      {antIconNames.map((name) => {
        const Icon = antIcons[name];
        return (
          <Select.Option key={name} value={name}>
            <span className="icon-option">
              {Icon ? <Icon /> : null}
              <span className="icon-option-name">{name}</span>
            </span>
          </Select.Option>
        );
      })}
    </Select>
  );
}
