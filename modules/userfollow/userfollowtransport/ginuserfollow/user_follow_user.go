package ginuserfollow

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/user/userstorage"
	"instago2/modules/userfollow/userfollowbiz"
	"instago2/modules/userfollow/userfollowmodel"
	"instago2/modules/userfollow/userfollowstorage"
	"net/http"
)

// UserFollowUser godoc
// @Summary      UserFollowUser
// @Description  User follow another user
// @Tags         users
// @Accept       json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true "user id of other user that current user want to follow (encoded to uuid)"
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /users/{id}/follow [POST]
func UserFollowUser(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data := userfollowmodel.Follow{
			FollowerId: requester.GetUserId(),
			UserId:     int(uid.GetLocalID()),
		}
		store := userfollowstorage.NewSQLStore(ctx.GetMainDBConnection())
		incStore := userstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := userfollowbiz.NewUserFollowBiz(store, incStore)
		if err := biz.UserFollowUser(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
