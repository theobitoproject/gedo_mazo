// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: gedo_mazo.proto

package gedo_mazo

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

// Request message for GedoMazo.GenerateDocumentFromTemplate
type GenerateDocumentFromTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the base template document that will be cloned
	TemplateDocument *Document `protobuf:"bytes,1,opt,name=template_document,json=templateDocument,proto3" json:"template_document,omitempty"`
	// folder in which the cloned document will be created
	OutputFolder *Folder `protobuf:"bytes,2,opt,name=output_folder,json=outputFolder,proto3" json:"output_folder,omitempty"`
	// document data that is provided for the cloned document, line name, etc
	// NOTE: id will be ignored since when the document is created,
	// the provider will create and assign a new one
	ClonedDocument *Document `protobuf:"bytes,3,opt,name=cloned_document,json=clonedDocument,proto3" json:"cloned_document,omitempty"`
	// all data that will be merged into the cloned document
	Data *MergingData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GenerateDocumentFromTemplateRequest) Reset() {
	*x = GenerateDocumentFromTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gedo_mazo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateDocumentFromTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateDocumentFromTemplateRequest) ProtoMessage() {}

func (x *GenerateDocumentFromTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gedo_mazo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateDocumentFromTemplateRequest.ProtoReflect.Descriptor instead.
func (*GenerateDocumentFromTemplateRequest) Descriptor() ([]byte, []int) {
	return file_gedo_mazo_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateDocumentFromTemplateRequest) GetTemplateDocument() *Document {
	if x != nil {
		return x.TemplateDocument
	}
	return nil
}

func (x *GenerateDocumentFromTemplateRequest) GetOutputFolder() *Folder {
	if x != nil {
		return x.OutputFolder
	}
	return nil
}

func (x *GenerateDocumentFromTemplateRequest) GetClonedDocument() *Document {
	if x != nil {
		return x.ClonedDocument
	}
	return nil
}

func (x *GenerateDocumentFromTemplateRequest) GetData() *MergingData {
	if x != nil {
		return x.Data
	}
	return nil
}

// Response message for GedoMazo.GenerateDocumentFromTemplate
type GenerateDocumentFromTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// document that was cloned from the template
	ClonedDocument *Document `protobuf:"bytes,1,opt,name=cloned_document,json=clonedDocument,proto3" json:"cloned_document,omitempty"`
}

func (x *GenerateDocumentFromTemplateResponse) Reset() {
	*x = GenerateDocumentFromTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gedo_mazo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateDocumentFromTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateDocumentFromTemplateResponse) ProtoMessage() {}

func (x *GenerateDocumentFromTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gedo_mazo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateDocumentFromTemplateResponse.ProtoReflect.Descriptor instead.
func (*GenerateDocumentFromTemplateResponse) Descriptor() ([]byte, []int) {
	return file_gedo_mazo_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateDocumentFromTemplateResponse) GetClonedDocument() *Document {
	if x != nil {
		return x.ClonedDocument
	}
	return nil
}

var File_gedo_mazo_proto protoreflect.FileDescriptor

var file_gedo_mazo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x65, 0x64, 0x6f, 0x5f, 0x6d, 0x61, 0x7a, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x67, 0x65, 0x64, 0x6f, 0x5f, 0x6d, 0x61, 0x7a, 0x6f, 0x1a, 0x0e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a,
	0x23, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x67, 0x65, 0x64, 0x6f, 0x5f, 0x6d, 0x61, 0x7a, 0x6f, 0x2e, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x10, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x67, 0x65, 0x64, 0x6f, 0x5f, 0x6d, 0x61, 0x7a, 0x6f, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x3c,
	0x0a, 0x0f, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x65, 0x64, 0x6f, 0x5f, 0x6d,
	0x61, 0x7a, 0x6f, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x63, 0x6c,
	0x6f, 0x6e, 0x65, 0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x65, 0x64,
	0x6f, 0x5f, 0x6d, 0x61, 0x7a, 0x6f, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x64, 0x0a, 0x24, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x0f, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x65, 0x64, 0x6f,
	0x5f, 0x6d, 0x61, 0x7a, 0x6f, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0e,
	0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x8b,
	0x01, 0x0a, 0x08, 0x47, 0x65, 0x64, 0x6f, 0x4d, 0x61, 0x7a, 0x6f, 0x12, 0x7f, 0x0a, 0x1c, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x46,
	0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x2e, 0x67, 0x65,
	0x64, 0x6f, 0x5f, 0x6d, 0x61, 0x7a, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x65,
	0x64, 0x6f, 0x5f, 0x6d, 0x61, 0x7a, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x48, 0x5a, 0x46,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x6f, 0x62,
	0x69, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x74, 0x6f, 0x61, 0x64, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x64, 0x6f, 0x5f, 0x6d, 0x61, 0x7a, 0x6f, 0x3b, 0x67, 0x65, 0x64,
	0x6f, 0x5f, 0x6d, 0x61, 0x7a, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gedo_mazo_proto_rawDescOnce sync.Once
	file_gedo_mazo_proto_rawDescData = file_gedo_mazo_proto_rawDesc
)

