package postlikestorage

import (
	"context"
	"instago2/common"
	"instago2/modules/postlike/postlikemodel"
)

func (s *sqlStore) Find(ctx context.Context, condition map[string]interface{}) (*postlikemodel.PostLikes, error) {
	db := s.db

	var oldData postlikemodel.PostLikes

	if err := db.
		Where(condition).
		First(&oldData).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &oldData, nil
}
