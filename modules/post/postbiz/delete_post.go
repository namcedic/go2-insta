package postbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/post/postmodel"
	"instago2/pubsub"
)

type DeletePostStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
	SoftDeleteData(
		ctx context.Context,
		data *postmodel.PostDelete,
	) error
}

type deletePostBiz struct {
	store  DeletePostStore
	pubsub pubsub.Pubsub
}

func NewDeletePostBiz(store DeletePostStore, pubsub pubsub.Pubsub) *deletePostBiz {
	return &deletePostBiz{store: store, pubsub: pubsub}
}

func (biz *deletePostBiz) DeletePost(ctx context.Context, data *postmodel.PostDelete) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": data.PostId, "user_id": data.UserId})

	if err != nil {
		return common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	biz.pubsub.Publish(ctx, common.TopicDeletePost, pubsub.NewMessage(data))

	if err := biz.store.SoftDeleteData(ctx, data); err != nil {
		return common.ErrCannotDeleteEntity(postmodel.EntityName, err)
	}

	return nil
}
