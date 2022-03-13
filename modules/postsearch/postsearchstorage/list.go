package postsearchstorage

import (
	"context"
	"instago2/common"
	"instago2/modules/post/postmodel"
	"instago2/modules/postsearch/postsearchmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	data *postsearchmodel.DataSearch,
	filter *postsearchmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]postmodel.Post, error) {
	var result []postmodel.Post

	db := s.db

	db = db.Table(postmodel.Post{}.TableName()).
		Where(conditions)
	if data.Caption != nil {
		db.Where("caption like ?", "%"+*data.Caption+"%")
	}
	if data.From != nil {
		db.Where("posts.created_at >=?", *data.From)
	}
	if data.To != nil {
		db.Where("posts.created_at <=?", *data.To)
	}
	db.Joins("left join users on posts.user_id = users.id")
	if data.SearchName != nil {
		db.Where("users.user_name like ? OR users.first_name like ? OR users.last_name like ? AND users.status in (1)",
			"%"+*data.SearchName+"%", "%"+*data.SearchName+"%", "%"+*data.SearchName+"%")
	}

	//db = db.Table(postmodel.Post{}.TableName()).
	//	Where(conditions).
	//	Where("caption like ? AND posts.created_at BETWEEN ? AND ? AND posts.status in (1)",
	//		*data.Caption, *data.From, *data.To).
	//	Joins("left join users on posts.user_id = users.id").
	//	Where("users.user_name like ? OR users.first_name like ? OR users.last_name like ? AND users.status in (1)",
	//		*data.SearchName, *data.SearchName, *data.SearchName)

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
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
