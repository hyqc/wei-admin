// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               v3.21.12
// source: common.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "admin";

export interface RespCode {
  code: number;
  msg: string;
  reason: string;
}

export interface ReqListBase {
  /** 分页偏移量 */
  pageSize: number;
  /** 页码 */
  pageNum: number;
  /** 排序字段 */
  sortField: string;
  /** 排序值 */
  sortType: string;
  /** 0：全部 1：启用 2：禁用 */
  enabled: number;
  /** 查询开始时间戳秒 */
  createStartTime: number;
  /** 查询结束时间戳秒 */
  createEndTime: number;
}

/** 接口列表返回结构 */
export interface ApiItem {
  id: number;
  path: string;
  key: string;
  name: string;
  enabled: boolean;
  permissionId: number;
  createdAt: string;
  updatedAt: string;
}

/** 账号列表返回结构 */
export interface AdminUserListItem {
  /** 管理员ID */
  adminId: number;
  /** 账号名称 */
  username: string;
  /** 昵称 */
  nickname: string;
  /** 邮箱 */
  email: string;
  /** 头像 */
  avatar: string;
  /** 登录次数 */
  loginTotal: number;
  /** 上次登录IP */
  lastLoginIp: string;
  /** 上次登录时间 */
  lastLoginTime: string;
  /** 是否启用 */
  isEnabled: boolean;
  /** 创建时间 */
  createdAt: string;
  /** 更新时间 */
  updatedAt: string;
  roles: AdminUserRoleItem[];
  isEnabledButtonDisabled: boolean;
}

/** 账号角色列表 */
export interface AdminUserRoleItem {
  /** 角色ID */
  roleId: number;
  roleName: string;
}

function createBaseRespCode(): RespCode {
  return { code: 0, msg: "", reason: "" };
}

