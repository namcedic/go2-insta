package commentlikestorage

import (
	"context"
	"instago2/common"
	"instago2/modules/commentlike/commentlikemodel"
)

func (s *sqlStore) Delete(ctx context.Context, commentId, userId int) error {
	db := s.db

	if err := db.Table(commentlikemodel.CommentLikes{}.TableName()).
		Where("comment_id = ? and user_id = ?", commentId, userId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) DeleteCommentList(ctx context.Context, postId int) error {
	db := s.db
	if err := db.Table(commentlikemodel.CommentLikes{}.TableName()).
		Where("post_id = ?", postId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
