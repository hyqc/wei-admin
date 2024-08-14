// package: admin
// file: common.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class ListBaseReq extends jspb.Message { 
    getPagesize(): number;
    setPagesize(value: number): ListBaseReq;
    getPagenum(): number;
    setPagenum(value: number): ListBaseReq;
    getSortfield(): string;
    setSortfield(value: string): ListBaseReq;
    getSorttype(): string;
    setSorttype(value: string): ListBaseReq;
    getEnabled(): number;
    setEnabled(value: number): ListBaseReq;
    getCreatestarttime(): number;
    setCreatestarttime(value: number): ListBaseReq;
    getCreateendtime(): number;
    setCreateendtime(value: number): ListBaseReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListBaseReq.AsObject;
    static toObject(includeInstance: boolean, msg: ListBaseReq): ListBaseReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListBaseReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListBaseReq;
    static deserializeBinaryFromReader(message: ListBaseReq, reader: jspb.BinaryReader): ListBaseReq;
}

export namespace ListBaseReq {
    export type AsObject = {
        pagesize: number,
        pagenum: number,
        sortfield: string,
        sorttype: string,
        enabled: number,
        createstarttime: number,
        createendtime: number,
    }
}
