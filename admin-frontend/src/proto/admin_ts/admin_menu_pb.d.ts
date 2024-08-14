// package: admin
// file: admin_menu.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as common_pb from "./common_pb";
import * as admin_permission_pb from "./admin_permission_pb";

export class MenuItem extends jspb.Message { 
    getKey(): string;
    setKey(value: string): MenuItem;
    getPath(): string;
    setPath(value: string): MenuItem;
    getName(): string;
    setName(value: string): MenuItem;
    getIcon(): string;
    setIcon(value: string): MenuItem;
    getComponent(): string;
    setComponent(value: string): MenuItem;
    getAuthority(): string;
    setAuthority(value: string): MenuItem;
    clearRoutesList(): void;
    getRoutesList(): Array<RouteItem>;
    setRoutesList(value: Array<RouteItem>): MenuItem;
    addRoutes(value?: RouteItem, index?: number): RouteItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuItem.AsObject;
    static toObject(includeInstance: boolean, msg: MenuItem): MenuItem.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuItem, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuItem;
    static deserializeBinaryFromReader(message: MenuItem, reader: jspb.BinaryReader): MenuItem;
}

export namespace MenuItem {
    export type AsObject = {
        key: string,
        path: string,
        name: string,
        icon: string,
        component: string,
        authority: string,
        routesList: Array<RouteItem.AsObject>,
    }
}

export class RouteItem extends jspb.Message { 
    getKey(): string;
    setKey(value: string): RouteItem;
    getPath(): string;
    setPath(value: string): RouteItem;
    getName(): string;
    setName(value: string): RouteItem;
    getIcon(): string;
    setIcon(value: string): RouteItem;
    getComponent(): string;
    setComponent(value: string): RouteItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RouteItem.AsObject;
    static toObject(includeInstance: boolean, msg: RouteItem): RouteItem.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RouteItem, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RouteItem;
    static deserializeBinaryFromReader(message: RouteItem, reader: jspb.BinaryReader): RouteItem;
}

export namespace RouteItem {
    export type AsObject = {
        key: string,
        path: string,
        name: string,
        icon: string,
        component: string,
    }
}

export class MenuListReq extends jspb.Message { 

    hasBase(): boolean;
    clearBase(): void;
    getBase(): common_pb.ListBaseReq | undefined;
    setBase(value?: common_pb.ListBaseReq): MenuListReq;
    getKey(): string;
    setKey(value: string): MenuListReq;
    getPath(): string;
    setPath(value: string): MenuListReq;
    getName(): string;
    setName(value: string): MenuListReq;
    getParentid(): number;
    setParentid(value: number): MenuListReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuListReq.AsObject;
    static toObject(includeInstance: boolean, msg: MenuListReq): MenuListReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuListReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuListReq;
    static deserializeBinaryFromReader(message: MenuListReq, reader: jspb.BinaryReader): MenuListReq;
}

export namespace MenuListReq {
    export type AsObject = {
        base?: common_pb.ListBaseReq.AsObject,
        key: string,
        path: string,
        name: string,
        parentid: number,
    }
}

export class MenuListRespData extends jspb.Message { 
    getTotal(): number;
    setTotal(value: number): MenuListRespData;
    clearRowsList(): void;
    getRowsList(): Array<MenuItem>;
    setRowsList(value: Array<MenuItem>): MenuListRespData;
    addRows(value?: MenuItem, index?: number): MenuItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuListRespData.AsObject;
    static toObject(includeInstance: boolean, msg: MenuListRespData): MenuListRespData.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuListRespData, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuListRespData;
    static deserializeBinaryFromReader(message: MenuListRespData, reader: jspb.BinaryReader): MenuListRespData;
}

export namespace MenuListRespData {
    export type AsObject = {
        total: number,
        rowsList: Array<MenuItem.AsObject>,
    }
}

export class MenuTreeItem extends jspb.Message { 
    getLevel(): number;
    setLevel(value: number): MenuTreeItem;
    getId(): number;
    setId(value: number): MenuTreeItem;
    getKey(): string;
    setKey(value: string): MenuTreeItem;
    getName(): string;
    setName(value: string): MenuTreeItem;
    getParentid(): number;
    setParentid(value: number): MenuTreeItem;
    getDescribe(): string;
    setDescribe(value: string): MenuTreeItem;
    getPath(): string;
    setPath(value: string): MenuTreeItem;
    getRedirect(): string;
    setRedirect(value: string): MenuTreeItem;
    getComponent(): string;
    setComponent(value: string): MenuTreeItem;
    getSort(): number;
    setSort(value: number): MenuTreeItem;
    getIcon(): string;
    setIcon(value: string): MenuTreeItem;
    getHidechildreninmenu(): boolean;
    setHidechildreninmenu(value: boolean): MenuTreeItem;
    getHideinmenu(): boolean;
    setHideinmenu(value: boolean): MenuTreeItem;
    getEnabled(): boolean;
    setEnabled(value: boolean): MenuTreeItem;
    getCreatetime(): number;
    setCreatetime(value: number): MenuTreeItem;
    getModifytime(): number;
    setModifytime(value: number): MenuTreeItem;
    clearChildrenList(): void;
    getChildrenList(): Array<MenuTreeItem>;
    setChildrenList(value: Array<MenuTreeItem>): MenuTreeItem;
    addChildren(value?: MenuTreeItem, index?: number): MenuTreeItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuTreeItem.AsObject;
    static toObject(includeInstance: boolean, msg: MenuTreeItem): MenuTreeItem.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuTreeItem, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuTreeItem;
    static deserializeBinaryFromReader(message: MenuTreeItem, reader: jspb.BinaryReader): MenuTreeItem;
}

