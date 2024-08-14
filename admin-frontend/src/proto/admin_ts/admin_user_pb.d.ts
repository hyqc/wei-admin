// package: admin
// file: admin_user.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as common_pb from "./common_pb";

export class AdminUserListReq extends jspb.Message { 

    hasBase(): boolean;
    clearBase(): void;
    getBase(): common_pb.ListBaseReq | undefined;
    setBase(value?: common_pb.ListBaseReq): AdminUserListReq;
    getUsername(): string;
    setUsername(value: string): AdminUserListReq;
    getNickname(): string;
    setNickname(value: string): AdminUserListReq;
    getEmail(): string;
    setEmail(value: string): AdminUserListReq;
    clearRoleidsList(): void;
    getRoleidsList(): Array<number>;
    setRoleidsList(value: Array<number>): AdminUserListReq;
    addRoleids(value: number, index?: number): number;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminUserListReq.AsObject;
    static toObject(includeInstance: boolean, msg: AdminUserListReq): AdminUserListReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminUserListReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminUserListReq;
    static deserializeBinaryFromReader(message: AdminUserListReq, reader: jspb.BinaryReader): AdminUserListReq;
}

export namespace AdminUserListReq {
    export type AsObject = {
        base?: common_pb.ListBaseReq.AsObject,
        username: string,
        nickname: string,
        email: string,
        roleidsList: Array<number>,
    }
}

export class AdminUserListItem extends jspb.Message { 
    getAdminid(): number;
    setAdminid(value: number): AdminUserListItem;
    getUsername(): string;
    setUsername(value: string): AdminUserListItem;
    getEmail(): string;
    setEmail(value: string): AdminUserListItem;
    getAvatar(): string;
    setAvatar(value: string): AdminUserListItem;
    getEnabled(): boolean;
    setEnabled(value: boolean): AdminUserListItem;
    getEnabledtext(): string;
    setEnabledtext(value: string): AdminUserListItem;
    getCreatedat(): string;
    setCreatedat(value: string): AdminUserListItem;
    getUpdatedat(): string;
    setUpdatedat(value: string): AdminUserListItem;
    getLogintotal(): number;
    setLogintotal(value: number): AdminUserListItem;
    getLastloginip(): string;
    setLastloginip(value: string): AdminUserListItem;
    getLastlogintime(): string;
    setLastlogintime(value: string): AdminUserListItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminUserListItem.AsObject;
    static toObject(includeInstance: boolean, msg: AdminUserListItem): AdminUserListItem.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminUserListItem, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminUserListItem;
    static deserializeBinaryFromReader(message: AdminUserListItem, reader: jspb.BinaryReader): AdminUserListItem;
}

export namespace AdminUserListItem {
    export type AsObject = {
        adminid: number,
        username: string,
        email: string,
        avatar: string,
        enabled: boolean,
        enabledtext: string,
        createdat: string,
        updatedat: string,
        logintotal: number,
        lastloginip: string,
        lastlogintime: string,
    }
}

export class AdminUserListResp extends jspb.Message { 
    getTotal(): number;
    setTotal(value: number): AdminUserListResp;
    clearRowsList(): void;
    getRowsList(): Array<AdminUserListItem>;
    setRowsList(value: Array<AdminUserListItem>): AdminUserListResp;
    addRows(value?: AdminUserListItem, index?: number): AdminUserListItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminUserListResp.AsObject;
    static toObject(includeInstance: boolean, msg: AdminUserListResp): AdminUserListResp.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminUserListResp, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminUserListResp;
    static deserializeBinaryFromReader(message: AdminUserListResp, reader: jspb.BinaryReader): AdminUserListResp;
}

export namespace AdminUserListResp {
    export type AsObject = {
        total: number,
        rowsList: Array<AdminUserListItem.AsObject>,
    }
}

export class AdminUserInfoReq extends jspb.Message { 
    getAdminid(): number;
    setAdminid(value: number): AdminUserInfoReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminUserInfoReq.AsObject;
    static toObject(includeInstance: boolean, msg: AdminUserInfoReq): AdminUserInfoReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminUserInfoReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminUserInfoReq;
    static deserializeBinaryFromReader(message: AdminUserInfoReq, reader: jspb.BinaryReader): AdminUserInfoReq;
}

export namespace AdminUserInfoReq {
    export type AsObject = {
        adminid: number,
    }
}

export class AdminUserInfo extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminUserInfo.AsObject;
    static toObject(includeInstance: boolean, msg: AdminUserInfo): AdminUserInfo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminUserInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminUserInfo;
    static deserializeBinaryFromReader(message: AdminUserInfo, reader: jspb.BinaryReader): AdminUserInfo;
}

