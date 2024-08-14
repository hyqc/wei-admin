// package: admin
// file: admin_role.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as common_pb from "./common_pb";

export class RoleItem extends jspb.Message { 
    getId(): number;
    setId(value: number): RoleItem;
    getName(): string;
    setName(value: string): RoleItem;
    getRolename(): string;
    setRolename(value: string): RoleItem;
    getCreateadminid(): number;
    setCreateadminid(value: number): RoleItem;
    getCreateadminname(): string;
    setCreateadminname(value: string): RoleItem;
    getEnabled(): boolean;
    setEnabled(value: boolean): RoleItem;
    getCreatedat(): string;
    setCreatedat(value: string): RoleItem;
    getUpdatedat(): string;
    setUpdatedat(value: string): RoleItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleItem.AsObject;
    static toObject(includeInstance: boolean, msg: RoleItem): RoleItem.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleItem, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleItem;
    static deserializeBinaryFromReader(message: RoleItem, reader: jspb.BinaryReader): RoleItem;
}

export namespace RoleItem {
    export type AsObject = {
        id: number,
        name: string,
        rolename: string,
        createadminid: number,
        createadminname: string,
        enabled: boolean,
        createdat: string,
        updatedat: string,
    }
}

export class RoleListReq extends jspb.Message { 

    hasBase(): boolean;
    clearBase(): void;
    getBase(): common_pb.ListBaseReq | undefined;
    setBase(value?: common_pb.ListBaseReq): RoleListReq;
    getId(): number;
    setId(value: number): RoleListReq;
    getName(): string;
    setName(value: string): RoleListReq;
    getEnabled(): number;
    setEnabled(value: number): RoleListReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleListReq.AsObject;
    static toObject(includeInstance: boolean, msg: RoleListReq): RoleListReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleListReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleListReq;
    static deserializeBinaryFromReader(message: RoleListReq, reader: jspb.BinaryReader): RoleListReq;
}

export namespace RoleListReq {
    export type AsObject = {
        base?: common_pb.ListBaseReq.AsObject,
        id: number,
        name: string,
        enabled: number,
    }
}

export class RoleListRespData extends jspb.Message { 
    getTotal(): number;
    setTotal(value: number): RoleListRespData;
    clearRowsList(): void;
    getRowsList(): Array<RoleItem>;
    setRowsList(value: Array<RoleItem>): RoleListRespData;
    addRows(value?: RoleItem, index?: number): RoleItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleListRespData.AsObject;
    static toObject(includeInstance: boolean, msg: RoleListRespData): RoleListRespData.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleListRespData, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleListRespData;
    static deserializeBinaryFromReader(message: RoleListRespData, reader: jspb.BinaryReader): RoleListRespData;
}

export namespace RoleListRespData {
    export type AsObject = {
        total: number,
        rowsList: Array<RoleItem.AsObject>,
    }
}

export class RoleAllReq extends jspb.Message { 
    getId(): number;
    setId(value: number): RoleAllReq;
    getName(): string;
    setName(value: string): RoleAllReq;
    getEnabled(): number;
    setEnabled(value: number): RoleAllReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleAllReq.AsObject;
    static toObject(includeInstance: boolean, msg: RoleAllReq): RoleAllReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleAllReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleAllReq;
    static deserializeBinaryFromReader(message: RoleAllReq, reader: jspb.BinaryReader): RoleAllReq;
}

export namespace RoleAllReq {
    export type AsObject = {
        id: number,
        name: string,
        enabled: number,
    }
}

export class RoleAllResp extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleAllResp.AsObject;
    static toObject(includeInstance: boolean, msg: RoleAllResp): RoleAllResp.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleAllResp, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleAllResp;
    static deserializeBinaryFromReader(message: RoleAllResp, reader: jspb.BinaryReader): RoleAllResp;
}

export namespace RoleAllResp {
    export type AsObject = {
    }
}

export class RoleAddReq extends jspb.Message { 
    getName(): string;
    setName(value: string): RoleAddReq;
    getDescribe(): string;
    setDescribe(value: string): RoleAddReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): RoleAddReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleAddReq.AsObject;
    static toObject(includeInstance: boolean, msg: RoleAddReq): RoleAddReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleAddReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleAddReq;
    static deserializeBinaryFromReader(message: RoleAddReq, reader: jspb.BinaryReader): RoleAddReq;
}

export namespace RoleAddReq {
    export type AsObject = {
        name: string,
        describe: string,
        enabled: boolean,
    }
}

export class RoleAddResp extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleAddResp.AsObject;
    static toObject(includeInstance: boolean, msg: RoleAddResp): RoleAddResp.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleAddResp, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleAddResp;
    static deserializeBinaryFromReader(message: RoleAddResp, reader: jspb.BinaryReader): RoleAddResp;
}

export namespace RoleAddResp {
    export type AsObject = {
    }
}

export class RoleInfoReq extends jspb.Message { 
    getId(): number;
    setId(value: number): RoleInfoReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleInfoReq.AsObject;
    static toObject(includeInstance: boolean, msg: RoleInfoReq): RoleInfoReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleInfoReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleInfoReq;
    static deserializeBinaryFromReader(message: RoleInfoReq, reader: jspb.BinaryReader): RoleInfoReq;
}

export namespace RoleInfoReq {
    export type AsObject = {
        id: number,
    }
}

