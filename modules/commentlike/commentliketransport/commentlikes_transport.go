package commentliketransport

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/comment/commentstorage"
	"instago2/modules/commentlike/commentlikebusiness"
	"instago2/modules/commentlike/commentlikemodel"
	"instago2/modules/commentlike/commentlikestorage"
	"net/http"
)

// CreateCommentLikes godoc
// @Summary      Create a commentlike
// @Description  Create by comment ID
// @Tags         comments
// @Accept       json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true "id of the comment that user want to like (encoded in uuid)"
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /comments/{id}/like [post]
func CreateCommentLikes(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		cid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := commentlikemodel.CommentLikes{
			CommentId: int(cid.GetLocalID()),
			UserId:    requester.GetUserId(),
		}

		store := commentlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := commentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := commentlikebusiness.NewCreateCommentLikesBiz(store, incStore)

		if err := biz.CreateCommentLikes(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
