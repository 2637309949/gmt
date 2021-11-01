// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/helloworld/helloworld.proto

package helloworld

import (
	fmt "fmt"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Article struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Creater              int64    `protobuf:"varint,2,opt,name=creater,proto3" json:"creater"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=create_time,proto3" json:"create_time"`
	Updater              int64    `protobuf:"varint,4,opt,name=updater,proto3" json:"updater"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,proto3" json:"update_time"`
	IsDelete             uint32   `protobuf:"varint,6,opt,name=is_delete,json=-,proto3" json:"is_delete"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name"`
	XUser                string   `protobuf:"bytes,9,opt,name=x_user,proto3" json:"x_user"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_d16ac54e3a89a4f6, []int{0}
}

func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Article) GetCreater() int64 {
	if m != nil {
		return m.Creater
	}
	return 0
}

func (m *Article) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *Article) GetUpdater() int64 {
	if m != nil {
		return m.Updater
	}
	return 0
}

func (m *Article) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *Article) GetIsDelete() uint32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *Article) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *Article) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Article) GetXUser() string {
	if m != nil {
		return m.XUser
	}
	return ""
}

type ArticleAddReq struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Creater              int64    `protobuf:"varint,2,opt,name=creater,proto3" json:"creater"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=create_time,proto3" json:"create_time"`
	Updater              int64    `protobuf:"varint,4,opt,name=updater,proto3" json:"updater"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,proto3" json:"update_time"`
	IsDelete             uint32   `protobuf:"varint,6,opt,name=is_delete,json=-,proto3" json:"is_delete"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name"`
	XUser                string   `protobuf:"bytes,9,opt,name=x_user,proto3" json:"x_user"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleAddReq) Reset()         { *m = ArticleAddReq{} }
func (m *ArticleAddReq) String() string { return proto.CompactTextString(m) }
func (*ArticleAddReq) ProtoMessage()    {}
func (*ArticleAddReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d16ac54e3a89a4f6, []int{1}
}

func (m *ArticleAddReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleAddReq.Unmarshal(m, b)
}
func (m *ArticleAddReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleAddReq.Marshal(b, m, deterministic)
}
func (m *ArticleAddReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleAddReq.Merge(m, src)
}
func (m *ArticleAddReq) XXX_Size() int {
	return xxx_messageInfo_ArticleAddReq.Size(m)
}
func (m *ArticleAddReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleAddReq.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleAddReq proto.InternalMessageInfo

func (m *ArticleAddReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleAddReq) GetCreater() int64 {
	if m != nil {
		return m.Creater
	}
	return 0
}

func (m *ArticleAddReq) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *ArticleAddReq) GetUpdater() int64 {
	if m != nil {
		return m.Updater
	}
	return 0
}

func (m *ArticleAddReq) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *ArticleAddReq) GetIsDelete() uint32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *ArticleAddReq) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ArticleAddReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArticleAddReq) GetXUser() string {
	if m != nil {
		return m.XUser
	}
	return ""
}

type ArticleAddRes struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Creater              int64    `protobuf:"varint,2,opt,name=creater,proto3" json:"creater"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=create_time,proto3" json:"create_time"`
	Updater              int64    `protobuf:"varint,4,opt,name=updater,proto3" json:"updater"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,proto3" json:"update_time"`
	IsDelete             uint32   `protobuf:"varint,6,opt,name=is_delete,json=-,proto3" json:"is_delete"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name"`
	XUser                string   `protobuf:"bytes,9,opt,name=x_user,proto3" json:"x_user"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleAddRes) Reset()         { *m = ArticleAddRes{} }
func (m *ArticleAddRes) String() string { return proto.CompactTextString(m) }
func (*ArticleAddRes) ProtoMessage()    {}
func (*ArticleAddRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d16ac54e3a89a4f6, []int{2}
}

func (m *ArticleAddRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleAddRes.Unmarshal(m, b)
}
func (m *ArticleAddRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleAddRes.Marshal(b, m, deterministic)
}
func (m *ArticleAddRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleAddRes.Merge(m, src)
}
func (m *ArticleAddRes) XXX_Size() int {
	return xxx_messageInfo_ArticleAddRes.Size(m)
}
func (m *ArticleAddRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleAddRes.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleAddRes proto.InternalMessageInfo

func (m *ArticleAddRes) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleAddRes) GetCreater() int64 {
	if m != nil {
		return m.Creater
	}
	return 0
}

func (m *ArticleAddRes) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *ArticleAddRes) GetUpdater() int64 {
	if m != nil {
		return m.Updater
	}
	return 0
}

func (m *ArticleAddRes) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *ArticleAddRes) GetIsDelete() uint32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *ArticleAddRes) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ArticleAddRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArticleAddRes) GetXUser() string {
	if m != nil {
		return m.XUser
	}
	return ""
}

type ArticleDelReq struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Creater              int64    `protobuf:"varint,2,opt,name=creater,proto3" json:"creater"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=create_time,proto3" json:"create_time"`
	Updater              int64    `protobuf:"varint,4,opt,name=updater,proto3" json:"updater"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,proto3" json:"update_time"`
	IsDelete             uint32   `protobuf:"varint,6,opt,name=is_delete,json=-,proto3" json:"is_delete"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name"`
	XUser                string   `protobuf:"bytes,9,opt,name=x_user,proto3" json:"x_user"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleDelReq) Reset()         { *m = ArticleDelReq{} }
func (m *ArticleDelReq) String() string { return proto.CompactTextString(m) }
func (*ArticleDelReq) ProtoMessage()    {}
func (*ArticleDelReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d16ac54e3a89a4f6, []int{3}
}

func (m *ArticleDelReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleDelReq.Unmarshal(m, b)
}
func (m *ArticleDelReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleDelReq.Marshal(b, m, deterministic)
}
func (m *ArticleDelReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleDelReq.Merge(m, src)
}
func (m *ArticleDelReq) XXX_Size() int {
	return xxx_messageInfo_ArticleDelReq.Size(m)
}
func (m *ArticleDelReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleDelReq.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleDelReq proto.InternalMessageInfo

func (m *ArticleDelReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleDelReq) GetCreater() int64 {
	if m != nil {
		return m.Creater
	}
	return 0
}

func (m *ArticleDelReq) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *ArticleDelReq) GetUpdater() int64 {
	if m != nil {
		return m.Updater
	}
	return 0
}

func (m *ArticleDelReq) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *ArticleDelReq) GetIsDelete() uint32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *ArticleDelReq) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ArticleDelReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArticleDelReq) GetXUser() string {
	if m != nil {
		return m.XUser
	}
	return ""
}

type ArticleDelRes struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Creater              int64    `protobuf:"varint,2,opt,name=creater,proto3" json:"creater"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=create_time,proto3" json:"create_time"`
	Updater              int64    `protobuf:"varint,4,opt,name=updater,proto3" json:"updater"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,proto3" json:"update_time"`
	IsDelete             uint32   `protobuf:"varint,6,opt,name=is_delete,json=-,proto3" json:"is_delete"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name"`
	XUser                string   `protobuf:"bytes,9,opt,name=x_user,proto3" json:"x_user"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleDelRes) Reset()         { *m = ArticleDelRes{} }
func (m *ArticleDelRes) String() string { return proto.CompactTextString(m) }
func (*ArticleDelRes) ProtoMessage()    {}
func (*ArticleDelRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d16ac54e3a89a4f6, []int{4}
}

func (m *ArticleDelRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleDelRes.Unmarshal(m, b)
}
func (m *ArticleDelRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleDelRes.Marshal(b, m, deterministic)
}
func (m *ArticleDelRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleDelRes.Merge(m, src)
}
func (m *ArticleDelRes) XXX_Size() int {
	return xxx_messageInfo_ArticleDelRes.Size(m)
}
func (m *ArticleDelRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleDelRes.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleDelRes proto.InternalMessageInfo

func (m *ArticleDelRes) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleDelRes) GetCreater() int64 {
	if m != nil {
		return m.Creater
	}
	return 0
}

func (m *ArticleDelRes) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *ArticleDelRes) GetUpdater() int64 {
	if m != nil {
		return m.Updater
	}
	return 0
}

func (m *ArticleDelRes) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *ArticleDelRes) GetIsDelete() uint32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *ArticleDelRes) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ArticleDelRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArticleDelRes) GetXUser() string {
	if m != nil {
		return m.XUser
	}
	return ""
}

type ArticleUpdateReq struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Creater              int64    `protobuf:"varint,2,opt,name=creater,proto3" json:"creater"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=create_time,proto3" json:"create_time"`
	Updater              int64    `protobuf:"varint,4,opt,name=updater,proto3" json:"updater"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,proto3" json:"update_time"`
	IsDelete             uint32   `protobuf:"varint,6,opt,name=is_delete,json=-,proto3" json:"is_delete"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name"`
	XUser                string   `protobuf:"bytes,9,opt,name=x_user,proto3" json:"x_user"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleUpdateReq) Reset()         { *m = ArticleUpdateReq{} }
func (m *ArticleUpdateReq) String() string { return proto.CompactTextString(m) }
func (*ArticleUpdateReq) ProtoMessage()    {}
func (*ArticleUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d16ac54e3a89a4f6, []int{5}
}

func (m *ArticleUpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleUpdateReq.Unmarshal(m, b)
}
func (m *ArticleUpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleUpdateReq.Marshal(b, m, deterministic)
}
func (m *ArticleUpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleUpdateReq.Merge(m, src)
}
func (m *ArticleUpdateReq) XXX_Size() int {
	return xxx_messageInfo_ArticleUpdateReq.Size(m)
}
func (m *ArticleUpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleUpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleUpdateReq proto.InternalMessageInfo

func (m *ArticleUpdateReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleUpdateReq) GetCreater() int64 {
	if m != nil {
		return m.Creater
	}
	return 0
}

func (m *ArticleUpdateReq) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *ArticleUpdateReq) GetUpdater() int64 {
	if m != nil {
		return m.Updater
	}
	return 0
}

func (m *ArticleUpdateReq) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *ArticleUpdateReq) GetIsDelete() uint32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *ArticleUpdateReq) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ArticleUpdateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArticleUpdateReq) GetXUser() string {
	if m != nil {
		return m.XUser
	}
	return ""
}

type ArticleUpdateRes struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Creater              int64    `protobuf:"varint,2,opt,name=creater,proto3" json:"creater"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=create_time,proto3" json:"create_time"`
	Updater              int64    `protobuf:"varint,4,opt,name=updater,proto3" json:"updater"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,proto3" json:"update_time"`
	IsDelete             uint32   `protobuf:"varint,6,opt,name=is_delete,json=-,proto3" json:"is_delete"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name"`
	XUser                string   `protobuf:"bytes,9,opt,name=x_user,proto3" json:"x_user"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleUpdateRes) Reset()         { *m = ArticleUpdateRes{} }
func (m *ArticleUpdateRes) String() string { return proto.CompactTextString(m) }
func (*ArticleUpdateRes) ProtoMessage()    {}
func (*ArticleUpdateRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d16ac54e3a89a4f6, []int{6}
}

func (m *ArticleUpdateRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleUpdateRes.Unmarshal(m, b)
}
func (m *ArticleUpdateRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleUpdateRes.Marshal(b, m, deterministic)
}
func (m *ArticleUpdateRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleUpdateRes.Merge(m, src)
}
func (m *ArticleUpdateRes) XXX_Size() int {
	return xxx_messageInfo_ArticleUpdateRes.Size(m)
}
func (m *ArticleUpdateRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleUpdateRes.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleUpdateRes proto.InternalMessageInfo

func (m *ArticleUpdateRes) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleUpdateRes) GetCreater() int64 {
	if m != nil {
		return m.Creater
	}
	return 0
}

func (m *ArticleUpdateRes) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *ArticleUpdateRes) GetUpdater() int64 {
	if m != nil {
		return m.Updater
	}
	return 0
}

func (m *ArticleUpdateRes) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *ArticleUpdateRes) GetIsDelete() uint32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *ArticleUpdateRes) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ArticleUpdateRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArticleUpdateRes) GetXUser() string {
	if m != nil {
		return m.XUser
	}
	return ""
}

type ArticleOneReq struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Creater              int64    `protobuf:"varint,2,opt,name=creater,proto3" json:"creater"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=create_time,proto3" json:"create_time"`
	Updater              int64    `protobuf:"varint,4,opt,name=updater,proto3" json:"updater"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,proto3" json:"update_time"`
	IsDelete             uint32   `protobuf:"varint,6,opt,name=is_delete,json=-,proto3" json:"is_delete"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name"`
	XUser                string   `protobuf:"bytes,9,opt,name=x_user,proto3" json:"x_user"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleOneReq) Reset()         { *m = ArticleOneReq{} }
func (m *ArticleOneReq) String() string { return proto.CompactTextString(m) }
func (*ArticleOneReq) ProtoMessage()    {}
func (*ArticleOneReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d16ac54e3a89a4f6, []int{7}
}

func (m *ArticleOneReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleOneReq.Unmarshal(m, b)
}
func (m *ArticleOneReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleOneReq.Marshal(b, m, deterministic)
}
func (m *ArticleOneReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleOneReq.Merge(m, src)
}
func (m *ArticleOneReq) XXX_Size() int {
	return xxx_messageInfo_ArticleOneReq.Size(m)
}
func (m *ArticleOneReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleOneReq.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleOneReq proto.InternalMessageInfo

func (m *ArticleOneReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleOneReq) GetCreater() int64 {
	if m != nil {
		return m.Creater
	}
	return 0
}

func (m *ArticleOneReq) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *ArticleOneReq) GetUpdater() int64 {
	if m != nil {
		return m.Updater
	}
	return 0
}

func (m *ArticleOneReq) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *ArticleOneReq) GetIsDelete() uint32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *ArticleOneReq) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ArticleOneReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArticleOneReq) GetXUser() string {
	if m != nil {
		return m.XUser
	}
	return ""
}

type ArticleOneRes struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Creater              int64    `protobuf:"varint,2,opt,name=creater,proto3" json:"creater"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=create_time,proto3" json:"create_time"`
	Updater              int64    `protobuf:"varint,4,opt,name=updater,proto3" json:"updater"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,proto3" json:"update_time"`
	IsDelete             uint32   `protobuf:"varint,6,opt,name=is_delete,json=-,proto3" json:"is_delete"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name"`
	XUser                string   `protobuf:"bytes,9,opt,name=x_user,proto3" json:"x_user"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleOneRes) Reset()         { *m = ArticleOneRes{} }
func (m *ArticleOneRes) String() string { return proto.CompactTextString(m) }
func (*ArticleOneRes) ProtoMessage()    {}
func (*ArticleOneRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d16ac54e3a89a4f6, []int{8}
}

func (m *ArticleOneRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleOneRes.Unmarshal(m, b)
}
func (m *ArticleOneRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleOneRes.Marshal(b, m, deterministic)
}
func (m *ArticleOneRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleOneRes.Merge(m, src)
}
func (m *ArticleOneRes) XXX_Size() int {
	return xxx_messageInfo_ArticleOneRes.Size(m)
}
func (m *ArticleOneRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleOneRes.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleOneRes proto.InternalMessageInfo

func (m *ArticleOneRes) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleOneRes) GetCreater() int64 {
	if m != nil {
		return m.Creater
	}
	return 0
}

func (m *ArticleOneRes) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *ArticleOneRes) GetUpdater() int64 {
	if m != nil {
		return m.Updater
	}
	return 0
}

func (m *ArticleOneRes) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *ArticleOneRes) GetIsDelete() uint32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *ArticleOneRes) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ArticleOneRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArticleOneRes) GetXUser() string {
	if m != nil {
		return m.XUser
	}
	return ""
}

type ArticlePageReq struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Creater              int64    `protobuf:"varint,2,opt,name=creater,proto3" json:"creater"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=create_time,proto3" json:"create_time"`
	Updater              int64    `protobuf:"varint,4,opt,name=updater,proto3" json:"updater"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,proto3" json:"update_time"`
	IsDelete             uint32   `protobuf:"varint,6,opt,name=is_delete,json=-,proto3" json:"is_delete"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name"`
	XUser                string   `protobuf:"bytes,9,opt,name=x_user,proto3" json:"x_user"`
	Page                 int64    `protobuf:"varint,10,opt,name=page,proto3" json:"page"`
	Size                 int64    `protobuf:"varint,11,opt,name=size,proto3" json:"size"`
	Order                string   `protobuf:"bytes,12,opt,name=order,proto3" json:"order"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticlePageReq) Reset()         { *m = ArticlePageReq{} }
func (m *ArticlePageReq) String() string { return proto.CompactTextString(m) }
func (*ArticlePageReq) ProtoMessage()    {}
func (*ArticlePageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d16ac54e3a89a4f6, []int{9}
}

func (m *ArticlePageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticlePageReq.Unmarshal(m, b)
}
func (m *ArticlePageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticlePageReq.Marshal(b, m, deterministic)
}
func (m *ArticlePageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticlePageReq.Merge(m, src)
}
func (m *ArticlePageReq) XXX_Size() int {
	return xxx_messageInfo_ArticlePageReq.Size(m)
}
func (m *ArticlePageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticlePageReq.DiscardUnknown(m)
}

var xxx_messageInfo_ArticlePageReq proto.InternalMessageInfo

func (m *ArticlePageReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticlePageReq) GetCreater() int64 {
	if m != nil {
		return m.Creater
	}
	return 0
}

func (m *ArticlePageReq) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *ArticlePageReq) GetUpdater() int64 {
	if m != nil {
		return m.Updater
	}
	return 0
}

func (m *ArticlePageReq) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *ArticlePageReq) GetIsDelete() uint32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *ArticlePageReq) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ArticlePageReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArticlePageReq) GetXUser() string {
	if m != nil {
		return m.XUser
	}
	return ""
}

func (m *ArticlePageReq) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ArticlePageReq) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ArticlePageReq) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

type ArticlePageRes struct {
	Page                 uint64     `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Size                 int64      `protobuf:"varint,2,opt,name=size,proto3" json:"size"`
	TotalRecord          int64      `protobuf:"varint,3,opt,name=total_record,proto3" json:"total_record"`
	TotalPage            int64      `protobuf:"varint,4,opt,name=total_page,proto3" json:"total_page"`
	Data                 []*Article `protobuf:"bytes,5,rep,name=data,proto3" json:"data"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ArticlePageRes) Reset()         { *m = ArticlePageRes{} }
func (m *ArticlePageRes) String() string { return proto.CompactTextString(m) }
func (*ArticlePageRes) ProtoMessage()    {}
func (*ArticlePageRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d16ac54e3a89a4f6, []int{10}
}

func (m *ArticlePageRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticlePageRes.Unmarshal(m, b)
}
func (m *ArticlePageRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticlePageRes.Marshal(b, m, deterministic)
}
func (m *ArticlePageRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticlePageRes.Merge(m, src)
}
func (m *ArticlePageRes) XXX_Size() int {
	return xxx_messageInfo_ArticlePageRes.Size(m)
}
func (m *ArticlePageRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticlePageRes.DiscardUnknown(m)
}

var xxx_messageInfo_ArticlePageRes proto.InternalMessageInfo

func (m *ArticlePageRes) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ArticlePageRes) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ArticlePageRes) GetTotalRecord() int64 {
	if m != nil {
		return m.TotalRecord
	}
	return 0
}

func (m *ArticlePageRes) GetTotalPage() int64 {
	if m != nil {
		return m.TotalPage
	}
	return 0
}

func (m *ArticlePageRes) GetData() []*Article {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Article)(nil), "helloworld.Article")
	proto.RegisterType((*ArticleAddReq)(nil), "helloworld.ArticleAddReq")
	proto.RegisterType((*ArticleAddRes)(nil), "helloworld.ArticleAddRes")
	proto.RegisterType((*ArticleDelReq)(nil), "helloworld.ArticleDelReq")
	proto.RegisterType((*ArticleDelRes)(nil), "helloworld.ArticleDelRes")
	proto.RegisterType((*ArticleUpdateReq)(nil), "helloworld.ArticleUpdateReq")
	proto.RegisterType((*ArticleUpdateRes)(nil), "helloworld.ArticleUpdateRes")
	proto.RegisterType((*ArticleOneReq)(nil), "helloworld.ArticleOneReq")
	proto.RegisterType((*ArticleOneRes)(nil), "helloworld.ArticleOneRes")
	proto.RegisterType((*ArticlePageReq)(nil), "helloworld.ArticlePageReq")
	proto.RegisterType((*ArticlePageRes)(nil), "helloworld.ArticlePageRes")
}

func init() { proto.RegisterFile("proto/helloworld/helloworld.proto", fileDescriptor_d16ac54e3a89a4f6) }

var fileDescriptor_d16ac54e3a89a4f6 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x56, 0x4d, 0x4e, 0xf3, 0x30,
	0x10, 0x95, 0x93, 0xb4, 0xfd, 0x3a, 0xfd, 0xd1, 0x27, 0x53, 0x21, 0xaf, 0x50, 0xc8, 0x86, 0x6c,
	0x28, 0x12, 0x9c, 0xa0, 0x12, 0x7b, 0x50, 0x24, 0xd6, 0x91, 0xa9, 0x47, 0x25, 0xc2, 0xad, 0x8b,
	0xe3, 0x0a, 0xc4, 0x7d, 0xd8, 0x73, 0x2f, 0x56, 0x48, 0x1c, 0x00, 0x65, 0x1c, 0x20, 0x45, 0xdc,
	0xc0, 0xbb, 0xf7, 0x5e, 0xe6, 0x4d, 0xfc, 0x66, 0x64, 0xc9, 0x70, 0xbc, 0xb5, 0xc6, 0x99, 0xb3,
	0x3b, 0xd4, 0xda, 0x3c, 0x1a, 0xab, 0x55, 0x07, 0xce, 0xe9, 0x1b, 0x87, 0x1f, 0x25, 0x7b, 0x63,
	0x30, 0x58, 0x58, 0x57, 0x2d, 0x35, 0xf2, 0x29, 0x44, 0x95, 0x12, 0x2c, 0x65, 0x79, 0x52, 0x44,
	0x95, 0xe2, 0x02, 0x06, 0x4b, 0x8b, 0xd2, 0xa1, 0x15, 0x51, 0xca, 0xf2, 0xb8, 0xf8, 0xa2, 0x3c,
	0x85, 0x91, 0x87, 0xa5, 0xab, 0xd6, 0x28, 0xe2, 0x94, 0xe5, 0xc3, 0xa2, 0x2b, 0x35, 0xde, 0xdd,
	0x56, 0x91, 0x37, 0xf1, 0xde, 0x96, 0x36, 0x5e, 0x0f, 0xbd, 0xb7, 0xe7, 0xbd, 0x1d, 0x89, 0xcf,
	0x60, 0x58, 0xd5, 0xa5, 0x42, 0x8d, 0x0e, 0x45, 0x3f, 0x65, 0xf9, 0xa4, 0x60, 0xa7, 0xfc, 0x10,
	0xfa, 0x16, 0xd7, 0xd2, 0xde, 0x8b, 0x01, 0x59, 0x5a, 0xc6, 0x39, 0x24, 0x1b, 0xb9, 0x46, 0xf1,
	0x8f, 0x54, 0xc2, 0x4d, 0xed, 0x53, 0xb9, 0xab, 0xd1, 0x8a, 0xa1, 0xaf, 0xf5, 0x2c, 0x7b, 0x67,
	0x30, 0x69, 0xd3, 0x2e, 0x94, 0x2a, 0xf0, 0x21, 0xc0, 0xcc, 0x75, 0x58, 0x99, 0x2f, 0x51, 0x07,
	0xb7, 0x67, 0xca, 0x1c, 0xc2, 0x9e, 0x3f, 0x18, 0xfc, 0x6f, 0x33, 0xdf, 0xd0, 0x0f, 0xc3, 0x58,
	0xf5, 0x1f, 0xb1, 0x03, 0xbb, 0xd5, 0x57, 0x1b, 0x0c, 0xee, 0x56, 0x53, 0xe6, 0x10, 0xf6, 0xfc,
	0x1a, 0xc1, 0xb4, 0xcd, 0x7c, 0x2d, 0x57, 0x61, 0x2c, 0xba, 0xa9, 0xdd, 0xca, 0x15, 0x0a, 0xa0,
	0x23, 0x11, 0x6e, 0xb4, 0xba, 0x7a, 0x46, 0x31, 0xf2, 0x5a, 0x83, 0xf9, 0x0c, 0x7a, 0xc6, 0x2a,
	0xb4, 0x62, 0x4c, 0x76, 0x4f, 0xb2, 0x17, 0xf6, 0x6b, 0x64, 0xf5, 0x77, 0x43, 0x3f, 0xb4, 0xfd,
	0x86, 0x51, 0xa7, 0x61, 0x06, 0x63, 0x67, 0x9c, 0xd4, 0xa5, 0xc5, 0xa5, 0xb1, 0x8a, 0x26, 0x16,
	0x17, 0x7b, 0x1a, 0x3f, 0x02, 0xf0, 0x9c, 0x3a, 0xfa, 0xa9, 0x75, 0x14, 0x7e, 0x02, 0x89, 0x92,
	0x4e, 0x8a, 0x5e, 0x1a, 0xe7, 0xa3, 0xf3, 0x83, 0x79, 0xe7, 0xc9, 0xd9, 0x9e, 0xaa, 0xa0, 0x82,
	0xdb, 0x3e, 0xbd, 0x40, 0x2f, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x50, 0xd0, 0xae, 0xa0, 0xa6,
	0x0a, 0x00, 0x00,
}
