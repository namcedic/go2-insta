package main

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"instago2/common"
	"instago2/component"
	"instago2/component/uploadprovider"
	_ "instago2/docs"
	"instago2/middleware"
	"instago2/modules/comment/transport/gincomment"
	"instago2/modules/comment/transport/ginreplycomment"
	"instago2/modules/commentlike/commentliketransport"
	"instago2/modules/post/posttransport/ginpost"
	"instago2/modules/postlike/postliketransport"
	"instago2/modules/postsearch/postsearchtransport/ginpostsearch"
	"instago2/modules/upload/uploadtransport/ginupload"
	"instago2/modules/user/usertransport/ginuser"
	"instago2/modules/userfollow/userfollowtransport/ginuserfollow"
	"instago2/pubsub/pblocal"
	"instago2/subscriber"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title           Swagger Insta API
// @version         1.0
// @description     This is a images social network.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth

func goDotEnvVariable(key string) string {

	err := godotenv.Load("md.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	dsn := goDotEnvVariable("DBCONNECTIONSTR")

	s3BucketName := goDotEnvVariable("S3BUCKETNAME")
	s3Region := goDotEnvVariable("S3REGION")
	s3APIKey := goDotEnvVariable("S3APIKEY")
	s3SecretKey := goDotEnvVariable("S3SECRETKEY")
	s3Domain := goDotEnvVariable("S3DOMAIN")
	secretKey := goDotEnvVariable("LOGINSECRETKEY")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()
	if err := runService(db, s3Provider, secretKey); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {

	appCtx := component.NewAppContext(db, upProvider, secretKey, pblocal.NewPubSub())

	if err := subscriber.NewEngine(appCtx).Start(); err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "true",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	v1.POST("/upload", middleware.RequiredAuth(appCtx), ginupload.Upload(appCtx))
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))

	users := v1.Group("/users", middleware.RequiredAuth(appCtx))
	{
		// get profile of current user
		users.GET("/profile", ginuser.GetProfile(appCtx))
		// update profile of current user
		users.PATCH("", ginuser.UpdateProfile(appCtx))

		// get profile of other user
		users.GET("/profile/:id", ginuser.GetOtherProfile(appCtx))
		// follow
		users.POST("/:id/follow", ginuserfollow.UserFollowUser(appCtx))
		// unfollow
		users.DELETE("/:id/unfollow", ginuserfollow.UserUnfollowUser(appCtx))
		// list followers
		users.GET("/follower", ginuserfollow.ListFollower(appCtx))
		// list following
		users.GET("/following", ginuserfollow.ListFollowing(appCtx))
		// search by name
		users.GET("/search/:searchKey", ginuser.SearchUserByName(appCtx))
	}

	posts := v1.Group("/posts", middleware.RequiredAuth(appCtx))
	{
		//Create post
		posts.POST("", ginpost.CreatePost(appCtx))
		//Update post
		posts.PATCH("/:id", ginpost.UpdatePost(appCtx))
		// create reply a comment
		posts.POST("/comments/replies", ginreplycomment.CreateReply(appCtx))

		// List comments of one post
		posts.GET("/:id/comments", ginreplycomment.ListComment(appCtx))
		// list all posts
		posts.GET("/explore", ginpost.ListPost(appCtx))
		// Get post
		posts.GET("/explore/:id", ginpost.GetPost(appCtx))
		// Get posts of all following list
		posts.GET("/explore/following", ginpost.ListFollowingPost(appCtx))
		// Like post
		posts.POST("/:id/like", postliketransport.CreatePostLikes(appCtx))
		// Unlike post
		posts.DELETE("/:id/unlike", postliketransport.UnlikePost(appCtx))
		// Comment post
		posts.POST("/:id/comment", gincomment.CreateComment(appCtx))
		// Delete post
		posts.DELETE("/:id", ginpost.DeletePost(appCtx))
		// Search posts by caption, search name and time
		posts.POST("/searches", ginpostsearch.ListPostSearch(appCtx))
	}

	comments := v1.Group("comments", middleware.RequiredAuth(appCtx))
	{
		// Like comment
		comments.POST("/:id/like", commentliketransport.CreateCommentLikes(appCtx))
		//Unlike comment
		comments.DELETE("/:id/unlike", commentliketransport.UnlikeComment(appCtx))
		//Delete comments
		comments.DELETE("/:id/:post_id", gincomment.DeleteComment(appCtx))

	}

	v1.GET("/encode-uid", func(c *gin.Context) {
		type reqData struct {
			DbType int `form:"type"`
			RealId int `form:"id"`
		}

		var d reqData
		c.ShouldBind(&d)

		c.JSON(http.StatusOK, gin.H{
			"id": common.NewUID(uint32(d.RealId), d.DbType, 1),
		})
	})

	return r.Run()
}
