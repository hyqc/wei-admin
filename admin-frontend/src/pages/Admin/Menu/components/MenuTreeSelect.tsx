import { ResponseAdminMenuListItemType } from '@/services/apis/admin/menu';
import { Select } from 'antd';
import { useEffect, useState } from 'react';

export type AddModalPropsType = {
  data: ResponseAdminMenuListItemType[];
  value?: number;
  disabled?: boolean;
  onChange?: (parentid: number) => void;
};

const MenuTreeSelect: React.FC<AddModalPropsType> = (props) => {
  const { data, value, disabled, onChange } = props;
  const [menuId, setMenuId] = useState<number>(value || 0);
  const triggerChange = (parentid: number) => {
    onChange?.(parentid);
  };

  function onMenuIdChange(parentid: number) {
    setMenuId(parentid);
    triggerChange(parentid);
  }

  function makeSelectOptionStyle(level: number, hasChildren: boolean) {
    const styleCss = { paddingLeft: `${level + 1}rem` };
    if (hasChildren) {
      styleCss['borderTop'] = '1px solid #F0F0F0';
    }
    return styleCss;
  }

  useEffect(() => {}, [data]);

  return (
    <Select
      value={value || menuId}
      showSearch
      optionFilterProp="children"
      filterOption={(input, option) =>
        (option!.children as unknown as string).toLowerCase().includes(input.toLowerCase())
      }
      onChange={onMenuIdChange}
      disabled={disabled}
    >
      {data?.map((item: ResponseAdminMenuListItemType) => (
        <Select.Option
          key={item.id}
          value={item.id}
          style={makeSelectOptionStyle(item.level || 0, item.hasChildren)}
        >
          {item.name}
        </Select.Option>
      ))}
    </Select>
  );
};

export default MenuTreeSelect;
