// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: admin_menu.proto

package admin_menu

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MenuItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`   // 菜单唯一键
	Path      string       `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"` // 菜单路由
	Name      string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"` // 菜单名称
	Icon      string       `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"` // 菜单图标
	Component string       `protobuf:"bytes,5,opt,name=component,proto3" json:"component,omitempty"`
	Authority string       `protobuf:"bytes,6,opt,name=authority,proto3" json:"authority,omitempty"`
	Routes    []*RouteItem `protobuf:"bytes,7,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *MenuItem) Reset() {
	*x = MenuItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuItem) ProtoMessage() {}

func (x *MenuItem) ProtoReflect() protoreflect.Message {
	mi := &file_admin_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuItem.ProtoReflect.Descriptor instead.
func (*MenuItem) Descriptor() ([]byte, []int) {
	return file_admin_menu_proto_rawDescGZIP(), []int{0}
}

func (x *MenuItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MenuItem) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MenuItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuItem) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *MenuItem) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *MenuItem) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *MenuItem) GetRoutes() []*RouteItem {
	if x != nil {
		return x.Routes
	}
	return nil
}

type RouteItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Path      string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Icon      string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	Component string `protobuf:"bytes,5,opt,name=component,proto3" json:"component,omitempty"`
}

func (x *RouteItem) Reset() {
	*x = RouteItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteItem) ProtoMessage() {}

func (x *RouteItem) ProtoReflect() protoreflect.Message {
	mi := &file_admin_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteItem.ProtoReflect.Descriptor instead.
func (*RouteItem) Descriptor() ([]byte, []int) {
	return file_admin_menu_proto_rawDescGZIP(), []int{1}
}

func (x *RouteItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RouteItem) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RouteItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteItem) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *RouteItem) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

var File_admin_menu_proto protoreflect.FileDescriptor

var file_admin_menu_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0xc3,
	0x01, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x65,
	0x6e, 0x75, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x22, 0x77, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x42, 0x23, 0x5a,
	0x21, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x65,
	0x6e, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_menu_proto_rawDescOnce sync.Once
	file_admin_menu_proto_rawDescData = file_admin_menu_proto_rawDesc
)

func file_admin_menu_proto_rawDescGZIP() []byte {
	file_admin_menu_proto_rawDescOnce.Do(func() {
		file_admin_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_menu_proto_rawDescData)
	})
	return file_admin_menu_proto_rawDescData
}

var file_admin_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_admin_menu_proto_goTypes = []interface{}{
	(*MenuItem)(nil),  // 0: admin_menu.menuItem
	(*RouteItem)(nil), // 1: admin_menu.routeItem
}
var file_admin_menu_proto_depIdxs = []int32{
	1, // 0: admin_menu.menuItem.routes:type_name -> admin_menu.routeItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_admin_menu_proto_init() }
func file_admin_menu_proto_init() {
	if File_admin_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_admin_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_menu_proto_goTypes,
		DependencyIndexes: file_admin_menu_proto_depIdxs,
		MessageInfos:      file_admin_menu_proto_msgTypes,
	}.Build()
	File_admin_menu_proto = out.File
	file_admin_menu_proto_rawDesc = nil
	file_admin_menu_proto_goTypes = nil
	file_admin_menu_proto_depIdxs = nil
}
