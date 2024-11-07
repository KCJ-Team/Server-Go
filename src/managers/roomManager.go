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
	Players map[string]*Player // 여러 플레이어를 map으로 관리
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

	// 새로운 Room 생성. Room(Player의 초기값)
	room := &Room{
		RoomId:  roomId,
		Players: make(map[string]*Player),
	}

	// Players 맵에 각 플레이어 추가
	room.Players[player1.playerId] = player1
	room.Players[player2.playerId] = player2

	// Room을 RoomManager의 맵에 추가
	rm.rooms[roomId] = room

	// 성공시 RoomInfo 응답을 전송
	GetNetworkManager().SendRoomInfoResponse(pb.MessageType_MATCHMAKING_START, Success, player1.conn, "룸 생성 성공, 멀티 게임 시작", room)
	GetNetworkManager().SendRoomInfoResponse(pb.MessageType_MATCHMAKING_START, Success, player2.conn, "룸 생성 성공, 멀티 게임 시작", room)

	return room, nil
}

// UpdatePlayerPosition: 특정 룸의 플레이어 위치를 업데이트하고 다른 플레이어들에게 알림
func (rm *RoomManager) UpdatePlayerPositionInRoom(roomUpdate *pb.RoomPlayerUpdate) error {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	// 룸 ID로 룸 찾기
	room, exists := rm.rooms[roomUpdate.RoomId]
	if !exists {
		return errors.New("room not found")
	}

	// 업데이트할 플레이어 ID로 룸에서 플레이어 찾기
	player, exists := room.Players[roomUpdate.PlayerInfo.PlayerId]
	if !exists {
		return errors.New("player not found in room")
	}

	// 업데이트 할 플레이어 위치 및 기타 정보 업데이트
	player.x = roomUpdate.PlayerInfo.X
	player.y = roomUpdate.PlayerInfo.Y
	player.z = roomUpdate.PlayerInfo.Z
	player.speed = roomUpdate.PlayerInfo.Speed
	player.hp = roomUpdate.PlayerInfo.Hp

	// 디버깅 로그 출력
	log.Printf("Updated position for player %s in room %s: (X: %f, Y: %f, Z: %f, Speed: %f, Health: %f)",
		player.playerId, roomUpdate.RoomId, player.x, player.y, player.z, player.speed, player.hp)

	// 다른 플레이어들에게 업데이트된 위치 정보 브로드캐스트
	for _, otherPlayer := range room.Players {
		if otherPlayer.playerId != player.playerId {
			GetNetworkManager().SendRoomPlayerUpdateResponse(pb.MessageType_PLAYER_POSITION_UPDATE, Success, otherPlayer.conn, "다른 플레이어들에게 위치 업데이트", room.RoomId, player)
		}
	}

	return nil
}

// UpdatePlayerHpInRoom : 특정 룸의 플레이어 HP를 업데이트하고 다른 플레이어들에게 알림
func (rm *RoomManager) UpdatePlayerHpInRoom(roomUpdate *pb.RoomPlayerUpdate) error {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	// 룸 ID로 룸 찾기
	room, exists := rm.rooms[roomUpdate.RoomId]
	if !exists {
		return errors.New("room not found")
	}

	// 업데이트할 플레이어 ID로 룸에서 플레이어 찾기
	player, exists := room.Players[roomUpdate.PlayerInfo.PlayerId]
	if !exists {
		return errors.New("player not found in room")
	}

	// 업데이트 할 플레이어 hp 업데이트
	player.hp = roomUpdate.PlayerInfo.Hp

	// 디버깅 로그 출력
	log.Printf("Updated hp for player %s in room %s: Hp: %f",
		player.playerId, roomUpdate.RoomId, player.hp)

	// 다른 플레이어들에게 업데이트된 정보 브로드캐스트
	for _, otherPlayer := range room.Players {
		if otherPlayer.playerId != player.playerId {
			GetNetworkManager().SendRoomPlayerUpdateResponse(pb.MessageType_PLAYER_HP_UPDATE, Success, otherPlayer.conn, "다른 플레이어들에게 HP 업데이트", room.RoomId, player)
		}
	}

	return nil
}

// ChangePlayerWeaponType: 특정 룸의 플레이어 무기 유형을 업데이트하고 다른 플레이어들에게 알림
func (rm *RoomManager) ChangePlayerWeaponType(roomUpdate *pb.RoomPlayerUpdate) error {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	// 룸 ID로 룸 찾기
	room, exists := rm.rooms[roomUpdate.RoomId]
	if !exists {
		return errors.New("room not found")
	}

	// 업데이트할 플레이어 ID로 룸에서 플레이어 찾기
	player, exists := room.Players[roomUpdate.PlayerInfo.PlayerId]
	if !exists {
		return errors.New("player not found in room")
	}

	// 업데이트 할 플레이어 무기 정보 업데이트
	player.weaponType = roomUpdate.PlayerInfo.PrefabWeaponType

	// 디버깅 로그 출력
	log.Printf("Player %s weapon type updated in Room %s: New Weapon Type = %d",
		player.playerId, room.RoomId, player.weaponType)

	// 다른 플레이어들에게 업데이트된 위치 정보 브로드캐스트
	for _, otherPlayer := range room.Players {
		if otherPlayer.playerId != player.playerId {
			GetNetworkManager().SendRoomPlayerUpdateResponse(pb.MessageType_PLAYER_POSITION_UPDATE, Success, otherPlayer.conn, "다른 플레이어들에게 위치 업데이트", room.RoomId, player)
		}
	}

	return nil
}

// UpdatePlayerAnimationParam: 특정 룸의 플레이어 애니메이션 파라미터를 업데이트하고 다른 플레이어들에게 알림
func (rm *RoomManager) UpdatePlayerAnimationParam(roomUpdate *pb.RoomPlayerUpdate) error {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	// 룸 ID로 룸 찾기
	room, exists := rm.rooms[roomUpdate.RoomId]
	if !exists {
		return errors.New("room not found")
	}

	// 업데이트할 플레이어 ID로 룸에서 플레이어 찾기
	player, exists := room.Players[roomUpdate.PlayerInfo.PlayerId]
	if !exists {
		return errors.New("player not found in room")
	}

	// 플레이어 애니메이션 파라미터 업데이트
	player.animParams.isRunning = roomUpdate.PlayerInfo.AnimParams.IsRunning
	player.animParams.isAim = roomUpdate.PlayerInfo.AnimParams.IsAim
	player.animParams.movementX = roomUpdate.PlayerInfo.AnimParams.MovementX
	player.animParams.movementY = roomUpdate.PlayerInfo.AnimParams.MovementY
	player.animParams.weaponType = roomUpdate.PlayerInfo.AnimParams.WeaponType

	// 디버깅 로그 출력
	log.Printf("Player %s animation parameters updated in Room %s:\n  IsRunning: %f\n  IsAim: %t\n  MovementX: %f\n  MovementY: %f\n  WeaponType: %d",
		player.playerId, room.RoomId,
		player.animParams.isRunning,
		player.animParams.isAim,
		player.animParams.movementX,
		player.animParams.movementY,
		player.animParams.weaponType)

	// 다른 플레이어들에게 업데이트된 위치 정보 브로드캐스트
	for _, otherPlayer := range room.Players {
		if otherPlayer.playerId != player.playerId {
			GetNetworkManager().SendRoomPlayerUpdateResponse(pb.MessageType_PLAYER_POSITION_UPDATE, Success, otherPlayer.conn, "다른 플레이어들에게 위치 업데이트", room.RoomId, player)
		}
	}

	return nil
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
