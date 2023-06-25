package entity

import (
	"gorm.io/gorm"
)

// Database entities

type Novel struct {
	gorm.Model
	BookId   string `json:"bookId"`
	AuthorId string `json:"authorId"`
}

type AuthorInfo struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarId    string `json:"avatarId"`
}

type ContentInfo struct {
	gorm.Model
	Name        string `json:"name"`
	CatalogId   string `json:"catalogId"`
	CoverId     string `json:"coverId"`
	Description string `json:"description"`
}

// Entities to store at local storage

type CatalogInfo struct {
	Chapters []CatalogInfo `json:"chapters"`
}

type ChapterInfo struct {
	Name      string `json:"name"`
	ChapterId string `json:"chapterId"`
}
