import { TablePaginationConfig } from "antd"

export const DefaultPagination: false | TablePaginationConfig = {
    showQuickJumper: true,
    showSizeChanger: true,
    pageSizeOptions: [1, 5, 10, 50, 100],
    position: ['bottomRight'],
    showTotal: (total: number, range: [number, number]) => (`共 ${total} 条`),
}