package userfollowstorage

import (
	"context"
	"instago2/common"
	"instago2/modules/userfollow/userfollowmodel"
)

func (s *sqlStore) Delete(ctx context.Context, followerId, userId int) error {
	db := s.db

	if err := db.Table(userfollowmodel.Follow{}.TableName()).
		Where("follower_id = ? and user_id = ?", followerId, userId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
