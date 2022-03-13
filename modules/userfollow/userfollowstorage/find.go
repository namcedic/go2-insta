package userfollowstorage

import (
	"context"
	"instago2/common"
	"instago2/modules/userfollow/userfollowmodel"
)

func (s *sqlStore) Find(ctx context.Context, condition map[string]interface{}) (*userfollowmodel.Follow, error) {
	db := s.db

	var oldData userfollowmodel.Follow

	if err := db.
		Where(condition).
		First(&oldData).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &oldData, nil
}
