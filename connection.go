package main

import (
	"encoding/binary"
	"log"
	"net"

	mg "server-go/src/managers"
	pb "server-go/src/pb"

	"google.golang.org/protobuf/proto"
)

// StartConnection: 서버 시작 및 연결 처리
func StartConnection(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

// handleConnection: 클라이언트 연결을 처리하고 엔티티별 서비스에 위임
func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		// 메시지 길이를 읽음 (4바이트, Little Endian)
		lengthBuf := make([]byte, 4)
		_, err := conn.Read(lengthBuf)
		if err != nil {
			log.Println("Failed to read message length:", err)
			return
		}
		length := binary.LittleEndian.Uint32(lengthBuf)

		// 메시지 본문 읽기
		messageBuf := make([]byte, length)
		_, err = conn.Read(messageBuf)
		if err != nil {
			log.Println("Failed to read message body:", err)
			return
		}

		// Protocol Buffers 메시지를 파싱
		var gameMessage pb.GameMessage
		if err := proto.Unmarshal(messageBuf, &gameMessage); err != nil {
			log.Println("Failed to unmarshal message:", err)
			return
		}

		// 메시지 처리 및 엔티티로 라우팅
		processMessage(&gameMessage, conn)
	}
}

// processMessage: 수신한 메시지의 유형을 식별하고 해당 엔티티로 전송
// 전송 모델을 먼저 본다음, MessageType을 보고 나눈다.
func processMessage(message *pb.GameMessage, conn net.Conn) {
	switch msg := message.Message.(type) {
	case *pb.GameMessage_PlayerInfo:
		playerInfo := msg.PlayerInfo

		// MessageType에 따라 추가적인 분리
		switch message.MessageType {
		case pb.MessageType_SESSION_LOGIN:
			mg.GetPlayerManager().Login(playerInfo.PlayerId, &conn)
		case pb.MessageType_SESSION_LOGOUT:
			mg.GetPlayerManager().Logout(playerInfo.PlayerId, &conn)
		case pb.MessageType_MATCHMAKING_START:
			mg.GetMatchmakingManager().StartMatchmaking(playerInfo.PlayerId, &conn)
		case pb.MessageType_MATCHMAKING_CANCEL:
			mg.GetMatchmakingManager().CancelMatchmaking(playerInfo.PlayerId)
		}
	case *pb.GameMessage_RoomPlayerUpdate:
		roomUpdate := msg.RoomPlayerUpdate

		switch message.MessageType {
		case pb.MessageType_PLAYER_POSITION_UPDATE:
			mg.GetRoomManager().UpdatePlayerPositionInRoom(roomUpdate)
		case pb.MessageType_PLAYER_HP_UPDATE:
			mg.GetRoomManager().UpdatePlayerHpInRoom(roomUpdate)
		case pb.MessageType_PLAYER_CHANGE_WEAPON:
			mg.GetRoomManager().ChangePlayerWeaponType(roomUpdate)
		case pb.MessageType_PLAYER_ANIMATION:
			mg.GetRoomManager().UpdatePlayerAnimationParam(roomUpdate)
		}
	default:
		log.Printf("Unexpected message type: %T", msg)
	}
}
