package commentmodel

import (
	"errors"
	"instago2/common"
	"strings"
)

const EntityName = "Comment"

type Comment struct {
	common.SQLModel  `json:",inline"`
	UserId           string             `json:"user_id" gorm:"column:user_id;"`
	PostId           string             `json:"post_id" gorm:"column:post_id;"`
	ParentId         string             `json:"parent_id" gorm:"column:parent_id;"`
	Content          string             `json:"content" gorm:"column:content;"`
	User             *common.SimpleUser `json:"user" gorm:"preload:false;"`
	CommentLikeCount int                `json:"comment_like_count" gorm:"column:comment_like_count;"`
}

// can call directly
func (u *Comment) GetCommentId() int {
	return u.Id
}

func (u *Comment) GetUserId() string {
	return u.UserId
}

func (u *Comment) GetPostId() string {
	return u.PostId
}

func (u *Comment) GetParentId() string {
	return u.ParentId
}

// can call directly

func (Comment) TableName() string {
	return "comments"
}

//func (u *Comment) Mask(isAdmin bool) {
//	u.GenUID(common.DbTypeComment)
//}

type CommentCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"user_id" gorm:"column:user_id;"`
	PostId          int    `json:"post_id" gorm:"column:post_id" form:"post_id"`
	ParentId        int    `json:"parent_id" gorm:"column:parent_id" form:"parent_id"`
	Content         string `json:"content" gorm:"column:content" form:"content"`
}

type CommentCreateRequest struct {
	UserId    int    `json:"user_id"`
	PostId    string `json:"post_id" form:"post_id"`
	CommentId string `json:"comment_id" form:"comment_id"`
	Content   string `json:"content" form:"content"`
}

type CommentDelete struct {
	CommentId int `json:"comment_id" form:"id"`
	PostId    int `json:"post_id" form:"post_id"`
	UserId    int `json:"user_id" form:"user_id"`
}

func (c *CommentDelete) GetPostIdToDecreaseCommentCount() int {
	return c.CommentId
}

func (CommentCreate) TableName() string {
	return Comment{}.TableName()
}

func (u *CommentCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeComment)
}

func (res *CommentCreate) Validate() error {
	res.Content = strings.TrimSpace(res.Content)

	if len(res.Content) == 0 {
		return ErrCannotReplyAReply
	}

	return nil
}

var (
	ErrCannotReplyAReply = common.NewCustomError(nil, "can not reply a reply", "ErrCannotReplyAReply")
)

func (data *Comment) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeComment)

	if u := data.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

var (
	ErrCannotReplyComment = common.NewCustomError(
		errors.New("Can not reply comment"),
		"Can not reply comment",
		"ErrCannotReplyComment",
	)
)
