// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: messages.proto

package pb

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

// 메시지 타입을 정의
type MessageType int32

const (
	MessageType_SESSION_LOGIN          MessageType = 0 // 세션 로그인 요청
	MessageType_SESSION_LOGOUT         MessageType = 1 // 세션 로그아웃 요청
	MessageType_PLAYER_QUERY           MessageType = 2 // 플레이어 정보 조회 요청
	MessageType_PLAYER_POSITION_UPDATE MessageType = 3 // 플레이어 위치 업데이트 요청
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "SESSION_LOGIN",
		1: "SESSION_LOGOUT",
		2: "PLAYER_QUERY",
		3: "PLAYER_POSITION_UPDATE",
	}
	MessageType_value = map[string]int32{
		"SESSION_LOGIN":          0,
		"SESSION_LOGOUT":         1,
		"PLAYER_QUERY":           2,
		"PLAYER_POSITION_UPDATE": 3,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

type GameMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType MessageType `protobuf:"varint,1,opt,name=messageType,proto3,enum=messages.MessageType" json:"messageType,omitempty"`
	// 다양한 메시지 유형을 포함하는 oneof 필드
	//
	// Types that are assignable to Message:
	//
	//	*GameMessage_PlayerRequest
	//	*GameMessage_Response
	//	*GameMessage_PlayerPosition
	Message isGameMessage_Message `protobuf_oneof:"message"`
}

func (x *GameMessage) Reset() {
	*x = GameMessage{}
	mi := &file_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameMessage) ProtoMessage() {}

func (x *GameMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameMessage.ProtoReflect.Descriptor instead.
func (*GameMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GameMessage) GetMessageType() MessageType {
	if x != nil {
		return x.MessageType
	}
	return MessageType_SESSION_LOGIN
}

func (m *GameMessage) GetMessage() isGameMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *GameMessage) GetPlayerRequest() *PlayerRequest {
	if x, ok := x.GetMessage().(*GameMessage_PlayerRequest); ok {
		return x.PlayerRequest
	}
	return nil
}

func (x *GameMessage) GetResponse() *Response {
	if x, ok := x.GetMessage().(*GameMessage_Response); ok {
		return x.Response
	}
	return nil
}

func (x *GameMessage) GetPlayerPosition() *PlayerPosition {
	if x, ok := x.GetMessage().(*GameMessage_PlayerPosition); ok {
		return x.PlayerPosition
	}
	return nil
}

type isGameMessage_Message interface {
	isGameMessage_Message()
}

type GameMessage_PlayerRequest struct {
	PlayerRequest *PlayerRequest `protobuf:"bytes,2,opt,name=playerRequest,proto3,oneof"`
}

type GameMessage_Response struct {
	Response *Response `protobuf:"bytes,3,opt,name=response,proto3,oneof"`
}

type GameMessage_PlayerPosition struct {
	PlayerPosition *PlayerPosition `protobuf:"bytes,4,opt,name=playerPosition,proto3,oneof"`
}

func (*GameMessage_PlayerRequest) isGameMessage_Message() {}

func (*GameMessage_Response) isGameMessage_Message() {}

func (*GameMessage_PlayerPosition) isGameMessage_Message() {}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x10, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x61,
	0x74, 0x68, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x02, 0x0a,
	0x0b, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x62, 0x0a, 0x0b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x47, 0x4f, 0x55, 0x54, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x51, 0x55, 0x45, 0x52,
	0x59, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x50, 0x4f,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x42,
	0x12, 0x5a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_messages_proto_goTypes = []any{
	(MessageType)(0),       // 0: messages.MessageType
	(*GameMessage)(nil),    // 1: messages.GameMessage
	(*PlayerRequest)(nil),  // 2: playerinfo.PlayerRequest
	(*Response)(nil),       // 3: messages.Response
	(*PlayerPosition)(nil), // 4: playerinfo.PlayerPosition
}
var file_messages_proto_depIdxs = []int32{
	0, // 0: messages.GameMessage.messageType:type_name -> messages.MessageType
	2, // 1: messages.GameMessage.playerRequest:type_name -> playerinfo.PlayerRequest
	3, // 2: messages.GameMessage.response:type_name -> messages.Response
	4, // 3: messages.GameMessage.playerPosition:type_name -> playerinfo.PlayerPosition
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	file_playerinfo_proto_init()
	file_pathinfo_proto_init()
	file_response_proto_init()
	file_messages_proto_msgTypes[0].OneofWrappers = []any{
		(*GameMessage_PlayerRequest)(nil),
		(*GameMessage_Response)(nil),
		(*GameMessage_PlayerPosition)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		EnumInfos:         file_messages_proto_enumTypes,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
