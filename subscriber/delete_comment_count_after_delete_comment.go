package subscriber

import (
	"context"
	"instago2/component"
	"instago2/modules/post/poststorage"
	"instago2/pubsub"
)

type HasPostIdFormDeleteComment interface {
	GetPostIdToDecreaseCommentCount() int
}

func DecreaseCommentCountAfterDeletePost(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Delete comment count after delete comment",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
			decreaseData := message.Data().(HasPostIdFormDeleteComment)
			return store.DecreaseCommentCount(ctx, decreaseData.GetPostIdToDecreaseCommentCount())
		},
	}
}
