package commentstorage

import (
	"context"
	"instago2/common"
	"instago2/modules/comment/commentmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *commentmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]commentmodel.Comment, error) {
	var result []commentmodel.Comment
	db := s.db
	db = db.Table(commentmodel.Comment{}.TableName()).
		Where(conditions).Where("status in (1)").Find(&result)

	if v := filter; v != nil {
		if v.PostId > 0 {
			db = db.Where("post_id = ?", v.PostId)
		}
	}

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
