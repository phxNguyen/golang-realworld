package articleservice

import (
	"context"
	"errors"
	"golang-cookie-blog/modules/article/articlemodel"
)

type DeleteArticleStore interface {
	FindArticleByCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*articlemodel.Article, error)
	DeleteArticle(ctx context.Context, cond map[string]interface{}) error
}

type deleteArticleService struct {
	store DeleteArticleStore
}

func NewDeleteArticleService(store DeleteArticleStore) *deleteArticleService {
	return &deleteArticleService{store: store}
}
func (service *deleteArticleService) DeleteArticleService(ctx context.Context, id int) error {

	// lấy data đã có trong DB
	oldData, err := service.store.FindArticleByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	// rồi sửa status thành 0 vì đây là soft delete
	if err := service.store.DeleteArticle(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
