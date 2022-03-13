package userstorage

import (
	"context"
	"gorm.io/gorm"
	"instago2/common"
	"instago2/modules/user/usermodel"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseFollowCount(ctx context.Context, id int, isCurrentUser bool) error {
	db := s.db

	if isCurrentUser {
		if err := db.Table(usermodel.User{}.TableName()).
			Where("id = ?", id).
			Update("following_count", gorm.Expr("following_count + ?", 1)).Error; err != nil {
			return common.ErrDB(err)
		}
	} else {
		if err := db.Table(usermodel.User{}.TableName()).
			Where("id = ?", id).
			Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return common.ErrDB(err)
		}
	}
	return nil
}

func (s *sqlStore) DecreaseFollowCount(ctx context.Context, id int, isCurrentUser bool) error {
	db := s.db

	if isCurrentUser {
		if err := db.Table(usermodel.User{}.TableName()).
			Where("id = ?", id).
			Update("following_count", gorm.Expr("following_count - ?", 1)).Error; err != nil {
			return common.ErrDB(err)
		}
	} else {
		if err := db.Table(usermodel.User{}.TableName()).
			Where("id = ?", id).
			Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return common.ErrDB(err)
		}
	}
	return nil
}
