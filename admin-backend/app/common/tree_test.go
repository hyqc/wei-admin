package common

import (
	"admin/proto/admin_proto"
	"fmt"
	"testing"
)

func TestMenuTree(t *testing.T) {
	// 示例数据和调用
	menusMap := map[int32][]*admin_proto.MenuTreeItem{
		1: {&admin_proto.MenuTreeItem{Id: 2}, &admin_proto.MenuTreeItem{Id: 3}},
		2: {&admin_proto.MenuTreeItem{Id: 4}, &admin_proto.MenuTreeItem{Id: 5}},
	}
	menusIds := make(map[int32]bool)
	for _, id := range []int32{1, 2, 3} {
		menusIds[id] = true
	}
	tree := MenuTree(menusMap, menusIds, 1, 1, 3)
	fmt.Println(tree)
}
