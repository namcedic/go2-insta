package ginpost

import (
	"instago2/common"
	"instago2/component"
	"instago2/modules/post/postbiz"
	"instago2/modules/post/poststorage"
	"instago2/modules/userfollow/userfollowstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListFollowingPost godoc
// @Summary      Get all posts of following
// @Description  Get all posts of following
// @Tags         posts
// @Accept       json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /posts/explore/following [get]
func ListFollowingPost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		var ids []int
		currentId := c.MustGet(common.CurrentUser).(common.Requester).GetUserId()

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		liStore := userfollowstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewListFollowingPostBiz(store, liStore)

		result, err := biz.ListFollowingPost(c.Request.Context(), ids, currentId, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
