package service

import (
	"blog-api/model"
	"blog-api/repository"
)

type IBlogService interface {
	CreateBlog(blog *model.Blog) error
	GetAllBlogs() ([]model.Blog, error)
	GetBlogById(id int) (model.Blog, error)
	UpdateBlog(blog *model.Blog) error
	DeleteBlog(blog *model.Blog) error
}

func NewBlogService(repo repository.IBlogRepository) IBlogService {
	return &BlogService{repo: repo}
}

type BlogService struct {
	repo repository.IBlogRepository
}

func (s *BlogService) CreateBlog(blog *model.Blog) error {
	return s.repo.CreateBlog(blog)
}

func (s *BlogService) GetAllBlogs() ([]model.Blog, error) {
	return s.repo.GetAllBlogs()
}

func (s *BlogService) GetBlogById(id int) (model.Blog, error) {
	return s.repo.GetBlogById(id)
}

func (s *BlogService) UpdateBlog(blog *model.Blog) error {
	return s.repo.UpdateBlog(blog)
}

func (s *BlogService) DeleteBlog(blog *model.Blog) error {
	return s.repo.DeleteBlog(blog)
}
