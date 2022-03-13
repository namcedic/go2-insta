package ginuserfollow

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/userfollow/userfollowbiz"
	"instago2/modules/userfollow/userfollowmodel"
	"instago2/modules/userfollow/userfollowstorage"
	"net/http"
)

// ListFollower godoc
// @Summary      ListFollower
// @Description  List follower of current user
// @Tags         users
// @Accept       json
// @Produce 	 json
// @Param 		 Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        name   query      string  false "username, lastname or firstname of other user that current user want to search in list follower"
// @Param        limit   query      string  false "limit records return in one page"
// @Param        cursor   query      string  false "used for paging purpose"
// @Success      200  {object}  []common.SimpleUser
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /users/follower [GET]
func ListFollower(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		userId := requester.GetUserId()
		var filter userfollowmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := userfollowstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := userfollowbiz.NewListFollowerBiz(store)
		result, err := biz.ListFollower(c.Request.Context(), userId, &filter, &paging)
		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
