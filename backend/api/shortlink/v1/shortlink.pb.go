// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: shortlink/v1/shortlink.proto

package shortlinkv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path  string            `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	State []*ShareableState `protobuf:"bytes,2,rep,name=state,proto3" json:"state,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortlink_v1_shortlink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_v1_shortlink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_shortlink_v1_shortlink_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateRequest) GetState() []*ShareableState {
	if x != nil {
		return x.State
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortlink_v1_shortlink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_v1_shortlink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_shortlink_v1_shortlink_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortlink_v1_shortlink_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_v1_shortlink_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_shortlink_v1_shortlink_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path  string            `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	State []*ShareableState `protobuf:"bytes,2,rep,name=state,proto3" json:"state,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortlink_v1_shortlink_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_v1_shortlink_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_shortlink_v1_shortlink_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GetResponse) GetState() []*ShareableState {
	if x != nil {
		return x.State
	}
	return nil
}

// ShareableState stores a key identifier that maps to state.
// This is analogous to a map, however we are not using a map here as that will restrict
// our ability to further expand on this message.
// For example adding a "revision" identifier in the future.
type ShareableState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string          `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	State *structpb.Value `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ShareableState) Reset() {
	*x = ShareableState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortlink_v1_shortlink_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareableState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareableState) ProtoMessage() {}

func (x *ShareableState) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_v1_shortlink_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareableState.ProtoReflect.Descriptor instead.
func (*ShareableState) Descriptor() ([]byte, []int) {
	return file_shortlink_v1_shortlink_proto_rawDescGZIP(), []int{4}
}

func (x *ShareableState) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ShareableState) GetState() *structpb.Value {
	if x != nil {
		return x.State
	}
	return nil
}

var File_shortlink_v1_shortlink_proto protoreflect.FileDescriptor

var file_shortlink_v1_shortlink_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x43, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x29, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x5c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63,
	0x68, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x63, 0x0a, 0x0e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x32, 0xf6, 0x01, 0x0a, 0x0c, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x50, 0x49, 0x12, 0x78, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63,
	0x68, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xaa, 0xe1,
	0x1c, 0x02, 0x08, 0x01, 0x12, 0x6c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x6c,
	0x75, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63,
	0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0xaa, 0xe1, 0x1c, 0x02,
	0x08, 0x02, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69,
	0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortlink_v1_shortlink_proto_rawDescOnce sync.Once
	file_shortlink_v1_shortlink_proto_rawDescData = file_shortlink_v1_shortlink_proto_rawDesc
)

func file_shortlink_v1_shortlink_proto_rawDescGZIP() []byte {
	file_shortlink_v1_shortlink_proto_rawDescOnce.Do(func() {
		file_shortlink_v1_shortlink_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortlink_v1_shortlink_proto_rawDescData)
	})
	return file_shortlink_v1_shortlink_proto_rawDescData
}

var file_shortlink_v1_shortlink_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_shortlink_v1_shortlink_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),  // 0: clutch.shortlink.v1.CreateRequest
	(*CreateResponse)(nil), // 1: clutch.shortlink.v1.CreateResponse
	(*GetRequest)(nil),     // 2: clutch.shortlink.v1.GetRequest
	(*GetResponse)(nil),    // 3: clutch.shortlink.v1.GetResponse
	(*ShareableState)(nil), // 4: clutch.shortlink.v1.ShareableState
	(*structpb.Value)(nil), // 5: google.protobuf.Value
}
var file_shortlink_v1_shortlink_proto_depIdxs = []int32{
	4, // 0: clutch.shortlink.v1.CreateRequest.state:type_name -> clutch.shortlink.v1.ShareableState
	4, // 1: clutch.shortlink.v1.GetResponse.state:type_name -> clutch.shortlink.v1.ShareableState
	5, // 2: clutch.shortlink.v1.ShareableState.state:type_name -> google.protobuf.Value
	0, // 3: clutch.shortlink.v1.ShortlinkAPI.Create:input_type -> clutch.shortlink.v1.CreateRequest
	2, // 4: clutch.shortlink.v1.ShortlinkAPI.Get:input_type -> clutch.shortlink.v1.GetRequest
	1, // 5: clutch.shortlink.v1.ShortlinkAPI.Create:output_type -> clutch.shortlink.v1.CreateResponse
	3, // 6: clutch.shortlink.v1.ShortlinkAPI.Get:output_type -> clutch.shortlink.v1.GetResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_shortlink_v1_shortlink_proto_init() }
func file_shortlink_v1_shortlink_proto_init() {
	if File_shortlink_v1_shortlink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortlink_v1_shortlink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_shortlink_v1_shortlink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_shortlink_v1_shortlink_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_shortlink_v1_shortlink_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_shortlink_v1_shortlink_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareableState); i {
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
			RawDescriptor: file_shortlink_v1_shortlink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortlink_v1_shortlink_proto_goTypes,
		DependencyIndexes: file_shortlink_v1_shortlink_proto_depIdxs,
		MessageInfos:      file_shortlink_v1_shortlink_proto_msgTypes,
	}.Build()
	File_shortlink_v1_shortlink_proto = out.File
	file_shortlink_v1_shortlink_proto_rawDesc = nil
	file_shortlink_v1_shortlink_proto_goTypes = nil
	file_shortlink_v1_shortlink_proto_depIdxs = nil
}
