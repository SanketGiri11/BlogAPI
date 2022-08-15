package main

import (
	"blogapi/controllers"
	"blogapi/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Start the application")
	//Database connection
	database.Connect()

	router := initRouter()
	router.Run(":9500")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	//Group all the api endpoint under blog path param
	api := router.Group("/blog")
	{
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/token", controllers.GenerateToken)
		api.POST("/createPost", controllers.CreatePosts)
		api.GET("/getPosts", controllers.AllPosts)
		api.GET("/getPosts/:id", controllers.DetailPosts)
		api.PUT("/updatePost/:id", controllers.UpdatePosts)
		api.GET("/uniquePosts", controllers.UniquePosts)
		api.DELETE("/deletePost/:id", controllers.DeletePosts)
		api.POST("/addComments", controllers.AddComments)
		api.GET("/allComments", controllers.AllComments)
		api.GET("/getComments/:id", controllers.DetailComments)
		api.PUT("/updatecomment/:id", controllers.UpdateComments)
		api.DELETE("/deleteComment/:id", controllers.DeleteCommnets)
		api.POST("/addTags", controllers.AddTags)
		api.GET("/allTags", controllers.AllTags)
		api.GET("/getTags/:id", controllers.DetailTags)
		api.PUT("/updateTags/:id", controllers.UpdateTags)
		api.DELETE("/deleteTag/:id", controllers.DeleteTags)
	}

	return router
}
