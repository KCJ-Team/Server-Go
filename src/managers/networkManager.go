package manager

import (
	"encoding/binary"
	"log"
	"net"

	pb "server-go/src/pb" // 프로토버퍼 메시지 패키지

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

// ResponseCode: 응답 코드 정의
type ResponseCode int32

const (
	Success ResponseCode = iota // 0 성공
	Error                       // 1 일반 오류
	// 상세 오류
	BadRequest // 2 잘못된 요청
	NotFound   // 3 서버에서 찾을 수 없음
	Conflict   // 4 충돌 (플레이어가 이미 로그인 상태)
)

// BasicResponse: 모든 응답 메시지에 기본적으로 포함될 구조체 정의
type BasicResponse struct {
	Code    ResponseCode // 응답 코드
	Message string       // 응답 메시지
}

type NetworkManager struct{}

var networkManager *NetworkManager

func GetNetworkManager() *NetworkManager {
	if networkManager == nil {
		networkManager = &NetworkManager{}
	}

	return networkManager
}

// SendResponseMessage: 응답 메시지를 생성하고 클라이언트에게 전송하는 함수
func (nm *NetworkManager) SendResponseMessage(messageType pb.MessageType, conn *net.Conn, code ResponseCode, message string) error {
	// 기본 응답 메시지 생성
	builder := NewResponseBuilder(messageType, pb.ResponseCode(code), message)
	response := builder.Build()

	// 직렬화된 텍스트 메시지 확인용 로그 추가
	log.Printf("Generated Response: %s", prototext.Format(response))

	// 메시지 전송
	return nm.SendPacketToClient(conn, response)
}

// // CreateResponseMessage: BasicResponse를 GameMessage 형태로 변환하는 함수
// func (nm *NetworkManager) CreateResponseMessage(messageType pb.MessageType, code ResponseCode, message string) *pb.GameMessage {
// 	return &pb.GameMessage{
// 		MessageType: messageType, // 메세지 타입
// 		Message: &pb.GameMessage_Response{
// 			Response: &pb.Response{
// 				Code:    pb.ResponseCode(code),
// 				Message: message,
// 			},
// 		},
// 	}
// }

// SendPacketToClient: 클라이언트에게 메시지를 전송하는 함수
func (nm *NetworkManager) SendPacketToClient(conn *net.Conn, message proto.Message) error {
	packet, err := MakePacket(message)
	if err != nil {
		return err
	}

	// 전송 실행
	if _, err := (*conn).Write(packet); err != nil {
		log.Printf("Failed to send packet to client (%s): %v", (*conn).RemoteAddr(), err)
		return err
	}

	//log.Printf("Packet successfully sent to client at %s", (*conn).RemoteAddr())
	//log.Printf("Packet length sent: %d, Packet data: %x", len(packet), packet)
	return nil
}

// 룸에 속한(멀티 게임중인) 플레이어들에게 패킷을 전송하는 메소드(클라가 X)
func (nm *NetworkManager) SendPacketToRoomPlayers(room *Room, message proto.Message) error {
	packet, err := MakePacket(message)

	if err != nil {
		return err
	}

	for _, player := range room.Players {
		if player != nil && player.conn != nil {
			if _, err := (*player.conn).Write(packet); err != nil {
				log.Printf("Failed to send packet to player %s (%s): %v", player.playerId, (*player.conn).RemoteAddr(), err)
			} else {
				log.Printf("Packet sent to player %s at %s", player.playerId, (*player.conn).RemoteAddr())
			}
		}
	}

	log.Printf("SendPacketToRoomPlayers 성공 Packet length sent: %d, Packet data: %x", len(packet), packet)

	return nil
}

// MakePacket: 길이 정보를 포함한 패킷 생성 및 직렬화된 패킷 반환.
func MakePacket(message proto.Message) ([]byte, error) {
	// 메시지 직렬화
	messageData, err := proto.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err)
		return nil, err
	}

	// 패킷에 길이 정보 포함
	lengthBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(lengthBuf, uint32(len(messageData)))
	lengthBuf = append(lengthBuf, messageData...)

	//log.Printf("Packet successfully created with length: %d", len(lengthBuf))

	return lengthBuf, nil
}

// Custom Response들..

// SendRoomInfoResponse: RoomInfo 메시지를 생성하고 클라이언트에게 전송
func (nm *NetworkManager) SendRoomInfoResponse(messageType pb.MessageType, code ResponseCode, conn *net.Conn, message string, room *Room) error {
	// RoomInfo 포함 응답 메시지 생성
	builder := NewResponseBuilder(messageType, pb.ResponseCode(code), message).WithRoomInfo(room)
	response := builder.Build()

	// 직렬화된 텍스트 메시지 확인용 로그 추가
	log.Printf("Generated RoomInfo Response: %s", prototext.Format(response))

	// 메시지 전송
	return nm.SendPacketToClient(conn, response)
}

// SendRoomPlayerUpdateResponse: RoomPlayerUpdate 메시지를 생성하고 클라이언트에게 전송
func (nm *NetworkManager) SendRoomPlayerUpdateResponse(messageType pb.MessageType, code ResponseCode, conn *net.Conn, message string, roomId string, player *Player) error {
	// RoomPlayerUpdate 포함 응답 메시지 생성
	builder := NewResponseBuilder(messageType, pb.ResponseCode(code), message).WithRoomPlayerUpdate(roomId, player)
	response := builder.Build()

	// 직렬화된 텍스트 메시지 확인용 로그 추가
	log.Printf("Generated RoomPlayerUpdate Response: %s", prototext.Format(response))

	// 메시지 전송
	return nm.SendPacketToClient(conn, response)
}
