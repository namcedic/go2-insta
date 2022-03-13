package poststorage

import (
	"context"
	"gorm.io/gorm"
	"instago2/common"
	"instago2/modules/post/postmodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	userId int,
	data *postmodel.PostUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Where("user_id = ?", userId).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(postmodel.PostUpdate{}.TableName()).Where("id = ?", id).
		Update("post_liked_count", gorm.Expr("post_liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreasePostLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(postmodel.PostUpdate{}.TableName()).Where("id = ?", id).
		Update("post_liked_count", gorm.Expr("post_liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) IncreaseCommentCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(postmodel.PostUpdate{}.TableName()).Where("id = ?", id).
		Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseCommentCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(postmodel.PostUpdate{}.TableName()).Where("id = ? AND status in (1)", id).
		Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
