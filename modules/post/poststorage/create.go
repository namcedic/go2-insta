package poststorage

import (
	"context"
	"instago2/common"
	"instago2/modules/post/postmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *postmodel.PostCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
