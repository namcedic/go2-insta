package userfollowbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/userfollow/userfollowmodel"
)

type ListFollowingStore interface {
	GetFollowingList(
		ctx context.Context,
		condition map[string]interface{},
		filter *userfollowmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listFollowingBiz struct {
	store ListFollowingStore
}

func NewListFollowingBiz(store ListFollowingStore) *listFollowingBiz {
	return &listFollowingBiz{
		store: store,
	}
}

func (biz *listFollowingBiz) ListFollowing(
	ctx context.Context,
	followerId int,
	filter *userfollowmodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetFollowingList(ctx, map[string]interface{}{"follower_id": followerId}, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(userfollowmodel.EntityName, err)
	}
	return users, nil
}
