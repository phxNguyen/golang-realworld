package articlestore

import (
	"context"
	"golang-realworld/common"
	"golang-realworld/modules/article/articlemodel"
)

func (s *articleStore) UpdateArticle(ctx context.Context, data *articlemodel.ArticleUpdate, cond map[string]interface{}) error {

	db := s.db

	if err := db.Table(articlemodel.Article{}.TableName()).
		Where(cond).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
