package commentstorage

import (
	"context"
	"gorm.io/gorm"
	"instago2/common"
	"instago2/modules/comment/commentmodel"
)

func (s *sqlStore) DecreaseCommentLikeCount(ctx context.Context, commentId int) error {
	db := s.db

	if err := db.Table(commentmodel.Comment{}.TableName()).
		Where("id = ?", commentId).
		Update("comment_liked_count", gorm.Expr("comment_liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) IncreaseCommentLikeCount(ctx context.Context, commentId int) error {
	db := s.db

	if err := db.Table(commentmodel.Comment{}.TableName()).
		Where("id = ?", commentId).
		Update("comment_liked_count", gorm.Expr("comment_liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
