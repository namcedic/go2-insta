package commentbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/comment/commentmodel"
)

type ListCommentStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *commentmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]commentmodel.Comment, error)
}

type listCommentBiz struct {
	store ListCommentStore
}

func NewListCommentBiz(store ListCommentStore) *listCommentBiz {
	return &listCommentBiz{store: store}
}

func (biz *listCommentBiz) ListComment(
	ctx context.Context,
	id int,
	filter *commentmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]commentmodel.Comment, error) {
	result, err := biz.store.ListDataByCondition(ctx, map[string]interface{}{"post_id": id}, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(commentmodel.EntityName, err)
	}

	return result, nil
}
