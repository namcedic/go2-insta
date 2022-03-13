package subscriber

import (
	"context"
	"instago2/component"
	"instago2/modules/comment/commentstorage"
	"instago2/modules/commentlike/commentlikestorage"
	"instago2/modules/postlike/postlikestorage"
	"instago2/pubsub"
)

type HasPostId interface {
	GetPostDeleteId() int
}

func DeleteCommentAfterDeletePost(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Delete comment after delete post",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := commentstorage.NewSQLStore(appCtx.GetMainDBConnection())
			deleteData := message.Data().(HasPostId)
			return store.SoftDeleteDataList(ctx, deleteData.GetPostDeleteId())
		},
	}
}

func DeletePostLikeAfterDeletePost(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Delete post like after delete post",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := postlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
			deleteData := message.Data().(HasPostId)
			return store.DeletePostLikeList(ctx, deleteData.GetPostDeleteId())
		},
	}
}

func DeleteCommentLikeAfterDeletePost(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Delete like comment after delete post",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := commentlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
			deleteData := message.Data().(HasPostId)
			return store.DeleteCommentList(ctx, deleteData.GetPostDeleteId())
		},
	}
}
