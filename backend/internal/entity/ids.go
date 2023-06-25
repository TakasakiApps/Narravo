package entity

import "gorm.io/gorm"

type IdRecord struct {
	gorm.Model
	Value string `json:"value"`
}
