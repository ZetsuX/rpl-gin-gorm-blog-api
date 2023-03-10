package service

import (
	"context"
	"go-blogrpl/entity"
	"go-blogrpl/repository"
)

type likeService struct {
	likeRepository repository.LikeRepository
}

type LikeService interface {
	// BlogLikes
	GetAllBlogLikes(ctx context.Context) ([]entity.BlogLike, error)

	// CommentLikes
}

func NewLikeService(likeR repository.LikeRepository) LikeService {
	return &likeService{likeRepository: likeR}
}

func (likeS *likeService) GetAllBlogLikes(ctx context.Context) ([]entity.BlogLike, error) {
	blikes, err := likeS.likeRepository.GetAllBlogLikes(ctx, nil)
	if err != nil {
		return []entity.BlogLike{}, err
	}
	return blikes, nil
}
