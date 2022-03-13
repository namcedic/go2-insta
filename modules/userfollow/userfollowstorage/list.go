package userfollowstorage

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"instago2/common"
	"instago2/modules/userfollow/userfollowmodel"
	"time"
)

const TimeLayout = "2006-01-02T15:04:05.999999"

// GetFollowerList
// Get a list of users who are following the current user
func (s *sqlStore) GetFollowerList(
	ctx context.Context,
	condition map[string]interface{},
	filter *userfollowmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	db := s.db
	var result []common.SimpleUser

	db = db.Table(userfollowmodel.Follow{}.TableName()).Where(condition)
	db = db.Joins("JOIN users ON users.id = followers.follower_id")
	if v := filter; v != nil {
		if v.Name != "" {
			db = db.Where("(last_name like ?", "%"+v.Name+"%").
				Or("first_name like ?)", "%"+v.Name+"%")
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(TimeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("followers.created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))

	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("followers.created_at desc").
		Select("users.id", "first_name", "last_name", "avatar", "followers.created_at", "users.status", "users.follower_count", "users.following_count", "users.post_count").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i, item := range result {
		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(TimeLayout))))
			paging.NextCursor = cursorStr
		}
	}
	return result, nil
}

// GetFollowingList
// Get a list of users who the current user is following
func (s *sqlStore) GetFollowingList(
	ctx context.Context,
	condition map[string]interface{},
	filter *userfollowmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	db := s.db
	var result []common.SimpleUser

	db = db.Table(userfollowmodel.Following{}.TableName()).Where(condition)
	db = db.Joins("JOIN users ON users.id = followers.user_id")
	if v := filter; v != nil {
		if v.Name != "" {
			db = db.Where("(last_name like ?", "%"+v.Name+"%").
				Or("first_name like ?)", "%"+v.Name+"%")
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(TimeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("followers.created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))

	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("followers.created_at desc").
		Select("users.id", "email", "first_name", "last_name", "avatar", "followers.created_at", "users.status", "users.follower_count", "users.following_count", "users.post_count").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i, item := range result {
		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(TimeLayout))))
			paging.NextCursor = cursorStr
		}
	}
	return result, nil
}

// GetFollowingList for Post module

func (s *sqlStore) GetFollowingListFromPost(ctx context.Context, currentId int) ([]int, error) {
	db := s.db
	var result []userfollowmodel.Following

	db = db.Table(userfollowmodel.Following{}.TableName()).Where("follower_id = ?", currentId).Find(&result)
	followingList := make([]int, len(result))
	for i := range result {
		followingList[i] = result[i].UserId
	}
	return followingList, nil
}
