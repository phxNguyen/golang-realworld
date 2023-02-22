package articleservice

import (
	"context"
	"golang-realworld/common"
	"golang-realworld/modules/article/articlemodel"
)

const EntityName = "Article"

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
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(articlemodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(articlemodel.EntityName, err)
	}
	if data.Status == 0 {
		return nil, common.ErrCannotGetEntity(articlemodel.EntityName, nil)
	}
	return data, err
}
