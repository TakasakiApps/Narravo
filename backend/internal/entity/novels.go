package entity

import (
	"gorm.io/gorm"
)

// Database entities

type PostNovel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CoverId     string `json:"coverId"`
	AuthorId    string `json:"authorId"`
}

var NovelTable = "novels"

type Novel struct {
	gorm.Model
	ID string `gorm:"primarykey" json:"id"`
	PostNovel
	// TODO: 增加小说审核逻辑
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
