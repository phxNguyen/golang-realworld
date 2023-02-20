package articleservice

import (
	"context"
	"errors"
	"golang-cookie-blog/modules/article/articlemodel"
)

type UpdateArticleStore interface {
	FindArticleByCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*articlemodel.Article, error)
	UpdateArticle(ctx context.Context, data *articlemodel.ArticleUpdate, cond map[string]interface{}) error
}

type updateArticleService struct {
	store UpdateArticleStore
}

func NewUpdateArticleService(store UpdateArticleStore) *updateArticleService {
	return &updateArticleService{store: store}
}
func (service *updateArticleService) UpdateArticleService(ctx context.Context, data *articlemodel.ArticleUpdate, id int) error {

	// lấy data đã có trong DB
	oldData, err := service.store.FindArticleByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	// rồi update với data mới
	if err := service.store.UpdateArticle(ctx, data, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
