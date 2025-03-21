package repository

import (
	"blog-api/model"

	"gorm.io/gorm"
)

type IBlogRepository interface {
	CreateBlog(blog *model.Blog) error
	GetAllBlogs() ([]model.Blog, error)
	GetBlogById(id int) (model.Blog, error)
	UpdateBlog(blog *model.Blog) error
	DeleteBlog(blog *model.Blog) error
}

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) IBlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) CreateBlog(blog *model.Blog) error {
	return r.db.Create(blog).Error
}

func (r *BlogRepository) GetAllBlogs() ([]model.Blog, error) {
	var blogs []model.Blog
	err := r.db.Find(&blogs).Error
	return blogs, err
}

func (r *BlogRepository) GetBlogById(id int) (model.Blog, error) {
	var blog model.Blog
	err := r.db.First(&blog, id).Error
	return blog, err
}

func (r *BlogRepository) UpdateBlog(blog *model.Blog) error {
	return r.db.Save(blog).Error
}

func (r *BlogRepository) DeleteBlog(blog *model.Blog) error {
	return r.db.Delete(blog).Error
}
