package database

import (
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"gorm.io/gorm"
)

type Driver interface {
	GetInstance() *gorm.DB

	UserCtrl
	NovelCtrl
	AuthorCtrl
	IdRecordCtrl
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
	// AddNovel adds a novel to the database.
	AddNovel(novel *entity.Novel) int64
	// UpdateNovel updates the information of a novel in the database and returns the number of affected rows.
	UpdateNovel(novel *entity.Novel) int64
	// QueryNovelById queries a novel from the database by its ID and returns the novel entity.
	QueryNovelById(novelId string) *entity.Novel
}

// AuthorCtrl management for Author
type AuthorCtrl interface {
	// AddAuthor adds a new author to the database and returns the number of affected rows.
	AddAuthor(author *entity.Author) int64

	// UpdateAuthor updates the information of an existing author in the database and returns the number of affected rows.
	UpdateAuthor(author *entity.Author) int64

	// QueryAuthorById retrieves the information of an author based on their ID and returns a pointer to an Author object.
	QueryAuthorById(authorId string) *entity.Author
}

// IdRecordCtrl management for IdRecord
type IdRecordCtrl interface {
	AddIdRecord(id string) int64
	DelIdRecord(id string) int64
	IsIdRecordExist(id string) bool
}
