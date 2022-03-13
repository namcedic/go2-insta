package commentliketransport

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/comment/commentstorage"
	"instago2/modules/commentlike/commentlikebusiness"
	"instago2/modules/commentlike/commentlikestorage"
	"net/http"
)

// UnlikeComment godoc
// @Summary      Delete a commentlike
// @Description  Delete by comment ID
// @Tags         comments
// @Accept       json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true "id of the comment that user want to unlike (encoded in uuid)"
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /comments/{id}/unlike [delete]
func UnlikeComment(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := commentlikestorage.NewSQLStore(ctx.GetMainDBConnection())
		decStore := commentstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := commentlikebusiness.NewUnlikeCommentBiz(store, decStore)

		if err := biz.UnlikeComment(c.Request.Context(), int(uid.GetLocalID()), requester.GetUserId()); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
