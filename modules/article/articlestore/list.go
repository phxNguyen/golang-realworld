package articlestore

import (
	"context"
	"golang-cookie-blog/common"
	"golang-cookie-blog/modules/article/articlemodel"
)

func (s *sqlStore) ListArticle(ctx context.Context, filter *articlemodel.Filter, paging *common.Paging, moreKeys ...string) ([]articlemodel.Article, error) {

	var data []articlemodel.Article
	db := s.db
	offset := (paging.Page - 1) * paging.Limit

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	// phai filter tung cai
	if hasLiked := filter.HasLiked; hasLiked > 0 {
		db = db.Where("has_liked = ?", hasLiked)
	}
	if err := db.Table(articlemodel.Article{}.TableName()).
		Limit(paging.Limit).
		Offset(offset).
		Count(&paging.Total).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
