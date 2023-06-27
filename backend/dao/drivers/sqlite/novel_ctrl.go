package sqlite

import "github.com/TakasakiApps/Narravo/backend/internal/entity"

func (s SQLite) AddNovel(novel *entity.Novel) int64 {
	result := s.GetInstance().Table(entity.NovelTable).Create(novel)

	return result.RowsAffected
}

func (s SQLite) UpdateNovel(novel *entity.Novel) int64 {
	result := s.GetInstance().Table(entity.NovelTable).Where("id = ?", novel.ID).Updates(novel)

	return result.RowsAffected
}

func (s SQLite) QueryNovelsLimit(offset int, limit int) []*entity.Novel {
	var novels []*entity.Novel
	s.GetInstance().Table(entity.NovelTable).Offset(offset).Limit(limit).Find(&novels)

	return novels
}

func (s SQLite) QueryNovelById(novelId string) *entity.Novel {
	var novels []*entity.Novel
	s.GetInstance().Table(entity.NovelTable).Where("id = ?", novelId).Find(&novels)

	if len(novels) == 0 {
		return nil
	}

	return novels[0]
}

func (s SQLite) CountNovel() (count int64) {
	s.GetInstance().Table(entity.NovelTable).Count(&count)
	return
}