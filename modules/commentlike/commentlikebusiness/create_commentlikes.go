package commentlikebusiness

import (
	"context"
	"errors"
	"instago2/common"
	"instago2/component/asyncjob"
	"instago2/modules/commentlike/commentlikemodel"
)

type CreateCommentLikes interface {
	Find(ctx context.Context, condition map[string]interface{}) (*commentlikemodel.CommentLikes, error)
	Create(ctx context.Context, data *commentlikemodel.CommentLikes) error
}

type IncreaseCommentLikeCountStore interface {
	IncreaseCommentLikeCount(ctx context.Context, commentId int) error
}

type CreateCommentLikesBiz struct {
	store    CreateCommentLikes
	incStore IncreaseCommentLikeCountStore
}

func NewCreateCommentLikesBiz(store CreateCommentLikes, incStore IncreaseCommentLikeCountStore) *CreateCommentLikesBiz {
	return &CreateCommentLikesBiz{store: store, incStore: incStore}
}

func (biz *CreateCommentLikesBiz) CreateCommentLikes(ctx context.Context, data *commentlikemodel.CommentLikes) error {
	if data, _ := biz.store.Find(ctx, map[string]interface{}{"comment_id": data.CommentId, "user_id": data.UserId}); data != nil {
		return commentlikemodel.ErrCannotLikeComment(errors.New("You can not like this comment"))
	}
	err := biz.store.Create(ctx, data)

	if err != nil {
		return common.ErrCannotCreateEntity("CommentLike", err)
	}

	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.incStore.IncreaseCommentLikeCount(ctx, data.CommentId)
		})

		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()
	return err
}
