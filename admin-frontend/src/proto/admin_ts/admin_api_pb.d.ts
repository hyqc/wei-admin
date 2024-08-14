// package: admin
// file: admin_api.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as common_pb from "./common_pb";

export class ApiListReq extends jspb.Message { 

    hasBase(): boolean;
    clearBase(): void;
    getBase(): common_pb.ListBaseReq | undefined;
    setBase(value?: common_pb.ListBaseReq): ApiListReq;
    getKey(): string;
    setKey(value: string): ApiListReq;
    getName(): string;
    setName(value: string): ApiListReq;
    getPath(): string;
    setPath(value: string): ApiListReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ApiListReq.AsObject;
    static toObject(includeInstance: boolean, msg: ApiListReq): ApiListReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ApiListReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ApiListReq;
    static deserializeBinaryFromReader(message: ApiListReq, reader: jspb.BinaryReader): ApiListReq;
}

export namespace ApiListReq {
    export type AsObject = {
        base?: common_pb.ListBaseReq.AsObject,
        key: string,
        name: string,
        path: string,
    }
}

export class ApiItem extends jspb.Message { 
    getId(): number;
    setId(value: number): ApiItem;
    getPath(): string;
    setPath(value: string): ApiItem;
    getKey(): string;
    setKey(value: string): ApiItem;
    getName(): string;
    setName(value: string): ApiItem;
    getEnabled(): boolean;
    setEnabled(value: boolean): ApiItem;
    getPermissionid(): number;
    setPermissionid(value: number): ApiItem;
    getCreatedat(): number;
    setCreatedat(value: number): ApiItem;
    getUpdatedat(): number;
    setUpdatedat(value: number): ApiItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ApiItem.AsObject;
    static toObject(includeInstance: boolean, msg: ApiItem): ApiItem.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ApiItem, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ApiItem;
    static deserializeBinaryFromReader(message: ApiItem, reader: jspb.BinaryReader): ApiItem;
}

export namespace ApiItem {
    export type AsObject = {
        id: number,
        path: string,
        key: string,
        name: string,
        enabled: boolean,
        permissionid: number,
        createdat: number,
        updatedat: number,
    }
}

export class ApiListResp extends jspb.Message { 
    getTotal(): number;
    setTotal(value: number): ApiListResp;
    clearRowsList(): void;
    getRowsList(): Array<ApiItem>;
    setRowsList(value: Array<ApiItem>): ApiListResp;
    addRows(value?: ApiItem, index?: number): ApiItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ApiListResp.AsObject;
    static toObject(includeInstance: boolean, msg: ApiListResp): ApiListResp.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ApiListResp, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ApiListResp;
    static deserializeBinaryFromReader(message: ApiListResp, reader: jspb.BinaryReader): ApiListResp;
}

export namespace ApiListResp {
    export type AsObject = {
        total: number,
        rowsList: Array<ApiItem.AsObject>,
    }
}

export class ApiAllReq extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ApiAllReq.AsObject;
    static toObject(includeInstance: boolean, msg: ApiAllReq): ApiAllReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ApiAllReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ApiAllReq;
    static deserializeBinaryFromReader(message: ApiAllReq, reader: jspb.BinaryReader): ApiAllReq;
}

export namespace ApiAllReq {
    export type AsObject = {
    }
}

export class ApiAddReq extends jspb.Message { 
    getPath(): string;
    setPath(value: string): ApiAddReq;
    getKey(): string;
    setKey(value: string): ApiAddReq;
    getName(): string;
    setName(value: string): ApiAddReq;
    getDescribe(): string;
    setDescribe(value: string): ApiAddReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): ApiAddReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ApiAddReq.AsObject;
    static toObject(includeInstance: boolean, msg: ApiAddReq): ApiAddReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ApiAddReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ApiAddReq;
    static deserializeBinaryFromReader(message: ApiAddReq, reader: jspb.BinaryReader): ApiAddReq;
}

export namespace ApiAddReq {
    export type AsObject = {
        path: string,
        key: string,
        name: string,
        describe: string,
        enabled: boolean,
    }
}

export class ApiInfoReq extends jspb.Message { 
    getId(): number;
    setId(value: number): ApiInfoReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ApiInfoReq.AsObject;
    static toObject(includeInstance: boolean, msg: ApiInfoReq): ApiInfoReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ApiInfoReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ApiInfoReq;
    static deserializeBinaryFromReader(message: ApiInfoReq, reader: jspb.BinaryReader): ApiInfoReq;
}

export namespace ApiInfoReq {
    export type AsObject = {
        id: number,
    }
}

export class ApiEditReq extends jspb.Message { 
    getId(): number;
    setId(value: number): ApiEditReq;
    getPath(): string;
    setPath(value: string): ApiEditReq;
    getKey(): string;
    setKey(value: string): ApiEditReq;
    getName(): string;
    setName(value: string): ApiEditReq;
    getDescribe(): string;
    setDescribe(value: string): ApiEditReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): ApiEditReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ApiEditReq.AsObject;
    static toObject(includeInstance: boolean, msg: ApiEditReq): ApiEditReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ApiEditReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ApiEditReq;
    static deserializeBinaryFromReader(message: ApiEditReq, reader: jspb.BinaryReader): ApiEditReq;
}

export namespace ApiEditReq {
    export type AsObject = {
        id: number,
        path: string,
        key: string,
        name: string,
        describe: string,
        enabled: boolean,
    }
}

export class ApiEnableReq extends jspb.Message { 
    getId(): number;
    setId(value: number): ApiEnableReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): ApiEnableReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ApiEnableReq.AsObject;
    static toObject(includeInstance: boolean, msg: ApiEnableReq): ApiEnableReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ApiEnableReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ApiEnableReq;
    static deserializeBinaryFromReader(message: ApiEnableReq, reader: jspb.BinaryReader): ApiEnableReq;
}

export namespace ApiEnableReq {
    export type AsObject = {
        id: number,
        enabled: boolean,
    }
}

export class ApiDeleteReq extends jspb.Message { 
    getId(): number;
    setId(value: number): ApiDeleteReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ApiDeleteReq.AsObject;
    static toObject(includeInstance: boolean, msg: ApiDeleteReq): ApiDeleteReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ApiDeleteReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ApiDeleteReq;
    static deserializeBinaryFromReader(message: ApiDeleteReq, reader: jspb.BinaryReader): ApiDeleteReq;
}

export namespace ApiDeleteReq {
    export type AsObject = {
        id: number,
    }
}
