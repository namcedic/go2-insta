package poststorage

import (
	"context"
	"gorm.io/gorm"
	"instago2/common"
	"instago2/modules/post/postmodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*postmodel.Post, error) {
	var result postmodel.Post

	db := s.db

	db = db.Preload("User")

	if err := db.Where(conditions).Where("status in (1)").
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
