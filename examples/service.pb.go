// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: examples/service.proto

package example

import (
	_ "github.com/nvx/go-proto-logfields"
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

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author string `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_examples_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_examples_service_proto_rawDescGZIP(), []int{0}
}

func (x *Note) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Note *Note  `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_examples_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_examples_service_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Request) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DidStuff    bool    `protobuf:"varint,1,opt,name=did_stuff,json=didStuff,proto3" json:"did_stuff,omitempty"`
	ChangedNote *Note   `protobuf:"bytes,2,opt,name=changed_note,json=changedNote,proto3" json:"changed_note,omitempty"`
	Notes       []*Note `protobuf:"bytes,3,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_examples_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_examples_service_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetDidStuff() bool {
	if x != nil {
		return x.DidStuff
	}
	return false
}

func (x *Response) GetChangedNote() *Note {
	if x != nil {
		return x.ChangedNote
	}
	return nil
}

func (x *Response) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

var File_examples_service_proto protoreflect.FileDescriptor

var file_examples_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x1a, 0x0f, 0x6c, 0x6f, 0x67, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xa2, 0xab, 0x1e, 0x08,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x22, 0x4c, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xa2, 0xab, 0x1e, 0x06, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x8c,
	0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x64,
	0x69, 0x64, 0x5f, 0x73, 0x74, 0x75, 0x66, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0c,
	0xa2, 0xab, 0x1e, 0x08, 0x0a, 0x06, 0x64, 0x69, 0x64, 0x5f, 0x69, 0x74, 0x52, 0x08, 0x64, 0x69,
	0x64, 0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x30, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x76, 0x78, 0x2f,
	0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6c, 0x6f, 0x67, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_examples_service_proto_rawDescOnce sync.Once
	file_examples_service_proto_rawDescData = file_examples_service_proto_rawDesc
)

func file_examples_service_proto_rawDescGZIP() []byte {
	file_examples_service_proto_rawDescOnce.Do(func() {
		file_examples_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_service_proto_rawDescData)
	})
	return file_examples_service_proto_rawDescData
}

var file_examples_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_examples_service_proto_goTypes = []interface{}{
	(*Note)(nil),     // 0: example.Note
	(*Request)(nil),  // 1: example.Request
	(*Response)(nil), // 2: example.Response
}
var file_examples_service_proto_depIdxs = []int32{
	0, // 0: example.Request.note:type_name -> example.Note
	0, // 1: example.Response.changed_note:type_name -> example.Note
	0, // 2: example.Response.notes:type_name -> example.Note
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_examples_service_proto_init() }
func file_examples_service_proto_init() {
	if File_examples_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
		file_examples_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_examples_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_examples_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_examples_service_proto_goTypes,
		DependencyIndexes: file_examples_service_proto_depIdxs,
		MessageInfos:      file_examples_service_proto_msgTypes,
	}.Build()
	File_examples_service_proto = out.File
	file_examples_service_proto_rawDesc = nil
	file_examples_service_proto_goTypes = nil
	file_examples_service_proto_depIdxs = nil
}
