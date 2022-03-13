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

// UpdateProfile godoc
// @Summary      UpdateProfile
// @Description  Users update their own account information
// @Tags         users
// @Accept       json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        user   body    usermodel.UserUpdate  true "information that user want to update"
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /users [PATCH]
func UpdateProfile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data usermodel.UserUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewUpdateProfileBiz(store)

		if err := biz.UpdateProfile(c.Request.Context(), requester.GetUserId(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
