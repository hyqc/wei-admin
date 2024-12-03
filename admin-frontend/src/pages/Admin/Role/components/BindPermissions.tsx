import { Checkbox, Col, Row } from 'antd';
import { useEffect, useState } from 'react';
import './index.less';
import { MenuModeItem, MenuPageItem, MenuPagePermissions } from '@/proto/admin_ts/admin_menu';

export type MenuPagePermissionsType = {
  disabled?: boolean;
  permissionIds?: number[];
  datasource: MenuModeItem[];
  onChange?: (values: number[]) => void;
};

const BindPermissions: React.FC<MenuPagePermissionsType> = (props) => {
  const { disabled, datasource, permissionIds, onChange } = props;
  const dataMap: Map<number, Map<number, Set<number>>> = new Map(); //key=modeId,value=pageAllMap
  const initMap: Map<number, Map<number, Set<number>>> = new Map(); //key=modeId,value=pageAllMap
  const permissionSetIds = new Set(permissionIds)

  const [menuPageData] = useState<MenuModeItem[]>([...datasource]);
  

  const typeMode = 'mode'
  const typePage = 'page'
  const typePermission = 'permission'



  datasource?.forEach((mode) => {
    const modeId = mode.modelId ?? 0
    const pageMap = new Map<number, Set<number>>()
    mode?.pages?.forEach(page => {
      const pageId = page.pageId ?? 0
      const pids = new Set<number>()
      page?.permissions?.forEach(p => {
        const pid = p.permissionId ?? 0
        if (pid) {
          pids.add(pid)
          if(permissionSetIds.has(pid)){
            if(!initMap.has(modeId)){
              initMap.set(modeId,new Map())
            }
            if(!initMap.get(modeId)?.has(pageId)){
              initMap.get(modeId)?.set(pageId,new Set())
            }
            initMap.get(modeId)?.get(pageId)?.add(pid)
          }
        }
      })
      if (pageId) {
        pageMap.set(pageId, pids)
      }
    })
    if (modeId) {
      dataMap.set(modeId, pageMap)
    }
  });

  const [checkedIds, setCheckedIds] = useState<Map<number, Map<number, Set<number>>>>(initMap);


  const triggerChange = (ids: number[]) => {
    onChange?.(ids);
  };

  useEffect(() => {
    const ids: number[] = []
    checkedIds?.forEach(pages => {
      pages?.forEach(pids => {
        pids.forEach(id=>{
          if(id){
            ids.push(id)
          }
        })
      })
    }
    )
    triggerChange(ids);
  }, [checkedIds]);

  useEffect(() => {
  }, [datasource]);

  const setChange = (type: string, target: any) => {
    switch (type) {
      case typeMode:
        if (target.checked) {
          //全选
          setCheckedIds((ov) => {
            let modeId = target.value ?? 0
            const nv = new Map(ov)
            if (modeId && dataMap.has(modeId)) {
              const pageMap: Map<number, Set<number>> = dataMap.get(modeId) ?? new Map()
              nv.set(modeId, pageMap)
            }
            return nv
          });
        } else {
          //全不选
          setCheckedIds((ov) => {
            let modeId = target.value ?? 0
            const nv = new Map(ov)
            nv.delete(modeId)
            return nv
          });
        }
        break;
      case typePage:
        if (target.checked) {
          //全选
          setCheckedIds((ov) => {
            let modeId = target.modeId ?? 0
            const nv = new Map(ov)
            if (modeId && dataMap.has(modeId)) {
              let pageId = target.value ?? 0
              const pageMap: Map<number, Set<number>> = dataMap.get(modeId) ?? new Map()
              if (pageId && pageMap.has(pageId)) {
                const pidsMap: Set<number> = pageMap.get(pageId) ?? new Set()
                if (nv.has(modeId)) {
                  nv.get(modeId)?.set(pageId, pidsMap)
                } else {
                  const tmpPage = new Map()
                  tmpPage.set(pageId, pidsMap)
                  nv.set(modeId, tmpPage)
                }
              }
            }
            return nv
          });
        } else {
          //全不选
          setCheckedIds((ov) => {
            let modeId = target.modeId ?? 0
            let pageId = target.value ?? 0
            const nv = new Map(ov)
            nv.get(modeId)?.delete(pageId)
            if(nv.get(modeId)?.size ===0){
              nv.delete(modeId)
            }
            return nv
          });
        }
        break;
      case typePermission:
        if (target.checked) {
          //天加
          setCheckedIds((ov) => {
            let modeId = target.modeId ?? 0
            let pageId = target.pageId ?? 0
            let pid = target.value ?? 0
            const nv = new Map(ov)
            if (nv.has(modeId)) {
              if (nv.get(modeId)?.has(pageId)) {
                nv.get(modeId)?.get(pageId)?.add(target.value)
              } else {
                const pm: Map<number, Set<number>> = nv.get(modeId) ?? new Map()
                const idm = pm.get(pageId) ?? new Set()
                idm.add(pid)
                pm.set(pageId, idm)
                nv.set(modeId, pm)
              }
            } else {
              const pm: Map<number, Set<number>> = new Map()
              const im: Set<number> = new Set()
              im.add(pid)
              pm.set(pageId, im)
              nv.set(modeId, pm)
            }
            return nv
          });
        } else {
          //删除
          setCheckedIds((ov) => {
            let modeId = target.modeId ?? 0
            let pageId = target.pageId ?? 0
            const nv = new Map(ov)
            nv.get(modeId)?.get(pageId)?.delete(target.value)
            if (nv.get(modeId)?.get(pageId)?.size === 0) {
              nv.get(modeId)?.delete(pageId)
              if (nv.get(modeId)?.size === 0) {
                nv.delete(modeId)
              }
            }
            return nv
          });
        }
        break;
    }
  };


  const modeChanged = (e: any) => {
    setChange(typeMode, { value: e.target.value, checked: e.target.checked, pageIds: e.target['data-pageIds'] })
  };

  const pageChanged = (e: any) => {
    setChange(typePage, { value: e.target.value, checked: e.target.checked, modeId: e.target['data-modeId'], permissionIds: e.target['data-permissionIds'] })
  };

  const permissionChanged = (e: any) => {
    setChange(typePermission, { value: e.target.value, checked: e.target.checked, modeId: e.target['data-modeId'], pageId: e.target['data-pageId'] })
  };

  const modeIndeterminate = (modeId: number): boolean => {
    const check = checkedIds.get(modeId)?.size ?? 0
    const length = dataMap.get(modeId)?.size ?? 0
    return check !== length && check !== 0
  }

  const pageIndeterminate = (modeId: number, pageId: number): boolean => {
    const check = checkedIds.get(modeId)?.get(pageId)?.size ?? 0
    const length = dataMap.get(modeId)?.get(pageId)?.size ?? 0
    return check !== length && check !== 0
  }

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
                  checked={checkedIds.has(mode.modelId || 0)}
                  indeterminate={modeIndeterminate(mode.modelId || 0)}
                >
                  {mode.modelName}
                </Checkbox>
              </Col>
              <Col span={18}>
                {
                  mode.pages?.map((page: MenuPageItem, index: number) => {
                    return (
                      <Row key={page.pageId} style={index % 2 === 0 ? { backgroundColor: "#fafafa" } : {}}>
                        <Col span={10}>
                          <Checkbox
                            onChange={pageChanged}
                            disabled={disabled}
                            value={page.pageId}
                            data-modeId={mode.modelId}
                            checked={checkedIds.get(mode.modelId || 0)?.has(page.pageId || 0)}
                            indeterminate={pageIndeterminate(mode.modelId || 0, page.pageId || 0)}
                          >
                            {page.pageName}
                          </Checkbox>
                        </Col>
                        <Col span={14}>
                          {
                            page?.permissions?.map((permission: MenuPagePermissions) => {
                              return (
                                <Checkbox
                                  key={permission.permissionId}
                                  onChange={permissionChanged}
                                  disabled={disabled}
                                  checked={checkedIds.get(mode.modelId || 0)?.get(page.pageId || 0)?.has(permission.permissionId || 0)}
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
