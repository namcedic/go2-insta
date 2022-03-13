package postlikebusiness

import (
	"context"
	"instago2/common"
	"instago2/component/asyncjob"
	"instago2/modules/postlike/postlikemodel"
)

type CreatePostLikes interface {
	Create(ctx context.Context, data *postlikemodel.PostLikes) error
}

type IncreaseLikedCountStore interface {
	IncreaseLikeCount(ctx context.Context, commentId int) error
}

type CreatePostLikesBiz struct {
	store    CreatePostLikes
	incStore IncreaseLikedCountStore
}

func NewCreatePostLikesBiz(store CreatePostLikes, incStore IncreaseLikedCountStore) *CreatePostLikesBiz {
	return &CreatePostLikesBiz{store: store, incStore: incStore}
}

func (biz *CreatePostLikesBiz) CreatePostLikes(ctx context.Context, data *postlikemodel.PostLikes) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return common.ErrCannotCreateEntity("Postlike", err)
	}

	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.incStore.IncreaseLikeCount(ctx, data.PostId)
		})

		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()

	return err
}
