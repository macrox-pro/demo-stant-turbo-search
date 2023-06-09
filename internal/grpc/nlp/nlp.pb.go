// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/proto/nlp.proto

package nlp

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

type Doc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Doc) Reset() {
	*x = Doc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_nlp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Doc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Doc) ProtoMessage() {}

func (x *Doc) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_nlp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Doc.ProtoReflect.Descriptor instead.
func (*Doc) Descriptor() ([]byte, []int) {
	return file_api_proto_nlp_proto_rawDescGZIP(), []int{0}
}

func (x *Doc) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Intent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Confidence float32 `protobuf:"fixed32,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *Intent) Reset() {
	*x = Intent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_nlp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Intent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Intent) ProtoMessage() {}

func (x *Intent) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_nlp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Intent.ProtoReflect.Descriptor instead.
func (*Intent) Descriptor() ([]byte, []int) {
	return file_api_proto_nlp_proto_rawDescGZIP(), []int{1}
}

func (x *Intent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Intent) GetConfidence() float32 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start       uint32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End         uint32 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Value       string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	NormalValue string `protobuf:"bytes,5,opt,name=normal_value,json=normalValue,proto3" json:"normal_value,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_nlp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_nlp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_api_proto_nlp_proto_rawDescGZIP(), []int{2}
}

func (x *Entity) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Entity) GetEnd() uint32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Entity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Entity) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Entity) GetNormalValue() string {
	if x != nil {
		return x.NormalValue
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string    `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Intent   *Intent   `protobuf:"bytes,2,opt,name=intent,proto3,oneof" json:"intent,omitempty"`
	Entities []*Entity `protobuf:"bytes,3,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_nlp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_nlp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_api_proto_nlp_proto_rawDescGZIP(), []int{3}
}

func (x *Result) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Result) GetIntent() *Intent {
	if x != nil {
		return x.Intent
	}
	return nil
}

func (x *Result) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

var File_api_proto_nlp_proto protoreflect.FileDescriptor

var file_api_proto_nlp_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6c, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x03, 0x44, 0x6f, 0x63, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x3c, 0x0a, 0x06, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x7d,
	0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x6f,
	0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x72, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x32, 0x1d, 0x0a, 0x03, 0x4e, 0x4c, 0x50, 0x12, 0x16, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x12, 0x04, 0x2e, 0x44, 0x6f, 0x63, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x42, 0x13, 0x5a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x6e, 0x6c, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_nlp_proto_rawDescOnce sync.Once
	file_api_proto_nlp_proto_rawDescData = file_api_proto_nlp_proto_rawDesc
)

func file_api_proto_nlp_proto_rawDescGZIP() []byte {
	file_api_proto_nlp_proto_rawDescOnce.Do(func() {
		file_api_proto_nlp_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_nlp_proto_rawDescData)
	})
	return file_api_proto_nlp_proto_rawDescData
}

var file_api_proto_nlp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_nlp_proto_goTypes = []interface{}{
	(*Doc)(nil),    // 0: Doc
	(*Intent)(nil), // 1: Intent
	(*Entity)(nil), // 2: Entity
	(*Result)(nil), // 3: Result
}
var file_api_proto_nlp_proto_depIdxs = []int32{
	1, // 0: Result.intent:type_name -> Intent
	2, // 1: Result.entities:type_name -> Entity
	0, // 2: NLP.Parse:input_type -> Doc
	3, // 3: NLP.Parse:output_type -> Result
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_nlp_proto_init() }
func file_api_proto_nlp_proto_init() {
	if File_api_proto_nlp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_nlp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Doc); i {
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
		file_api_proto_nlp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Intent); i {
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
		file_api_proto_nlp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_api_proto_nlp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
	file_api_proto_nlp_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_nlp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_nlp_proto_goTypes,
		DependencyIndexes: file_api_proto_nlp_proto_depIdxs,
		MessageInfos:      file_api_proto_nlp_proto_msgTypes,
	}.Build()
	File_api_proto_nlp_proto = out.File
	file_api_proto_nlp_proto_rawDesc = nil
	file_api_proto_nlp_proto_goTypes = nil
	file_api_proto_nlp_proto_depIdxs = nil
}
