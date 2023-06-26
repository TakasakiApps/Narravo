package sqlite

import "github.com/TakasakiApps/Narravo/backend/internal/entity"

func (s SQLite) AddAuthor(author *entity.Author) int64 {
	//TODO implement me
	panic("implement me")
}

func (s SQLite) UpdateAuthor(author *entity.Author) int64 {
	//TODO implement me
	panic("implement me")
}

func (s SQLite) QueryAuthorById(authorId string) *entity.Author {
	var authors []*entity.Author
	s.GetInstance().Table(entity.AuthorTable).Where("id = ?", authorId).Find(&authors)

	if len(authors) == 0 {
		return nil
	}

	return authors[0]
}
