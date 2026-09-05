import { useEffect, useState } from 'react';
import { Button, Empty, Form, Input, Modal, Select, Spin, Switch, message } from 'antd';
import { ApiOutlined, PlusOutlined } from '@ant-design/icons';
import { getAdminMenuPermissions } from '@/api/admin/menu';
import { addAdminMenuPermissions } from '@/api/admin/permission';
import PermissionApiBindModal from './PermissionApiBindModal';
import {
  DEFAULT_PERMISSION_TYPES,
  DEFAULT_PERMISSION_TEMPLATE_TYPES,
  PERMISSION_TYPE_OPTIONS,
  handleKey,
} from '@/views/admin/permission/components/common';
import type { MenuTreeItem, MenuPermissionItem } from '@/types/admin_menu';

interface Props {
  open: boolean;
  detailData?: MenuTreeItem;
  onClose: () => void;
}

/** 权限点行；auto 标记行 key 未被人工改动时随类型联动；名称始终跟随动作类型自动生成、不可修改 */
type PermissionRow = MenuPermissionItem & { auto?: boolean };

const keyPattern = new RegExp('^([A-Z][a-zA-Z0-9]*)+$');

/** 菜单操作权限配置（每个操作按钮对应一个权限点，唯一键与前端按钮权限码一致） */
export default function PermissionsSaveModal({ open, detailData, onClose }: Props) {
  const [loading, setLoading] = useState(false);
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [permissions, setPermissions] = useState<PermissionRow[]>([]);
  /** 接口绑定弹窗：当前编辑的权限点行下标 */
  const [apiBindOpen, setApiBindOpen] = useState(false);
  const [apiBindIndex, setApiBindIndex] = useState(-1);
  const currentRow = apiBindIndex >= 0 ? permissions[apiBindIndex] : undefined;

  useEffect(() => {
    if (open && detailData) {
      setPermissions([]);
      setLoading(true);
      const menuId = detailData.id;
      const menuName = detailData.name || '';
      const menuPath = detailData.path || '';
      getAdminMenuPermissions({ menuId })
        .then((res) => {
          const existing = res.data.permissions || [];
          if (existing.length > 0) {
            // 已有配置：名称/唯一键原样回显（名称只读，保留历史语义；切换动作类型时名称按规则重算）；接口绑定随行带回用于回显
            setPermissions(
              existing.map((item) => ({
                ...item,
                apiIds: item.apiIds || [],
                typeName: DEFAULT_PERMISSION_TYPES.find((t) => t.key === item.type)?.name || item.type || '',
                auto: false,
              })),
            );
            return;
          }
          // 菜单尚未配置权限时，按默认模板生成查看/新增/编辑/删除四类基础动作
          setPermissions(
            DEFAULT_PERMISSION_TEMPLATE_TYPES.map((type) => ({
              menuId,
              type: type.key,
              typeName: type.name,
              name: `${menuName}${type.name}`,
              key: handleKey(menuPath, type.key),
              enabled: true,
              describe: '',
              apiIds: [],
              auto: true,
            })),
          );
        })
        .finally(() => setLoading(false));
    }
  }, [open, detailData]);

  const updateRow = (index: number, patch: Partial<PermissionRow>) => {
    setPermissions((prev) => prev.map((p, i) => (i === index ? { ...p, ...patch } : p)));
  };

  const addRow = () => {
    const menuName = detailData?.name || '';
    const menuPath = detailData?.path || '';
    setPermissions((prev) => [
      ...prev,
      {
        menuId: detailData?.id,
        type: 'view',
        typeName: '查看',
        name: `${menuName}查看`,
        key: handleKey(menuPath, 'view'),
        enabled: true,
        describe: '',
        apiIds: [],
        auto: true,
      },
    ]);
  };

  const openApiBind = (index: number) => {
    setApiBindIndex(index);
    setApiBindOpen(true);
  };

  const onApiBindOk = (apiIds: number[]) => {
    if (currentRow) {
      updateRow(apiBindIndex, { apiIds });
    }
    setApiBindOpen(false);
  };

  const removeRow = (index: number) => {
    setPermissions((prev) => prev.filter((_, i) => i !== index));
  };

  const onTypeChange = (index: number, value: string) => {
    setPermissions((prev) =>
      prev.map((p, i) => {
        if (i !== index) return p;
        const typeName = DEFAULT_PERMISSION_TYPES.find((t) => t.key === value)?.name || value;
        const next: PermissionRow = { ...p, type: value, typeName };
        // 权限名称跟随动作类型自动生成、不可手工修改；key 未被人工改过时随类型联动
        next.name = `${detailData?.name || ''}${typeName}`;
        if (p.auto) {
          next.key = handleKey(detailData?.path || '', value);
        }
        return next;
      }),
    );
  };

  const markManual = (index: number) => {
    updateRow(index, { auto: false });
  };

  const handleOk = () => {
    if (!detailData) return;
    for (const item of permissions) {
      const name = (item.name || '').trim();
      const key = (item.key || '').trim();
      if (!name || !key) {
        message.warning('请完整填写权限名称与唯一键');
        return;
      }
      if (!keyPattern.test(key)) {
        message.warning(`唯一键 ${key} 请按照驼峰法命名`);
        return;
      }
      item.name = name;
      item.key = key;
    }
    setConfirmLoading(true);
    const list = permissions.map((item) => ({
      menuId: detailData.id,
      menuName: detailData.name,
      menuPath: detailData.path,
      id: item.id,
      name: item.name,
      key: item.key,
      type: item.type,
      describe: item.describe,
      enabled: item.enabled,
      apiIds: item.apiIds || [],
    }));
    addAdminMenuPermissions(list)
      .then((res) => {
        message.success(res.msg, 2);
        onClose();
      })
      .finally(() => setConfirmLoading(false));
  };

  return (
    <Modal open={open} title="菜单操作权限配置" width={1120} confirmLoading={confirmLoading} maskClosable={false} okText="保存" cancelText="取消" onOk={handleOk} onCancel={onClose}>
      <Spin spinning={loading}>
        <Form labelCol={{ span: 3 }} wrapperCol={{ span: 20 }}>
          <Form.Item label="菜单">
            <Input value={detailData?.name} disabled style={{ width: 260 }} />
            <div style={{ color: 'rgba(0,0,0,0.45)', fontSize: 12, lineHeight: 1.5, marginTop: 4 }}>
              权限点代表该页面上的一个可授权操作：唯一键与前端按钮权限码一一对应，并绑定该操作需要访问的接口；授权后按钮可见且对应接口可访问。保存后权限点与接口绑定一次性生效
              <br />
              动作类型固定为「查看 / 编辑 / 删除」三类（新增、绑定、重置等写操作统一归到“编辑”），权限名称由「菜单名 + 动作类型」自动生成且不可修改，唯一键按类型自动生成、可再手工修改
            </div>
          </Form.Item>
          <Form.Item label="权限动作">
            {permissions.length > 0 && (
              <div style={{ display: 'flex', alignItems: 'center', gap: 12, marginBottom: 8, color: 'rgba(0,0,0,0.45)', fontSize: 12 }}>
                <span style={{ flex: '0 0 112px' }}>动作类型</span>
                <span style={{ flex: '0 0 170px' }}>权限名称</span>
                <span style={{ flex: '1 1 auto', minWidth: 250 }}>唯一键</span>
                <span style={{ flex: '0 0 60px' }}>状态</span>
                <span style={{ flex: '0 0 140px' }}>接口</span>
              </div>
            )}
            {permissions.map((item, index) => (
              <div key={index} style={{ display: 'flex', alignItems: 'center', gap: 12, marginBottom: 8 }}>
                <Select
                  style={{ width: 112 }}
                  options={PERMISSION_TYPE_OPTIONS}
                  value={item.type}
                  onChange={(value) => onTypeChange(index, value)}
                />
                <Input
                  style={{ width: 170 }}
                  value={item.name}
                  readOnly
                  placeholder="跟随动作类型自动生成"
                />
                <Input
                  style={{ flex: '1 1 auto', minWidth: 250 }}
                  placeholder="如 AdminUserResetPwd"
                  value={item.key}
                  onChange={(e) => {
                    markManual(index);
                    updateRow(index, { key: e.target.value });
                  }}
                />
                <Switch
                  checked={item.enabled}
                  checkedChildren="启用"
                  unCheckedChildren="禁用"
                  onChange={(checked) => updateRow(index, { enabled: checked })}
                />
                <Button
                  size="small"
                  type={item.apiIds?.length ? 'primary' : 'default'}
                  icon={<ApiOutlined />}
                  onClick={() => openApiBind(index)}
                >
                  {item.apiIds?.length ? `已绑 ${item.apiIds.length}` : '绑定接口'}
                </Button>
                <Button type="link" size="small" danger onClick={() => removeRow(index)}>
                  删除
                </Button>
              </div>
            ))}
            <div style={{ marginBottom: 8 }}>
              <Button type="dashed" block icon={<PlusOutlined />} onClick={addRow}>
                添加权限动作
              </Button>
            </div>
            {permissions.length === 0 && <Empty description="暂无权限配置，可点击上方按钮添加" />}
          </Form.Item>
        </Form>
      </Spin>
      <PermissionApiBindModal
        open={apiBindOpen}
        apiIds={currentRow?.apiIds}
        permissionName={currentRow?.name}
        onOk={onApiBindOk}
        onClose={() => setApiBindOpen(false)}
      />
    </Modal>
  );
}
