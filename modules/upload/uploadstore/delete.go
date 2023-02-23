package uploadstore

import (
	"context"
	"golang-realworld/common"
)

func (s *UploadStore) DeleteImage(ctx context.Context, ids []int) error {

	db := s.db

	if err := db.Table(common.Image{}.TableName()).Where("id in (?)", ids).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