export namespace AdminUserInfo {
    export type AsObject = {
    }
}

export class AdminUserAddReq extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): AdminUserAddReq;
    getNickname(): string;
    setNickname(value: string): AdminUserAddReq;
    getPassword(): string;
    setPassword(value: string): AdminUserAddReq;
    getConfirmpassword(): string;
    setConfirmpassword(value: string): AdminUserAddReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): AdminUserAddReq;
    getEmail(): string;
    setEmail(value: string): AdminUserAddReq;
    getAvatar(): string;
    setAvatar(value: string): AdminUserAddReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminUserAddReq.AsObject;
    static toObject(includeInstance: boolean, msg: AdminUserAddReq): AdminUserAddReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminUserAddReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminUserAddReq;
    static deserializeBinaryFromReader(message: AdminUserAddReq, reader: jspb.BinaryReader): AdminUserAddReq;
}

export namespace AdminUserAddReq {
    export type AsObject = {
        username: string,
        nickname: string,
        password: string,
        confirmpassword: string,
        enabled: boolean,
        email: string,
        avatar: string,
    }
}

export class AdminUserEditReq extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): AdminUserEditReq;
    getNickname(): string;
    setNickname(value: string): AdminUserEditReq;
    getPassword(): string;
    setPassword(value: string): AdminUserEditReq;
    getConfirmpassword(): string;
    setConfirmpassword(value: string): AdminUserEditReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): AdminUserEditReq;
    getEmail(): string;
    setEmail(value: string): AdminUserEditReq;
    getAvatar(): string;
    setAvatar(value: string): AdminUserEditReq;
    getAdminid(): number;
    setAdminid(value: number): AdminUserEditReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminUserEditReq.AsObject;
    static toObject(includeInstance: boolean, msg: AdminUserEditReq): AdminUserEditReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminUserEditReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminUserEditReq;
    static deserializeBinaryFromReader(message: AdminUserEditReq, reader: jspb.BinaryReader): AdminUserEditReq;
}

export namespace AdminUserEditReq {
    export type AsObject = {
        username: string,
        nickname: string,
        password: string,
        confirmpassword: string,
        enabled: boolean,
        email: string,
        avatar: string,
        adminid: number,
    }
}

export class AdminUserEnabledReq extends jspb.Message { 
    getAdminid(): number;
    setAdminid(value: number): AdminUserEnabledReq;
    getEnabled(): boolean;
    setEnabled(value: boolean): AdminUserEnabledReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminUserEnabledReq.AsObject;
    static toObject(includeInstance: boolean, msg: AdminUserEnabledReq): AdminUserEnabledReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminUserEnabledReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminUserEnabledReq;
    static deserializeBinaryFromReader(message: AdminUserEnabledReq, reader: jspb.BinaryReader): AdminUserEnabledReq;
}

export namespace AdminUserEnabledReq {
    export type AsObject = {
        adminid: number,
        enabled: boolean,
    }
}

export class AdminUserDeleteReq extends jspb.Message { 
    getAdminid(): number;
    setAdminid(value: number): AdminUserDeleteReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminUserDeleteReq.AsObject;
    static toObject(includeInstance: boolean, msg: AdminUserDeleteReq): AdminUserDeleteReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminUserDeleteReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminUserDeleteReq;
    static deserializeBinaryFromReader(message: AdminUserDeleteReq, reader: jspb.BinaryReader): AdminUserDeleteReq;
}

export namespace AdminUserDeleteReq {
    export type AsObject = {
        adminid: number,
    }
}

export class AdminUserBindRolesReq extends jspb.Message { 
    getAdminid(): number;
    setAdminid(value: number): AdminUserBindRolesReq;
    clearRoleidsList(): void;
    getRoleidsList(): Array<number>;
    setRoleidsList(value: Array<number>): AdminUserBindRolesReq;
    addRoleids(value: number, index?: number): number;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AdminUserBindRolesReq.AsObject;
    static toObject(includeInstance: boolean, msg: AdminUserBindRolesReq): AdminUserBindRolesReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AdminUserBindRolesReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AdminUserBindRolesReq;
    static deserializeBinaryFromReader(message: AdminUserBindRolesReq, reader: jspb.BinaryReader): AdminUserBindRolesReq;
}

export namespace AdminUserBindRolesReq {
    export type AsObject = {
        adminid: number,
        roleidsList: Array<number>,
    }
}
