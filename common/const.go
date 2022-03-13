package common

const (
	DbTypePost    = 1
	DbTypeComment = 2
	DbTypeUser    = 3
)
const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
