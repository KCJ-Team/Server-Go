package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"server-go/src/database"

	"gorm.io/gorm"
)

const (
	protocol = "tcp"
	port     = "8888"
)

var DB *gorm.DB // 전역 변수로 DB 인스턴스 선언

func main() {
	// 환경변수에서 DSN 가져오기
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN environment variable is required but not set")
	}

	// 데이터베이스 연결
	database.InitDB(dsn) // DB 초기화

	fmt.Println("Database connected successfully")

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
