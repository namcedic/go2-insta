package gincomment

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/comment/commentbiz"
	"instago2/modules/comment/commentmodel"
	"instago2/modules/comment/commentstorage"
	"instago2/modules/post/poststorage"
	"net/http"
)

// CreateComment godoc
// @Summary      Create a comment
// @Description  Create by post ID
// @Tags         posts
// @Accept       json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true "id of the post that user want to comment (encoded in uuid)"
// @Param        content   body    commentmodel.CommentCreate  true "content of comment"
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /posts/{id}/comment [post]
func CreateComment(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		pid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var content commentmodel.CommentCreate
		if err := c.ShouldBind(&content); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data := commentmodel.CommentCreate{
			PostId:  int(pid.GetLocalID()),
			UserId:  requester.GetUserId(),
			Content: content.Content,
		}

		store := commentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := commentbiz.NewCreateCommentBiz(store, incStore)

		if err := biz.CreateComment(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
