package ginpost

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/post/postbiz"
	"instago2/modules/post/postmodel"
	"instago2/modules/post/poststorage"
	"net/http"
)

// UpdatePost godoc
// @Summary      UpdatePost
// @Description  User updates post
// @Tags         posts
// @Accept       json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      int  true "id of the post that user want to delete (encoded in uuid)"
// @Param        post   body    postmodel.PostUpdate  true "information of the post that user want to edit"
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /posts/{id} [PATCH]
func UpdatePost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		//uid, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		userId := requester.GetUserId()
		var data postmodel.PostUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewUpdatePostBiz(store)

		if err := biz.UpdatePost(c.Request.Context(), int(uid.GetLocalID()), userId, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
