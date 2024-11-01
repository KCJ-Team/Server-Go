package manager

import (
	"errors"
	"fmt"
	"log"
	"server-go/src/pb"
	"sync"
)

var roomManager *RoomManager

// Room 구조체: 각 방에 대한 정보를 관리
type Room struct {
	RoomId  string
	Player1 *Player
	Player2 *Player
}

// RoomManager 구조체: 방을 관리하는 싱글톤 매니저
type RoomManager struct {
	rooms map[string]*Room
	mutex sync.Mutex
}

// GetRoomManager: 싱글톤 패턴으로 RoomManager 생성 및 반환
func GetRoomManager() *RoomManager {
	if roomManager == nil {
		roomManager = &RoomManager{
			rooms: make(map[string]*Room),
		}
	}
	return roomManager
}

// CreateRoom: 매치 성공 후 두 명의 플레이어로 방을 생성
func (rm *RoomManager) CreateRoom(player1, player2 *Player) (*Room, error) {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	// Room ID 생성 (플레이어 ID를 조합하여 방 ID 생성, 구분자는 "|")
	roomId := fmt.Sprintf("%s|%s", player1.playerId, player2.playerId)

	// 방이 이미 존재하는지 확인
	if _, exists := rm.rooms[roomId]; exists {
		return nil, errors.New("Room already exists")
	}

	// 새로운 Room 생성
	room := &Room{
		RoomId:  roomId,
		Player1: player1,
		Player2: player2,
	}

	// Room을 RoomManager의 맵에 추가
	rm.rooms[roomId] = room

	// 성공시 RoomInfo 응답을 전송
	GetNetworkManager().SendRoomInfoResponse(pb.MessageType_MATCHMAKING_START, Success, player1.conn, "룸 생성 성공, 멀티 게임 시작", room)
	GetNetworkManager().SendRoomInfoResponse(pb.MessageType_MATCHMAKING_START, Success, player2.conn, "룸 생성 성공, 멀티 게임 시작", room)

	return room, nil
}

// DeleteRoom: 특정 Room을 삭제. 나중에 게임 종료시 룸을 파괴해야한다.
func (rm *RoomManager) DeleteRoom(roomId string) error {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	// 방이 존재하는지 확인
	if _, exists := rm.rooms[roomId]; !exists {
		return errors.New("Room not found")
	}

	// 방 삭제
	delete(rm.rooms, roomId)
	log.Printf("Room %s deleted", roomId)

	return nil
}
