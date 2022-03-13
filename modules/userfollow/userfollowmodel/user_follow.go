package userfollowmodel

import (
	"fmt"
	"instago2/common"
	"time"
)

const EntityName = "UserFollow"

type Follow struct {
	FollowerId int                `json:"follower_id" gorm:"column:follower_id"`
	UserId     int                `json:"user_id" gorm:"column:user_id"`
	CreatedAt  *time.Time         `json:"created_at" gorm:"column:created_at"`
	User       *common.SimpleUser `json:"user" gorm:"preload:false;foreignKey:FollowerId"`
}

func (Follow) TableName() string {
	return "followers"
}

type Following struct {
	FollowerId int                `json:"follower_id" gorm:"column:follower_id"`
	UserId     int                `json:"user_id" gorm:"column:user_id"`
	CreatedAt  *time.Time         `json:"created_at" gorm:"column:created_at"`
	User       *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (Following) TableName() string {
	return "followers"
}

func ErrCannotFollowUser(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot follow this user"),
		fmt.Sprintf("ErrCannotFollowUser"),
	)
}
func ErrCannotFollowOwnAccount(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot follow your own account"),
		fmt.Sprintf("ErrCannotFollowOwnAccount"),
	)
}

func ErrCannotUnfollowUser(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot unfollow this user"),
		fmt.Sprintf("ErrCannotUnfollowUser"),
	)
}

func ErrCannotUnfollowOwnAccount(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot unfollow your own account"),
		fmt.Sprintf("ErrCannotFollowOwnAccount"),
	)
}
