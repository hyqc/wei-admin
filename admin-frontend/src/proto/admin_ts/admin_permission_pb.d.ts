// package: admin
// file: admin_permission.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as common_pb from "./common_pb";
import * as admin_api_pb from "./admin_api_pb";

export class PermissionListItem extends jspb.Message { 
    getId(): number;
    setId(value: number): PermissionListItem;
    getMenuid(): number;
    setMenuid(value: number): PermissionListItem;
    getMenuname(): string;
    setMenuname(value: string): PermissionListItem;
    getMenupath(): string;
    setMenupath(value: string): PermissionListItem;
    clearApisList(): void;
    getApisList(): Array<admin_api_pb.ApiItem>;
    setApisList(value: Array<admin_api_pb.ApiItem>): PermissionListItem;
    addApis(value?: admin_api_pb.ApiItem, index?: number): admin_api_pb.ApiItem;
    getKey(): string;
    setKey(value: string): PermissionListItem;
    getName(): string;
    setName(value: string): PermissionListItem;
    getDescribe(): string;
    setDescribe(value: string): PermissionListItem;
    getType(): string;
    setType(value: string): PermissionListItem;
    getTypetext(): string;
    setTypetext(value: string): PermissionListItem;
    getEnabled(): boolean;
    setEnabled(value: boolean): PermissionListItem;
    getCreatedat(): string;
    setCreatedat(value: string): PermissionListItem;
    getUpdatedat(): string;
    setUpdatedat(value: string): PermissionListItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionListItem.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionListItem): PermissionListItem.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionListItem, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionListItem;
    static deserializeBinaryFromReader(message: PermissionListItem, reader: jspb.BinaryReader): PermissionListItem;
}

export namespace PermissionListItem {
    export type AsObject = {
        id: number,
        menuid: number,
        menuname: string,
        menupath: string,
        apisList: Array<admin_api_pb.ApiItem.AsObject>,
        key: string,
        name: string,
        describe: string,
        type: string,
        typetext: string,
        enabled: boolean,
        createdat: string,
        updatedat: string,
    }
}

export class PermissionListReq extends jspb.Message { 

    hasBase(): boolean;
    clearBase(): void;
    getBase(): common_pb.ListBaseReq | undefined;
    setBase(value?: common_pb.ListBaseReq): PermissionListReq;
    getMenuid(): number;
    setMenuid(value: number): PermissionListReq;
    getKey(): string;
    setKey(value: string): PermissionListReq;
    getName(): string;
    setName(value: string): PermissionListReq;
    getType(): string;
    setType(value: string): PermissionListReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionListReq.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionListReq): PermissionListReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionListReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionListReq;
    static deserializeBinaryFromReader(message: PermissionListReq, reader: jspb.BinaryReader): PermissionListReq;
}

export namespace PermissionListReq {
    export type AsObject = {
        base?: common_pb.ListBaseReq.AsObject,
        menuid: number,
        key: string,
        name: string,
        type: string,
    }
}

export class PermissionListRespData extends jspb.Message { 
    getTotal(): number;
    setTotal(value: number): PermissionListRespData;
    clearRowsList(): void;
    getRowsList(): Array<PermissionListItem>;
    setRowsList(value: Array<PermissionListItem>): PermissionListRespData;
    addRows(value?: PermissionListItem, index?: number): PermissionListItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionListRespData.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionListRespData): PermissionListRespData.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionListRespData, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionListRespData;
    static deserializeBinaryFromReader(message: PermissionListRespData, reader: jspb.BinaryReader): PermissionListRespData;
}

export namespace PermissionListRespData {
    export type AsObject = {
        total: number,
        rowsList: Array<PermissionListItem.AsObject>,
    }
}

