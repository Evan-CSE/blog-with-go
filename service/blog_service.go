package service

import (
	"blog-api/model"
	"blog-api/repository"
)

type IBlogService interface {
	CreateBlog(blog *model.Blog) model.Result[model.Blog]
	GetAllBlogs() model.Result[[]model.Blog]
	GetBlogById(id int) model.Result[model.Blog]
	UpdateBlog(blog *model.Blog) model.Result[model.Blog]
	DeleteBlog(blog *model.Blog) model.Result[model.Blog]
}

func NewBlogService(repo repository.IBlogRepository) IBlogService {
	return &BlogService{repo: repo}
}

type BlogService struct {
	repo repository.IBlogRepository
}

func (s *BlogService) CreateBlog(blog *model.Blog) model.Result[model.Blog] {
	result := s.repo.CreateBlog(blog)
	return result
}

func (s *BlogService) GetAllBlogs() model.Result[[]model.Blog] {
	result := s.repo.GetAllBlogs()
	return result
}

func (s *BlogService) GetBlogById(id int) model.Result[model.Blog] {
	result := s.repo.GetBlogById(id)
	return result
}

func (s *BlogService) UpdateBlog(blog *model.Blog) model.Result[model.Blog] {
	result := s.repo.UpdateBlog(blog)
	return result
}

func (s *BlogService) DeleteBlog(blog *model.Blog) model.Result[model.Blog] {
	result := s.repo.DeleteBlog(blog)
	return result
}
