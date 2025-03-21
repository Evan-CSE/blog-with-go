package router

import (
	"blog-api/handler"
	"blog-api/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(blogService service.IBlogService) *gin.Engine {
	r := gin.Default()
	handler := handler.NewBlogHandler(blogService)
	r.POST("/blog", handler.CreateBlog)
	r.GET("/blog", handler.GetAllBlogs)
	r.GET("/blog/:id", handler.GetBlogById)
	r.PUT("/blog/:id", handler.UpdateBlog)
	r.DELETE("/blog/:id", handler.DeleteBlog)
	return r
}
