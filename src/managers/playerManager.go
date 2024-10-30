package manager

import (
	"encoding/binary"
	"fmt"
	"log"

	pb "server-go/src/pb"

	"net"

	"google.golang.org/protobuf/proto"
)

// playerManager 전역 변수로 PlayerManager 싱글톤 패턴 구현
var playerManager *PlayerManager

// Player: 개별 플레이어의 정보를 담는 구조체
type Player struct {
	ID   int       // 플레이어 고유 ID
	Name string    // 플레이어 이름
	Age  int       // 플레이어 나이
	Conn *net.Conn // 플레이어의 네트워크 연결
	X    float32   // 플레이어의 X 좌표
	Y    float32   // 플레이어의 Y 좌표
	Z    float32   // 플레이어의 Z 좌표
}

// PlayerManager: 플레이어 목록을 관리하고, ID를 자동 할당하는 구조체
type PlayerManager struct {
	players map[string]*Player // 플레이어 이름을 키로 하여 Player 포인터를 저장하는 맵
	nextID  int                // 새로운 플레이어 ID를 생성할 때 사용
}

// GetPlayerManager: 싱글톤 패턴으로 PlayerManager를 생성 및 반환
func GetPlayerManager() *PlayerManager {
	if playerManager == nil {
		playerManager = &PlayerManager{
			players: make(map[string]*Player),
			nextID:  1,
		}
	}
	return playerManager
}

func (pm *PlayerManager) PlayerPosition(p *pb.GameMessage_PlayerPosition) {
	fmt.Printf("Received Player Position: PlayerID=%s, X=%f, Y=%f, Z=%f\n", p.PlayerPosition.PlayerId, p.PlayerPosition.X, p.PlayerPosition.Y, p.PlayerPosition.Z)
}

// // AddPlayer: 새로운 플레이어를 추가하는 함수
// //	플레이어의 기본 좌표 설정과 자기 자신 및 다른 플레이어들에게 입장 알림을 보냄
// func (pm *PlayerManager) AddPlayer(name string, age int, conn *net.Conn) *Player {
// 	player := Player{
// 		ID:   pm.nextID, // 고유 ID 할당
// 		Name: name,      // 이름 설정
// 		Age:  age,       // 나이 설정
// 		Conn: conn,      // 네트워크 연결 설정
// 	}

// 	// 플레이어를 맵에 추가하고 ID 증가
// 	pm.players[name] = &player
// 	pm.nextID++

// 	// 플레이어의 초기 좌표를 설정
// 	player.X, player.Y, player.Z = 0, 0, 0

// 	// 현재 플레이어에게 자신의 스폰 정보를 전송
// 	myPlayerSpawn := &pb.GameMessage{
// 		Message: &pb.GameMessage_SpawnMyPlayer{
// 			SpawnMyPlayer: &pb.SpawnMyPlayer{
// 				X: player.X,
// 				Y: player.Y,
// 				Z: player.Z,
// 			},
// 		},
// 	}

// 	// 네비게이션 경로 테스트를 위한 메시지 생성
// 	pathTest := &pb.GameMessage{
// 		Message: &pb.GameMessage_PathTest{
// 			PathTest: &pb.PathTest{},
// 		},
// 	}

// 	// 네비게이션 경로 계산
// 	path, err := GetNavMeshManager().PathFinding(-230, 0, -291, 235, 0, 180)
// 	if err == nil {
// 		// 경로에 있는 좌표를 pathTest 메시지에 추가
// 		for _, path := range path.PathList {
// 			pathTest.GetPathTest().Paths = append(pathTest.GetPathTest().Paths, &pb.NavV3{
// 				X: float32(path.X),
// 				Y: float32(path.Y),
// 				Z: float32(path.Z),
// 			})
// 		}
// 		// 경로 데이터를 플레이어에게 전송
// 		response := GetNetworkManager().MakePacket(pathTest)
// 		(*player.Conn).Write(response)
// 	}

// 	// 스폰 메시지를 패킷화하여 플레이어에게 전송
// 	response := GetNetworkManager().MakePacket(myPlayerSpawn)
// 	(*player.Conn).Write(response)

