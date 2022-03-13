package ginpost

import (
	"instago2/common"
	"instago2/component"
	"instago2/modules/post/postbiz"
	"instago2/modules/post/poststorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListPost godoc
// @Summary      Get all post
// @Description  Get all post
// @Tags         posts
// @Accept       json
// @Produce 	 json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /posts/explore [GET]
func ListPost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter common.Paging

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewListPostBiz(store)

		result, err := biz.ListPost(c.Request.Context(), &paging)

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