export namespace MenuTreeItem {
    export type AsObject = {
        level: number,
        id: number,
        key: string,
        name: string,
        parentid: number,
        describe: string,
        path: string,
        redirect: string,
        component: string,
        sort: number,
        icon: string,
        hidechildreninmenu: boolean,
        hideinmenu: boolean,
        enabled: boolean,
        createtime: number,
        modifytime: number,
        childrenList: Array<MenuTreeItem.AsObject>,
    }
}

export class MenuTreeRespData extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuTreeRespData.AsObject;
    static toObject(includeInstance: boolean, msg: MenuTreeRespData): MenuTreeRespData.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuTreeRespData, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuTreeRespData;
    static deserializeBinaryFromReader(message: MenuTreeRespData, reader: jspb.BinaryReader): MenuTreeRespData;
}

export namespace MenuTreeRespData {
    export type AsObject = {
    }
}

export class MenuAddReq extends jspb.Message { 
    getKey(): string;
    setKey(value: string): MenuAddReq;
    getPath(): string;
    setPath(value: string): MenuAddReq;
    getName(): string;
    setName(value: string): MenuAddReq;
    getParentid(): number;
    setParentid(value: number): MenuAddReq;
    getDescribe(): string;
    setDescribe(value: string): MenuAddReq;
    getRedirect(): string;
    setRedirect(value: string): MenuAddReq;
    getIcon(): string;
    setIcon(value: string): MenuAddReq;
    getHidechildreninmenu(): boolean;
    setHidechildreninmenu(value: boolean): MenuAddReq;
    getHideinmenu(): boolean;
    setHideinmenu(value: boolean): MenuAddReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): MenuAddReq;
    getSort(): number;
    setSort(value: number): MenuAddReq;
    getComponent(): string;
    setComponent(value: string): MenuAddReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuAddReq.AsObject;
    static toObject(includeInstance: boolean, msg: MenuAddReq): MenuAddReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuAddReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuAddReq;
    static deserializeBinaryFromReader(message: MenuAddReq, reader: jspb.BinaryReader): MenuAddReq;
}

export namespace MenuAddReq {
    export type AsObject = {
        key: string,
        path: string,
        name: string,
        parentid: number,
        describe: string,
        redirect: string,
        icon: string,
        hidechildreninmenu: boolean,
        hideinmenu: boolean,
        enabled: boolean,
        sort: number,
        component: string,
    }
}

export class MenuInfoReq extends jspb.Message { 
    getMenuid(): number;
    setMenuid(value: number): MenuInfoReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuInfoReq.AsObject;
    static toObject(includeInstance: boolean, msg: MenuInfoReq): MenuInfoReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuInfoReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuInfoReq;
    static deserializeBinaryFromReader(message: MenuInfoReq, reader: jspb.BinaryReader): MenuInfoReq;
}

export namespace MenuInfoReq {
    export type AsObject = {
        menuid: number,
    }
}

export class MenuInfoRespData extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuInfoRespData.AsObject;
    static toObject(includeInstance: boolean, msg: MenuInfoRespData): MenuInfoRespData.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuInfoRespData, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuInfoRespData;
    static deserializeBinaryFromReader(message: MenuInfoRespData, reader: jspb.BinaryReader): MenuInfoRespData;
}

export namespace MenuInfoRespData {
    export type AsObject = {
    }
}

export class MenuEditReq extends jspb.Message { 
    getMenuid(): number;
    setMenuid(value: number): MenuEditReq;
    getKey(): string;
    setKey(value: string): MenuEditReq;
    getPath(): string;
    setPath(value: string): MenuEditReq;
    getName(): string;
    setName(value: string): MenuEditReq;
    getParentid(): number;
    setParentid(value: number): MenuEditReq;
    getDescribe(): string;
    setDescribe(value: string): MenuEditReq;
    getRedirect(): string;
    setRedirect(value: string): MenuEditReq;
    getIcon(): string;
    setIcon(value: string): MenuEditReq;
    getHidechildreninmenu(): boolean;
    setHidechildreninmenu(value: boolean): MenuEditReq;
    getHideinmenu(): boolean;
    setHideinmenu(value: boolean): MenuEditReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): MenuEditReq;
    getSort(): number;
    setSort(value: number): MenuEditReq;
    getComponent(): string;
    setComponent(value: string): MenuEditReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuEditReq.AsObject;
    static toObject(includeInstance: boolean, msg: MenuEditReq): MenuEditReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuEditReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuEditReq;
    static deserializeBinaryFromReader(message: MenuEditReq, reader: jspb.BinaryReader): MenuEditReq;
}

