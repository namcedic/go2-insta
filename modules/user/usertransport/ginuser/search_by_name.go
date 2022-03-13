package ginuser

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/user/userbiz"
	"instago2/modules/user/userstorage"
	"net/http"
)

// SearchUserByName godoc
// @Summary      SearchUserByName
// @Description  Search other user by username, first name or last name
// @Tags         users
// @Accept       json
// @Produce 	 json
// @Param 		 Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        searchKey   path      string  true "username, lastname or firstname of other user that current user want to search"
// @Success      200  {object}  []common.SimpleUser
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /users/search/{searchKey} [GET]
func SearchUserByName(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()
		searchKey := c.Param("searchKey")
		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.SearchUserByNameBiz(store)

		data, err := biz.SearchUser(c.Request.Context(), &paging, searchKey)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
			if i == len(data)-1 {
				paging.NextCursor = data[i].FakeId.String()
			}
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))

	}
}
