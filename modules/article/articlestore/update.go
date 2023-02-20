package articlestore

import (
	"context"
	"golang-cookie-blog/modules/article/articlemodel"
)

func (s *sqlStore) UpdateArticle(ctx context.Context, data *articlemodel.ArticleUpdate, cond map[string]interface{}) error {

	db := s.db

	if err := db.Table(articlemodel.Article{}.TableName()).
		Where(cond).
		Updates(data).Error; err != nil {
		return err
	}

	return nil
}