export namespace MenuEditReq {
    export type AsObject = {
        menuid: number,
        key: string,
        path: string,
        name: string,
        parentid: number,
        describe: string,
        redirect: string,
        icon: string,
        hidechildreninmenu: boolean,
        hideinmenu: boolean,
        enabled: boolean,
        sort: number,
        component: string,
    }
}

export class MenuEnableReq extends jspb.Message { 
    getMenuid(): number;
    setMenuid(value: number): MenuEnableReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): MenuEnableReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuEnableReq.AsObject;
    static toObject(includeInstance: boolean, msg: MenuEnableReq): MenuEnableReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuEnableReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuEnableReq;
    static deserializeBinaryFromReader(message: MenuEnableReq, reader: jspb.BinaryReader): MenuEnableReq;
}

export namespace MenuEnableReq {
    export type AsObject = {
        menuid: number,
        enabled: boolean,
    }
}

export class MenuDeleteReq extends jspb.Message { 
    getMenuid(): number;
    setMenuid(value: number): MenuDeleteReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuDeleteReq.AsObject;
    static toObject(includeInstance: boolean, msg: MenuDeleteReq): MenuDeleteReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuDeleteReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuDeleteReq;
    static deserializeBinaryFromReader(message: MenuDeleteReq, reader: jspb.BinaryReader): MenuDeleteReq;
}

export namespace MenuDeleteReq {
    export type AsObject = {
        menuid: number,
    }
}

export class MenuPermissionsReq extends jspb.Message { 
    getMenuid(): number;
    setMenuid(value: number): MenuPermissionsReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuPermissionsReq.AsObject;
    static toObject(includeInstance: boolean, msg: MenuPermissionsReq): MenuPermissionsReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuPermissionsReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuPermissionsReq;
    static deserializeBinaryFromReader(message: MenuPermissionsReq, reader: jspb.BinaryReader): MenuPermissionsReq;
}

export namespace MenuPermissionsReq {
    export type AsObject = {
        menuid: number,
    }
}

export class MenuPermissions extends jspb.Message { 

    hasMenu(): boolean;
    clearMenu(): void;
    getMenu(): MenuTreeItem | undefined;
    setMenu(value?: MenuTreeItem): MenuPermissions;
    clearPermissionsList(): void;
    getPermissionsList(): Array<admin_permission_pb.PermissionApiItem>;
    setPermissionsList(value: Array<admin_permission_pb.PermissionApiItem>): MenuPermissions;
    addPermissions(value?: admin_permission_pb.PermissionApiItem, index?: number): admin_permission_pb.PermissionApiItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuPermissions.AsObject;
    static toObject(includeInstance: boolean, msg: MenuPermissions): MenuPermissions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuPermissions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuPermissions;
    static deserializeBinaryFromReader(message: MenuPermissions, reader: jspb.BinaryReader): MenuPermissions;
}

export namespace MenuPermissions {
    export type AsObject = {
        menu?: MenuTreeItem.AsObject,
        permissionsList: Array<admin_permission_pb.PermissionApiItem.AsObject>,
    }
}

export class MenuPagesReq extends jspb.Message { 
    getAll(): boolean;
    setAll(value: boolean): MenuPagesReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuPagesReq.AsObject;
    static toObject(includeInstance: boolean, msg: MenuPagesReq): MenuPagesReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuPagesReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuPagesReq;
    static deserializeBinaryFromReader(message: MenuPagesReq, reader: jspb.BinaryReader): MenuPagesReq;
}

export namespace MenuPagesReq {
    export type AsObject = {
        all: boolean,
    }
}

export class MenuPagesRespData extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuPagesRespData.AsObject;
    static toObject(includeInstance: boolean, msg: MenuPagesRespData): MenuPagesRespData.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuPagesRespData, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuPagesRespData;
    static deserializeBinaryFromReader(message: MenuPagesRespData, reader: jspb.BinaryReader): MenuPagesRespData;
}

export namespace MenuPagesRespData {
    export type AsObject = {
    }
}

export class MenuModeReq extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuModeReq.AsObject;
    static toObject(includeInstance: boolean, msg: MenuModeReq): MenuModeReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuModeReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuModeReq;
    static deserializeBinaryFromReader(message: MenuModeReq, reader: jspb.BinaryReader): MenuModeReq;
}

export namespace MenuModeReq {
    export type AsObject = {
    }
}

export class MenuModeRespData extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MenuModeRespData.AsObject;
    static toObject(includeInstance: boolean, msg: MenuModeRespData): MenuModeRespData.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MenuModeRespData, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MenuModeRespData;
    static deserializeBinaryFromReader(message: MenuModeRespData, reader: jspb.BinaryReader): MenuModeRespData;
}

export namespace MenuModeRespData {
    export type AsObject = {
    }
}
