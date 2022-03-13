package commentstorage

import (
	"context"
	"gorm.io/gorm"
	"instago2/common"
	"instago2/modules/comment/commentmodel"
)

func (s *sqlStore) FindCommentIsAllowed(ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*commentmodel.CommentCreate, error) {
	db := s.db.Table(commentmodel.CommentCreate{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var allowReply commentmodel.CommentCreate

	if err := db.Where(conditions).First(&allowReply).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &allowReply, nil
}

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	data *commentmodel.CommentDelete,
	moreKeys ...string,
) (*commentmodel.Comment, error) {
	var result commentmodel.Comment

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where("id = ? AND user_id = ? AND status in (1)", data.CommentId, data.UserId).
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
