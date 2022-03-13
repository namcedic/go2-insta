package userstorage

import (
	"context"
	"gorm.io/gorm"
	"instago2/common"
)

func (s *sqlStore) FindOtherProfile(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*common.SimpleUser, error) {
	db := s.db.Table(common.SimpleUser{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var profile common.SimpleUser

	if err := db.Where(conditions).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &profile, nil
}