export class PermissionAddReq extends jspb.Message { 
    getMenuid(): number;
    setMenuid(value: number): PermissionAddReq;
    getKey(): string;
    setKey(value: string): PermissionAddReq;
    getName(): string;
    setName(value: string): PermissionAddReq;
    getDescribe(): string;
    setDescribe(value: string): PermissionAddReq;
    getType(): string;
    setType(value: string): PermissionAddReq;
    getRedirect(): string;
    setRedirect(value: string): PermissionAddReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): PermissionAddReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionAddReq.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionAddReq): PermissionAddReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionAddReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionAddReq;
    static deserializeBinaryFromReader(message: PermissionAddReq, reader: jspb.BinaryReader): PermissionAddReq;
}

export namespace PermissionAddReq {
    export type AsObject = {
        menuid: number,
        key: string,
        name: string,
        describe: string,
        type: string,
        redirect: string,
        enabled: boolean,
    }
}

export class PermissionInfoReq extends jspb.Message { 
    getId(): number;
    setId(value: number): PermissionInfoReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionInfoReq.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionInfoReq): PermissionInfoReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionInfoReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionInfoReq;
    static deserializeBinaryFromReader(message: PermissionInfoReq, reader: jspb.BinaryReader): PermissionInfoReq;
}

export namespace PermissionInfoReq {
    export type AsObject = {
        id: number,
    }
}

export class PermissionInfo extends jspb.Message { 
    getId(): number;
    setId(value: number): PermissionInfo;
    getMenuid(): number;
    setMenuid(value: number): PermissionInfo;
    getMenuname(): string;
    setMenuname(value: string): PermissionInfo;
    getMenupath(): string;
    setMenupath(value: string): PermissionInfo;
    clearApisList(): void;
    getApisList(): Array<admin_api_pb.ApiItem>;
    setApisList(value: Array<admin_api_pb.ApiItem>): PermissionInfo;
    addApis(value?: admin_api_pb.ApiItem, index?: number): admin_api_pb.ApiItem;
    getKey(): string;
    setKey(value: string): PermissionInfo;
    getName(): string;
    setName(value: string): PermissionInfo;
    getDescribe(): string;
    setDescribe(value: string): PermissionInfo;
    getType(): string;
    setType(value: string): PermissionInfo;
    getTypetext(): string;
    setTypetext(value: string): PermissionInfo;
    getRedirect(): string;
    setRedirect(value: string): PermissionInfo;
    getEnabled(): boolean;
    setEnabled(value: boolean): PermissionInfo;
    getCreatedat(): string;
    setCreatedat(value: string): PermissionInfo;
    getUpdatedat(): string;
    setUpdatedat(value: string): PermissionInfo;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionInfo.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionInfo): PermissionInfo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionInfo;
    static deserializeBinaryFromReader(message: PermissionInfo, reader: jspb.BinaryReader): PermissionInfo;
}

export namespace PermissionInfo {
    export type AsObject = {
        id: number,
        menuid: number,
        menuname: string,
        menupath: string,
        apisList: Array<admin_api_pb.ApiItem.AsObject>,
        key: string,
        name: string,
        describe: string,
        type: string,
        typetext: string,
        redirect: string,
        enabled: boolean,
        createdat: string,
        updatedat: string,
    }
}

export class PermissionEditReq extends jspb.Message { 
    getId(): number;
    setId(value: number): PermissionEditReq;
    getMenuid(): number;
    setMenuid(value: number): PermissionEditReq;
    getKey(): string;
    setKey(value: string): PermissionEditReq;
    getName(): string;
    setName(value: string): PermissionEditReq;
    getDescribe(): string;
    setDescribe(value: string): PermissionEditReq;
    getType(): string;
    setType(value: string): PermissionEditReq;
    getRedirect(): string;
    setRedirect(value: string): PermissionEditReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): PermissionEditReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionEditReq.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionEditReq): PermissionEditReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionEditReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionEditReq;
    static deserializeBinaryFromReader(message: PermissionEditReq, reader: jspb.BinaryReader): PermissionEditReq;
}

export namespace PermissionEditReq {
    export type AsObject = {
        id: number,
        menuid: number,
        key: string,
        name: string,
        describe: string,
        type: string,
        redirect: string,
        enabled: boolean,
    }
}

