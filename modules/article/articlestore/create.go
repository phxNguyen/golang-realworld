package articlestore

import (
	"context"
	"golang-realworld/common"
	"golang-realworld/modules/article/articlemodel"
)

func (s *articleStore) CreateArticle(ctx context.Context, data *articlemodel.ArticleCreate) error {

	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
