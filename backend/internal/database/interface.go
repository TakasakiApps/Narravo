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
	// AddUser adds a new user to the database and returns the number of rows affected.
	AddUser(user *entity.User) int64

	// UpdateUser updates the information of a user specified by the input parameter `user`.
	UpdateUser(user *entity.User) int64

	// QueryUserByName searches the database for a user with the given username and returns it if found.
	QueryUserByName(username string) *entity.User

	// CountUser counts the number of users in the database and returns the count.
	CountUser() (count int64)
}

// NovelCtrl management for Novel
type NovelCtrl interface {
	// AddNovel adds a novel to the database.
	AddNovel(novel *entity.Novel) int64

	// UpdateNovel updates the information of a novel in the database and returns the number of affected rows.
	UpdateNovel(novel *entity.Novel) int64

	// QueryNovelsLimit queries a list of novels from the database with certain range of offset and limit, and returns the list of novel entities.
	QueryNovelsLimit(offset int, limit int) []*entity.Novel

	// QueryNovelById queries a novel from the database by its ID and returns the novel entity.
	QueryNovelById(novelId string) *entity.Novel

	// CountNovel counts the number of novels in the database and returns the count.
	CountNovel() (count int64)
}

// AuthorCtrl management for Author
type AuthorCtrl interface {
	// AddAuthor adds a new author to the database and returns the number of affected rows.
	AddAuthor(author *entity.Author) int64

	// UpdateAuthor updates the information of an existing author in the database and returns the number of affected rows.
	UpdateAuthor(author *entity.Author) int64

	// QueryAuthorsLimit retrieves a limited number of authors from the database based on the given offset and limit, and returns a slice of Author objects.
	QueryAuthorsLimit(offset int, limit int) []*entity.Author

	// QueryAuthorById retrieves the information of an author based on their ID and returns a pointer to an Author object.
	QueryAuthorById(authorId string) *entity.Author

	// CountAuthor counts the number of authors in the database and returns the count.
	CountAuthor() (count int64)
}

// IdRecordCtrl management for IdRecord
type IdRecordCtrl interface {
	AddIdRecord(id string) int64
	DelIdRecord(id string) int64
	IsIdRecordExist(id string) bool
}
