package commentlikebusiness

import (
	"context"
	"errors"
	"instago2/common"
	"instago2/component/asyncjob"
	"instago2/modules/commentlike/commentlikemodel"
)

type UnlikeCommentStore interface {
	Find(ctx context.Context, condition map[string]interface{}) (*commentlikemodel.CommentLikes, error)
	Delete(ctx context.Context, commentId, userId int) error
}

type unlikeCommentBiz struct {
	store    UnlikeCommentStore
	decStore DecreaseLikedCountStore
}
type DecreaseLikedCountStore interface {
	DecreaseCommentLikeCount(ctx context.Context, commentId int) error
}

func NewUnlikeCommentBiz(store UnlikeCommentStore, decStore DecreaseLikedCountStore) *unlikeCommentBiz {
	return &unlikeCommentBiz{
		store:    store,
		decStore: decStore,
	}
}

func (biz *unlikeCommentBiz) UnlikeComment(ctx context.Context, commentId, userId int) error {
	if data, _ := biz.store.Find(ctx, map[string]interface{}{"comment_id": commentId, "user_id": userId}); data == nil {
		return commentlikemodel.ErrCannotUnlikeComment(errors.New("you did not like this comment"))
	}
	if err := biz.store.Delete(ctx, commentId, userId); err != nil {
		return commentlikemodel.ErrCannotUnlikeComment(err)
	}

	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.decStore.DecreaseCommentLikeCount(ctx, commentId)
		})

		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()
	return nil
}
