package service

import (
	"context"
	"go-blogrpl/entity"
	"go-blogrpl/repository"
)

type blogService struct {
	blogRepository repository.BlogRepository
}

type BlogService interface {
	GetAllBlogs(ctx context.Context) ([]entity.Blog, error)
}

func NewBlogService(blogR repository.BlogRepository) BlogService {
	return &blogService{blogRepository: blogR}
}

func (blogS *blogService) GetAllBlogs(ctx context.Context) ([]entity.Blog, error) {
	blogs, err := blogS.blogRepository.GetAllBlogs(ctx, nil)
	if err != nil {
		return []entity.Blog{}, err
	}
	return blogs, nil
}
