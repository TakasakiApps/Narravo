package sqlite

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
)

func (s SQLite) AddAuthor(author *entity.Author) int64 {
	result := s.GetInstance().Table(entity.AuthorTable).Create(author)

	return result.RowsAffected
}

func (s SQLite) UpdateAuthor(author *entity.Author) int64 {
	result := s.GetInstance().Table(entity.AuthorTable).Where("id = ?", author.ID).Updates(author)

	return result.RowsAffected
}

func (s SQLite) QueryAuthorsLimit(offset int, limit int) []*entity.Author {
	var authors []*entity.Author
	s.GetInstance().Table(entity.AuthorTable).Offset(offset).Limit(limit).Find(&authors)

	return authors
}

func (s SQLite) QueryAuthorById(authorId string) *entity.Author {
	var authors []*entity.Author
	s.GetInstance().Table(entity.AuthorTable).Where("id = ?", authorId).Find(&authors)

	if len(authors) == 0 {
		return nil
	}

	return authors[0]
}

func (s SQLite) SearchAuthorByName(keyword string, offset int, limit int) []*entity.Author {
	var authors []*entity.Author
	s.GetInstance().Table(entity.AuthorTable).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
		Offset(offset).Limit(limit).Find(&authors)

	return authors
}

func (s SQLite) CountAuthor() (count int64) {
	s.GetInstance().Table(entity.AuthorTable).Count(&count)
	return
}
