package userbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/user/usermodel"
)

type SearchUserByName interface {
	FindUserByName(
		ctx context.Context,
		paging *common.Paging,
		searchKey string,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type searchUserByNameBiz struct {
	store SearchUserByName
}

func SearchUserByNameBiz(store SearchUserByName) *searchUserByNameBiz {
	return &searchUserByNameBiz{store: store}
}

func (biz *searchUserByNameBiz) SearchUser(ctx context.Context,
	paging *common.Paging,
	searchKey string) ([]common.SimpleUser, error) {

	data, err := biz.store.FindUserByName(ctx, paging, searchKey)

	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	return data, err
}
