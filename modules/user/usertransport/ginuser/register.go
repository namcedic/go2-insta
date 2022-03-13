package ginuser

import (
	"instago2/common"
	"instago2/component"
	"instago2/component/hasher"
	"instago2/modules/user/userbiz"
	"instago2/modules/user/usermodel"
	"instago2/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary      Register
// @Description  User create new account
// @Tags         users
// @Accept       json
// @Param        user   body    usermodel.UserCreate  true "information of user that need to create an account"
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /register [POST]
func Register(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
