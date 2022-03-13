package postlikebusiness

import (
	"context"
	"errors"
	"instago2/common"
	"instago2/component/asyncjob"
	"instago2/modules/postlike/postlikemodel"
)

type UnlikePostStore interface {
	Find(ctx context.Context, condition map[string]interface{}) (*postlikemodel.PostLikes, error)
	Delete(ctx context.Context, postId, userId int) error
}

type unlikePostBiz struct {
	store    UnlikePostStore
	decStore DecreasePostLikeCountStore
}
type DecreasePostLikeCountStore interface {
	DecreasePostLikeCount(ctx context.Context, postId int) error
}

func NewUnlikePostBiz(store UnlikePostStore, decStore DecreasePostLikeCountStore) *unlikePostBiz {
	return &unlikePostBiz{
		store:    store,
		decStore: decStore,
	}
}

func (biz *unlikePostBiz) UnlikePost(ctx context.Context, postId, userId int) error {
	if data, _ := biz.store.Find(ctx, map[string]interface{}{"post_id": postId, "user_id": userId}); data == nil {
		return postlikemodel.ErrCannotUnlikePost(errors.New("you did not like this post"))
	}
	if err := biz.store.Delete(ctx, postId, userId); err != nil {
		return postlikemodel.ErrCannotUnlikePost(err)
	}

	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.decStore.DecreasePostLikeCount(ctx, postId)
		})

		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()
	return nil
}
