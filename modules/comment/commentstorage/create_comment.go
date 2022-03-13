package commentstorage

import (
	"context"
	"instago2/common"
	"instago2/modules/comment/commentmodel"
)

func (s *sqlStore) CreateComment(ctx context.Context, data *commentmodel.CommentCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
