// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.6
// source: conf/conf.proto

package conf

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

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   *Server   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data     *Data     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Registry *Registry `protobuf:"bytes,3,opt,name=registry,proto3" json:"registry,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetRegistry() *Registry {
	if x != nil {
		return x.Registry
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatabaseSource string `protobuf:"bytes,1,opt,name=database_source,json=databaseSource,proto3" json:"database_source,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Data) GetDatabaseSource() string {
	if x != nil {
		return x.DatabaseSource
	}
	return ""
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http *Http `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc *Grpc `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Server) GetHttp() *Http {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Grpc {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Http struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout int32  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Http) Reset() {
	*x = Http{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http) ProtoMessage() {}

func (x *Http) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http.ProtoReflect.Descriptor instead.
func (*Http) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Http) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Http) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Http) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type Grpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout int32  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Grpc) Reset() {
	*x = Grpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grpc) ProtoMessage() {}

func (x *Grpc) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grpc.ProtoReflect.Descriptor instead.
func (*Grpc) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Grpc) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Grpc) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Grpc) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Etcd *Etcd `protobuf:"bytes,1,opt,name=etcd,proto3" json:"etcd,omitempty"`
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{5}
}

func (x *Registry) GetEtcd() *Etcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

type Etcd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *Etcd) Reset() {
	*x = Etcd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Etcd) ProtoMessage() {}

func (x *Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Etcd.ProtoReflect.Descriptor instead.
func (*Etcd) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{6}
}

func (x *Etcd) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   *Server   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data     *Data     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Registry *Registry `protobuf:"bytes,3,opt,name=registry,proto3" json:"registry,omitempty"`
}

func (x *Conf) Reset() {
	*x = Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf) ProtoMessage() {}

func (x *Conf) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conf.ProtoReflect.Descriptor instead.
func (*Conf) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{7}
}

func (x *Conf) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Conf) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Conf) GetRegistry() *Registry {
	if x != nil {
		return x.Registry
	}
	return nil
}

var File_conf_conf_proto protoreflect.FileDescriptor

var file_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x22, 0x7d, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x22, 0x2f, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27,
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x48, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x04, 0x68, 0x74, 0x74,
	0x70, 0x12, 0x1e, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x52, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x22, 0x4e, 0x0a, 0x04, 0x48, 0x74, 0x74, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x22, 0x4e, 0x0a, 0x04, 0x47, 0x72, 0x70, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x22, 0x2a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x1e, 0x0a,
	0x04, 0x65, 0x74, 0x63, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x22, 0x22, 0x0a,
	0x04, 0x45, 0x74, 0x63, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x22, 0x78, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x2a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x42, 0x0f, 0x5a, 0x0d, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_conf_proto_rawDescOnce sync.Once
	file_conf_conf_proto_rawDescData = file_conf_conf_proto_rawDesc
)

func file_conf_conf_proto_rawDescGZIP() []byte {
	file_conf_conf_proto_rawDescOnce.Do(func() {
		file_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_conf_proto_rawDescData)
	})
	return file_conf_conf_proto_rawDescData
}

var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_conf_conf_proto_goTypes = []any{
	(*Bootstrap)(nil), // 0: conf.Bootstrap
	(*Data)(nil),      // 1: conf.Data
	(*Server)(nil),    // 2: conf.Server
	(*Http)(nil),      // 3: conf.Http
	(*Grpc)(nil),      // 4: conf.Grpc
	(*Registry)(nil),  // 5: conf.Registry
	(*Etcd)(nil),      // 6: conf.Etcd
	(*Conf)(nil),      // 7: conf.Conf
}
var file_conf_conf_proto_depIdxs = []int32{
	2, // 0: conf.Bootstrap.server:type_name -> conf.Server
	1, // 1: conf.Bootstrap.data:type_name -> conf.Data
	5, // 2: conf.Bootstrap.registry:type_name -> conf.Registry
	3, // 3: conf.Server.http:type_name -> conf.Http
	4, // 4: conf.Server.grpc:type_name -> conf.Grpc
	6, // 5: conf.Registry.etcd:type_name -> conf.Etcd
	2, // 6: conf.Conf.server:type_name -> conf.Server
	1, // 7: conf.Conf.data:type_name -> conf.Data
	5, // 8: conf.Conf.registry:type_name -> conf.Registry
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_conf_conf_proto_init() }
func file_conf_conf_proto_init() {
	if File_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_conf_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Bootstrap); i {
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
		file_conf_conf_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Data); i {
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
		file_conf_conf_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Server); i {
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
		file_conf_conf_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Http); i {
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
		file_conf_conf_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Grpc); i {
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
		file_conf_conf_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Registry); i {
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
		file_conf_conf_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Etcd); i {
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
		file_conf_conf_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Conf); i {
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
			RawDescriptor: file_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_conf_proto_goTypes,
		DependencyIndexes: file_conf_conf_proto_depIdxs,
		MessageInfos:      file_conf_conf_proto_msgTypes,
	}.Build()
	File_conf_conf_proto = out.File
	file_conf_conf_proto_rawDesc = nil
	file_conf_conf_proto_goTypes = nil
	file_conf_conf_proto_depIdxs = nil
}