func file_gedo_mazo_proto_rawDescGZIP() []byte {
	file_gedo_mazo_proto_rawDescOnce.Do(func() {
		file_gedo_mazo_proto_rawDescData = protoimpl.X.CompressGZIP(file_gedo_mazo_proto_rawDescData)
	})
	return file_gedo_mazo_proto_rawDescData
}

var file_gedo_mazo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gedo_mazo_proto_goTypes = []interface{}{
	(*GenerateDocumentFromTemplateRequest)(nil),  // 0: gedo_mazo.GenerateDocumentFromTemplateRequest
	(*GenerateDocumentFromTemplateResponse)(nil), // 1: gedo_mazo.GenerateDocumentFromTemplateResponse
	(*Document)(nil),    // 2: gedo_mazo.Document
	(*Folder)(nil),      // 3: gedo_mazo.Folder
	(*MergingData)(nil), // 4: gedo_mazo.MergingData
}
var file_gedo_mazo_proto_depIdxs = []int32{
	2, // 0: gedo_mazo.GenerateDocumentFromTemplateRequest.template_document:type_name -> gedo_mazo.Document
	3, // 1: gedo_mazo.GenerateDocumentFromTemplateRequest.output_folder:type_name -> gedo_mazo.Folder
	2, // 2: gedo_mazo.GenerateDocumentFromTemplateRequest.cloned_document:type_name -> gedo_mazo.Document
	4, // 3: gedo_mazo.GenerateDocumentFromTemplateRequest.data:type_name -> gedo_mazo.MergingData
	2, // 4: gedo_mazo.GenerateDocumentFromTemplateResponse.cloned_document:type_name -> gedo_mazo.Document
	0, // 5: gedo_mazo.GedoMazo.GenerateDocumentFromTemplate:input_type -> gedo_mazo.GenerateDocumentFromTemplateRequest
	1, // 6: gedo_mazo.GedoMazo.GenerateDocumentFromTemplate:output_type -> gedo_mazo.GenerateDocumentFromTemplateResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_gedo_mazo_proto_init() }
func file_gedo_mazo_proto_init() {
	if File_gedo_mazo_proto != nil {
		return
	}
	file_entities_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gedo_mazo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateDocumentFromTemplateRequest); i {
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
		file_gedo_mazo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateDocumentFromTemplateResponse); i {
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
			RawDescriptor: file_gedo_mazo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gedo_mazo_proto_goTypes,
		DependencyIndexes: file_gedo_mazo_proto_depIdxs,
		MessageInfos:      file_gedo_mazo_proto_msgTypes,
	}.Build()
	File_gedo_mazo_proto = out.File
	file_gedo_mazo_proto_rawDesc = nil
	file_gedo_mazo_proto_goTypes = nil
	file_gedo_mazo_proto_depIdxs = nil
}
