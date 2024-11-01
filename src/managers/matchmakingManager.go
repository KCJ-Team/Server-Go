package manager

import (
	"errors"
	"log"
	"net"
	"server-go/src/pb"
	"sync"
)

var matchmakingManager *MatchmakingManager

// Matchmaking 구조체: 각 플레이어의 매치메이킹 상태를 관리
type Matchmaking struct {
	Player    *Player // 매치메이킹 중인 플레이어 정보
	IsMatched bool    // 매칭 성공 여부
}

// MatchmakingManager 구조체: 매치메이킹 대기열과 매치메이킹 취소 관리
type MatchmakingManager struct {
	matchmakingQueue []*Matchmaking // 매치메이킹 대기열
	mutex            sync.Mutex     // 요청을 순차적으로 처리하기 위한 뮤텍스
}

// GetMatchmakingManager: 싱글톤 패턴으로 MatchmakingManager 생성 및 반환
func GetMatchmakingManager() *MatchmakingManager {
	if matchmakingManager == nil {
		matchmakingManager = &MatchmakingManager{
			matchmakingQueue: make([]*Matchmaking, 0), // 빈 슬라이스로 초기화
		}
	}
	return matchmakingManager
}

// MatchMaking을 시작하는 메소드
func (mm *MatchmakingManager) StartMatchmaking(playerId string, conn *net.Conn) error {
	mm.mutex.Lock()
	defer mm.mutex.Unlock()

	// 플레이어매니저에서 특정 플레이어를 갖고와야함..
	player := GetPlayerManager().players[playerId]
	log.Printf("Player %s retrieved from PlayerManager.", playerId)

	// 매치메이킹 생성
	matchmaking := &Matchmaking{
		Player:    player,
		IsMatched: false,
	}

	// 매치메이킹 대기열에 추가
	mm.matchmakingQueue = append(mm.matchmakingQueue, matchmaking)
	log.Printf("Player %s added to matchmaking queue. Queue size: %d", playerId, len(mm.matchmakingQueue))

	// 두 명이 쌓일 때마다 매칭 시도
	if len(mm.matchmakingQueue) >= 2 {
		log.Println("Queue has two or more players. Attempting to match...")
		return mm.TryMatch()
	} else { // 한 명일 경우 대기중 상태
		log.Printf("Player %s is waiting for more players to join.", playerId)
		// 클라이언트에 대기중이라고 알려준다.
		GetNetworkManager().SendResponseMessage(pb.MessageType_MATCHMAKING_START, player.conn, Success, "MatchMaking 대기열 등록 성공, 대기중..")
	}

	return nil
}

// TryMatch는 큐에 두 명 이상의 플레이어가 있을 때 매칭을 진행
func (mm *MatchmakingManager) TryMatch() error {
	// 두 명의 플레이어를 매칭
	match1 := mm.matchmakingQueue[0]
	match2 := mm.matchmakingQueue[1]

	// 매칭이 성공하면 취소 불가능하도록 설정
	match1.IsMatched = true
	match2.IsMatched = true

	// 매칭된 플레이어들을 큐에서 제거
	mm.matchmakingQueue = mm.matchmakingQueue[2:]

	// 매칭 성공 응답을 클라이언트들 에게 전송
	GetNetworkManager().SendResponseMessage(pb.MessageType_MATCHMAKING_START, match1.Player.conn, Success, "Match Success! Create Room(Will Start MultiGame Soon..)")
	GetNetworkManager().SendResponseMessage(pb.MessageType_MATCHMAKING_START, match2.Player.conn, Success, "Match Success! Create Room(Will Start MultiGame Soon..)")

	log.Printf("Match started between Player %s and Player %s", match1.Player.playerId, match2.Player.playerId)

	// 방 생성 로직 호출
	room, err := GetRoomManager().CreateRoom(match1.Player, match2.Player)
	if err != nil {
		log.Printf("Failed to create room: %v", err)
		return err
	}

	log.Printf("Room %s created for Player %s and Player %s", room.RoomId, match1.Player.playerId, match2.Player.playerId)
	return nil
}

// CancelMatchmaking: 특정 플레이어의 매치메이킹을 취소
func (mm *MatchmakingManager) CancelMatchmaking(playerId string) error {
	mm.mutex.Lock()
	defer mm.mutex.Unlock()

	// 매치메이킹 대기열에서 해당 플레이어 찾기 => 큐는 근데 어차피 2개밖에 없음.
	for i, matchmaking := range mm.matchmakingQueue {
		if matchmaking.Player.playerId == playerId {
			// 매칭이 이미 완료된 경우 취소 불가
			if matchmaking.IsMatched {
				return errors.New("Cannot cancel, matchmaking already completed")
			}

			// 대기열에서 해당 플레이어 제거
			mm.matchmakingQueue = append(mm.matchmakingQueue[:i], mm.matchmakingQueue[i+1:]...)
			log.Printf("Player %s deleted at matchmaking queue. Queue size: %d", playerId, len(mm.matchmakingQueue))

			// 클라이언트에게 취소됨을 알려준다.
			GetNetworkManager().SendResponseMessage(pb.MessageType_MATCHMAKING_CANCEL, matchmaking.Player.conn, Success, "MatchMaking 취소, 대기열에서 삭제 성공")

			return nil
		}
	}

	return errors.New("Player not found in matchmaking queue")
}