export class PermissionEnableReq extends jspb.Message { 
    getId(): number;
    setId(value: number): PermissionEnableReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): PermissionEnableReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionEnableReq.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionEnableReq): PermissionEnableReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionEnableReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionEnableReq;
    static deserializeBinaryFromReader(message: PermissionEnableReq, reader: jspb.BinaryReader): PermissionEnableReq;
}

export namespace PermissionEnableReq {
    export type AsObject = {
        id: number,
        enabled: boolean,
    }
}

export class PermissionDeleteReq extends jspb.Message { 
    getId(): number;
    setId(value: number): PermissionDeleteReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionDeleteReq.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionDeleteReq): PermissionDeleteReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionDeleteReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionDeleteReq;
    static deserializeBinaryFromReader(message: PermissionDeleteReq, reader: jspb.BinaryReader): PermissionDeleteReq;
}

export namespace PermissionDeleteReq {
    export type AsObject = {
        id: number,
    }
}

export class PermissionBindApisReq extends jspb.Message { 
    getPermissionid(): number;
    setPermissionid(value: number): PermissionBindApisReq;
    clearApiidsList(): void;
    getApiidsList(): Array<number>;
    setApiidsList(value: Array<number>): PermissionBindApisReq;
    addApiids(value: number, index?: number): number;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionBindApisReq.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionBindApisReq): PermissionBindApisReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionBindApisReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionBindApisReq;
    static deserializeBinaryFromReader(message: PermissionBindApisReq, reader: jspb.BinaryReader): PermissionBindApisReq;
}

export namespace PermissionBindApisReq {
    export type AsObject = {
        permissionid: number,
        apiidsList: Array<number>,
    }
}

export class PermissionBindMenuReq extends jspb.Message { 
    getMenuid(): number;
    setMenuid(value: number): PermissionBindMenuReq;
    clearPermissionsList(): void;
    getPermissionsList(): Array<PermissionAddReq>;
    setPermissionsList(value: Array<PermissionAddReq>): PermissionBindMenuReq;
    addPermissions(value?: PermissionAddReq, index?: number): PermissionAddReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionBindMenuReq.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionBindMenuReq): PermissionBindMenuReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionBindMenuReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionBindMenuReq;
    static deserializeBinaryFromReader(message: PermissionBindMenuReq, reader: jspb.BinaryReader): PermissionBindMenuReq;
}

export namespace PermissionBindMenuReq {
    export type AsObject = {
        menuid: number,
        permissionsList: Array<PermissionAddReq.AsObject>,
    }
}

export class PermissionApiItem extends jspb.Message { 
    getId(): number;
    setId(value: number): PermissionApiItem;
    getMenuid(): number;
    setMenuid(value: number): PermissionApiItem;
    getKey(): string;
    setKey(value: string): PermissionApiItem;
    getType(): string;
    setType(value: string): PermissionApiItem;
    getTypetext(): string;
    setTypetext(value: string): PermissionApiItem;
    getName(): string;
    setName(value: string): PermissionApiItem;
    clearApisList(): void;
    getApisList(): Array<admin_api_pb.ApiItem>;
    setApisList(value: Array<admin_api_pb.ApiItem>): PermissionApiItem;
    addApis(value?: admin_api_pb.ApiItem, index?: number): admin_api_pb.ApiItem;
    getEnabled(): boolean;
    setEnabled(value: boolean): PermissionApiItem;
    getDescribe(): string;
    setDescribe(value: string): PermissionApiItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionApiItem.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionApiItem): PermissionApiItem.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionApiItem, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionApiItem;
    static deserializeBinaryFromReader(message: PermissionApiItem, reader: jspb.BinaryReader): PermissionApiItem;
}

export namespace PermissionApiItem {
    export type AsObject = {
        id: number,
        menuid: number,
        key: string,
        type: string,
        typetext: string,
        name: string,
        apisList: Array<admin_api_pb.ApiItem.AsObject>,
        enabled: boolean,
        describe: string,
    }
}
