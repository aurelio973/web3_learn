package main

import (
	"blog/handler"
	"blog/middleware"
	"blog/model"

	"github.com/gin-gonic/gin"
)

var db = model.InitDB()

func main() {
	r := gin.Default()

	// 公共路由
	public := r.Group("/api")
	{
		public.POST("/register", handler.Register)
		public.POST("/login", handler.Login)
		public.GET("/posts", handler.GetAllPosts)
		public.GET("/posts/:id", handler.GetPost)
		public.GET("/posts/:id/comments", handler.GetCommentsByPost)
	}

	// 需要认证的路由
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/posts", handler.CreatePost)
		protected.PUT("/posts/:id", handler.UpdatePost)
		protected.DELETE("/posts/:id", handler.DeletePost)
		protected.POST("/posts/:id/comments", handler.CreateComment)
	}

	r.Run(":8080")
}
