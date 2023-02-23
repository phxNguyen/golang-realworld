package uploadstore

import (
	"context"
	"golang-realworld/common"
)

func (s *UploadStore) UpdateImage(ctx context.Context, id int) error {

	db := s.db

	if err := db.Table(common.Image{}.TableName()).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
