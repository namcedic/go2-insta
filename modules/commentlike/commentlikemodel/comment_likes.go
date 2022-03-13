package commentlikemodel

import (
	"fmt"
	"instago2/common"
)

type CommentLikes struct {
	CommentId int                `json:"comment_id" gorm:"column:comment_id;"`
	UserId    int                `json:"user_id" gorm:"column:user_id;"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (u *CommentLikes) GetUserId() int {
	return u.UserId
}
func (u *CommentLikes) GetCommentId() int {
	return u.CommentId
}

func (CommentLikes) TableName() string {
	return "comment_likes"
}
func ErrCannotLikeComment(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this comment"),
		fmt.Sprintf("ErrCannotLikeComment"),
	)
}

func ErrCannotUnlikeComment(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot unlike this comment"),
		fmt.Sprintf("ErrCannotUnlikeComment"),
	)
}
