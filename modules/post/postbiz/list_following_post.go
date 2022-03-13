package postbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/post/postmodel"
)

type ListFollowingPostStore interface {
	ListDataFollowingByCondition(ctx context.Context,
		ids []int,
		paging *common.Paging,
		moreKeys ...string,
	) ([]postmodel.Post, error)
}
type GetFollowingListFromPostStore interface {
	GetFollowingListFromPost(ctx context.Context, currentId int) ([]int, error)
}

type listFollowingPostBiz struct {
	store   ListFollowingPostStore
	liStore GetFollowingListFromPostStore
}

func NewListFollowingPostBiz(store ListFollowingPostStore, liStore GetFollowingListFromPostStore) *listFollowingPostBiz {
	return &listFollowingPostBiz{store: store, liStore: liStore}
}

func (biz *listFollowingPostBiz) ListFollowingPost(
	ctx context.Context,
	ids []int,
	currentId int,
	paging *common.Paging,
	moreKeys ...string,
) ([]postmodel.Post, error) {

	ids, err := biz.liStore.GetFollowingListFromPost(ctx, currentId)
	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	result, err := biz.store.ListDataFollowingByCondition(ctx, ids, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	return result, nil
}
