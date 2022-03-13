package userfollowstorage

import (
	"context"
	"instago2/common"
	"instago2/modules/userfollow/userfollowmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *userfollowmodel.Follow) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
