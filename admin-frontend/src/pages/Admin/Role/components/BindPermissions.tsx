import { Checkbox, Col, Row } from 'antd';
import { useEffect, useState } from 'react';
import {
  AdminMenuModePagesItemType,
  AdminMenuModePagesPermissionsItemType,
  ResponseAdminMenuModeTypeData,
} from '@/services/apis/admin/menu';
import './index.less';
export type NoticeModalPropsType = {
  reload?: boolean;
};

export type EditModalPropsType = {};

export type ModeRowType = {
  umid: string;
  upid: string;
  modelId: number;
  modelName: string;
  pageId: number;
  pageName: string;
  permissions: AdminMenuModePagesPermissionsItemType[];
};

export type MenuPagePermissionsType = {
  disabled?: boolean;
  permissionIds?: number[];
  datasource: ResponseAdminMenuModeTypeData[];
  onChange?: (values: number[]) => void;
};

const permissionStyle = { width: '260px' };
const menuPageStyle = { width: '160px' };

const BindPermissions: React.FC<MenuPagePermissionsType> = (props) => {
  const { disabled, permissionIds, datasource, onChange } = props;
  const [modeData, setMenuModeData] = useState<ModeRowType[]>();

  const [dataStatusMap, setDataStatusMap] = useState<any>({});
  const [dataMap, setDataMap] = useState<any>({});
  const [childrenDataMap, setChildrenDataMap] = useState<any>({});
  const [checkedPermissions, setCheckedPermissions] = useState<string[]>([]);
  const [checkedIds, setCheckedIds] = useState<number[]>([]);

  const triggerChange = (ids: number[]) => {
    onChange?.(ids);
  };

  function checkboxChange(v: any[]) {}

  function checkedKeys(checked: boolean, value: string, childrenDataMap: any, sets: Set<string>) {
    const s: string[] = value.split('-');
    if (s.length === 3) {
      // 权限
      const pageId = s[0] + '-' + s[1];
      const modelId = s[0];
      const peridsKeys: string[] = Object.keys(childrenDataMap[modelId][pageId]);
      if (checked) {
        sets.add(value);
      } else {
        sets.delete(value);
      }
      const pageSets = new Set([...sets, ...peridsKeys]);
      if (pageSets.size === sets.size) {
        sets.add(pageId);
        const pageKeys = Object.keys(childrenDataMap[modelId]);
        const modelSets: Set<string> = new Set([...sets, ...pageKeys]);
        if (modelSets.size === sets.size) {
          sets.add(modelId);
        } else {
          sets.delete(modelId);
        }
      } else {
        sets.delete(pageId);
        sets.delete(modelId);
      }
    } else if (s.length === 2) {
      // 页面
      const modelId = s[0];
      const pageId = value;
      const peridsKeys: string[] = Object.keys(childrenDataMap[modelId][pageId]);
      const pageKeys: string[] = Object.keys(childrenDataMap[modelId]);
      if (checked) {
        sets.add(value);
        peridsKeys.forEach((pid) => {
          sets.add(pid);
        });
        const modelSets: Set<string> = new Set([...sets, ...pageKeys]);
        if (modelSets.size === sets.size) {
          sets.add(modelId);
        } else {
          sets.delete(modelId);
        }
      } else {
        sets.delete(value);
        sets.delete(modelId);
        peridsKeys.forEach((pid) => {
          sets.delete(pid);
        });
      }
    } else {
      // 模块
      const modelId = s[0];
      const pageKeys = Object.keys(childrenDataMap[modelId]);
      if (checked) {
        sets.add(value);
        pageKeys.forEach((pageId) => {
          sets.add(pageId);
          Object.keys(childrenDataMap[modelId][pageId]).forEach((pid) => {
            sets.add(pid);
          });
        });
      } else {
        sets.delete(value);
        pageKeys.forEach((pageId) => {
          sets.delete(pageId);
          Object.keys(childrenDataMap[modelId][pageId]).forEach((pid) => {
            sets.delete(pid);
          });
        });
      }
    }
    return sets;
  }

  function makeCheckedIds() {
    const tmpIds: number[] = checkedPermissions
      .filter((id) => {
        return id.split('-').length === 3;
      })
      .map((id) => {
        return parseInt(id.split('-')[2]);
      });
    const idSets: Set<number> = new Set(tmpIds);
    const ids: number[] = [...idSets].sort();
    setCheckedIds(ids);
  }

  function changed(e: any) {
    const checked = e.target.checked;
    const value = e.target.value;
    let tmpSet = new Set([...checkedPermissions]);
    tmpSet = checkedKeys(checked, value, childrenDataMap, tmpSet);
    setCheckedPermissions([...tmpSet]);
  }

  function changeIndeterminate(idSets: Set<string>, modelMap: any) {
    let resultMap = {};
    Object.keys(modelMap).forEach((modelId: string) => {
      resultMap[modelId] = false;
      if (!idSets.has(modelId)) {
        // 未选中模块，后代可能选中
        const pagesMap = modelMap[modelId];
        Object.keys(pagesMap).forEach((pageId: string) => {
          resultMap[pageId] = false;
          if (!idSets.has(pageId)) {
            // 未选中页面，后代不是全选
            const res: boolean = Object.keys(pagesMap[pageId]).some((pid: string) => {
              return idSets.has(pid);
            });
            if (res) {
              resultMap[pageId] = true;
              resultMap[modelId] = true;
            }
          } else {
            // 选中页面，后代是全选
            resultMap[modelId] = true;
          }
        });
      } else {
        // 选中模块，后代全选中
        const pagesMap = modelMap[modelId];
        Object.keys(pagesMap).forEach((pageId: string) => {
          resultMap[pageId] = false;
          Object.keys(pagesMap[pageId]).forEach((pid: string) => {
            resultMap[pid] = false;
          });
        });
      }
    });
    return resultMap;
  }

  function initCheckedKeys(ids: number[], dataMap: any, childrenDataMap: any): Set<string> {
    const keys: string[] = Object.keys(dataMap);
    let resultKeys: Set<string> = new Set();
    ids.forEach((id: number) => {
      keys.forEach((key) => {
        const tmp = key.split('-');
        if (tmp.length === 3 && tmp[2] === id + '') {
          let tmpSet = new Set([...resultKeys]);
          tmpSet = checkedKeys(true, key, childrenDataMap, tmpSet);
          resultKeys = new Set([...resultKeys, ...tmpSet]);
        }
      });
    });
    return resultKeys;
  }

  useEffect(() => {
    let result: ModeRowType[] = [];
    let tmpChildrenDataMap = {};
    let tmpDataStatusMap = {};
    let tmpDataMap = {};
    datasource?.map((item: ResponseAdminMenuModeTypeData, index: number) => {
      const umid = item.modelId + '';
      const tmpItem = {
        modelId: item.modelId,
        modelName: item.modelName,
        umid,
      };
      tmpChildrenDataMap[umid] = {};
      tmpDataStatusMap[umid] = false;
      if (item?.pages) {
        item.pages.forEach((pageItem: AdminMenuModePagesItemType, pageIndex: number) => {
          const upid: string = tmpItem.umid + '-' + pageItem.pageId;
          tmpChildrenDataMap[umid][upid] = {};
          tmpDataStatusMap[upid] = false;
          tmpDataMap[upid] = umid;
          let tmpPageItem: ModeRowType = {
            ...tmpItem,
            pageId: pageItem.pageId,
            pageName: pageItem.pageName,
            permissions: pageItem.permissions?.map((p) => {
              const ukid: string = tmpItem.umid + '-' + pageItem.pageId + '-' + p.permissionId;
              tmpChildrenDataMap[umid][upid][ukid] = p.permissionId;
              tmpDataMap[ukid] = upid;
              return { ...p, ukid };
            }),
            upid,
          };
          if (pageIndex > 0) {
            tmpPageItem.modelName = '';
          }
          result.push(tmpPageItem);
        });
      }
    });
    setDataMap(tmpDataMap);
    setMenuModeData(result);
    setChildrenDataMap(tmpChildrenDataMap);
  }, [datasource]);

  useEffect(() => {
    const setKeys: Set<string> = initCheckedKeys(permissionIds || [], dataMap, childrenDataMap);
    setCheckedPermissions([...setKeys]);
  }, [dataMap, childrenDataMap]);

  useEffect(() => {}, [dataStatusMap]);

  useEffect(() => {
    makeCheckedIds();
    const keySets: Set<string> = new Set(checkedPermissions);
    //计算选中状态
    const tmpStatusMap = changeIndeterminate(keySets, childrenDataMap);
    const statusMap = { ...dataStatusMap, ...tmpStatusMap };
    setDataStatusMap(statusMap);
  }, [checkedPermissions]);

  useEffect(() => {
    triggerChange(checkedIds);
  }, [checkedIds]);

  return (
    <div className="menu-page-permissions">
      <Row wrap={false} className="title">
        <Col style={menuPageStyle}>模块</Col>
        <Col style={menuPageStyle}>页面</Col>
        <Col style={permissionStyle}>权限</Col>
      </Row>
      <Checkbox.Group onChange={checkboxChange} value={checkedPermissions}>
        {modeData?.map((item: ModeRowType) => {
          return (
            <Row wrap={false} key={item.pageId}>
              <Col style={menuPageStyle}>
                {item.modelName ? (
                  <Checkbox
                    disabled={disabled}
                    indeterminate={dataStatusMap[item.umid]}
                    onChange={changed}
                    value={item.umid}
                  >
                    {item.modelName}
                  </Checkbox>
                ) : (
                  ''
                )}
              </Col>
              <Col style={menuPageStyle}>
                <Checkbox
                  disabled={disabled}
                  indeterminate={dataStatusMap[item.upid]}
                  onChange={changed}
                  value={item.upid}
                >
                  {item.pageName}
                </Checkbox>
              </Col>
              <Col style={permissionStyle}>
                {item?.permissions?.map((permissionItem: AdminMenuModePagesPermissionsItemType) => (
                  <Checkbox
                    disabled={disabled}
                    onChange={changed}
                    key={permissionItem.permissionId}
                    value={permissionItem.ukid}
                  >
                    {permissionItem.permissionText}
                  </Checkbox>
                ))}
              </Col>
            </Row>
          );
        })}
      </Checkbox.Group>
    </div>
  );
};

export default BindPermissions;
