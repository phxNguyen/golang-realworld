package articleservice

import (
	"context"
	"golang-realworld/modules/article/articlemodel"
)

// CreateArticleStore can du lieu cua storage thi goi qua interface
type CreateArticleStore interface {
	CreateArticle(ctx context.Context, data *articlemodel.ArticleCreate) error
}

type createArticleService struct {
	store CreateArticleStore
}

// NewCreateArticleService là factory function
// dùng để khởi tạo một CreateArticleService và phụ thuộc (DI) vào storage layer
func NewCreateArticleService(store CreateArticleStore) *createArticleService {
	return &createArticleService{store: store}
}

func (service *createArticleService) CreateArticle(ctx context.Context, data *articlemodel.ArticleCreate) error {

	if err := data.Validate(); err != nil {
		return err
	}

	err := service.store.CreateArticle(ctx, data)
	return err
}
