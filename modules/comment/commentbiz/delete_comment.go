package commentbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/comment/commentmodel"
	"instago2/pubsub"
)

type DeleteCommentStore interface {
	FindDataByCondition(
		ctx context.Context,
		data *commentmodel.CommentDelete,
		moreKeys ...string,
	) (*commentmodel.Comment, error)
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteCommentBiz struct {
	store  DeleteCommentStore
	pubsub pubsub.Pubsub
}

func NewDeleteCommentBiz(store DeleteCommentStore, pubsub pubsub.Pubsub) *deleteCommentBiz {
	return &deleteCommentBiz{store: store, pubsub: pubsub}
}

func (biz *deleteCommentBiz) DeleteComment(ctx context.Context, data *commentmodel.CommentDelete) error {
	oldData, err := biz.store.FindDataByCondition(ctx, data)

	if err != nil {
		return common.ErrCannotGetEntity(commentmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(commentmodel.EntityName, nil)
	}

	biz.pubsub.Publish(ctx, common.TopicDeleteComment, pubsub.NewMessage(data))

	if err := biz.store.SoftDeleteData(ctx, data.CommentId); err != nil {
		return common.ErrCannotDeleteEntity(commentmodel.EntityName, err)
	}

	return nil
}
