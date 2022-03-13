package usermodel

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"instago2/common"
	"instago2/component/tokenprovider"
	"regexp"
	"time"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	UserName        string        `json:"user_name" gorm:"column:user_name;"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	DateOfBirth     *time.Time    `json:"date_of_birth" gorm:"column:date_of_birth"`
	Phone           *string       `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	FollowerCount   int           `json:"follower_count" gorm:"column:follower_count;"`
	FollowingCount  int           `json:"following_count" gorm:"column:following_count;"`
	PostCount       int           `json:"post_count" gorm:"column:post_count;"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	UserName        string        `json:"user_name" gorm:"column:user_name;"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"password" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	DateOfBirth     *time.Time    `json:"date_of_birth" gorm:"column:date_of_birth"`
	Phone           *string       `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

func (u UserCreate) Validate() error {
	return validation.ValidateStruct(&u,
		// Email cannot be empty, and must be a valid email
		validation.Field(&u.Email, validation.Required, is.Email),
		// LastName cannot be empty, and contain only letter
		validation.Field(&u.LastName, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z\\s]*$"))),
		// FirstName cannot be empty, and contain only letter
		validation.Field(&u.FirstName, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z\\s]*$"))),
		// Phone cannot be empty, and must be a valid phone number
		validation.Field(&u.Phone, validation.Required, validation.Match(regexp.MustCompile("^(([+]?\\d{2})|\\d?)[\\s-]?[0-9]{2}[\\s-]?[0-9]{3}[\\s-]?[0-9]{4}$"))),
	)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

func (u UserLogin) Validate() error {
	return validation.ValidateStruct(&u,
		// Email cannot be empty, and must be a valid email
		validation.Field(&u.Email, validation.Required, is.Email),
		// Password cannot be empty
		validation.Field(&u.Password, validation.Required),
	)
}

type UserUpdate struct {
	common.SQLModel `json:",inline"`
	UserName        string        `json:"user_name" gorm:"column:user_name;"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"password" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	DateOfBirth     *time.Time    `json:"date_of_birth" gorm:"column:date_of_birth"`
	Phone           *string       `json:"phone" gorm:"column:phone;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

func (u *UserUpdate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)

	ErrUserNameExisted = common.NewCustomError(
		errors.New("username has already existed"),
		"username has already existed",
		"ErrUserNameExisted",
	)
)
