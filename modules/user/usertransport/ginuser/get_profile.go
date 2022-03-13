package ginuser

import (
	"instago2/common"
	"instago2/component"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile godoc
// @Summary      GetProfile
// @Description  Get profile of current user
// @Tags         users
// @Accept       json
// @Produce 	 json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success      200  {object}  usermodel.User
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /users/profile [GET]
func GetProfile(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		//data := c.MustGet(common.CurrentUser)
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
