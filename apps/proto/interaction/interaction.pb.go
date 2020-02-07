// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interaction.proto

package interaction

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type InteractionStatusCode int32

const (
	InteractionStatusCode_Unknown InteractionStatusCode = 0
	InteractionStatusCode_Ok      InteractionStatusCode = 1
	InteractionStatusCode_Failed  InteractionStatusCode = 2
)

var InteractionStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var InteractionStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x InteractionStatusCode) String() string {
	return proto.EnumName(InteractionStatusCode_name, int32(x))
}

func (InteractionStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{0}
}

type IdType int32

const (
	IdType_IRDI    IdType = 0
	IdType_IRI     IdType = 1
	IdType_Custom  IdType = 2
	IdType_IdShort IdType = 3
)

var IdType_name = map[int32]string{
	0: "IRDI",
	1: "IRI",
	2: "Custom",
	3: "IdShort",
}

var IdType_value = map[string]int32{
	"IRDI":    0,
	"IRI":     1,
	"Custom":  2,
	"IdShort": 3,
}

func (x IdType) String() string {
	return proto.EnumName(IdType_name, int32(x))
}

func (IdType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{1}
}

type InteractionStatus struct {
	Code                 InteractionStatusCode `protobuf:"varint,1,opt,name=Code,proto3,enum=interaction.InteractionStatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *InteractionStatus) Reset()         { *m = InteractionStatus{} }
func (m *InteractionStatus) String() string { return proto.CompactTextString(m) }
func (*InteractionStatus) ProtoMessage()    {}
func (*InteractionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{0}
}

func (m *InteractionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InteractionStatus.Unmarshal(m, b)
}
func (m *InteractionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InteractionStatus.Marshal(b, m, deterministic)
}
func (m *InteractionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InteractionStatus.Merge(m, src)
}
func (m *InteractionStatus) XXX_Size() int {
	return xxx_messageInfo_InteractionStatus.Size(m)
}
func (m *InteractionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_InteractionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_InteractionStatus proto.InternalMessageInfo

func (m *InteractionStatus) GetCode() InteractionStatusCode {
	if m != nil {
		return m.Code
	}
	return InteractionStatusCode_Unknown
}

type InteractionMessage struct {
	Frame                *Frame   `protobuf:"bytes,1,opt,name=frame,proto3" json:"frame,omitempty"`
	InteractionElements  []byte   `protobuf:"bytes,2,opt,name=interactionElements,proto3" json:"interactionElements,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InteractionMessage) Reset()         { *m = InteractionMessage{} }
func (m *InteractionMessage) String() string { return proto.CompactTextString(m) }
func (*InteractionMessage) ProtoMessage()    {}
func (*InteractionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{1}
}

func (m *InteractionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InteractionMessage.Unmarshal(m, b)
}
func (m *InteractionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InteractionMessage.Marshal(b, m, deterministic)
}
func (m *InteractionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InteractionMessage.Merge(m, src)
}
func (m *InteractionMessage) XXX_Size() int {
	return xxx_messageInfo_InteractionMessage.Size(m)
}
func (m *InteractionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InteractionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InteractionMessage proto.InternalMessageInfo

func (m *InteractionMessage) GetFrame() *Frame {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *InteractionMessage) GetInteractionElements() []byte {
	if m != nil {
		return m.InteractionElements
	}
	return nil
}

type Frame struct {
	SemanticProtocol     string              `protobuf:"bytes,1,opt,name=semanticProtocol,proto3" json:"semanticProtocol,omitempty"`
	Type                 string              `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	MessageId            string              `protobuf:"bytes,3,opt,name=messageId,proto3" json:"messageId,omitempty"`
	ReplyBy              uint32              `protobuf:"varint,4,opt,name=replyBy,proto3" json:"replyBy,omitempty"`
	Receiver             *ConversationMember `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Sender               *ConversationMember `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	ConversationId       string              `protobuf:"bytes,7,opt,name=conversationId,proto3" json:"conversationId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{2}
}

func (m *Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Frame.Unmarshal(m, b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
}
func (m *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(m, src)
}
func (m *Frame) XXX_Size() int {
	return xxx_messageInfo_Frame.Size(m)
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

func (m *Frame) GetSemanticProtocol() string {
	if m != nil {
		return m.SemanticProtocol
	}
	return ""
}

func (m *Frame) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Frame) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *Frame) GetReplyBy() uint32 {
	if m != nil {
		return m.ReplyBy
	}
	return 0
}

func (m *Frame) GetReceiver() *ConversationMember {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *Frame) GetSender() *ConversationMember {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Frame) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

type ConversationMember struct {
	Identifier           *Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Role                 *Role       `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ConversationMember) Reset()         { *m = ConversationMember{} }
func (m *ConversationMember) String() string { return proto.CompactTextString(m) }
func (*ConversationMember) ProtoMessage()    {}
func (*ConversationMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{3}
}

func (m *ConversationMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversationMember.Unmarshal(m, b)
}
func (m *ConversationMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversationMember.Marshal(b, m, deterministic)
}
func (m *ConversationMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversationMember.Merge(m, src)
}
func (m *ConversationMember) XXX_Size() int {
	return xxx_messageInfo_ConversationMember.Size(m)
}
func (m *ConversationMember) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversationMember.DiscardUnknown(m)
}

var xxx_messageInfo_ConversationMember proto.InternalMessageInfo

func (m *ConversationMember) GetIdentifier() *Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *ConversationMember) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type Role struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{4}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Identifier struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IdType               IdType   `protobuf:"varint,2,opt,name=idType,proto3,enum=interaction.IdType" json:"idType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{5}
}

func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (m *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(m, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Identifier) GetIdType() IdType {
	if m != nil {
		return m.IdType
	}
	return IdType_IRDI
}

func init() {
	proto.RegisterEnum("interaction.InteractionStatusCode", InteractionStatusCode_name, InteractionStatusCode_value)
	proto.RegisterEnum("interaction.IdType", IdType_name, IdType_value)
	proto.RegisterType((*InteractionStatus)(nil), "interaction.InteractionStatus")
	proto.RegisterType((*InteractionMessage)(nil), "interaction.InteractionMessage")
	proto.RegisterType((*Frame)(nil), "interaction.Frame")
	proto.RegisterType((*ConversationMember)(nil), "interaction.ConversationMember")
	proto.RegisterType((*Role)(nil), "interaction.Role")
	proto.RegisterType((*Identifier)(nil), "interaction.Identifier")
}

func init() { proto.RegisterFile("interaction.proto", fileDescriptor_43fb9bc5eb13a2f0) }

var fileDescriptor_43fb9bc5eb13a2f0 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5f, 0x6b, 0xd3, 0x50,
	0x14, 0x5f, 0xd2, 0x34, 0x5d, 0x4f, 0xb5, 0xa4, 0x67, 0x88, 0x61, 0xc8, 0x56, 0x02, 0x4a, 0xa9,
	0x30, 0xa4, 0x8a, 0x13, 0x7c, 0xb3, 0x3a, 0x08, 0x22, 0xca, 0xad, 0x7b, 0xf4, 0x21, 0xcb, 0x3d,
	0xd3, 0xcb, 0x92, 0x7b, 0xcb, 0xcd, 0x5d, 0xa5, 0x1f, 0xc6, 0x2f, 0xe6, 0xa7, 0x91, 0xdc, 0xa4,
	0x6b, 0xb2, 0x6e, 0xe2, 0x83, 0x6f, 0x39, 0xe7, 0xf7, 0x3b, 0xbf, 0xf3, 0xf7, 0x06, 0x46, 0x42,
	0x1a, 0xd2, 0x49, 0x6a, 0x84, 0x92, 0x27, 0x4b, 0xad, 0x8c, 0xc2, 0x41, 0xc3, 0x15, 0x7d, 0x84,
	0x51, 0xbc, 0x35, 0x17, 0x26, 0x31, 0xd7, 0x05, 0xbe, 0x06, 0x6f, 0xae, 0x38, 0x85, 0xce, 0xd8,
	0x99, 0x0c, 0x67, 0xd1, 0x49, 0x53, 0x63, 0x87, 0x5d, 0x32, 0x99, 0xe5, 0x47, 0x4b, 0xc0, 0x06,
	0xfc, 0x89, 0x8a, 0x22, 0xf9, 0x4e, 0x38, 0x81, 0xee, 0xa5, 0x4e, 0xf2, 0x4a, 0x6e, 0x30, 0xc3,
	0x96, 0xdc, 0x59, 0x89, 0xb0, 0x8a, 0x80, 0x2f, 0xe0, 0xa0, 0x81, 0x7d, 0xc8, 0x28, 0x27, 0x69,
	0x8a, 0xd0, 0x1d, 0x3b, 0x93, 0x07, 0xec, 0x2e, 0x28, 0xfa, 0xe5, 0x42, 0xd7, 0x4a, 0xe0, 0x14,
	0x82, 0x82, 0xf2, 0x44, 0x1a, 0x91, 0x7e, 0x29, 0xdb, 0x4c, 0x55, 0x66, 0x13, 0xf6, 0xd9, 0x8e,
	0x1f, 0x11, 0x3c, 0xb3, 0x5e, 0x92, 0x15, 0xee, 0x33, 0xfb, 0x8d, 0x4f, 0xa0, 0x9f, 0x57, 0x05,
	0xc7, 0x3c, 0xec, 0x58, 0x60, 0xeb, 0xc0, 0x10, 0x7a, 0x9a, 0x96, 0xd9, 0xfa, 0xdd, 0x3a, 0xf4,
	0xc6, 0xce, 0xe4, 0x21, 0xdb, 0x98, 0xf8, 0x16, 0xf6, 0x35, 0xa5, 0x24, 0x56, 0xa4, 0xc3, 0xae,
	0x6d, 0xf0, 0xb8, 0xd5, 0xe0, 0x5c, 0xc9, 0x15, 0xe9, 0x22, 0xa9, 0x26, 0x92, 0x5f, 0x90, 0x66,
	0x37, 0x01, 0x78, 0x0a, 0x7e, 0x41, 0x92, 0x93, 0x0e, 0xfd, 0x7f, 0x0b, 0xad, 0xe9, 0xf8, 0x0c,
	0x86, 0x69, 0x03, 0x8d, 0x79, 0xd8, 0xb3, 0x25, 0xdf, 0xf2, 0x46, 0x06, 0x70, 0x57, 0x05, 0x4f,
	0x01, 0x04, 0x27, 0x69, 0xc4, 0xa5, 0x20, 0x5d, 0xaf, 0xe5, 0x71, 0x7b, 0xcb, 0x37, 0x30, 0x6b,
	0x50, 0xf1, 0x29, 0x78, 0x5a, 0x65, 0xd5, 0xe0, 0x06, 0xb3, 0x51, 0x2b, 0x84, 0xa9, 0x8c, 0x98,
	0x85, 0xa3, 0x43, 0xf0, 0x4a, 0xab, 0x9c, 0xb3, 0xdc, 0x2c, 0xbe, 0xcf, 0xec, 0x77, 0x14, 0x03,
	0x6c, 0xc5, 0x71, 0x08, 0xae, 0xe0, 0x35, 0xee, 0x0a, 0x8e, 0xcf, 0xc1, 0x17, 0xfc, 0xeb, 0x66,
	0x37, 0xc3, 0xd9, 0xc1, 0xad, 0xaa, 0x4a, 0x88, 0xd5, 0x94, 0xe9, 0x1b, 0x78, 0x74, 0xe7, 0x35,
	0xe2, 0x00, 0x7a, 0xe7, 0xf2, 0x4a, 0xaa, 0x9f, 0x32, 0xd8, 0x43, 0x1f, 0xdc, 0xcf, 0x57, 0x81,
	0x83, 0x00, 0xfe, 0x59, 0x22, 0x32, 0xe2, 0x81, 0x3b, 0x7d, 0x05, 0x7e, 0xa5, 0x85, 0xfb, 0xe0,
	0xc5, 0xec, 0x7d, 0x1c, 0xec, 0x61, 0x0f, 0x3a, 0x31, 0x8b, 0x2b, 0xe2, 0xfc, 0xba, 0x30, 0x2a,
	0x0f, 0xdc, 0x52, 0x29, 0xe6, 0x8b, 0x1f, 0x4a, 0x9b, 0xa0, 0x33, 0xfb, 0xed, 0xb4, 0xee, 0x7b,
	0x41, 0x7a, 0x25, 0x52, 0x42, 0x82, 0xa3, 0xf3, 0x65, 0xa6, 0x12, 0xbe, 0x7b, 0xfb, 0x0b, 0xa3,
	0x29, 0xc9, 0xf1, 0xf8, 0xbe, 0x17, 0x54, 0xd3, 0x0e, 0x8f, 0xfe, 0xfe, 0xc4, 0xa2, 0xbd, 0x89,
	0x83, 0xdf, 0x20, 0xbc, 0x2f, 0xcd, 0x7f, 0x48, 0x70, 0xe1, 0xdb, 0x9f, 0xc3, 0xcb, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xac, 0x8a, 0x6d, 0x0a, 0x31, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InteractionServiceClient is the client API for InteractionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InteractionServiceClient interface {
	UploadInteractionMessageStream(ctx context.Context, opts ...grpc.CallOption) (InteractionService_UploadInteractionMessageStreamClient, error)
	UploadInteractionMessage(ctx context.Context, in *InteractionMessage, opts ...grpc.CallOption) (*InteractionStatus, error)
}

type interactionServiceClient struct {
	cc *grpc.ClientConn
}

func NewInteractionServiceClient(cc *grpc.ClientConn) InteractionServiceClient {
	return &interactionServiceClient{cc}
}

func (c *interactionServiceClient) UploadInteractionMessageStream(ctx context.Context, opts ...grpc.CallOption) (InteractionService_UploadInteractionMessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_InteractionService_serviceDesc.Streams[0], "/interaction.InteractionService/UploadInteractionMessageStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &interactionServiceUploadInteractionMessageStreamClient{stream}
	return x, nil
}

type InteractionService_UploadInteractionMessageStreamClient interface {
	Send(*InteractionMessage) error
	CloseAndRecv() (*InteractionStatus, error)
	grpc.ClientStream
}

type interactionServiceUploadInteractionMessageStreamClient struct {
	grpc.ClientStream
}

func (x *interactionServiceUploadInteractionMessageStreamClient) Send(m *InteractionMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *interactionServiceUploadInteractionMessageStreamClient) CloseAndRecv() (*InteractionStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(InteractionStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *interactionServiceClient) UploadInteractionMessage(ctx context.Context, in *InteractionMessage, opts ...grpc.CallOption) (*InteractionStatus, error) {
	out := new(InteractionStatus)
	err := c.cc.Invoke(ctx, "/interaction.InteractionService/UploadInteractionMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionServiceServer is the server API for InteractionService service.
type InteractionServiceServer interface {
	UploadInteractionMessageStream(InteractionService_UploadInteractionMessageStreamServer) error
	UploadInteractionMessage(context.Context, *InteractionMessage) (*InteractionStatus, error)
}

// UnimplementedInteractionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInteractionServiceServer struct {
}

func (*UnimplementedInteractionServiceServer) UploadInteractionMessageStream(srv InteractionService_UploadInteractionMessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadInteractionMessageStream not implemented")
}
func (*UnimplementedInteractionServiceServer) UploadInteractionMessage(ctx context.Context, req *InteractionMessage) (*InteractionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadInteractionMessage not implemented")
}

func RegisterInteractionServiceServer(s *grpc.Server, srv InteractionServiceServer) {
	s.RegisterService(&_InteractionService_serviceDesc, srv)
}

func _InteractionService_UploadInteractionMessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InteractionServiceServer).UploadInteractionMessageStream(&interactionServiceUploadInteractionMessageStreamServer{stream})
}

type InteractionService_UploadInteractionMessageStreamServer interface {
	SendAndClose(*InteractionStatus) error
	Recv() (*InteractionMessage, error)
	grpc.ServerStream
}

type interactionServiceUploadInteractionMessageStreamServer struct {
	grpc.ServerStream
}

func (x *interactionServiceUploadInteractionMessageStreamServer) SendAndClose(m *InteractionStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *interactionServiceUploadInteractionMessageStreamServer) Recv() (*InteractionMessage, error) {
	m := new(InteractionMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _InteractionService_UploadInteractionMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InteractionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).UploadInteractionMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interaction.InteractionService/UploadInteractionMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).UploadInteractionMessage(ctx, req.(*InteractionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _InteractionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interaction.InteractionService",
	HandlerType: (*InteractionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadInteractionMessage",
			Handler:    _InteractionService_UploadInteractionMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadInteractionMessageStream",
			Handler:       _InteractionService_UploadInteractionMessageStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "interaction.proto",
}
