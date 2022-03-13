package postliketransport

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/post/poststorage"
	"instago2/modules/postlike/postlikebusiness"
	"instago2/modules/postlike/postlikemodel"
	"instago2/modules/postlike/postlikestorage"
	"net/http"
)

// CreatePostLikes godoc
// @Summary      Create a postlike
// @Description  Create by post ID
// @Tags         posts
// @Accept       json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true "id of the post that user want to like (encoded in uuid)"
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /posts/{id}/like [post]
func CreatePostLikes(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		pid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := postlikemodel.PostLikes{
			PostId: int(pid.GetLocalID()),
			UserId: requester.GetUserId(),
		}

		store := postlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postlikebusiness.NewCreatePostLikesBiz(store, incStore)

		if err := biz.CreatePostLikes(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
