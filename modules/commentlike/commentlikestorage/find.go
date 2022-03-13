package commentlikestorage

import (
	"context"
	"instago2/common"
	"instago2/modules/commentlike/commentlikemodel"
)

func (s *sqlStore) Find(ctx context.Context, condition map[string]interface{}) (*commentlikemodel.CommentLikes, error) {
	db := s.db

	var oldData commentlikemodel.CommentLikes

	if err := db.
		Where(condition).
		First(&oldData).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &oldData, nil
}
