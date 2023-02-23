package articlestore

import (
	"context"
	"golang-realworld/common"
	"golang-realworld/modules/article/articlemodel"
)

func (s *articleStore) DeleteArticle(ctx context.Context, cond map[string]interface{}) error {

	db := s.db

	if err := db.Table(articlemodel.Article{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
