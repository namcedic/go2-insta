package postbiz

import (
	"context"
	"instago2/modules/post/postmodel"
)

type CreatePost interface {
	Create(ctx context.Context, data *postmodel.PostCreate) error
}

type createPostBiz struct {
	store CreatePost
}

func NewCreatePostBiz(store CreatePost) *createPostBiz {
	return &createPostBiz{store: store}
}

func (biz *createPostBiz) CreatePost(ctx context.Context, data *postmodel.PostCreate) error {

	err := biz.store.Create(ctx, data)

	return err
}
