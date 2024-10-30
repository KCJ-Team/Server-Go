package main

import (
	"fmt"
	"log"
	"net"
)

// 프로토콜과 포트를 상수로 정의 => 나중에 환경변수로 설정...
const (
	protocol = "tcp"
	port     = "8888"
)

func main() {
	// 서버 주소를 형식화
	address := fmt.Sprintf(":%s", port)

	// TCP 서버 설정
	listener, err := net.Listen(protocol, address)
	if err != nil {
		log.Fatalf("Failed to listen on %s:%s: %v", protocol, port, err)
	}
	defer listener.Close()

	log.Printf("Server is listening on %s:%s...", protocol, port)

	StartConnection(listener) // 서버 시작
}
