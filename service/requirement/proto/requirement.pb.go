// Code generated by protoc-gen-go. DO NOT EDIT.
// source: requirement.proto

/*
Package requirement is a generated protocol buffer package.

It is generated from these files:
	requirement.proto

It has these top-level messages:
	PublishRequest
	PublishResponse
	QueryRequest
	QueryResponse
	QueryData
	RequirementData
*/
package requirement

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PublishRequest struct {
	Version     uint32 `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint32 `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint32 `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64 `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       string `protobuf:"bytes,8,opt,name=param" json:"param"`
	SigAlg      uint32 `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
	Signature   string `protobuf:"bytes,10,opt,name=signature" json:"signature"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (m *PublishRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PublishRequest) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PublishRequest) GetCursorNum() uint32 {
	if m != nil {
		return m.CursorNum
	}
	return 0
}

func (m *PublishRequest) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *PublishRequest) GetLifetime() uint64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *PublishRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *PublishRequest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *PublishRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *PublishRequest) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func (m *PublishRequest) GetSigAlg() uint32 {
	if m != nil {
		return m.SigAlg
	}
	return 0
}

func (m *PublishRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type PublishResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *PublishResponse) Reset()                    { *m = PublishResponse{} }
func (m *PublishResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()               {}
func (*PublishResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PublishResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PublishResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *PublishResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryRequest struct {
	PageSize  uint32 `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size"`
	PageNum   uint32 `protobuf:"varint,2,opt,name=page_num,json=pageNum" json:"page_num"`
	ReqId     string `protobuf:"bytes,3,opt,name=req_id,json=reqId" json:"req_id"`
	ReqType   uint64 `protobuf:"varint,4,opt,name=req_type,json=reqType" json:"req_type"`
	Username  string `protobuf:"bytes,5,opt,name=username" json:"username"`
	Random    string `protobuf:"bytes,6,opt,name=random" json:"random"`
	Signature string `protobuf:"bytes,7,opt,name=signature" json:"signature"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueryRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *QueryRequest) GetPageNum() uint32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryRequest) GetReqId() string {
	if m != nil {
		return m.ReqId
	}
	return ""
}

func (m *QueryRequest) GetReqType() uint64 {
	if m != nil {
		return m.ReqType
	}
	return 0
}

func (m *QueryRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryRequest) GetRandom() string {
	if m != nil {
		return m.Random
	}
	return ""
}

func (m *QueryRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type QueryResponse struct {
	Code int32      `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *QueryData `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string     `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueryResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *QueryResponse) GetData() *QueryData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *QueryResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryData struct {
	PageNum  uint32             `protobuf:"varint,1,opt,name=page_num,json=pageNum" json:"page_num"`
	RowCount uint32             `protobuf:"varint,2,opt,name=row_count,json=rowCount" json:"row_count"`
	Row      []*RequirementData `protobuf:"bytes,3,rep,name=row" json:"row"`
}

func (m *QueryData) Reset()                    { *m = QueryData{} }
func (m *QueryData) String() string            { return proto.CompactTextString(m) }
func (*QueryData) ProtoMessage()               {}
func (*QueryData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryData) GetPageNum() uint32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryData) GetRowCount() uint32 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func (m *QueryData) GetRow() []*RequirementData {
	if m != nil {
		return m.Row
	}
	return nil
}

type RequirementData struct {
	Username        string `protobuf:"bytes,1,opt,name=username" json:"username"`
	RequirementId   string `protobuf:"bytes,2,opt,name=requirement_id,json=requirementId" json:"requirement_id"`
	RequirementName string `protobuf:"bytes,3,opt,name=requirement_name,json=requirementName" json:"requirement_name"`
	ReqType         uint64 `protobuf:"varint,4,opt,name=req_type,json=reqType" json:"req_type"`
	FeatureTag      uint64 `protobuf:"varint,5,opt,name=feature_tag,json=featureTag" json:"feature_tag"`
	SamplePath      string `protobuf:"bytes,6,opt,name=sample_path,json=samplePath" json:"sample_path"`
	SampleHash      string `protobuf:"bytes,7,opt,name=sample_hash,json=sampleHash" json:"sample_hash"`
	ExpireTime      uint64 `protobuf:"varint,8,opt,name=expire_time,json=expireTime" json:"expire_time"`
	Price           uint64 `protobuf:"varint,9,opt,name=price" json:"price"`
	Description     string `protobuf:"bytes,10,opt,name=description" json:"description"`
	PublishDate     uint64 `protobuf:"varint,11,opt,name=publish_date,json=publishDate" json:"publish_date"`
}

func (m *RequirementData) Reset()                    { *m = RequirementData{} }
func (m *RequirementData) String() string            { return proto.CompactTextString(m) }
func (*RequirementData) ProtoMessage()               {}
func (*RequirementData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RequirementData) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RequirementData) GetRequirementId() string {
	if m != nil {
		return m.RequirementId
	}
	return ""
}

