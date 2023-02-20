package articleservice

import (
	"context"
	"golang-cookie-blog/modules/article/articlemodel"
)

type FindArticleStore interface {
	FindArticleByCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*articlemodel.Article, error)
}

type getArticleService struct {
	store FindArticleStore
}

func NewGetArticle(store FindArticleStore) *getArticleService {
	return &getArticleService{store: store}
}

func (service *getArticleService) FindArticleById(ctx context.Context, id int) (*articlemodel.Article, error) {

	data, err := service.store.FindArticleByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
