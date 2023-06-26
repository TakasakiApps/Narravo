package entity

import "gorm.io/gorm"

var IdRecordTable = "id_records"

type IdRecord struct {
	gorm.Model
	Value string `json:"value"`
}
