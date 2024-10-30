// services/player_service.go
package services

import (
	"errors"
	"server-go/src/database"
	"server-go/src/managers"

	"gorm.io/gorm"
)

// PlayerService: 플레이어 관련 데이터베이스 로직을 관리하는 구조체
type PlayerService struct{}

// NewPlayerService: PlayerService 인스턴스를 생성하는 함수
// 서비스가 필요한 곳에서 이 함수로 새 인스턴스를 반환받아 사용
func NewPlayerService() *PlayerService {
	return &PlayerService{}
}

// CreatePlayer: 새로운 플레이어 정보를 데이터베이스에 추가하는 함수
func (ps *PlayerService) CreatePlayer(player *managers.Player) error {
	result := database.DB.Create(&player)
	return result.Error
}

// GetPlayerByID: ID로 특정 플레이어 정보를 조회하는 함수
func (ps *PlayerService) GetPlayerByID(playerID int) (*managers.Player, error) {
	var player managers.Player
	result := database.DB.First(&player, playerID) // playerID를 조건으로 첫 번째 레코드 조회, player 데이터 반환

	// 에러가 발생하고 그 에러가 RecordNotFound일 경우
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("player not found")
	}
	return &player, result.Error
}

// UpdatePlayer: 특정 플레이어의 정보를 업데이트하는 함수
func (ps *PlayerService) UpdatePlayer(player *managers.Player) error {
	result := database.DB.Save(&player) // 업데이트
	return result.Error
}

// DeletePlayer: ID로 특정 플레이어 정보를 삭제하는 함수
func (ps *PlayerService) DeletePlayer(playerID int) error {
	result := database.DB.Delete(&managers.Player{}, playerID)
	return result.Error
}
