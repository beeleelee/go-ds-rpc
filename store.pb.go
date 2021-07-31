// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.10.0
// source: store.proto

package dsrpc

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

type PutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{0}
}

func (x *PutRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type ErrReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *ErrReply) Reset() {
	*x = ErrReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrReply) ProtoMessage() {}

func (x *ErrReply) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrReply.ProtoReflect.Descriptor instead.
func (*ErrReply) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{1}
}

func (x *ErrReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type BoolReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BoolReply) Reset() {
	*x = BoolReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolReply) ProtoMessage() {}

func (x *BoolReply) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolReply.ProtoReflect.Descriptor instead.
func (*BoolReply) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{2}
}

func (x *BoolReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type SizeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *SizeReply) Reset() {
	*x = SizeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeReply) ProtoMessage() {}

func (x *SizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeReply.ProtoReflect.Descriptor instead.
func (*SizeReply) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{3}
}

func (x *SizeReply) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type QueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res []byte `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *QueryReply) Reset() {
	*x = QueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryReply) ProtoMessage() {}

func (x *QueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryReply.ProtoReflect.Descriptor instead.
func (*QueryReply) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{4}
}

func (x *QueryReply) GetRes() []byte {
	if x != nil {
		return x.Res
	}
	return nil
}

type StoreKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *StoreKey) Reset() {
	*x = StoreKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreKey) ProtoMessage() {}

func (x *StoreKey) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreKey.ProtoReflect.Descriptor instead.
func (*StoreKey) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{5}
}

func (x *StoreKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type StoreValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StoreValue) Reset() {
	*x = StoreValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreValue) ProtoMessage() {}

func (x *StoreValue) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreValue.ProtoReflect.Descriptor instead.
func (*StoreValue) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{6}
}

func (x *StoreValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type StoreQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Q []byte `protobuf:"bytes,1,opt,name=q,proto3" json:"q,omitempty"`
}

func (x *StoreQuery) Reset() {
	*x = StoreQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreQuery) ProtoMessage() {}

func (x *StoreQuery) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreQuery.ProtoReflect.Descriptor instead.
func (*StoreQuery) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{7}
}

func (x *StoreQuery) GetQ() []byte {
	if x != nil {
		return x.Q
	}
	return nil
}

var File_store_proto protoreflect.FileDescriptor

var file_store_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64,
	0x73, 0x72, 0x70, 0x63, 0x22, 0x34, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1c, 0x0a, 0x08, 0x45, 0x72,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x25, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x1f, 0x0a, 0x09, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x1e, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x65, 0x73,
	0x22, 0x1c, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x22,
	0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x1a, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x71, 0x32, 0x9e,
	0x02, 0x0a, 0x07, 0x4b, 0x56, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x50, 0x75,
	0x74, 0x12, 0x11, 0x2e, 0x64, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x64, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x0f, 0x2e, 0x64, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4b,
	0x65, 0x79, 0x1a, 0x0f, 0x2e, 0x64, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x64,
	0x73, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x1a, 0x11, 0x2e,
	0x64, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x00, 0x12, 0x2a, 0x0a, 0x03, 0x48, 0x61, 0x73, 0x12, 0x0f, 0x2e, 0x64, 0x73, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x1a, 0x10, 0x2e, 0x64, 0x73, 0x72,
	0x70, 0x63, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2e,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0f, 0x2e, 0x64, 0x73, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x1a, 0x10, 0x2e, 0x64, 0x73, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2f,
	0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x11, 0x2e, 0x64, 0x73, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x11, 0x2e, 0x64, 0x73, 0x72,
	0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x08, 0x5a, 0x06, 0x2f, 0x64, 0x73, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_store_proto_rawDescOnce sync.Once
	file_store_proto_rawDescData = file_store_proto_rawDesc
)

func file_store_proto_rawDescGZIP() []byte {
	file_store_proto_rawDescOnce.Do(func() {
		file_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_proto_rawDescData)
	})
	return file_store_proto_rawDescData
}

var file_store_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_store_proto_goTypes = []interface{}{
	(*PutRequest)(nil), // 0: dsrpc.PutRequest
	(*ErrReply)(nil),   // 1: dsrpc.ErrReply
	(*BoolReply)(nil),  // 2: dsrpc.BoolReply
	(*SizeReply)(nil),  // 3: dsrpc.SizeReply
	(*QueryReply)(nil), // 4: dsrpc.QueryReply
	(*StoreKey)(nil),   // 5: dsrpc.StoreKey
	(*StoreValue)(nil), // 6: dsrpc.StoreValue
	(*StoreQuery)(nil), // 7: dsrpc.StoreQuery
}
var file_store_proto_depIdxs = []int32{
	0, // 0: dsrpc.KVStore.Put:input_type -> dsrpc.PutRequest
	5, // 1: dsrpc.KVStore.Delete:input_type -> dsrpc.StoreKey
	5, // 2: dsrpc.KVStore.Get:input_type -> dsrpc.StoreKey
	5, // 3: dsrpc.KVStore.Has:input_type -> dsrpc.StoreKey
	5, // 4: dsrpc.KVStore.GetSize:input_type -> dsrpc.StoreKey
	7, // 5: dsrpc.KVStore.Query:input_type -> dsrpc.StoreQuery
	1, // 6: dsrpc.KVStore.Put:output_type -> dsrpc.ErrReply
	1, // 7: dsrpc.KVStore.Delete:output_type -> dsrpc.ErrReply
	6, // 8: dsrpc.KVStore.Get:output_type -> dsrpc.StoreValue
	2, // 9: dsrpc.KVStore.Has:output_type -> dsrpc.BoolReply
	3, // 10: dsrpc.KVStore.GetSize:output_type -> dsrpc.SizeReply
	4, // 11: dsrpc.KVStore.Query:output_type -> dsrpc.QueryReply
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_store_proto_init() }
func file_store_proto_init() {
	if File_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRequest); i {
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
		file_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrReply); i {
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
		file_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolReply); i {
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
		file_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SizeReply); i {
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
		file_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryReply); i {
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
		file_store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreKey); i {
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
		file_store_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreValue); i {
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
		file_store_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreQuery); i {
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
			RawDescriptor: file_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_store_proto_goTypes,
		DependencyIndexes: file_store_proto_depIdxs,
		MessageInfos:      file_store_proto_msgTypes,
	}.Build()
	File_store_proto = out.File
	file_store_proto_rawDesc = nil
	file_store_proto_goTypes = nil
	file_store_proto_depIdxs = nil
}
