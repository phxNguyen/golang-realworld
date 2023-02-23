package articlestore

import (
	"context"
	"golang-realworld/common"
	"golang-realworld/modules/article/articlemodel"
)

func (s *articleStore) ListArticle(ctx context.Context, filter *articlemodel.Filter, paging *common.Paging, moreKeys ...string) ([]articlemodel.Article, error) {

	var data []articlemodel.Article
	db := s.db
	offset := (paging.Page - 1) * paging.Limit

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	// conditions
	db = db.Table(articlemodel.Article{}.TableName()).Not("status = ?", 0)

	// filter
	if hasLiked := filter.HasLiked; hasLiked > 0 {
		db = db.Where("has_liked = ?", hasLiked)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	// query
	if err := db.Table(articlemodel.Article{}.TableName()).
		Limit(paging.Limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
