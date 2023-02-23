package articlestore

import "gorm.io/gorm"

type articleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) *articleStore {
	return &articleStore{db: db}
}
