package articlestore

import (
	"context"
	"golang-cookie-blog/modules/article/articlemodel"
)

func (s *sqlStore) CreateArticle(ctx context.Context, data *articlemodel.ArticleCreate) error {

	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
