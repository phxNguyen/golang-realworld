package uploadstore

import "gorm.io/gorm"

type UploadStore struct {
	db *gorm.DB
}

func NewUploadStore(db *gorm.DB) *UploadStore {
	return &UploadStore{
		db: db,
	}
}
