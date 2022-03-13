package userbiz

import (
	"context"
	"instago2/common"
)

type GetOtherProfile interface {
	FindOtherProfile(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*common.SimpleUser, error)
}

type getOtherProfileBiz struct {
	store GetOtherProfile
}

func NewGetOtherProfileBiz(store GetOtherProfile) *getOtherProfileBiz {
	return &getOtherProfileBiz{store: store}
}

func (biz *getOtherProfileBiz) GetProfile(ctx context.Context, id int) (*common.SimpleUser, error) {
	data, err := biz.store.FindOtherProfile(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity("Can not found the user profile!", err)
	}
	return data, err
}
