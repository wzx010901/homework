package routes

import (
	"blog-system/handlers"
	"blog-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			//注册
			auth.POST("/register", handlers.Register)
			//登录
			auth.POST("/login", handlers.Login)
		}

		posts := api.Group("/posts")
		posts.Use(middleware.AuthMiddleware())
		{
			//获取文章列表
			posts.GET("", handlers.GetPosts)
			//获取文章
			posts.GET("/:id", handlers.GetPost)
			//创建文章
			posts.POST("", handlers.CreatePost)
			//更新文章
			posts.PUT("/:id", handlers.UpdatePost)
			//删除文章
			posts.DELETE("/:id", handlers.DeletePost)
		}

		comments := api.Group("/comments")
		comments.Use(middleware.AuthMiddleware())
		{
			//获取某个文章的评论
			comments.GET("/post/:postID", handlers.GetCommentsByPost)
			//创建评论
			comments.POST("", handlers.CreateComment)
		}

		users := api.Group("/users")
		users.Use(middleware.AuthMiddleware())
		{
			//获取当前登录用户信息
			users.GET("/getCurrentUser", handlers.GetCurrentUser)
		}
	}

	return router
}
