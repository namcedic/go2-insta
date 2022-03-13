package commentbiz

import (
	"context"
	"instago2/common"
	"instago2/component/asyncjob"
	"instago2/modules/comment/commentmodel"
)

type CreateComment interface {
	CreateComment(ctx context.Context, data *commentmodel.CommentCreate) error
}

type IncreaseCommentCountStore interface {
	IncreaseCommentCount(ctx context.Context, commentId int) error
}

type CreateCommentBiz struct {
	store    CreateComment
	incStore IncreaseCommentCountStore
}

func (biz *CreateCommentBiz) CreateComment(ctx context.Context, data *commentmodel.CommentCreate) error {
	err := biz.store.CreateComment(ctx, data)

	if err != nil {
		return common.ErrCannotCreateEntity("Can not create comment!", err)
	}

	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.incStore.IncreaseCommentCount(ctx, data.PostId)
		})

		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()
	return err
}

func NewCreateCommentBiz(store CreateComment, incStore IncreaseCommentCountStore) *CreateCommentBiz {
	return &CreateCommentBiz{store: store, incStore: incStore}
}
