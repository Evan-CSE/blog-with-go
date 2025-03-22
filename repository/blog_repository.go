package repository

import (
	"blog-api/model"

	"gorm.io/gorm"
)

type IBlogRepository interface {
	CreateBlog(blog *model.Blog) model.Result[model.Blog]
	GetAllBlogs() model.Result[[]model.Blog]
	GetBlogById(id int) model.Result[model.Blog]
	UpdateBlog(blog *model.Blog) model.Result[model.Blog]
	DeleteBlog(blog *model.Blog) model.Result[model.Blog]
}

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) IBlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) CreateBlog(blog *model.Blog) model.Result[model.Blog] {
	if err := r.db.Create(blog).Error; err != nil {
		return *model.Failure[model.Blog](err.Error(), "")
	}
	return *model.Success(*blog)
}

func (r *BlogRepository) GetAllBlogs() model.Result[[]model.Blog] {
	var blogs []model.Blog
	if err := r.db.Find(&blogs).Error; err != nil {
		return *model.Failure[[]model.Blog](err.Error(), "")
	}
	return *model.Success(blogs)
}

func (r *BlogRepository) GetBlogById(id int) model.Result[model.Blog] {
	var blog model.Blog
	if err := r.db.First(&blog, id).Error; err != nil {
		return *model.Failure[model.Blog](err.Error(), "")
	}
	return *model.Success(blog)
}

func (r *BlogRepository) UpdateBlog(blog *model.Blog) model.Result[model.Blog] {
	if err := r.db.Save(blog).Error; err != nil {
		return *model.Failure[model.Blog](err.Error(), "")
	}
	return *model.Success(*blog)
}

func (r *BlogRepository) DeleteBlog(blog *model.Blog) model.Result[model.Blog] {
	if err := r.db.Delete(blog).Error; err != nil {
		return *model.Failure[model.Blog](err.Error(), "")
	}
	return *model.Success(*blog)
}
