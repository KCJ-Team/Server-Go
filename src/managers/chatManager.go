package manager

// Protocol Buffers로 생성된 메시지 패키지

// Protocol Buffers를 위한 proto 패키지

// chatManager 전역 변수로 ChatManager 싱글톤 패턴 구현
var chatManager *ChatManager

// ChatManager는 플레이어들의 목록을 관리하고, 채팅 메시지를 브로드캐스트하는 구조체
type ChatManager struct {
	players map[int]Player // 플레이어 ID와 Player 정보를 저장하는 맵
	nextID  int            // 새로운 플레이어 ID를 생성할 때 사용
}

// GetChatManager: ChatManager의 싱글톤 인스턴스를 반환
func GetChatManager() *ChatManager {
	if chatManager == nil {
		// ChatManager 인스턴스를 초기화
		chatManager = &ChatManager{
			players: make(map[int]Player), // 플레이어 정보를 저장할 맵 초기화
			nextID:  1,
		}
	}
	return chatManager
}

// Broadcast: 주어진 이름과 내용을 포함한 채팅 메시지를 모든 플레이어에게 전송
func (pm *ChatManager) Broadcast(name string, content string) {
	// // 채팅 메시지를 포함하는 GameMessage 생성
	// gameMessage := &pb.GameMessage{
	// 	Message: &pb.GameMessage_Chat{
	// 		Chat: &pb.ChatMessage{
	// 			Sender:  name,    // 메시지 보낸 사람의 이름
	// 			Content: content, // 채팅 내용
	// 		},
	// 	},
	// }

	// // Protocol Buffers 형식으로 직렬화하여 네트워크 전송을 위한 데이터 준비
	// response, err := proto.Marshal(gameMessage)
	// if err != nil {
	// 	log.Printf("Failed to marshal response: %v", err)
	// 	return // 직렬화 실패 시 함수 종료
	// }

	// // PlayerManager의 모든 플레이어에게 메시지를 브로드캐스트
	// for _, player := range GetPlayerManager().ListPlayers() {
	// 	// 패킷의 길이를 4바이트로 설정하고 Little Endian 형식으로 변환
	// 	lengthBuf := make([]byte, 4)
	// 	binary.LittleEndian.PutUint32(lengthBuf, uint32(len(response)))

	// 	// 메시지 길이와 메시지 본문을 결합하여 패킷 완성
	// 	lengthBuf = append(lengthBuf, response...)

	// 	// 각 플레이어의 네트워크 연결에 패킷 전송
	// 	(*player.Conn).Write(lengthBuf)
	// }
}
