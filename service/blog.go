package service

import (
	"context"
	"errors"
	"go-blogrpl/dto"
	"go-blogrpl/entity"
	"go-blogrpl/repository"
	"reflect"

	"github.com/jinzhu/copier"
)

type blogService struct {
	blogRepository repository.BlogRepository
}

type BlogService interface {
	GetAllBlogs(ctx context.Context) ([]entity.Blog, error)
	GetBlogBySlug(ctx context.Context, slug string) (entity.Blog, error)
	CreateNewBlog(ctx context.Context, blogDTO dto.BlogPostRequest, userId uint64) (entity.Blog, error)
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

func (blogS *blogService) GetBlogBySlug(ctx context.Context, slug string) (entity.Blog, error) {
	blog, err := blogS.blogRepository.GetBlogBySlug(ctx, nil, slug)
	if err != nil {
		return entity.Blog{}, err
	}
	return blog, nil
}

func (blogS *blogService) CreateNewBlog(ctx context.Context, blogDTO dto.BlogPostRequest, userId uint64) (entity.Blog, error) {
	// Copy BlogDTO to empty newly created blog var
	var blog entity.Blog
	blogDTO.UserID = userId

	copier.Copy(&blog, &blogDTO)

	// Check for duplicate Blog slug
	blogCheck, err := blogS.blogRepository.GetBlogBySlug(ctx, nil, blogDTO.Slug)
	if err != nil {
		return entity.Blog{}, err
	}

	// Check if duplicate is found
	if !(reflect.DeepEqual(blogCheck, entity.Blog{})) {
		return entity.Blog{}, errors.New("slug already used")
	}

	// create new blog
	newBlog, err := blogS.blogRepository.CreateNewBlog(ctx, nil, blog)
	if err != nil {
		return entity.Blog{}, err
	}
	return newBlog, nil
}
