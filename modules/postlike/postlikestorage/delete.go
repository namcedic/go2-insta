package postlikestorage

import (
	"context"
	"instago2/common"
	"instago2/modules/postlike/postlikemodel"
)

func (s *sqlStore) Delete(ctx context.Context, postId, userId int) error {
	db := s.db

	if err := db.Table(postlikemodel.PostLikes{}.TableName()).
		Where("post_id = ? and user_id = ?", postId, userId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) DeletePostLikeList(ctx context.Context, postId int) error {
	db := s.db

	if err := db.Table(postlikemodel.PostLikes{}.TableName()).
		Where("post_id = ?", postId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
