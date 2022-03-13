package ginuser

import (
	"instago2/common"
	"instago2/component"
	"instago2/component/hasher"
	"instago2/component/tokenprovider/jwt"
	"instago2/modules/user/userbiz"
	"instago2/modules/user/usermodel"
	"instago2/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary      Login
// @Description  User login into account
// @Tags         users
// @Accept       json
// @Param        user   body    usermodel.UserLogin  true "account's information of user that want to log in"
// @Success      200  {object}  tokenprovider.Token
// @Failure      400  {object}  common.AppError
// @Failure      404  {object}  common.AppError
// @Failure      500  {object}  common.AppError
// @Router       /login [POST]
func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*7)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
