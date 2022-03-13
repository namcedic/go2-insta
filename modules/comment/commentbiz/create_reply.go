package commentbiz

import (
	"context"
	"instago2/modules/comment/commentmodel"
)

type CreateReplyStore interface {
	FindCommentIsAllowed(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*commentmodel.CommentCreate, error)
	Create(ctx context.Context, data *commentmodel.CommentCreate) error
}

type createReplyBiz struct {
	store CreateReplyStore
}

func NewCreateReplyBiz(store CreateReplyStore) *createReplyBiz {
	return &createReplyBiz{store: store}
}

func (biz *createReplyBiz) CreateReply(ctx context.Context, data *commentmodel.CommentCreate) error {

	if err := data.Validate(); err != nil {
		return err
	}
	allowReply, _ := biz.store.FindCommentIsAllowed(ctx,
		map[string]interface{}{"id": data.ParentId, "parent_id": nil})

	if allowReply == nil {
		return commentmodel.ErrCannotReplyAReply
	}
	err := biz.store.Create(ctx, data)

	return err
}
