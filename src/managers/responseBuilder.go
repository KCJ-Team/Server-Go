package manager

import (
	"server-go/src/pb" // 프로토버퍼 패키지
)

// ResponseBuilder: 빌더패턴
type ResponseBuilder struct {
	gameMessage *pb.GameMessage
}

// NewResponseBuilder: ResponseBuilder 생성자. Response들에 기본적으로 포함되어야할 코드와 타입, 메세지
func NewResponseBuilder(messageType pb.MessageType, code pb.ResponseCode, message string) *ResponseBuilder {
	return &ResponseBuilder{
		gameMessage: &pb.GameMessage{
			MessageType: messageType,
			Message: &pb.GameMessage_Response{
				Response: &pb.Response{
					Code:    code,
					Message: message,
				},
			},
		},
	}
}

// WithRoomInfo: RoomInfo 데이터를 추가. 기본빌더에 붙여서 생성
func (rb *ResponseBuilder) WithRoomInfo(room *Room) *ResponseBuilder {
	roomInfo := &pb.RoomInfo{
		RoomId: room.RoomId,
		Player1Info: &pb.PlayerInfo{
			PlayerId: room.Player1.playerId,
			X:        room.Player1.x,
			Y:        room.Player1.y,
			Z:        room.Player1.z,
			Rx:       room.Player1.rx,
			Ry:       room.Player1.ry,
			Rz:       room.Player1.rz,
			Speed:    room.Player1.speed,
			Health:   room.Player1.health,
		},
		Player2Info: &pb.PlayerInfo{
			PlayerId: room.Player2.playerId,
			X:        room.Player2.x,
			Y:        room.Player2.y,
			Z:        room.Player2.z,
			Rx:       room.Player2.rx,
			Ry:       room.Player2.ry,
			Rz:       room.Player2.rz,
			Speed:    room.Player2.speed,
			Health:   room.Player2.health,
		},
	}
	rb.gameMessage.Message.(*pb.GameMessage_Response).Response.Data = &pb.Response_RoomInfo{RoomInfo: roomInfo}
	return rb
}

// Build: 최종 GameMessage 반환
func (rb *ResponseBuilder) Build() *pb.GameMessage {
	return rb.gameMessage
}
