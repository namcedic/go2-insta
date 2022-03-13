package userfollowbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/userfollow/userfollowmodel"
)

type ListFollowerStore interface {
	GetFollowerList(
		ctx context.Context,
		condition map[string]interface{},
		filter *userfollowmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listFollowerBiz struct {
	store ListFollowerStore
}

func NewListFollowerBiz(store ListFollowerStore) *listFollowerBiz {
	return &listFollowerBiz{
		store: store,
	}
}

func (biz *listFollowerBiz) ListFollower(
	ctx context.Context,
	userId int,
	filter *userfollowmodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetFollowerList(ctx, map[string]interface{}{"user_id": userId}, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(userfollowmodel.EntityName, err)
	}
	return users, nil
}
