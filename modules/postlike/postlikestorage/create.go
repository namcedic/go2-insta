package postlikestorage

import (
	"context"
	"instago2/common"
	"instago2/modules/postlike/postlikemodel"
)

func (s *sqlStore) Create(ctx context.Context, data *postlikemodel.PostLikes) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
