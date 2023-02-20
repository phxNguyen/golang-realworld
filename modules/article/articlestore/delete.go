package articlestore

import (
	"context"
	"golang-cookie-blog/modules/article/articlemodel"
)

func (s *sqlStore) DeleteArticle(ctx context.Context, cond map[string]interface{}) error {

	db := s.db

	if err := db.Table(articlemodel.Article{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return err
	}

	return nil
}