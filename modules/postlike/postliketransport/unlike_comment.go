package postliketransport

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/post/poststorage"
	"instago2/modules/postlike/postlikebusiness"
	"instago2/modules/postlike/postlikestorage"
	"net/http"
)

// UnlikePost godoc
// @Summary      Delete a postlike
// @Description  Delete by post ID
// @Tags         posts
// @Accept       json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true "id of the post that user want to unlike (encoded in uuid)"
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /posts/{id}/unlike [delete]
func UnlikePost(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := postlikestorage.NewSQLStore(ctx.GetMainDBConnection())
		decStore := poststorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := postlikebusiness.NewUnlikePostBiz(store, decStore)

		if err := biz.UnlikePost(c.Request.Context(), int(uid.GetLocalID()), requester.GetUserId()); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
