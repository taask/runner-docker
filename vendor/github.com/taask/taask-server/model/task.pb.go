// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package model

import (
	fmt "fmt"
	simplcrypto "github.com/cohix/simplcrypto"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Task struct {
	Meta                 *TaskMeta            `protobuf:"bytes,1,opt,name=Meta,proto3" json:"Meta,omitempty"`
	UUID                 string               `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Kind                 string               `protobuf:"bytes,3,opt,name=Kind,proto3" json:"Kind,omitempty"`
	Status               string               `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	EncBody              *simplcrypto.Message `protobuf:"bytes,5,opt,name=EncBody,proto3" json:"EncBody,omitempty"`
	EncResult            *simplcrypto.Message `protobuf:"bytes,6,opt,name=EncResult,proto3" json:"EncResult,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetMeta() *TaskMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Task) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Task) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Task) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Task) GetEncBody() *simplcrypto.Message {
	if m != nil {
		return m.EncBody
	}
	return nil
}

func (m *Task) GetEncResult() *simplcrypto.Message {
	if m != nil {
		return m.EncResult
	}
	return nil
}

type TaskMeta struct {
	Annotations          []string             `protobuf:"bytes,1,rep,name=Annotations,proto3" json:"Annotations,omitempty"`
	RunnerUUID           string               `protobuf:"bytes,2,opt,name=RunnerUUID,proto3" json:"RunnerUUID,omitempty"`
	ChildRunnerUUID      string               `protobuf:"bytes,3,opt,name=ChildRunnerUUID,proto3" json:"ChildRunnerUUID,omitempty"`
	MasterEncTaskKey     *simplcrypto.Message `protobuf:"bytes,4,opt,name=MasterEncTaskKey,proto3" json:"MasterEncTaskKey,omitempty"`
	RunnerEncTaskKey     *simplcrypto.Message `protobuf:"bytes,5,opt,name=RunnerEncTaskKey,proto3" json:"RunnerEncTaskKey,omitempty"`
	ClientEncTaskKey     *simplcrypto.Message `protobuf:"bytes,6,opt,name=ClientEncTaskKey,proto3" json:"ClientEncTaskKey,omitempty"`
	RetrySeconds         int32                `protobuf:"varint,7,opt,name=RetrySeconds,proto3" json:"RetrySeconds,omitempty"`
	TimeoutSeconds       int32                `protobuf:"varint,8,opt,name=TimeoutSeconds,proto3" json:"TimeoutSeconds,omitempty"`
	Version              int32                `protobuf:"varint,9,opt,name=Version,proto3" json:"Version,omitempty"`
	GroupUUID            string               `protobuf:"bytes,10,opt,name=GroupUUID,proto3" json:"GroupUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskMeta) Reset()         { *m = TaskMeta{} }
func (m *TaskMeta) String() string { return proto.CompactTextString(m) }
func (*TaskMeta) ProtoMessage()    {}
func (*TaskMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{1}
}

func (m *TaskMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskMeta.Unmarshal(m, b)
}
func (m *TaskMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskMeta.Marshal(b, m, deterministic)
}
func (m *TaskMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskMeta.Merge(m, src)
}
func (m *TaskMeta) XXX_Size() int {
	return xxx_messageInfo_TaskMeta.Size(m)
}
func (m *TaskMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskMeta.DiscardUnknown(m)
}

var xxx_messageInfo_TaskMeta proto.InternalMessageInfo

func (m *TaskMeta) GetAnnotations() []string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *TaskMeta) GetRunnerUUID() string {
	if m != nil {
		return m.RunnerUUID
	}
	return ""
}

func (m *TaskMeta) GetChildRunnerUUID() string {
	if m != nil {
		return m.ChildRunnerUUID
	}
	return ""
}

func (m *TaskMeta) GetMasterEncTaskKey() *simplcrypto.Message {
	if m != nil {
		return m.MasterEncTaskKey
	}
	return nil
}

func (m *TaskMeta) GetRunnerEncTaskKey() *simplcrypto.Message {
	if m != nil {
		return m.RunnerEncTaskKey
	}
	return nil
}

func (m *TaskMeta) GetClientEncTaskKey() *simplcrypto.Message {
	if m != nil {
		return m.ClientEncTaskKey
	}
	return nil
}

func (m *TaskMeta) GetRetrySeconds() int32 {
	if m != nil {
		return m.RetrySeconds
	}
	return 0
}

func (m *TaskMeta) GetTimeoutSeconds() int32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *TaskMeta) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TaskMeta) GetGroupUUID() string {
	if m != nil {
		return m.GroupUUID
	}
	return ""
}

