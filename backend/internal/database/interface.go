package database

import (
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"gorm.io/gorm"
)

type Driver interface {
	GetInstance() *gorm.DB

	UserCtrl
	NovelCtrl
}

// UserCtrl management for User
type UserCtrl interface {
	// AddUser adds a new user to the database and returns the number of rows affected
	AddUser(user *entity.User) int64
	// UpdateUser updates the information of a user specified by the input parameter `user`.
	UpdateUser(user *entity.User) int64
	// QueryUserByName QueryUser searches the database for a user with the given username and returns it if found
	QueryUserByName(username string) *entity.User
}

// NovelCtrl management for Novel
type NovelCtrl interface {
	// AddNovel adds a novel to the system and returns the ID of the new novel.
	AddNovel(novel *entity.Novel) int64
	// UpdateNovel updates the information of a novel in the system and returns the number of affected rows.
	UpdateNovel(novel *entity.Novel) int64
	// QueryNovelById queries a novel from the system by its ID and returns the novel entity.
	QueryNovelById(novelId string) *entity.Novel
}
