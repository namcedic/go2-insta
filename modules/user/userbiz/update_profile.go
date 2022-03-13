package userbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/user/usermodel"
)

type UpdateUserProfile interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *usermodel.UserUpdate,
	) error
}

type updateProfileBiz struct {
	store UpdateUserProfile
}

func NewUpdateProfileBiz(store UpdateUserProfile) *updateProfileBiz {
	return &updateProfileBiz{store: store}
}

func (biz *updateProfileBiz) UpdateProfile(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	oldData, err := biz.store.FindUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
