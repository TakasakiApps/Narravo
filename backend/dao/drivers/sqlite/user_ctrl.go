package sqlite

import "github.com/TakasakiApps/Narravo/backend/internal/entity"

func (s SQLite) QueryUser(username string) *entity.User {
	var users []*entity.User
	s.GetInstance().Model(&entity.User{}).Where("name = ?", username).Find(&users)

	if len(users) == 0 {
		return nil
	}

	return users[0]
}

func (s SQLite) AddUser(user *entity.User) int64 {
	result := s.GetInstance().Model(&entity.User{}).Create(user)

	return result.RowsAffected
}
