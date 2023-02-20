package articlestore

import (
	"context"
	"golang-cookie-blog/modules/article/articlemodel"
)

func (s *sqlStore) FindArticleByCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*articlemodel.Article, error) {

	var data articlemodel.Article

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
