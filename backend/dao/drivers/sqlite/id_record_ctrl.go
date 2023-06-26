package sqlite

import "github.com/TakasakiApps/Narravo/backend/internal/entity"

func (s SQLite) AddIdRecord(id string) int64 {
	//TODO implement me
	panic("implement me")
}

func (s SQLite) DelIdRecord(id string) int64 {
	//TODO implement me
	panic("implement me")
}

func (s SQLite) IsIdRecordExist(id string) bool {
	var idRecords []*entity.IdRecord
	s.GetInstance().Table(entity.IdRecordTable).Where("value = ?", id).Find(&idRecords)

	return len(idRecords) > 0
}
