package postmodel

import (
	"instago2/common"
)

const EntityName = "Post"

type Post struct {
	common.SQLModel `json:",inline"`
	UserId          int                `json:"user_id" gorm:"column:user_id;"`
	Img             *common.Image      `json:"img" gorm:"column:img;"`
	Caption         *string            `json:"caption" gorm:"column:caption;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
	PostLikedCount  int                `json:"post_liked_count" gorm:"post_liked_count;"`
	CommentCount    int                `json:"comment_count" gorm:"comment_count;"`
}

// can call directly, can delete these function
func (u *Post) GetPostId() int {
	return u.Id
}

func (u *Post) GetUserId() int {
	return u.UserId
}

func (Post) TableName() string {
	return "posts"
}

func (u *Post) Mask(isAdmin bool) {
	u.GenUID(common.DbTypePost)
}

func (u *Post) GetCaption() *string {
	return u.Caption
}

// can call directly, can delete this field

type PostCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int           `json:"user_id" gorm:"column:user_id;"`
	Img             *common.Image `json:"img" gorm:"column:img;"`
	Caption         string        `json:"caption" gorm:"column:caption;"`
}

func (PostCreate) TableName() string {
	return "posts"
}

type PostUpdate struct {
	Img     *common.Image `json:"img" gorm:"column:img;"`
	Caption *string       `json:"caption" gorm:"column:caption;"`
}

func (PostUpdate) TableName() string {
	return "posts"
}

type PostDelete struct {
	PostId int `json:"id" gorm:"id;" form:"id" `
	UserId int `json:"user_id" gorm:"user_id;" form:"user_id"`
}

func (p *PostDelete) GetPostDeleteId() int {
	return p.PostId
}

func (p *PostDelete) GetUserId() int {
	return p.UserId
}

func (PostDelete) TableName() string {
	return "posts"
}
