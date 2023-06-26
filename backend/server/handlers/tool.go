package handlers

import (
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
)

// GenerateUniqueID generate a unique ID.
func GenerateUniqueID() string {
	id := utils.GenerateStandardID()
	if dao.GetInstance().IsIdRecordExist(id) {
		return GenerateUniqueID()
	}
	return id
}
