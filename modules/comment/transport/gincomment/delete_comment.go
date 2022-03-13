package gincomment

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/comment/commentbiz"
	"instago2/modules/comment/commentmodel"
	"instago2/modules/comment/commentstorage"
	"net/http"
)

// DeleteComment godoc
// @Summary      Delete a comment
// @Description  Delete a comment by comment ID
// @Tags         comments
// @Accept       json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true "id of the comment that user want to delete (encoded in uuid)"
// @Param        post_id   path      string  true "post_id of the comment that user want to delete (encoded in uuid)"
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /comments/{id}/{post_id} [delete]
func DeleteComment(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		pid, err := common.FromBase58(c.Param("post_id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data := commentmodel.CommentDelete{
			CommentId: int(uid.GetLocalID()),
			PostId:    int(pid.GetLocalID()),
			UserId:    requester.GetUserId(),
		}
		store := commentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := commentbiz.NewDeleteCommentBiz(store, appCtx.GetPubsub())

		if err := biz.DeleteComment(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
