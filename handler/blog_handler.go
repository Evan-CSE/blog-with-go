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

	if result := h.blogService.CreateBlog(&newBlog); !result.IsSuccessful {
		c.JSON(500, result)
	} else {
		c.JSON(200, result)
	}
}

func (h *BlogHandler) GetAllBlogs(c *gin.Context) {
	if result := h.blogService.GetAllBlogs(); !result.IsSuccessful {
		c.JSON(500, result)
	} else {
		c.JSON(200, result)
	}
}

func (h *BlogHandler) GetBlogById(c *gin.Context) {

	id, error := strconv.Atoi(c.Param("id"))

	if error != nil {
		c.JSON(500, gin.H{"error": error.Error()})
		return
	}

	if result := h.blogService.GetBlogById(id); !result.IsSuccessful {
		c.JSON(404, result)
	} else {
		c.JSON(200, result)
	}
}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	var blog model.Blog

	err := c.BindJSON(&blog)

	if err != nil {
		c.JSON(400, model.Failure[model.Blog](err.Error(), ""))
		return
	}

	id, error := strconv.Atoi(c.Param("id"))

	if error != nil {
		c.JSON(500, model.Failure[model.Blog](error.Error(), ""))
		return
	}

	blog.Id = id

	if result := h.blogService.UpdateBlog(&blog); !result.IsSuccessful {
		c.JSON(500, result)
	} else {
		c.JSON(200, result)
	}
}

func (h *BlogHandler) DeleteBlog(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))

	if error != nil {
		c.JSON(500, model.Failure[model.Blog](error.Error(), ""))
		return
	}

	if result := h.blogService.DeleteBlog(&model.Blog{Id: id}); !result.IsSuccessful {
		c.JSON(500, result)
	} else {
		c.JSON(200, result)
	}
}
