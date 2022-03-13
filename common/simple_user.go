package common

type SimpleUser struct {
	SQLModel       `json:",inline"`
	LastName       string `json:"last_name" gorm:"column:last_name;"`
	FirstName      string `json:"first_name" gorm:"column:first_name;"`
	Avatar         *Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	FollowerCount  int    `json:"follower_count" gorm:"column:follower_count;"`
	FollowingCount int    `json:"following_count" gorm:"column:following_count;"`
	PostCount      int    `json:"post_count" gorm:"column:post_count;"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)
}