type TaskUpdate struct {
	UUID                 string               `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Status               string               `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	EncResult            *simplcrypto.Message `protobuf:"bytes,3,opt,name=EncResult,proto3" json:"EncResult,omitempty"`
	RunnerEncTaskKey     *simplcrypto.Message `protobuf:"bytes,4,opt,name=RunnerEncTaskKey,proto3" json:"RunnerEncTaskKey,omitempty"`
	RunnerUUID           string               `protobuf:"bytes,5,opt,name=RunnerUUID,proto3" json:"RunnerUUID,omitempty"`
	ChildRunnerUUID      string               `protobuf:"bytes,6,opt,name=ChildRunnerUUID,proto3" json:"ChildRunnerUUID,omitempty"`
	RetrySeconds         int32                `protobuf:"varint,7,opt,name=RetrySeconds,proto3" json:"RetrySeconds,omitempty"`
	Version              int32                `protobuf:"varint,8,opt,name=Version,proto3" json:"Version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskUpdate) Reset()         { *m = TaskUpdate{} }
func (m *TaskUpdate) String() string { return proto.CompactTextString(m) }
func (*TaskUpdate) ProtoMessage()    {}
func (*TaskUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{2}
}

func (m *TaskUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskUpdate.Unmarshal(m, b)
}
func (m *TaskUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskUpdate.Marshal(b, m, deterministic)
}
func (m *TaskUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskUpdate.Merge(m, src)
}
func (m *TaskUpdate) XXX_Size() int {
	return xxx_messageInfo_TaskUpdate.Size(m)
}
func (m *TaskUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_TaskUpdate proto.InternalMessageInfo

func (m *TaskUpdate) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *TaskUpdate) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TaskUpdate) GetEncResult() *simplcrypto.Message {
	if m != nil {
		return m.EncResult
	}
	return nil
}

func (m *TaskUpdate) GetRunnerEncTaskKey() *simplcrypto.Message {
	if m != nil {
		return m.RunnerEncTaskKey
	}
	return nil
}

func (m *TaskUpdate) GetRunnerUUID() string {
	if m != nil {
		return m.RunnerUUID
	}
	return ""
}

func (m *TaskUpdate) GetChildRunnerUUID() string {
	if m != nil {
		return m.ChildRunnerUUID
	}
	return ""
}

func (m *TaskUpdate) GetRetrySeconds() int32 {
	if m != nil {
		return m.RetrySeconds
	}
	return 0
}

func (m *TaskUpdate) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*Task)(nil), "taask.server.model.Task")
	proto.RegisterType((*TaskMeta)(nil), "taask.server.model.TaskMeta")
	proto.RegisterType((*TaskUpdate)(nil), "taask.server.model.TaskUpdate")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x55, 0xd6, 0x34, 0x6d, 0xbe, 0x21, 0x40, 0x96, 0x98, 0x2c, 0x34, 0xa1, 0xa8, 0x08, 0x94,
	0x0b, 0xc9, 0xc4, 0x7e, 0xc1, 0x36, 0x2a, 0x84, 0xa6, 0x5e, 0xbc, 0x95, 0x03, 0x37, 0x2f, 0xb1,
	0x56, 0x6b, 0x89, 0x1d, 0xd9, 0x0e, 0xa2, 0xff, 0x8c, 0x3b, 0xff, 0x83, 0xdf, 0x82, 0x6c, 0xaf,
	0xcc, 0x5d, 0x46, 0x55, 0x2e, 0x55, 0xbe, 0xa7, 0xe7, 0xd7, 0xe7, 0xe7, 0x67, 0x03, 0x18, 0xaa,
	0xef, 0x8a, 0x4e, 0x49, 0x23, 0x11, 0x32, 0xd4, 0x0e, 0x9a, 0xa9, 0xef, 0x4c, 0x15, 0xad, 0xac,
	0x59, 0xf3, 0xfa, 0xe4, 0x96, 0x9b, 0x55, 0x7f, 0x53, 0x54, 0xb2, 0x2d, 0x2b, 0xb9, 0xe2, 0x3f,
	0x4a, 0xcd, 0xdb, 0xae, 0xa9, 0xd4, 0xba, 0x33, 0xb2, 0x74, 0xeb, 0xca, 0x96, 0x69, 0x4d, 0x6f,
	0x99, 0x57, 0x99, 0xfd, 0x8e, 0x20, 0xbe, 0xa6, 0xfa, 0x0e, 0x9d, 0x40, 0xbc, 0x60, 0x86, 0xe2,
	0x28, 0x8b, 0xf2, 0xc3, 0x8f, 0xc7, 0xc5, 0x50, 0xbd, 0xb0, 0x3c, 0xcb, 0x21, 0x8e, 0x89, 0x10,
	0xc4, 0xcb, 0xe5, 0x97, 0x4f, 0xf8, 0x20, 0x8b, 0xf2, 0x94, 0xb8, 0x6f, 0x8b, 0x5d, 0x72, 0x51,
	0xe3, 0x91, 0xc7, 0xec, 0x37, 0x3a, 0x82, 0xe4, 0xca, 0x50, 0xd3, 0x6b, 0x1c, 0x3b, 0xf4, 0x7e,
	0x42, 0x25, 0x4c, 0xe6, 0xa2, 0x3a, 0x97, 0xf5, 0x1a, 0x8f, 0xdd, 0x9f, 0xbe, 0x2a, 0x9c, 0xdb,
	0xc2, 0xdb, 0x2d, 0x16, 0xde, 0x28, 0xd9, 0xb0, 0xd0, 0x29, 0xa4, 0x73, 0x51, 0x11, 0xa6, 0xfb,
	0xc6, 0xe0, 0x64, 0xd7, 0x92, 0x07, 0xde, 0xec, 0xd7, 0x08, 0xa6, 0x1b, 0xe3, 0x28, 0x83, 0xc3,
	0x33, 0x21, 0xa4, 0xa1, 0x86, 0x4b, 0xa1, 0x71, 0x94, 0x8d, 0xf2, 0x94, 0x84, 0x10, 0x7a, 0x03,
	0x40, 0x7a, 0x21, 0x98, 0x0a, 0xb6, 0x16, 0x20, 0x28, 0x87, 0x17, 0x17, 0x2b, 0xde, 0xd4, 0x01,
	0xc9, 0xef, 0xf5, 0x31, 0x8c, 0xce, 0xe0, 0xe5, 0x82, 0x6a, 0xc3, 0xd4, 0x5c, 0x54, 0xd6, 0xc0,
	0x25, 0x5b, 0xbb, 0x00, 0xfe, 0x69, 0x7a, 0x40, 0xb7, 0x12, 0x5e, 0x30, 0x90, 0xd8, 0x19, 0xd5,
	0x80, 0x6e, 0x25, 0x2e, 0x1a, 0xce, 0x84, 0x09, 0x24, 0x76, 0x46, 0x37, 0xa0, 0xa3, 0x19, 0x3c,
	0x23, 0xcc, 0xa8, 0xf5, 0x15, 0xab, 0xa4, 0xa8, 0x35, 0x9e, 0x64, 0x51, 0x3e, 0x26, 0x5b, 0x18,
	0x7a, 0x0f, 0xcf, 0xaf, 0x79, 0xcb, 0x64, 0x6f, 0x36, 0xac, 0xa9, 0x63, 0x3d, 0x42, 0x11, 0x86,
	0xc9, 0x57, 0xa6, 0x34, 0x97, 0x02, 0xa7, 0x8e, 0xb0, 0x19, 0xd1, 0x31, 0xa4, 0x9f, 0x95, 0xec,
	0x3b, 0x17, 0x29, 0xb8, 0x48, 0x1f, 0x80, 0xd9, 0xcf, 0x03, 0x00, 0xeb, 0x67, 0xd9, 0xd5, 0xd4,
	0xb0, 0xbf, 0xd5, 0x8b, 0x82, 0xea, 0x1d, 0x41, 0xa2, 0x7d, 0xcd, 0xfc, 0xa9, 0xdd, 0x4f, 0xdb,
	0xad, 0x19, 0xed, 0xd7, 0x9a, 0x27, 0x93, 0x8f, 0xff, 0x2f, 0xf9, 0xed, 0x26, 0x8d, 0xf7, 0x69,
	0x52, 0xf2, 0x74, 0x93, 0xf6, 0x39, 0x80, 0x20, 0xd8, 0xe9, 0x56, 0xb0, 0xe7, 0xef, 0xbe, 0xbd,
	0x0d, 0x5e, 0x05, 0x77, 0xad, 0xfd, 0xef, 0x07, 0x7f, 0xb9, 0x4b, 0x77, 0xb9, 0x6f, 0x12, 0xf7,
	0x1e, 0x9c, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x00, 0xe5, 0x44, 0x23, 0x63, 0x04, 0x00, 0x00,
}