package articlestore

import (
	"context"
	"golang-realworld/common"
	"golang-realworld/modules/article/articlemodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindArticleByCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*articlemodel.Article, error) {

	var data articlemodel.Article

	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := s.db.Where(cond).First(&data).Debug().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)

	}
	return &data, nil
}
