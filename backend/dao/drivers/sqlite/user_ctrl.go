package sqlite

import "github.com/TakasakiApps/Narravo/backend/internal/entity"

func (s SQLite) AddUser(user *entity.User) int64 {
	result := s.GetInstance().Model(&entity.User{}).Create(user)

	return result.RowsAffected
}

func (s SQLite) UpdateUser(user *entity.User) int64 {
	result := s.GetInstance().Model(&entity.User{}).Where("id = ?", user.ID).Updates(user)

	return result.RowsAffected
}

func (s SQLite) QueryUserByName(username string) *entity.User {
	var users []*entity.User
	s.GetInstance().Model(&entity.User{}).Where("name = ?", username).Find(&users)

	if len(users) == 0 {
		return nil
	}

	return users[0]
}
