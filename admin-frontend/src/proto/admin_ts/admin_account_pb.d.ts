// package: admin
// file: admin_account.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as admin_menu_pb from "./admin_menu_pb";

export class AdminInfo extends jspb.Message { 
    getAdminid(): number;
    setAdminid(value: number): AdminInfo;
    getUsername(): string;
    setUsername(value: string): AdminInfo;
    getNickname(): string;
    setNickname(value: string): AdminInfo;
    getAvatar(): string;
    setAvatar(value: string): AdminInfo;
    getEmail(): string;
    setEmail(value: string): AdminInfo;
    getCreatetime(): string;
    setCreatetime(value: string): AdminInfo;
    getModifytime(): string;
    setModifytime(value: string): AdminInfo;
    getLastlogintime(): string;
    setLastlogintime(value: string): AdminInfo;
    getLastloginip(): string;
    setLastloginip(value: string): AdminInfo;
    getLogintotal(): number;
    setLogintotal(value: number): AdminInfo;
    getEnabled(): boolean;
    setEnabled(value: boolean): AdminInfo;
    getToken(): string;
    setToken(value: string): AdminInfo;
    getExpire(): number;
    setExpire(value: number): AdminInfo;
    clearMenusList(): void;
    getMenusList(): Array<admin_menu_pb.MenuItem>;
    setMenusList(value: Array<admin_menu_pb.MenuItem>): AdminInfo;
    addMenus(value?: admin_menu_pb.MenuItem, index?: number): admin_menu_pb.MenuItem;

    getPermissionsMap(): jspb.Map<string, string>;
    clearPermissionsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminInfo.AsObject;
    static toObject(includeInstance: boolean, msg: AdminInfo): AdminInfo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminInfo;
    static deserializeBinaryFromReader(message: AdminInfo, reader: jspb.BinaryReader): AdminInfo;
}

export namespace AdminInfo {
    export type AsObject = {
        adminid: number,
        username: string,
        nickname: string,
        avatar: string,
        email: string,
        createtime: string,
        modifytime: string,
        lastlogintime: string,
        lastloginip: string,
        logintotal: number,
        enabled: boolean,
        token: string,
        expire: number,
        menusList: Array<admin_menu_pb.MenuItem.AsObject>,

        permissionsMap: Array<[string, string]>,
    }
}

export class LoginReq extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): LoginReq;
    getPassword(): string;
    setPassword(value: string): LoginReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LoginReq.AsObject;
    static toObject(includeInstance: boolean, msg: LoginReq): LoginReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LoginReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LoginReq;
    static deserializeBinaryFromReader(message: LoginReq, reader: jspb.BinaryReader): LoginReq;
}

export namespace LoginReq {
    export type AsObject = {
        username: string,
        password: string,
    }
}

export class AccountDetailReq extends jspb.Message { 
    getRefreshtoken(): boolean;
    setRefreshtoken(value: boolean): AccountDetailReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AccountDetailReq.AsObject;
    static toObject(includeInstance: boolean, msg: AccountDetailReq): AccountDetailReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AccountDetailReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AccountDetailReq;
    static deserializeBinaryFromReader(message: AccountDetailReq, reader: jspb.BinaryReader): AccountDetailReq;
}

export namespace AccountDetailReq {
    export type AsObject = {
        refreshtoken: boolean,
    }
}

export class AccountEditReq extends jspb.Message { 
    getNickname(): string;
    setNickname(value: string): AccountEditReq;
    getAvatar(): string;
    setAvatar(value: string): AccountEditReq;
    getEmail(): string;
    setEmail(value: string): AccountEditReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AccountEditReq.AsObject;
    static toObject(includeInstance: boolean, msg: AccountEditReq): AccountEditReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AccountEditReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AccountEditReq;
    static deserializeBinaryFromReader(message: AccountEditReq, reader: jspb.BinaryReader): AccountEditReq;
}

export namespace AccountEditReq {
    export type AsObject = {
        nickname: string,
        avatar: string,
        email: string,
    }
}

export class AccountEditResp extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AccountEditResp.AsObject;
    static toObject(includeInstance: boolean, msg: AccountEditResp): AccountEditResp.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AccountEditResp, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AccountEditResp;
    static deserializeBinaryFromReader(message: AccountEditResp, reader: jspb.BinaryReader): AccountEditResp;
}

export namespace AccountEditResp {
    export type AsObject = {
    }
}

export class AccountPasswordEditReq extends jspb.Message { 
    getOldpassword(): string;
    setOldpassword(value: string): AccountPasswordEditReq;
    getPassword(): string;
    setPassword(value: string): AccountPasswordEditReq;
    getConfirmpassword(): string;
    setConfirmpassword(value: string): AccountPasswordEditReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AccountPasswordEditReq.AsObject;
    static toObject(includeInstance: boolean, msg: AccountPasswordEditReq): AccountPasswordEditReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AccountPasswordEditReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AccountPasswordEditReq;
    static deserializeBinaryFromReader(message: AccountPasswordEditReq, reader: jspb.BinaryReader): AccountPasswordEditReq;
}

export namespace AccountPasswordEditReq {
    export type AsObject = {
        oldpassword: string,
        password: string,
        confirmpassword: string,
    }
}

export class AccountPasswordEditResp extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AccountPasswordEditResp.AsObject;
    static toObject(includeInstance: boolean, msg: AccountPasswordEditResp): AccountPasswordEditResp.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AccountPasswordEditResp, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AccountPasswordEditResp;
    static deserializeBinaryFromReader(message: AccountPasswordEditResp, reader: jspb.BinaryReader): AccountPasswordEditResp;
}

export namespace AccountPasswordEditResp {
    export type AsObject = {
    }
}

export class AccountPermissionReq extends jspb.Message { 
    getMenuid(): number;
    setMenuid(value: number): AccountPermissionReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AccountPermissionReq.AsObject;
    static toObject(includeInstance: boolean, msg: AccountPermissionReq): AccountPermissionReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AccountPermissionReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AccountPermissionReq;
    static deserializeBinaryFromReader(message: AccountPermissionReq, reader: jspb.BinaryReader): AccountPermissionReq;
}

export namespace AccountPermissionReq {
    export type AsObject = {
        menuid: number,
    }
}

export class AccountPermissionResp extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AccountPermissionResp.AsObject;
    static toObject(includeInstance: boolean, msg: AccountPermissionResp): AccountPermissionResp.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AccountPermissionResp, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AccountPermissionResp;
    static deserializeBinaryFromReader(message: AccountPermissionResp, reader: jspb.BinaryReader): AccountPermissionResp;
}

export namespace AccountPermissionResp {
    export type AsObject = {
    }
}
