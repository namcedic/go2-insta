package commentstorage

import (
	"context"
	"instago2/common"
	"instago2/modules/comment/commentmodel"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(commentmodel.Comment{}.TableName()).
		Where("id = ? OR parent_id = ?", id, id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) SoftDeleteDataList(
	ctx context.Context,
	postId int,
) error {
	db := s.db

	if err := db.Table(commentmodel.Comment{}.TableName()).
		Where("post_id = ?", postId).Where("status in (1)").Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