// 	// 새로운 플레이어 입장을 다른 플레이어들에게 알림
// 	otherPlayerSpawnPacket := &pb.GameMessage{
// 		Message: &pb.GameMessage_SpawnOtherPlayer{
// 			SpawnOtherPlayer: &pb.SpawnOtherPlayer{
// 				PlayerId: name,
// 				X:        player.X,
// 				Y:        player.Y,
// 				Z:        player.Z,
// 			},
// 		},
// 	}

// 	for _, p := range pm.players {
// 		if p.Name == name {
// 			continue // 현재 플레이어 제외
// 		}
// 		response = GetNetworkManager().MakePacket(otherPlayerSpawnPacket)
// 		(*p.Conn).Write(response)
// 	}

// 	// 현재 접속한 플레이어에게 다른 모든 플레이어의 위치 정보를 전송
// 	for _, p := range pm.players {
// 		if p.Name == name {
// 			continue
// 		}
// 		otherPlayerSpawnPacket := &pb.GameMessage{
// 			Message: &pb.GameMessage_SpawnOtherPlayer{
// 				SpawnOtherPlayer: &pb.SpawnOtherPlayer{
// 					PlayerId: p.Name,
// 					X:        p.X,
// 					Y:        p.Y,
// 					Z:        p.Z,
// 				},
// 			},
// 		}
// 		response = GetNetworkManager().MakePacket(otherPlayerSpawnPacket)
// 		(*player.Conn).Write(response)
// 	}

// 	return &player
// }

// MovePlayer: 플레이어의 위치를 업데이트하고 다른 플레이어들에게 위치 변경 알림
func (pm *PlayerManager) MovePlayer(p *pb.GameMessage_PlayerPosition) {
	// 해당 플레이어의 좌표 업데이트
	pm.players[p.PlayerPosition.PlayerId].X = p.PlayerPosition.X
	pm.players[p.PlayerPosition.PlayerId].Y = p.PlayerPosition.Y
	pm.players[p.PlayerPosition.PlayerId].Z = p.PlayerPosition.Z

	// 위치 변경 메시지 생성
	response, err := proto.Marshal(&pb.GameMessage{
		Message: p,
	})
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return
	}

	// 위치 변경 메시지를 모든 플레이어에게 전송 (자기 자신 제외)
	for _, player := range pm.players {
		if player.Name == p.PlayerPosition.PlayerId {
			continue
		}
		lengthBuf := make([]byte, 4)
		binary.LittleEndian.PutUint32(lengthBuf, uint32(len(response)))
		lengthBuf = append(lengthBuf, response...)
		(*player.Conn).Write(lengthBuf)
	}
}

// // GetPlayer: ID로 플레이어 정보를 조회
// func (pm *PlayerManager) GetPlayer(id string) (*Player, error) {
// 	player, exists := pm.players[id]
// 	if !exists {
// 		return nil, errors.New("player not found")
// 	}
// 	return player, nil
// }

// // RemovePlayer: 플레이어를 ID로 제거하고, 다른 플레이어들에게 알림
// func (pm *PlayerManager) RemovePlayer(id string) error {
// 	if _, exists := pm.players[id]; !exists {
// 		return errors.New("player not found")
// 	}
// 	delete(pm.players, id)

// 	// 로그아웃 메시지 생성
// 	logoutPacket := &pb.GameMessage{
// 		Message: &pb.GameMessage_Logout{
// 			Logout: &pb.LogoutMessage{
// 				PlayerId: id,
// 			},
// 		},
// 	}

// 	// 모든 플레이어에게 로그아웃 알림 전송
// 	response := GetNetworkManager().MakePacket(logoutPacket)
// 	for _, p := range pm.players {
// 		(*p.Conn).Write(response)
// 	}

// 	return nil
// }

// // ListPlayers: 현재 접속 중인 모든 플레이어 리스트 반환
// func (pm *PlayerManager) ListPlayers() []*Player {
// 	playerList := []*Player{}
// 	for _, player := range pm.players {
// 		playerList = append(playerList, player)
// 	}
// 	return playerList
// }
