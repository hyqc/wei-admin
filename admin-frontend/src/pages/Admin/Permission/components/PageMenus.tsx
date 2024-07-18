import { ResponseAdminMenuListItemType } from '@/services/apis/admin/menu';
import { Select } from 'antd';
import { useEffect, useState } from 'react';

export type AddModalPropsType = {
  data: ResponseAdminMenuListItemType[];
  value?: number;
  disabled?: boolean;
  onChange?: (parentid: number) => void;
};

const PageMenus: React.FC<AddModalPropsType> = (props) => {
  const { data, value, disabled, onChange } = props;
  const [menuId, setMenuId] = useState<number>(value || 0);
  const triggerChange = (id: number) => {
    onChange?.(id);
  };

  function onMenuIdChange(id: number) {
    setMenuId(id);
    triggerChange(id);
  }

  useEffect(() => {}, [data]);

  return (
    <Select
      showSearch
      style={{ offset: 0, width: 120 }}
      filterOption={(input, option) =>
        (option!.children as unknown as string).toLowerCase().includes(input.toLowerCase())
      }
      onChange={onMenuIdChange}
      disabled={disabled}
      value={value || menuId}
    >
      {props?.children}
      {data?.map((item: ResponseAdminMenuListItemType) => {
        return (
          <Select.Option key={item.id + ''} value={item.id}>
            {item.name}
          </Select.Option>
        );
      })}
    </Select>
  );
};

export default PageMenus;
