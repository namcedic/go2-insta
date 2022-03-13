package userstorage

import (
	"context"
	"gorm.io/gorm"
	"instago2/common"
	"instago2/modules/user/usermodel"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}

func (s *sqlStore) FindUserByName(
	ctx context.Context,
	paging *common.Paging,
	searchKey string,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var user []common.SimpleUser

	db := s.db.Table(usermodel.User{}.TableName())

	db = db.Where("user_name LIKE ? OR first_name LIKE ? OR last_name LIKE ?",
		"%"+searchKey+"%", "%"+searchKey+"%", "%"+searchKey+"%").
		Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&user).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return user, nil
}
