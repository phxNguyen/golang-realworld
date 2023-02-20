package articleservice

import (
	"context"
	"golang-cookie-blog/common"
	"golang-cookie-blog/modules/article/articlemodel"
)

type ListArticleStore interface {
	ListArticle(ctx context.Context, filter *articlemodel.Filter, paging *common.Paging, moreKey ...string) ([]articlemodel.Article, error)
}

type listArticleService struct {
	store ListArticleStore
}

func NewListArticleService(store ListArticleStore) *listArticleService {
	return &listArticleService{
		store: store,
	}
}

func (service *listArticleService) ListArticleService(
	ctx context.Context,
	filter *articlemodel.Filter,
	paging *common.Paging,
	moreKey ...string) ([]articlemodel.Article, error) {

	data, err := service.store.ListArticle(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
