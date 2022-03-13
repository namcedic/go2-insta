package ginuser

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/user/userbiz"
	"instago2/modules/user/usermodel"
	"instago2/modules/user/userstorage"
	"net/http"
)

// GetOtherProfile godoc
// @Summary      GetOtherProfile
// @Description  Get profile of another user by id
// @Tags         users
// @Accept       json
// @Produce 	 json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true "user id of other user that current user want to see profile (encoded to uuid)"
// @Success      200  {object}  common.SimpleUser
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /users/profile/{id} [GET]
func GetOtherProfile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrEntityNotFound(usermodel.EntityName, err))
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewGetOtherProfileBiz(store)

		data, err := biz.GetProfile(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
