package manager

import (
	"encoding/binary"
	"log"
	pb "server-go/src/pb" // Protocol Buffers로 생성된 메시지 패키지

	"google.golang.org/protobuf/proto" // Protocol Buffers를 위한 proto 패키지
)

// NetworkManager는 네트워크 패킷을 생성하고 관리하는 구조체
type NetworkManager struct {
}

// networkManager 전역 변수로 NetworkManager 싱글톤 패턴 구현
var networkManager *NetworkManager

// GetNetworkManager: 싱글톤 패턴을 사용하여 NetworkManager 인스턴스를 반환
func GetNetworkManager() *NetworkManager {
	if networkManager == nil {
		networkManager = &NetworkManager{}
	}
	return networkManager
}

// MakePacket: 주어진 GameMessage를 네트워크 패킷으로 변환하는 함수
func (nm *NetworkManager) MakePacket(msg *pb.GameMessage) []byte {
	// GameMessage를 Protocol Buffers 형식으로 직렬화
	response, err := proto.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return nil // 직렬화 실패 시 nil 반환
	}

	// 패킷의 길이를 4바이트 버퍼에 Little Endian 형식으로 저장
	lengthBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(lengthBuf, uint32(len(response)))

	// 패킷 길이 버퍼와 실제 메시지 데이터를 결합하여 최종 패킷 생성
	lengthBuf = append(lengthBuf, response...)

	return lengthBuf // 패킷 반환
}
