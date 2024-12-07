import { history, useModel } from '@umijs/max';
import { Tabs } from 'antd';
import React, { useEffect, useState } from 'react';

export type ContentType = {
    style?: React.CSSProperties;
    cardStyle?: React.CSSProperties;
    children?: React.ReactNode;
};


const PageContainer: React.FC<ContentType> = (props: any) => {
    const { style, cardStyle, children } = props;
    const storeTabListKey = "tabmenus"

    const { initialState } = useModel('@@initialState');
    const { currentUser } = initialState || {};
    const { menus } = currentUser || { menus: {} };

    const data = JSON.parse(localStorage.getItem(storeTabListKey) || "[]");
    const [tabList, setTabList] = useState<any[]>([...data]);
    const [curTab, setCurTab] = useState<any>({key: location.pathname});
    

    useEffect(() => { 
        console.log('tabs= 1', tabList, curTab)
    }, [tabList,curTab])

    useEffect(() => { 
        let currentTab:any = {}
        for (let key in menus) {
            if (menus[key]?.path === location.pathname) {
                currentTab = {
                    label: menus[key].name,
                    key: menus[key]?.path,
                    children:  <>{children}</>
                }
                break
            }
        }
        console.log('tabs= 3', currentTab)
        if (Object.keys(currentTab).length > 0) {
            setTabList((old) => {
                const nv = [...old]
                let isExist = false
                const storeData = old.map(item=>{
                    if(item.key === currentTab.key){
                        isExist = true
                        setCurTab(currentTab)
                    }
                    return {label:item.label,key:item.key}
                })
                if(!isExist){
                    nv.push(currentTab)
                    //storeData.push({label: currentTab.label, key: currentTab.key})
                }
                //localStorage.setItem(storeTabListKey, JSON.stringify(storeData))
                return nv
            })
        }
    }, [ location.pathname ])

    const onClick = (key: string) => {
        setCurTab(()=>{
            return {...tabList.find(item=>item.key === key), children}
        })
        // history.replace({
        //     pathname: key,
        // });
    }

    return (
        <div style={style}>
            
            <Tabs type="editable-card" items={tabList} activeKey={curTab.key} onTabClick={onClick} />
            {/* {children} */}
        </div>
    )
};

export default PageContainer;