func (m *RequirementData) GetRequirementName() string {
	if m != nil {
		return m.RequirementName
	}
	return ""
}

func (m *RequirementData) GetReqType() uint64 {
	if m != nil {
		return m.ReqType
	}
	return 0
}

func (m *RequirementData) GetFeatureTag() uint64 {
	if m != nil {
		return m.FeatureTag
	}
	return 0
}

func (m *RequirementData) GetSamplePath() string {
	if m != nil {
		return m.SamplePath
	}
	return ""
}

func (m *RequirementData) GetSampleHash() string {
	if m != nil {
		return m.SampleHash
	}
	return ""
}

func (m *RequirementData) GetExpireTime() uint64 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func (m *RequirementData) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *RequirementData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RequirementData) GetPublishDate() uint64 {
	if m != nil {
		return m.PublishDate
	}
	return 0
}

func init() {
	proto.RegisterType((*PublishRequest)(nil), "PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "PublishResponse")
	proto.RegisterType((*QueryRequest)(nil), "QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "QueryResponse")
	proto.RegisterType((*QueryData)(nil), "QueryData")
	proto.RegisterType((*RequirementData)(nil), "RequirementData")
}

func init() { proto.RegisterFile("requirement.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x4b, 0x6e, 0xdb, 0x3a,
	0x14, 0x86, 0xa3, 0xf8, 0x21, 0xeb, 0x28, 0x8e, 0x7d, 0x89, 0xfb, 0xd0, 0x4d, 0xfa, 0x70, 0x05,
	0x14, 0x70, 0x27, 0x1a, 0xa4, 0x2b, 0x28, 0x9a, 0x41, 0x83, 0x16, 0x41, 0xaa, 0xa6, 0x63, 0x81,
	0x91, 0x4e, 0x64, 0x02, 0x92, 0x28, 0x93, 0x54, 0xd3, 0x64, 0x37, 0xdd, 0x4e, 0x77, 0xd4, 0x59,
	0xc1, 0x87, 0x1d, 0xd9, 0x68, 0x3b, 0xe3, 0xff, 0xf3, 0x1c, 0x52, 0xfa, 0xc8, 0x9f, 0xf0, 0x97,
	0xc0, 0x75, 0xc7, 0x04, 0xd6, 0xd8, 0xa8, 0xa4, 0x15, 0x5c, 0xf1, 0xf8, 0xdb, 0x21, 0x1c, 0x5f,
	0x75, 0x37, 0x15, 0x93, 0xab, 0x14, 0xd7, 0x1d, 0x4a, 0x45, 0x22, 0xf0, 0xbf, 0xa0, 0x90, 0x8c,
	0x37, 0x91, 0xb7, 0xf0, 0x96, 0xd3, 0x74, 0x23, 0xc9, 0x53, 0x80, 0xbc, 0x13, 0x92, 0x8b, 0xac,
	0xe9, 0xea, 0xe8, 0xd0, 0x4c, 0x06, 0xd6, 0xb9, 0xec, 0x6a, 0xf2, 0x02, 0x8e, 0xdc, 0x74, 0x45,
	0x6f, 0xb0, 0x8a, 0x06, 0xa6, 0x20, 0xb4, 0xde, 0x07, 0x6d, 0x91, 0x13, 0x98, 0x54, 0xec, 0x16,
	0x15, 0xab, 0x31, 0x1a, 0x2e, 0xbc, 0xe5, 0x30, 0xdd, 0x6a, 0xf2, 0x2f, 0x8c, 0x25, 0x36, 0x05,
	0x8a, 0x68, 0xb4, 0xf0, 0x96, 0x41, 0xea, 0x94, 0xee, 0xc9, 0x79, 0xa3, 0x04, 0xcd, 0x55, 0x34,
	0x36, 0x33, 0x5b, 0xad, 0x7b, 0x6a, 0x54, 0x2b, 0x5e, 0x44, 0xbe, 0xed, 0xb1, 0x8a, 0xfc, 0x0d,
	0xa3, 0x96, 0x0a, 0x5a, 0x47, 0x13, 0x63, 0x5b, 0x41, 0xfe, 0x03, 0x5f, 0xb2, 0x32, 0xa3, 0x55,
	0x19, 0x05, 0xe6, 0xdb, 0xc6, 0x92, 0x95, 0x6f, 0xaa, 0x92, 0x3c, 0x81, 0x40, 0xb2, 0xb2, 0xa1,
	0xaa, 0x13, 0x18, 0x81, 0x69, 0x79, 0x34, 0xe2, 0xf7, 0x30, 0xdb, 0x22, 0x92, 0x2d, 0x6f, 0x24,
	0x12, 0x02, 0xc3, 0x9c, 0x17, 0xe8, 0x00, 0x99, 0xb1, 0xf6, 0x0a, 0xaa, 0xa8, 0xe1, 0x12, 0xa4,
	0x66, 0x4c, 0xe6, 0x30, 0xa8, 0x65, 0x69, 0x48, 0x04, 0xa9, 0x1e, 0xc6, 0xdf, 0x3d, 0x38, 0xfa,
	0xd8, 0xa1, 0xb8, 0xdf, 0xe0, 0x3e, 0x85, 0xa0, 0xa5, 0x25, 0x66, 0x92, 0x3d, 0x6c, 0xd6, 0x9b,
	0x68, 0xe3, 0x13, 0x7b, 0x40, 0xf2, 0x3f, 0x98, 0x71, 0x8f, 0xb7, 0xaf, 0xb5, 0xa6, 0xfd, 0x0f,
	0x8c, 0x05, 0xae, 0x33, 0x56, 0xb8, 0xd5, 0x47, 0x02, 0xd7, 0x17, 0x85, 0xee, 0xd0, 0xb6, 0xba,
	0x6f, 0x37, 0x84, 0x7d, 0x81, 0xeb, 0xeb, 0xfb, 0x16, 0x35, 0xc8, 0x4e, 0xa2, 0x68, 0x68, 0x8d,
	0x0e, 0xf1, 0x56, 0x6b, 0x90, 0x82, 0x36, 0x05, 0xaf, 0x1d, 0x62, 0xa7, 0x76, 0xc9, 0xf8, 0xfb,
	0x64, 0x3e, 0xc3, 0xd4, 0xfd, 0xcb, 0x2f, 0xb8, 0x8c, 0x1c, 0x97, 0x67, 0x3d, 0x2e, 0xe1, 0x19,
	0x24, 0xa6, 0xe3, 0x9c, 0x2a, 0xfa, 0x5b, 0x46, 0x25, 0x04, 0xdb, 0xa2, 0x1d, 0x04, 0xde, 0x2e,
	0x82, 0x53, 0x08, 0x04, 0xbf, 0xcb, 0x72, 0xde, 0x35, 0xca, 0xe1, 0x99, 0x08, 0x7e, 0xf7, 0x56,
	0x6b, 0x12, 0xc3, 0x40, 0xf0, 0xbb, 0x68, 0xb0, 0x18, 0x2c, 0xc3, 0xb3, 0x79, 0x92, 0x3e, 0x5e,
	0x7d, 0xb3, 0xb7, 0x9e, 0x8c, 0x7f, 0x1c, 0xc2, 0x6c, 0x6f, 0x62, 0x87, 0x92, 0xb7, 0x47, 0xe9,
	0x25, 0x1c, 0xf7, 0x22, 0xa4, 0xd9, 0xdb, 0xc3, 0x9e, 0xf6, 0xdc, 0x8b, 0x82, 0xbc, 0x82, 0x79,
	0xbf, 0xcc, 0x2c, 0x65, 0x7f, 0x6f, 0xd6, 0xf3, 0x2f, 0xf5, 0x8a, 0x7f, 0x38, 0xae, 0xe7, 0x10,
	0xde, 0xa2, 0xe1, 0x9c, 0x29, 0x5a, 0x9a, 0x13, 0x1b, 0xa6, 0xe0, 0xac, 0x6b, 0x5a, 0xea, 0x02,
	0x49, 0xeb, 0xb6, 0xc2, 0xac, 0xa5, 0x6a, 0xe5, 0x0e, 0x0e, 0xac, 0x75, 0x45, 0xd5, 0xaa, 0x57,
	0xb0, 0xa2, 0x72, 0xe5, 0x8e, 0xcf, 0x15, 0xbc, 0xa3, 0xd2, 0x14, 0xe0, 0xd7, 0x96, 0xe9, 0x1d,
	0x74, 0x22, 0x27, 0x76, 0x0b, 0x6b, 0x5d, 0xeb, 0x4c, 0xea, 0x1c, 0x09, 0x96, 0xa3, 0xc9, 0xcb,
	0x30, 0xb5, 0x82, 0x2c, 0x20, 0x2c, 0x50, 0xe6, 0x82, 0xb5, 0x4a, 0xbf, 0x12, 0x36, 0x30, 0x7d,
	0x4b, 0x3f, 0x05, 0xad, 0x8d, 0x4c, 0x56, 0x50, 0x85, 0x51, 0x68, 0xda, 0x43, 0xe7, 0x9d, 0x53,
	0x85, 0x67, 0x25, 0x84, 0x3d, 0xf4, 0x24, 0x01, 0xdf, 0x85, 0x8c, 0xcc, 0x92, 0xdd, 0x17, 0xe9,
	0x64, 0x9e, 0xec, 0xe5, 0x2f, 0x3e, 0x20, 0x4b, 0x18, 0x99, 0x3b, 0x42, 0xa6, 0x49, 0x3f, 0x4e,
	0x27, 0xc7, 0xc9, 0xce, 0x8d, 0x8c, 0x0f, 0x6e, 0xc6, 0xe6, 0xa5, 0x7b, 0xfd, 0x33, 0x00, 0x00,
	0xff, 0xff, 0xe3, 0xd7, 0x10, 0x89, 0xfe, 0x04, 0x00, 0x00,
}
