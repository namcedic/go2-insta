package userfollowbiz

import (
	"context"
	"errors"
	"instago2/common"
	"instago2/component/asyncjob"
	"instago2/modules/userfollow/userfollowmodel"
)

type UserFollowUserStore interface {
	Find(ctx context.Context, condition map[string]interface{}) (*userfollowmodel.Follow, error)
	Create(ctx context.Context, data *userfollowmodel.Follow) error
}

type userFollowUserBiz struct {
	store    UserFollowUserStore
	incStore IncreaseFollowStore
}

type IncreaseFollowStore interface {
	IncreaseFollowCount(ctx context.Context, id int, isFollowing bool) error
}

func NewUserFollowBiz(store UserFollowUserStore, incStore IncreaseFollowStore) *userFollowUserBiz {
	return &userFollowUserBiz{
		store:    store,
		incStore: incStore,
	}
}
func (biz *userFollowUserBiz) UserFollowUser(
	ctx context.Context,
	data *userfollowmodel.Follow,
) error {

	if data.FollowerId == data.UserId {
		return userfollowmodel.ErrCannotFollowOwnAccount(errors.New("you cannot follow your own account"))
	}
	if data, _ := biz.store.Find(ctx, map[string]interface{}{"follower_id": data.FollowerId, "user_id": data.UserId}); data != nil {
		return userfollowmodel.ErrCannotFollowUser(errors.New("you had already followed this user"))
	}
	if err := biz.store.Create(ctx, data); err != nil {
		return userfollowmodel.ErrCannotFollowUser(err)
	}
	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.incStore.IncreaseFollowCount(ctx, data.FollowerId, true)
		})

		job1 := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.incStore.IncreaseFollowCount(ctx, data.UserId, false)
		})

		_ = asyncjob.NewGroup(true, job, job1).Run(ctx)
	}()
	return nil
}
