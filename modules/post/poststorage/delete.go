package poststorage

import (
	"context"
	"instago2/common"
	"instago2/modules/post/postmodel"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	data *postmodel.PostDelete,
) error {
	db := s.db

	if err := db.Table(postmodel.Post{}.TableName()).
		Where("id = ? AND user_id = ?", data.PostId, data.UserId).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