export const RespCode: MessageFns<RespCode> = {
  encode(message: RespCode, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.msg !== "") {
      writer.uint32(18).string(message.msg);
    }
    if (message.reason !== "") {
      writer.uint32(26).string(message.reason);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): RespCode {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRespCode();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.code = reader.int32();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.msg = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.reason = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RespCode {
    return {
      code: isSet(object.code) ? globalThis.Number(object.code) : 0,
      msg: isSet(object.msg) ? globalThis.String(object.msg) : "",
      reason: isSet(object.reason) ? globalThis.String(object.reason) : "",
    };
  },

  toJSON(message: RespCode): unknown {
    const obj: any = {};
    if (message.code !== 0) {
      obj.code = Math.round(message.code);
    }
    if (message.msg !== "") {
      obj.msg = message.msg;
    }
    if (message.reason !== "") {
      obj.reason = message.reason;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RespCode>, I>>(base?: I): RespCode {
    return RespCode.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RespCode>, I>>(object: I): RespCode {
    const message = createBaseRespCode();
    message.code = object.code ?? 0;
    message.msg = object.msg ?? "";
    message.reason = object.reason ?? "";
    return message;
  },
};

function createBaseReqListBase(): ReqListBase {
  return { pageSize: 0, pageNum: 0, sortField: "", sortType: "", enabled: 0, createStartTime: 0, createEndTime: 0 };
}

export const ReqListBase: MessageFns<ReqListBase> = {
  encode(message: ReqListBase, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.pageSize !== 0) {
      writer.uint32(8).int32(message.pageSize);
    }
    if (message.pageNum !== 0) {
      writer.uint32(16).int32(message.pageNum);
    }
    if (message.sortField !== "") {
      writer.uint32(26).string(message.sortField);
    }
    if (message.sortType !== "") {
      writer.uint32(34).string(message.sortType);
    }
    if (message.enabled !== 0) {
      writer.uint32(40).int32(message.enabled);
    }
    if (message.createStartTime !== 0) {
      writer.uint32(48).int64(message.createStartTime);
    }
    if (message.createEndTime !== 0) {
      writer.uint32(56).int64(message.createEndTime);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ReqListBase {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReqListBase();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.pageSize = reader.int32();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.pageNum = reader.int32();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.sortField = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.sortType = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 40) {
            break;
          }

          message.enabled = reader.int32();
          continue;
        }
        case 6: {
          if (tag !== 48) {
            break;
          }

          message.createStartTime = longToNumber(reader.int64());
          continue;
        }
        case 7: {
          if (tag !== 56) {
            break;
          }

          message.createEndTime = longToNumber(reader.int64());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReqListBase {
    return {
      pageSize: isSet(object.pageSize) ? globalThis.Number(object.pageSize) : 0,
      pageNum: isSet(object.pageNum) ? globalThis.Number(object.pageNum) : 0,
      sortField: isSet(object.sortField) ? globalThis.String(object.sortField) : "",
      sortType: isSet(object.sortType) ? globalThis.String(object.sortType) : "",
      enabled: isSet(object.enabled) ? globalThis.Number(object.enabled) : 0,
      createStartTime: isSet(object.createStartTime) ? globalThis.Number(object.createStartTime) : 0,
      createEndTime: isSet(object.createEndTime) ? globalThis.Number(object.createEndTime) : 0,
    };
  },

  toJSON(message: ReqListBase): unknown {
    const obj: any = {};
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    if (message.pageNum !== 0) {
      obj.pageNum = Math.round(message.pageNum);
    }
    if (message.sortField !== "") {
      obj.sortField = message.sortField;
    }
    if (message.sortType !== "") {
      obj.sortType = message.sortType;
    }
    if (message.enabled !== 0) {
      obj.enabled = Math.round(message.enabled);
    }
    if (message.createStartTime !== 0) {
      obj.createStartTime = Math.round(message.createStartTime);
    }
    if (message.createEndTime !== 0) {
      obj.createEndTime = Math.round(message.createEndTime);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ReqListBase>, I>>(base?: I): ReqListBase {
    return ReqListBase.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ReqListBase>, I>>(object: I): ReqListBase {
    const message = createBaseReqListBase();
    message.pageSize = object.pageSize ?? 0;
    message.pageNum = object.pageNum ?? 0;
    message.sortField = object.sortField ?? "";
    message.sortType = object.sortType ?? "";
    message.enabled = object.enabled ?? 0;
    message.createStartTime = object.createStartTime ?? 0;
    message.createEndTime = object.createEndTime ?? 0;
    return message;
  },
};

function createBaseApiItem(): ApiItem {
  return { id: 0, path: "", key: "", name: "", enabled: false, permissionId: 0, createdAt: "", updatedAt: "" };
}

export const ApiItem: MessageFns<ApiItem> = {
  encode(message: ApiItem, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.path !== "") {
      writer.uint32(18).string(message.path);
    }
    if (message.key !== "") {
      writer.uint32(26).string(message.key);
    }
    if (message.name !== "") {
      writer.uint32(34).string(message.name);
    }
    if (message.enabled !== false) {
      writer.uint32(40).bool(message.enabled);
    }
    if (message.permissionId !== 0) {
      writer.uint32(48).int32(message.permissionId);
    }
    if (message.createdAt !== "") {
      writer.uint32(58).string(message.createdAt);
    }
    if (message.updatedAt !== "") {
      writer.uint32(66).string(message.updatedAt);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ApiItem {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseApiItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.id = reader.int32();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.path = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.key = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.name = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 40) {
            break;
          }

          message.enabled = reader.bool();
          continue;
        }
        case 6: {
          if (tag !== 48) {
            break;
          }

          message.permissionId = reader.int32();
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.createdAt = reader.string();
          continue;
        }
        case 8: {
          if (tag !== 66) {
            break;
          }

          message.updatedAt = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ApiItem {
    return {
      id: isSet(object.id) ? globalThis.Number(object.id) : 0,
      path: isSet(object.path) ? globalThis.String(object.path) : "",
      key: isSet(object.key) ? globalThis.String(object.key) : "",
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      enabled: isSet(object.enabled) ? globalThis.Boolean(object.enabled) : false,
      permissionId: isSet(object.permissionId) ? globalThis.Number(object.permissionId) : 0,
      createdAt: isSet(object.createdAt) ? globalThis.String(object.createdAt) : "",
      updatedAt: isSet(object.updatedAt) ? globalThis.String(object.updatedAt) : "",
    };
  },

  toJSON(message: ApiItem): unknown {
    const obj: any = {};
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    if (message.path !== "") {
      obj.path = message.path;
    }
    if (message.key !== "") {
      obj.key = message.key;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.enabled !== false) {
      obj.enabled = message.enabled;
    }
    if (message.permissionId !== 0) {
      obj.permissionId = Math.round(message.permissionId);
    }
    if (message.createdAt !== "") {
      obj.createdAt = message.createdAt;
    }
    if (message.updatedAt !== "") {
      obj.updatedAt = message.updatedAt;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ApiItem>, I>>(base?: I): ApiItem {
    return ApiItem.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ApiItem>, I>>(object: I): ApiItem {
    const message = createBaseApiItem();
    message.id = object.id ?? 0;
    message.path = object.path ?? "";
    message.key = object.key ?? "";
    message.name = object.name ?? "";
    message.enabled = object.enabled ?? false;
    message.permissionId = object.permissionId ?? 0;
    message.createdAt = object.createdAt ?? "";
    message.updatedAt = object.updatedAt ?? "";
    return message;
  },
};

function createBaseAdminUserListItem(): AdminUserListItem {
  return {
    adminId: 0,
    username: "",
    nickname: "",
    email: "",
    avatar: "",
    loginTotal: 0,
    lastLoginIp: "",
    lastLoginTime: "",
    isEnabled: false,
    createdAt: "",
    updatedAt: "",
    roles: [],
    isEnabledButtonDisabled: false,
  };
}

export const AdminUserListItem: MessageFns<AdminUserListItem> = {
  encode(message: AdminUserListItem, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.adminId !== 0) {
      writer.uint32(8).int32(message.adminId);
    }
    if (message.username !== "") {
      writer.uint32(18).string(message.username);
    }
    if (message.nickname !== "") {
      writer.uint32(26).string(message.nickname);
    }
    if (message.email !== "") {
      writer.uint32(34).string(message.email);
    }
    if (message.avatar !== "") {
      writer.uint32(42).string(message.avatar);
    }
    if (message.loginTotal !== 0) {
      writer.uint32(48).int32(message.loginTotal);
    }
    if (message.lastLoginIp !== "") {
      writer.uint32(58).string(message.lastLoginIp);
    }
    if (message.lastLoginTime !== "") {
      writer.uint32(66).string(message.lastLoginTime);
    }
    if (message.isEnabled !== false) {
      writer.uint32(72).bool(message.isEnabled);
    }
    if (message.createdAt !== "") {
      writer.uint32(82).string(message.createdAt);
    }
    if (message.updatedAt !== "") {
      writer.uint32(90).string(message.updatedAt);
    }
    for (const v of message.roles) {
      AdminUserRoleItem.encode(v!, writer.uint32(98).fork()).join();
    }
    if (message.isEnabledButtonDisabled !== false) {
      writer.uint32(104).bool(message.isEnabledButtonDisabled);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): AdminUserListItem {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAdminUserListItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.adminId = reader.int32();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.username = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.nickname = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.email = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.avatar = reader.string();
          continue;
        }
        case 6: {
          if (tag !== 48) {
            break;
          }

          message.loginTotal = reader.int32();
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.lastLoginIp = reader.string();
          continue;
        }
        case 8: {
          if (tag !== 66) {
            break;
          }

          message.lastLoginTime = reader.string();
          continue;
        }
        case 9: {
          if (tag !== 72) {
            break;
          }

          message.isEnabled = reader.bool();
          continue;
        }
        case 10: {
          if (tag !== 82) {
            break;
          }

          message.createdAt = reader.string();
          continue;
        }
        case 11: {
          if (tag !== 90) {
            break;
          }

          message.updatedAt = reader.string();
          continue;
        }
        case 12: {
          if (tag !== 98) {
            break;
          }

          message.roles.push(AdminUserRoleItem.decode(reader, reader.uint32()));
          continue;
        }
        case 13: {
          if (tag !== 104) {
            break;
          }

          message.isEnabledButtonDisabled = reader.bool();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AdminUserListItem {
    return {
      adminId: isSet(object.adminId) ? globalThis.Number(object.adminId) : 0,
      username: isSet(object.username) ? globalThis.String(object.username) : "",
      nickname: isSet(object.nickname) ? globalThis.String(object.nickname) : "",
      email: isSet(object.email) ? globalThis.String(object.email) : "",
      avatar: isSet(object.avatar) ? globalThis.String(object.avatar) : "",
      loginTotal: isSet(object.loginTotal) ? globalThis.Number(object.loginTotal) : 0,
      lastLoginIp: isSet(object.lastLoginIp) ? globalThis.String(object.lastLoginIp) : "",
      lastLoginTime: isSet(object.lastLoginTime) ? globalThis.String(object.lastLoginTime) : "",
      isEnabled: isSet(object.isEnabled) ? globalThis.Boolean(object.isEnabled) : false,
      createdAt: isSet(object.createdAt) ? globalThis.String(object.createdAt) : "",
      updatedAt: isSet(object.updatedAt) ? globalThis.String(object.updatedAt) : "",
      roles: globalThis.Array.isArray(object?.roles) ? object.roles.map((e: any) => AdminUserRoleItem.fromJSON(e)) : [],
      isEnabledButtonDisabled: isSet(object.isEnabledButtonDisabled)
        ? globalThis.Boolean(object.isEnabledButtonDisabled)
        : false,
    };
  },

  toJSON(message: AdminUserListItem): unknown {
    const obj: any = {};
    if (message.adminId !== 0) {
      obj.adminId = Math.round(message.adminId);
    }
    if (message.username !== "") {
      obj.username = message.username;
    }
    if (message.nickname !== "") {
      obj.nickname = message.nickname;
    }
    if (message.email !== "") {
      obj.email = message.email;
    }
    if (message.avatar !== "") {
      obj.avatar = message.avatar;
    }
    if (message.loginTotal !== 0) {
      obj.loginTotal = Math.round(message.loginTotal);
    }
    if (message.lastLoginIp !== "") {
      obj.lastLoginIp = message.lastLoginIp;
    }
    if (message.lastLoginTime !== "") {
      obj.lastLoginTime = message.lastLoginTime;
    }
    if (message.isEnabled !== false) {
      obj.isEnabled = message.isEnabled;
    }
    if (message.createdAt !== "") {
      obj.createdAt = message.createdAt;
    }
    if (message.updatedAt !== "") {
      obj.updatedAt = message.updatedAt;
    }
    if (message.roles?.length) {
      obj.roles = message.roles.map((e) => AdminUserRoleItem.toJSON(e));
    }
    if (message.isEnabledButtonDisabled !== false) {
      obj.isEnabledButtonDisabled = message.isEnabledButtonDisabled;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AdminUserListItem>, I>>(base?: I): AdminUserListItem {
    return AdminUserListItem.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<AdminUserListItem>, I>>(object: I): AdminUserListItem {
    const message = createBaseAdminUserListItem();
    message.adminId = object.adminId ?? 0;
    message.username = object.username ?? "";
    message.nickname = object.nickname ?? "";
    message.email = object.email ?? "";
    message.avatar = object.avatar ?? "";
    message.loginTotal = object.loginTotal ?? 0;
    message.lastLoginIp = object.lastLoginIp ?? "";
    message.lastLoginTime = object.lastLoginTime ?? "";
    message.isEnabled = object.isEnabled ?? false;
    message.createdAt = object.createdAt ?? "";
    message.updatedAt = object.updatedAt ?? "";
    message.roles = object.roles?.map((e) => AdminUserRoleItem.fromPartial(e)) || [];
    message.isEnabledButtonDisabled = object.isEnabledButtonDisabled ?? false;
    return message;
  },
};

function createBaseAdminUserRoleItem(): AdminUserRoleItem {
  return { roleId: 0, roleName: "" };
}

export const AdminUserRoleItem: MessageFns<AdminUserRoleItem> = {
  encode(message: AdminUserRoleItem, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.roleId !== 0) {
      writer.uint32(8).int32(message.roleId);
    }
    if (message.roleName !== "") {
      writer.uint32(18).string(message.roleName);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): AdminUserRoleItem {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAdminUserRoleItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.roleId = reader.int32();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.roleName = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AdminUserRoleItem {
    return {
      roleId: isSet(object.roleId) ? globalThis.Number(object.roleId) : 0,
      roleName: isSet(object.roleName) ? globalThis.String(object.roleName) : "",
    };
  },

  toJSON(message: AdminUserRoleItem): unknown {
    const obj: any = {};
    if (message.roleId !== 0) {
      obj.roleId = Math.round(message.roleId);
    }
    if (message.roleName !== "") {
      obj.roleName = message.roleName;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AdminUserRoleItem>, I>>(base?: I): AdminUserRoleItem {
    return AdminUserRoleItem.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<AdminUserRoleItem>, I>>(object: I): AdminUserRoleItem {
    const message = createBaseAdminUserRoleItem();
    message.roleId = object.roleId ?? 0;
    message.roleName = object.roleName ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(int64: { toString(): string }): number {
  const num = globalThis.Number(int64.toString());
  if (num > globalThis.Number.MAX_SAFE_INTEGER) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  if (num < globalThis.Number.MIN_SAFE_INTEGER) {
    throw new globalThis.Error("Value is smaller than Number.MIN_SAFE_INTEGER");
  }
  return num;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
  fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
