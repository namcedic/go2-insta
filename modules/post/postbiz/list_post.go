package postbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/post/postmodel"
)

type ListPostStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimplePost, error)
}

type listPostBiz struct {
	store ListPostStore
}

func NewListPostBiz(store ListPostStore) *listPostBiz {
	return &listPostBiz{store: store}
}

func (biz *listPostBiz) ListPost(
	ctx context.Context,
	paging *common.Paging,
) ([]common.SimplePost, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	return result, nil
}
