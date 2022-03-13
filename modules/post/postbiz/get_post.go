package postbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/post/postmodel"
)

type GetPostStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
}

type getPostBiz struct {
	store GetPostStore
}

func NewGetCategoryBiz(store GetPostStore) *getPostBiz {
	return &getPostBiz{store: store}
}

func (biz *getPostBiz) GetPost(ctx context.Context, id int) (*postmodel.Post, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	return data, err
}
