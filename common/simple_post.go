package common

type SimplePost struct {
	SQLModel       `json:",inline"`
	Img            *Image `json:"img" gorm:"column:img;"`
	CommentCount   int    `json:"comment_count" gorm:"comment_count;"`
	PostLikedCount int    `json:"post_liked_count" gorm:"post_liked_count;"`
}

func (SimplePost) TableName() string {
	return "posts"
}

func (u *SimplePost) Mask(isAdmin bool) {
	u.GenUID(DbTypePost)
}
