package commentlikestorage

import (
	"context"
	"instago2/common"
	"instago2/modules/commentlike/commentlikemodel"
)

func (s *sqlStore) Create(ctx context.Context, data *commentlikemodel.CommentLikes) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
