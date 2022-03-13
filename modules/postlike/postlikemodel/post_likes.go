package postlikemodel

import (
	"fmt"
	"instago2/common"
)

type PostLikes struct {
	PostId int                `json:"post_id" gorm:"column:post_id;"`
	UserId int                `json:"user_id" gorm:"column:user_id;"`
	User   *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (u *PostLikes) GetUserId() int {
	return u.UserId
}
func (u *PostLikes) GetPostId() int {
	return u.PostId
}

func (PostLikes) TableName() string {
	return "post_likes"
}

func ErrCannotUnlikePost(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot unlike this post"),
		fmt.Sprintf("ErrCannotUnlikePost"),
	)
}
