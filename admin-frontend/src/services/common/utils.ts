import { ReqListBase } from "@/proto/admin_ts/common";
import React from "react";

export const handlePagination = (pageNum?: number, pageSize?: number):ReqListBase=>{
    return {
        pageNum: pageNum || 1,
        pageSize: pageSize || 10,
    }
}

export const DefaultPagination = {
    showQuickJumper: true,
    showSizeChanger: true,
    pageSizeOptions: [1, 5, 10, 50, 100],
    position: ['bottomRight'],
    showTotal: (total: number, range: [number, number]) => (`共 ${total} 条`),
  }