export class RoleInfoRespData extends jspb.Message { 
    getId(): number;
    setId(value: number): RoleInfoRespData;
    getName(): string;
    setName(value: string): RoleInfoRespData;
    getDescribe(): string;
    setDescribe(value: string): RoleInfoRespData;
    getEnabled(): boolean;
    setEnabled(value: boolean): RoleInfoRespData;
    getCreateadminid(): number;
    setCreateadminid(value: number): RoleInfoRespData;
    getCreateadminname(): string;
    setCreateadminname(value: string): RoleInfoRespData;
    getModifyadminid(): number;
    setModifyadminid(value: number): RoleInfoRespData;
    getModifyadminname(): string;
    setModifyadminname(value: string): RoleInfoRespData;
    getCreatedat(): string;
    setCreatedat(value: string): RoleInfoRespData;
    getUpdatedat(): string;
    setUpdatedat(value: string): RoleInfoRespData;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleInfoRespData.AsObject;
    static toObject(includeInstance: boolean, msg: RoleInfoRespData): RoleInfoRespData.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleInfoRespData, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleInfoRespData;
    static deserializeBinaryFromReader(message: RoleInfoRespData, reader: jspb.BinaryReader): RoleInfoRespData;
}

export namespace RoleInfoRespData {
    export type AsObject = {
        id: number,
        name: string,
        describe: string,
        enabled: boolean,
        createadminid: number,
        createadminname: string,
        modifyadminid: number,
        modifyadminname: string,
        createdat: string,
        updatedat: string,
    }
}

export class RoleEditReq extends jspb.Message { 
    getId(): number;
    setId(value: number): RoleEditReq;
    getName(): string;
    setName(value: string): RoleEditReq;
    getDescribe(): string;
    setDescribe(value: string): RoleEditReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): RoleEditReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleEditReq.AsObject;
    static toObject(includeInstance: boolean, msg: RoleEditReq): RoleEditReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleEditReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleEditReq;
    static deserializeBinaryFromReader(message: RoleEditReq, reader: jspb.BinaryReader): RoleEditReq;
}

export namespace RoleEditReq {
    export type AsObject = {
        id: number,
        name: string,
        describe: string,
        enabled: boolean,
    }
}

export class RoleEnableReq extends jspb.Message { 
    getId(): number;
    setId(value: number): RoleEnableReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): RoleEnableReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleEnableReq.AsObject;
    static toObject(includeInstance: boolean, msg: RoleEnableReq): RoleEnableReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleEnableReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleEnableReq;
    static deserializeBinaryFromReader(message: RoleEnableReq, reader: jspb.BinaryReader): RoleEnableReq;
}

export namespace RoleEnableReq {
    export type AsObject = {
        id: number,
        enabled: boolean,
    }
}

export class RoleDeleteReq extends jspb.Message { 
    getId(): number;
    setId(value: number): RoleDeleteReq;
    clearPermissionidsList(): void;
    getPermissionidsList(): Array<number>;
    setPermissionidsList(value: Array<number>): RoleDeleteReq;
    addPermissionids(value: number, index?: number): number;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleDeleteReq.AsObject;
    static toObject(includeInstance: boolean, msg: RoleDeleteReq): RoleDeleteReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleDeleteReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleDeleteReq;
    static deserializeBinaryFromReader(message: RoleDeleteReq, reader: jspb.BinaryReader): RoleDeleteReq;
}

export namespace RoleDeleteReq {
    export type AsObject = {
        id: number,
        permissionidsList: Array<number>,
    }
}

export class RoleBindPermissionsReq extends jspb.Message { 
    getId(): number;
    setId(value: number): RoleBindPermissionsReq;
    clearPermissionidsList(): void;
    getPermissionidsList(): Array<number>;
    setPermissionidsList(value: Array<number>): RoleBindPermissionsReq;
    addPermissionids(value: number, index?: number): number;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoleBindPermissionsReq.AsObject;
    static toObject(includeInstance: boolean, msg: RoleBindPermissionsReq): RoleBindPermissionsReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoleBindPermissionsReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoleBindPermissionsReq;
    static deserializeBinaryFromReader(message: RoleBindPermissionsReq, reader: jspb.BinaryReader): RoleBindPermissionsReq;
}

export namespace RoleBindPermissionsReq {
    export type AsObject = {
        id: number,
        permissionidsList: Array<number>,
    }
}

export class RolePermissionItem extends jspb.Message { 
    getRoleid(): number;
    setRoleid(value: number): RolePermissionItem;
    getPermissionid(): number;
    setPermissionid(value: number): RolePermissionItem;
    getPermissionname(): string;
    setPermissionname(value: string): RolePermissionItem;
    getPermissionkey(): string;
    setPermissionkey(value: string): RolePermissionItem;
    getPermissiontype(): string;
    setPermissiontype(value: string): RolePermissionItem;
    getPermissiontypetext(): string;
    setPermissiontypetext(value: string): RolePermissionItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RolePermissionItem.AsObject;
    static toObject(includeInstance: boolean, msg: RolePermissionItem): RolePermissionItem.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RolePermissionItem, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RolePermissionItem;
    static deserializeBinaryFromReader(message: RolePermissionItem, reader: jspb.BinaryReader): RolePermissionItem;
}

export namespace RolePermissionItem {
    export type AsObject = {
        roleid: number,
        permissionid: number,
        permissionname: string,
        permissionkey: string,
        permissiontype: string,
        permissiontypetext: string,
    }
}

export class RolePermissionsReq extends jspb.Message { 
    getId(): number;
    setId(value: number): RolePermissionsReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RolePermissionsReq.AsObject;
    static toObject(includeInstance: boolean, msg: RolePermissionsReq): RolePermissionsReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RolePermissionsReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RolePermissionsReq;
    static deserializeBinaryFromReader(message: RolePermissionsReq, reader: jspb.BinaryReader): RolePermissionsReq;
}

export namespace RolePermissionsReq {
    export type AsObject = {
        id: number,
    }
}
