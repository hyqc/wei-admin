import { ReqListBase } from "@/proto/admin_ts/common";
import { PageInfoType } from "../apis/types";

export const handlePagination = (pageNum?: number, pageSize?: number):ReqListBase=>{
    return {
        pageNum: pageNum || 1,
        pageSize: pageSize || 10,
    }
}

export const searchResetPageInfo = (pageSize?: number):PageInfoType=>{
    return {
        pageNum: 1,
        pageSize: pageSize || 10,
        total: 0,
    }
}

