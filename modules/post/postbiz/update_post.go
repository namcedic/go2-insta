package postbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/post/postmodel"
)

type UpdatePost interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
	UpdateData(
		ctx context.Context,
		id int,
		userId int,
		data *postmodel.PostUpdate,
	) error
}

type updatePostBiz struct {
	store UpdatePost
}

func NewUpdatePostBiz(store UpdatePost) *updatePostBiz {
	return &updatePostBiz{store: store}
}

func (biz *updatePostBiz) UpdatePost(ctx context.Context, id int, userId int, data *postmodel.PostUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id, "user_id": userId})

	if err != nil {
		return common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	if err := biz.store.UpdateData(ctx, id, userId, data); err != nil {
		return common.ErrCannotUpdateEntity(postmodel.EntityName, err)
	}

	return nil
}
