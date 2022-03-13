package postsearchbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/post/postmodel"
	"instago2/modules/postsearch/postsearchmodel"
)

type ListPostStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		data *postsearchmodel.DataSearch,
		filter *postsearchmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]postmodel.Post, error)
}

type listPostBiz struct {
	store ListPostStore
}

func NewListPostBiz(store ListPostStore) *listPostBiz {
	return &listPostBiz{store: store}
}

func (biz *listPostBiz) ListSearch(
	ctx context.Context,
	data *postsearchmodel.DataSearch,
	filter *postsearchmodel.Filter,
	paging *common.Paging,
) ([]postmodel.Post, error) {
	var resultNil []postmodel.Post
	if data.SearchName == nil && data.Caption == nil && data.From == nil && data.To == nil {
		return resultNil, nil
	}

	//if data.Caption == nil {
	//	capt := "%%"
	//	data.Caption = &capt
	//} else {
	//	*data.Caption = "%" + *data.Caption + "%"
	//}
	//
	//if data.SearchName == nil {
	//	searchName := "%%"
	//	data.SearchName = &searchName
	//} else {
	//	*data.SearchName = "%" + *data.SearchName + "%"
	//}
	//
	//if data.From == nil {
	//	data.From = &time.Time{}
	//}
	//if data.To == nil {
	//	timeNow := time.Now()
	//	data.To = &timeNow
	//}

	result, err := biz.store.ListDataByCondition(ctx, nil, data, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	return result, nil
}
