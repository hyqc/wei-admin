import { ReqListBase } from "@/proto/admin_ts/common";

export const handlePagination = (pageNum?: number, pageSize?: number):ReqListBase=>{
    return {
        pageNum: pageNum || 1,
        pageSize: pageSize || 10,
    }
}