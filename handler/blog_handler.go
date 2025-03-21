package handler

import (
	"blog-api/model"
	"blog-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	blogService service.IBlogService
}

func NewBlogHandler(blogService service.IBlogService) *BlogHandler {
	return &BlogHandler{blogService: blogService}
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var newBlog model.Blog

	if err := h.blogService.CreateBlog(&newBlog); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "Blog created successfully"})
	}
}

func (h *BlogHandler) GetAllBlogs(c *gin.Context) {
	if blogs, err := h.blogService.GetAllBlogs(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"blogs": blogs})
	}
}

func (h *BlogHandler) GetBlogById(c *gin.Context) {

	id, error := strconv.Atoi(c.Param("id"))

	if error != nil {
		c.JSON(500, gin.H{"error": error.Error()})
		return
	}

	if blog, err := h.blogService.GetBlogById(id); err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"blog": blog})
	}
}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	var blog model.Blog

	err := c.BindJSON(&blog)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, error := strconv.Atoi(c.Param("id"))

	if error != nil {
		c.JSON(500, gin.H{"error": error.Error()})
		return
	}

	blog.Id = id

	if err := h.blogService.UpdateBlog(&blog); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "Blog updated successfully"})
	}
}

func (h *BlogHandler) DeleteBlog(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))

	if error != nil {
		c.JSON(500, gin.H{"error": error.Error()})
		return
	}

	if err := h.blogService.DeleteBlog(&model.Blog{Id: id}); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "Blog deleted successfully"})
	}
}
