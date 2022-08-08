// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: M3/Metadata/Metadata.proto

package metadata

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

type FieldType int32

const (
	FieldType_string FieldType = 0
	FieldType_number FieldType = 1
)

// Enum value maps for FieldType.
var (
	FieldType_name = map[int32]string{
		0: "string",
		1: "number",
	}
	FieldType_value = map[string]int32{
		"string": 0,
		"number": 1,
	}
)

func (x FieldType) Enum() *FieldType {
	p := new(FieldType)
	*p = x
	return p
}

func (x FieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_M3_Metadata_Metadata_proto_enumTypes[0].Descriptor()
}

func (FieldType) Type() protoreflect.EnumType {
	return &file_M3_Metadata_Metadata_proto_enumTypes[0]
}

func (x FieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldType.Descriptor instead.
func (FieldType) EnumDescriptor() ([]byte, []int) {
	return file_M3_Metadata_Metadata_proto_rawDescGZIP(), []int{0}
}

type SortDirection int32

const (
	SortDirection_Ascending  SortDirection = 0
	SortDirection_Descending SortDirection = 1
)

// Enum value maps for SortDirection.
var (
	SortDirection_name = map[int32]string{
		0: "Ascending",
		1: "Descending",
	}
	SortDirection_value = map[string]int32{
		"Ascending":  0,
		"Descending": 1,
	}
)

func (x SortDirection) Enum() *SortDirection {
	p := new(SortDirection)
	*p = x
	return p
}

func (x SortDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_M3_Metadata_Metadata_proto_enumTypes[1].Descriptor()
}

func (SortDirection) Type() protoreflect.EnumType {
	return &file_M3_Metadata_Metadata_proto_enumTypes[1]
}

func (x SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortDirection.Descriptor instead.
func (SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_M3_Metadata_Metadata_proto_rawDescGZIP(), []int{1}
}

type TableMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string                `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Columns     []*ColumnMetadata     `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	PrimaryKey  *Index                `protobuf:"bytes,4,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
	Indexes     []*Index              `protobuf:"bytes,5,rep,name=indexes,proto3" json:"indexes,omitempty"`
	ForeignKeys []*ForeignKeyRelation `protobuf:"bytes,6,rep,name=foreign_keys,json=foreignKeys,proto3" json:"foreign_keys,omitempty"`
}

func (x *TableMetadata) Reset() {
	*x = TableMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M3_Metadata_Metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableMetadata) ProtoMessage() {}

func (x *TableMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_M3_Metadata_Metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableMetadata.ProtoReflect.Descriptor instead.
func (*TableMetadata) Descriptor() ([]byte, []int) {
	return file_M3_Metadata_Metadata_proto_rawDescGZIP(), []int{0}
}

func (x *TableMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TableMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TableMetadata) GetColumns() []*ColumnMetadata {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *TableMetadata) GetPrimaryKey() *Index {
	if x != nil {
		return x.PrimaryKey
	}
	return nil
}

func (x *TableMetadata) GetIndexes() []*Index {
	if x != nil {
		return x.Indexes
	}
	return nil
}

func (x *TableMetadata) GetForeignKeys() []*ForeignKeyRelation {
	if x != nil {
		return x.ForeignKeys
	}
	return nil
}

type ColumnMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          FieldType `protobuf:"varint,2,opt,name=type,proto3,enum=dolittle.integrations.m3.metadata.FieldType" json:"type,omitempty"`
	FieldLength   int32     `protobuf:"varint,3,opt,name=field_length,json=fieldLength,proto3" json:"field_length,omitempty"`
	DecimalPlaces int32     `protobuf:"varint,4,opt,name=decimal_places,json=decimalPlaces,proto3" json:"decimal_places,omitempty"`
}

func (x *ColumnMetadata) Reset() {
	*x = ColumnMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M3_Metadata_Metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnMetadata) ProtoMessage() {}

func (x *ColumnMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_M3_Metadata_Metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnMetadata.ProtoReflect.Descriptor instead.
func (*ColumnMetadata) Descriptor() ([]byte, []int) {
	return file_M3_Metadata_Metadata_proto_rawDescGZIP(), []int{1}
}

func (x *ColumnMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ColumnMetadata) GetType() FieldType {
	if x != nil {
		return x.Type
	}
	return FieldType_string
}

func (x *ColumnMetadata) GetFieldLength() int32 {
	if x != nil {
		return x.FieldLength
	}
	return 0
}

func (x *ColumnMetadata) GetDecimalPlaces() int32 {
	if x != nil {
		return x.DecimalPlaces
	}
	return 0
}

type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Columns []*IndexColumn `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	Unique  bool           `protobuf:"varint,3,opt,name=unique,proto3" json:"unique,omitempty"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M3_Metadata_Metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_M3_Metadata_Metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_M3_Metadata_Metadata_proto_rawDescGZIP(), []int{2}
}

func (x *Index) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Index) GetColumns() []*IndexColumn {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Index) GetUnique() bool {
	if x != nil {
		return x.Unique
	}
	return false
}

type IndexColumn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Direction SortDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=dolittle.integrations.m3.metadata.SortDirection" json:"direction,omitempty"`
}

func (x *IndexColumn) Reset() {
	*x = IndexColumn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M3_Metadata_Metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexColumn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexColumn) ProtoMessage() {}

func (x *IndexColumn) ProtoReflect() protoreflect.Message {
	mi := &file_M3_Metadata_Metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexColumn.ProtoReflect.Descriptor instead.
func (*IndexColumn) Descriptor() ([]byte, []int) {
	return file_M3_Metadata_Metadata_proto_rawDescGZIP(), []int{3}
}

func (x *IndexColumn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IndexColumn) GetDirection() SortDirection {
	if x != nil {
		return x.Direction
	}
	return SortDirection_Ascending
}

type ForeignKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string                `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Key   []*ForeignKeyRelation `protobuf:"bytes,2,rep,name=key,proto3" json:"key,omitempty"`
}

func (x *ForeignKey) Reset() {
	*x = ForeignKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M3_Metadata_Metadata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForeignKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForeignKey) ProtoMessage() {}

func (x *ForeignKey) ProtoReflect() protoreflect.Message {
	mi := &file_M3_Metadata_Metadata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForeignKey.ProtoReflect.Descriptor instead.
func (*ForeignKey) Descriptor() ([]byte, []int) {
	return file_M3_Metadata_Metadata_proto_rawDescGZIP(), []int{4}
}

func (x *ForeignKey) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *ForeignKey) GetKey() []*ForeignKeyRelation {
	if x != nil {
		return x.Key
	}
	return nil
}

type ForeignKeyRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Column        string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	ForeignColumn string `protobuf:"bytes,2,opt,name=foreign_column,json=foreignColumn,proto3" json:"foreign_column,omitempty"`
}

func (x *ForeignKeyRelation) Reset() {
	*x = ForeignKeyRelation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M3_Metadata_Metadata_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForeignKeyRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForeignKeyRelation) ProtoMessage() {}

func (x *ForeignKeyRelation) ProtoReflect() protoreflect.Message {
	mi := &file_M3_Metadata_Metadata_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForeignKeyRelation.ProtoReflect.Descriptor instead.
func (*ForeignKeyRelation) Descriptor() ([]byte, []int) {
	return file_M3_Metadata_Metadata_proto_rawDescGZIP(), []int{5}
}

func (x *ForeignKeyRelation) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *ForeignKeyRelation) GetForeignColumn() string {
	if x != nil {
		return x.ForeignColumn
	}
	return ""
}

var File_M3_Metadata_Metadata_proto protoreflect.FileDescriptor

var file_M3_Metadata_Metadata_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x4d, 0x33, 0x2f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x64, 0x6f,
	0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x33, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0xfb, 0x02, 0x0a, 0x0d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x6d, 0x33, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x07, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x73, 0x12, 0x49, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x6d, 0x33, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12,
	0x42, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x33, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x5f, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x6d, 0x33, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x6f,
	0x72, 0x65, 0x69, 0x67, 0x6e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x4b, 0x65, 0x79, 0x73, 0x22, 0xb0, 0x01,
	0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x33, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x63,
	0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x73,
	0x22, 0x7d, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a,
	0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x33, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x22,
	0x71, 0x0a, 0x0b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x33,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x6b, 0x0a, 0x0a, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x4b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x47, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x33, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x53, 0x0a, 0x12, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x43, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x2a, 0x23, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x01, 0x2a, 0x2e, 0x0a, 0x0d, 0x53, 0x6f, 0x72,
	0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x73,
	0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x65, 0x73,
	0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x42, 0x64, 0x5a, 0x34, 0x67, 0x6f, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x33, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xaa, 0x02, 0x2b, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2e, 0x4d, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_M3_Metadata_Metadata_proto_rawDescOnce sync.Once
	file_M3_Metadata_Metadata_proto_rawDescData = file_M3_Metadata_Metadata_proto_rawDesc
)

func file_M3_Metadata_Metadata_proto_rawDescGZIP() []byte {
	file_M3_Metadata_Metadata_proto_rawDescOnce.Do(func() {
		file_M3_Metadata_Metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_M3_Metadata_Metadata_proto_rawDescData)
	})
	return file_M3_Metadata_Metadata_proto_rawDescData
}

var file_M3_Metadata_Metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_M3_Metadata_Metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_M3_Metadata_Metadata_proto_goTypes = []interface{}{
	(FieldType)(0),             // 0: dolittle.integrations.m3.metadata.FieldType
	(SortDirection)(0),         // 1: dolittle.integrations.m3.metadata.SortDirection
	(*TableMetadata)(nil),      // 2: dolittle.integrations.m3.metadata.TableMetadata
	(*ColumnMetadata)(nil),     // 3: dolittle.integrations.m3.metadata.ColumnMetadata
	(*Index)(nil),              // 4: dolittle.integrations.m3.metadata.Index
	(*IndexColumn)(nil),        // 5: dolittle.integrations.m3.metadata.IndexColumn
	(*ForeignKey)(nil),         // 6: dolittle.integrations.m3.metadata.ForeignKey
	(*ForeignKeyRelation)(nil), // 7: dolittle.integrations.m3.metadata.ForeignKeyRelation
}
var file_M3_Metadata_Metadata_proto_depIdxs = []int32{
	3, // 0: dolittle.integrations.m3.metadata.TableMetadata.columns:type_name -> dolittle.integrations.m3.metadata.ColumnMetadata
	4, // 1: dolittle.integrations.m3.metadata.TableMetadata.primary_key:type_name -> dolittle.integrations.m3.metadata.Index
	4, // 2: dolittle.integrations.m3.metadata.TableMetadata.indexes:type_name -> dolittle.integrations.m3.metadata.Index
	7, // 3: dolittle.integrations.m3.metadata.TableMetadata.foreign_keys:type_name -> dolittle.integrations.m3.metadata.ForeignKeyRelation
	0, // 4: dolittle.integrations.m3.metadata.ColumnMetadata.type:type_name -> dolittle.integrations.m3.metadata.FieldType
	5, // 5: dolittle.integrations.m3.metadata.Index.columns:type_name -> dolittle.integrations.m3.metadata.IndexColumn
	1, // 6: dolittle.integrations.m3.metadata.IndexColumn.direction:type_name -> dolittle.integrations.m3.metadata.SortDirection
	7, // 7: dolittle.integrations.m3.metadata.ForeignKey.key:type_name -> dolittle.integrations.m3.metadata.ForeignKeyRelation
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_M3_Metadata_Metadata_proto_init() }
func file_M3_Metadata_Metadata_proto_init() {
	if File_M3_Metadata_Metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_M3_Metadata_Metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableMetadata); i {
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
		file_M3_Metadata_Metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnMetadata); i {
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
		file_M3_Metadata_Metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
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
		file_M3_Metadata_Metadata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexColumn); i {
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
		file_M3_Metadata_Metadata_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForeignKey); i {
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
		file_M3_Metadata_Metadata_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForeignKeyRelation); i {
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
			RawDescriptor: file_M3_Metadata_Metadata_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_M3_Metadata_Metadata_proto_goTypes,
		DependencyIndexes: file_M3_Metadata_Metadata_proto_depIdxs,
		EnumInfos:         file_M3_Metadata_Metadata_proto_enumTypes,
		MessageInfos:      file_M3_Metadata_Metadata_proto_msgTypes,
	}.Build()
	File_M3_Metadata_Metadata_proto = out.File
	file_M3_Metadata_Metadata_proto_rawDesc = nil
	file_M3_Metadata_Metadata_proto_goTypes = nil
	file_M3_Metadata_Metadata_proto_depIdxs = nil
}
