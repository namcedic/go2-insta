package userfollowbiz

import (
	"context"
	"errors"
	"instago2/common"
	"instago2/component/asyncjob"
	"instago2/modules/userfollow/userfollowmodel"
)

type UserUnfollowUserStore interface {
	Find(ctx context.Context, condition map[string]interface{}) (*userfollowmodel.Follow, error)
	Delete(ctx context.Context, followerId, userId int) error
}

type userUnfollowUserBiz struct {
	store    UserUnfollowUserStore
	decStore DecreaseFollowStore
}
type DecreaseFollowStore interface {
	DecreaseFollowCount(ctx context.Context, id int, isFollowing bool) error
}

func NewUserUnfollowUserBiz(store UserUnfollowUserStore, decStore DecreaseFollowStore) *userUnfollowUserBiz {
	return &userUnfollowUserBiz{
		store:    store,
		decStore: decStore,
	}
}

func (biz *userUnfollowUserBiz) UserUnfollowUser(ctx context.Context, followerId, userId int) error {
	if followerId == userId {
		return userfollowmodel.ErrCannotUnfollowOwnAccount(errors.New("you cannot unfollow your own account"))
	}
	if data, _ := biz.store.Find(ctx, map[string]interface{}{"follower_id": followerId, "user_id": userId}); data == nil {
		return userfollowmodel.ErrCannotUnfollowUser(errors.New("you did not follow this user"))
	}
	if err := biz.store.Delete(ctx, followerId, userId); err != nil {
		return userfollowmodel.ErrCannotUnfollowUser(err)
	}

	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.decStore.DecreaseFollowCount(ctx, followerId, true)
		})

		job1 := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.decStore.DecreaseFollowCount(ctx, userId, false)
		})

		_ = asyncjob.NewGroup(true, job, job1).Run(ctx)
	}()
	return nil
}
