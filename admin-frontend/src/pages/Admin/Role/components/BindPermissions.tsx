import { Checkbox, Col, Row } from 'antd';
import { useEffect, useState } from 'react';
import './index.less';
import { MenuModeItem, MenuPageItem, MenuPagePermissions } from '@/proto/admin_ts/admin_menu';
import permission from 'mock/admin/permission';

export type MenuPagePermissionsType = {
  disabled?: boolean;
  permissionIds?: number[];
  datasource: MenuModeItem[];
  onChange?: (values: number[]) => void;
};


const modeStyle = { width: '20rem' };
const pageStyle = { width: '20rem' };
const permissionStyle = { width: '30rem' };

const BindPermissions: React.FC<MenuPagePermissionsType> = (props) => {
  const { disabled, datasource, permissionIds, onChange } = props;
  const modeAllMap: Map<number, Map<number, number[]>> = new Map();
  const pageAllMap: Map<number, number[]> = new Map();

  const [menuPageData, setMenuPageData] = useState<MenuModeItem[]>([...datasource]);
  const [checkedIds, setCheckedIds] = useState<number[]>([]);

  const [modeCheckedIds, setModeCheckedIds] = useState<Set<number>>(new Set());
  const [pageCheckedIds, setPageCheckedIds] = useState<Map<number, number[]>>(new Map());
  const [permissionCheckedIds, setPermissionCheckedIds] = useState<Map<number, number[]>>(new Map());

  const triggerChange = (ids: number[]) => {
    onChange?.(ids);
  };

  useEffect(() => {
    triggerChange(checkedIds);
  }, [checkedIds]);

  const typeMode = 'mode'
  const typePage = 'page'
  const typePermission = 'permission'

  const isIndeterminate = (type: string, id: number, nv: Set<number>): boolean => {
    switch (type) {
      case typeMode:
        return false
      case typePage:
        return modeAllMap.get(id)?.size == nv.size
      case typePermission:
        return pageAllMap.get(id)?.length === nv.size
      default:
        return false
    }
  }

  const checkboxChangedDataCall = (type: string, e: any, old: Map<number, number[]>) => {
    const nv = new Set(old);
    if (e.checked) {
      nv.add(e.value);
    } else {
      nv.delete(e.value);
    }
    const pageId = e.pageId ?? 0
    const modeId = e.modelId ?? 0
    // 触发操作
    switch (type) {
      case typePermission:
        //触发页面操作
        pageAllMap.get(pageId)?.forEach(permissionId => {

        })
        break;
      case typePage:
        pageAllMap.get(e.value)?.forEach(permissionId => {

        })
        break;
      case typeMode:
        modeAllMap.get(e.value)?.forEach(pageId => {

        })
        break;
    }
    return nv
  }

  const setChange = (type: string, target: any) => {
    switch (type) {
      case typeMode:
        setModeCheckedIds((old) => {
          const nv = new Set(old);
          if (target.checked) {
            nv.add(target.value);
          } else {
            nv.delete(target.value);
          }
          return nv
        });
        break;
      case typePage:
        setPageCheckedIds((old) => {
          const nv = new Map(old)
          if (target.checked) {
            nv.set(target.value, target['data-permissionIds']);
          } else {
            nv.delete(target.value);
          }
          return nv
        });
      case typePermission:
        setPermissionCheckedIds((old) => {
          const nv = new Map(old)
          if (target.checked) {
            nv.set(target['data-pageId'], target.value);
          } else {
            nv.delete(target.value);
          }
          return nv
        });
        break;
    }
  };

  useEffect(() => {
    // permissionIds?.forEach((id) => {
    //   setChange(id, typePermission, true);
    // });
  }, [permissionIds]);

  useEffect(() => {
    datasource?.forEach((mode) => {
      const pageMap = new Map();
      mode?.pages?.map((page: MenuPageItem) => {
        let permissionIds = page?.permissions?.map((item) => item.permissionId ?? 0).filter((item) => item > 0) || [];
        if (page.pageId) {
          pageAllMap.set(page.pageId, permissionIds);
          pageMap.set(page.pageId, permissionIds);
        }
        return page;
      }) || [];
      if (mode.modelId) {
        modeAllMap.set(mode.modelId, pageMap);
      }
    });

  }, [datasource]);

  const modeChanged = (e: any) => {
    console.log('模块点击target：', e.target.checked)
    console.log('模块点击e：', e.target)
    setChange(typeMode, { value: e.target.value, checked: e.target.checked, pageIds: e.target['data-pageIds'] })
  };

  const pageChanged = (e: any) => {
    console.log('模块点击target：', e.target.checked)
    console.log('模块点击e：', e.target)
    setChange(typePage, { value: e.target.value, checked: e.target.checked, modeId: e.target['data-modeId'], permissionIds: e.target['data-permissionIds'] })
  };

  const permissionChanged = (e: any) => {
    console.log('模块点击target：', e.target.checked)
    console.log('模块点击e：', e.target)
    setChange(typePermission, { value: e.target.value, checked: e.target.checked, modeId: e.target['data-modeId'], pageId: e.target['data-pageId'] })
  };

  return (
    <div className="menu-page-permissions">
      <Row className="title">
        <Col span={6}>模块</Col>
        <Col span={8}>页面</Col>
        <Col span={10}>权限</Col>
      </Row>
      {
        menuPageData?.map((mode: MenuModeItem) => {
          return (
            <Row key={mode.modelId}>
              <Col span={6}>
                <Checkbox
                  disabled={disabled}
                  onChange={modeChanged}
                  value={mode.modelId}
                  checked={modeCheckedIds.has(mode.modelId || 0)}
                  data-pageIds={modeAllMap.get(mode.modelId || 0)}
                // indeterminate={modeAllMap.get(mode.modelId || 0)?.length === pageCheckedIds.size}
                >
                  {mode.modelName}
                </Checkbox>
              </Col>
              <Col span={18}>
                {
                  mode.pages?.map((page: MenuPageItem, index: number) => {
                    return (
                      <Row style={index % 2 === 0 ? { backgroundColor: "#fafafa" } : {}}>
                        <Col span={10}>
                          <Checkbox
                            onChange={pageChanged}
                            disabled={disabled}
                            value={page.pageId}
                            data-modeId={mode.modelId}
                            data-permissionIds={pageAllMap.get(page.pageId || 0)}
                            checked={pageCheckedIds.has(page.pageId || 0)}
                            indeterminate={pageAllMap.get(page.pageId || 0)?.length === permissionCheckedIds.size}
                          >
                            {page.pageName}
                          </Checkbox>
                        </Col>
                        <Col span={14}>
                          {
                            page?.permissions?.map((permission: MenuPagePermissions) => {
                              return (
                                <Checkbox
                                  onChange={permissionChanged}
                                  disabled={disabled}
                                  checked={permissionCheckedIds.has(permission.permissionId || 0)}
                                  value={permission.permissionId}
                                  data-modeId={mode.modelId}
                                  data-pageId={page.pageId}
                                >
                                  {permission.permissionTypeName}
                                </Checkbox>
                              );
                            })
                          }
                        </Col>
                      </Row>
                    );
                  })
                }
              </Col>
            </Row>
          );
        })
      }

    </div>
  );
};

export default BindPermissions;
