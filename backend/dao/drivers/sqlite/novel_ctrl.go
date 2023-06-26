package sqlite

import "github.com/TakasakiApps/Narravo/backend/internal/entity"

func (s SQLite) AddNovel(novel *entity.Novel) int64 {
	//TODO implement me
	panic("implement me")
}

func (s SQLite) UpdateNovel(novel *entity.Novel) int64 {
	//TODO implement me
	panic("implement me")
}

func (s SQLite) QueryNovelById(novelId string) *entity.Novel {
	var novels []*entity.Novel
	s.GetInstance().Table(entity.NovelTable).Where("id = ?", novelId).Find(&novels)

	if len(novels) == 0 {
		return nil
	}

	return novels[0]
}
