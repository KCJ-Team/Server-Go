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

// WithRoomInfo: RoomInfo 데이터를 추가. 방에 모든 플레이어 정보. 기본빌더에 붙여서 생성
func (rb *ResponseBuilder) WithRoomInfo(room *Room) *ResponseBuilder {
	// RoomInfo 객체 생성 및 players 리스트 초기화
	roomInfo := &pb.RoomInfo{
		RoomId:  room.RoomId,
		Players: make([]*pb.PlayerInfo, 0, len(room.Players)), // 플레이어 리스트 초기화
	}

	// Room의 Players 배열을 순회하며 PlayerInfo 객체를 생성하여 추가
	for _, player := range room.Players {
		if player != nil {
			playerInfo := &pb.PlayerInfo{
				PlayerId: player.playerId,
				X:        player.x,
				Y:        player.y,
				Z:        player.z,
				Rx:       player.rx,
				Ry:       player.ry,
				Rz:       player.rz,
				Speed:    player.speed,
				Health:   player.health,
				// PrefabType: player.prefabType, // prefabType이 있다면 추가
			}
			roomInfo.Players = append(roomInfo.Players, playerInfo)
		}
	}

	rb.gameMessage.Message.(*pb.GameMessage_Response).Response.Data = &pb.Response_RoomInfo{RoomInfo: roomInfo}
	return rb
}

// WithRoomPlayerUpdate: RoomPlayerUpdate 데이터를 추가. 방의 특정 플레이어의 업데이트 정보만 포함
func (rb *ResponseBuilder) WithRoomPlayerUpdate(roomId string, player *Player) *ResponseBuilder {
	// RoomPlayerUpdate 메시지 생성
	roomPlayerUpdate := &pb.RoomPlayerUpdate{
		RoomId: roomId,
		PlayerInfo: &pb.PlayerInfo{
			PlayerId: player.playerId,
			X:        player.x,
			Y:        player.y,
			Z:        player.z,
			Rx:       player.rx,
			Ry:       player.ry,
			Rz:       player.rz,
			Speed:    player.speed,
			Health:   player.health,
			// PrefabType: player.prefabType, // prefabType이 있다면 추가
		},
	}

	// GameMessage의 Message 필드에 RoomPlayerUpdate를 추가
	rb.gameMessage.Message = &pb.GameMessage_RoomPlayerUpdate{
		RoomPlayerUpdate: roomPlayerUpdate,
	}
	return rb
}

// Build: 최종 GameMessage 반환
func (rb *ResponseBuilder) Build() *pb.GameMessage {
	return rb.gameMessage
}
