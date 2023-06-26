package entity

import (
	"gorm.io/gorm"
)

// Database entities

var NovelTable = "novels"

type Novel struct {
	gorm.Model
	ID          string `gorm:"primarykey" json:"id"`
	Name        string `json:"name"`
	CoverId     string `json:"coverId"`
	Description string `json:"description"`

	AuthorId  string `json:"authorId"`
	CatalogId string `json:"catalogId"`
}

var AuthorTable = "authors"

type Author struct {
	gorm.Model
	ID          string `gorm:"primarykey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarId    string `json:"avatarId"`
}

// Entities to store at local storage

type CatalogInfo struct {
	Chapters []CatalogInfo `json:"chapters"`
}

type ChapterInfo struct {
	Name      string `json:"name"`
	ChapterId string `json:"chapterId"`
}
