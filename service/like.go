package service

import (
	"context"
	"go-blogrpl/entity"
	"go-blogrpl/repository"
	"reflect"
)

type likeService struct {
	likeRepository repository.LikeRepository
	blogRepository repository.BlogRepository
}

type LikeService interface {
	// BlogLikes
	GetAllBlogLikes(ctx context.Context) ([]entity.BlogLike, error)
	ChangeLikeForBlog(ctx context.Context, blogId uint64, userId uint64) (string, error)

	// CommentLikes
	GetAllCommentLikes(ctx context.Context) ([]entity.CommentLike, error)
}

func NewLikeService(likeR repository.LikeRepository, blogR repository.BlogRepository) LikeService {
	return &likeService{likeRepository: likeR, blogRepository: blogR}
}

func (likeS *likeService) GetAllBlogLikes(ctx context.Context) ([]entity.BlogLike, error) {
	blikes, err := likeS.likeRepository.GetAllBlogLikes(ctx, nil)
	if err != nil {
		return []entity.BlogLike{}, err
	}
	return blikes, nil
}

func (likeS *likeService) ChangeLikeForBlog(ctx context.Context, blogId uint64, userId uint64) (string, error) {
	bl := entity.BlogLike{BlogID: blogId, UserID: userId}

	status, blRes, err := likeS.likeRepository.CheckBlogLike(ctx, nil, bl, blogId, userId)
	if err != nil {
		return "failed posting like to blog", err
	}

	switch status {
	case 1:
		_, err = likeS.likeRepository.CreateNewBlogLike(ctx, nil, bl)
	case 2:
		check, _ := likeS.likeRepository.GetBlogLikeByID(ctx, nil, blRes.ID)

		if reflect.DeepEqual(check, entity.BlogLike{}) {
			_, err = likeS.likeRepository.RestoreBlogLike(ctx, nil, blRes)
		} else {
			err = likeS.likeRepository.DeleteBlogLike(ctx, nil, blRes.ID)
			if err != nil {
				return err.Error(), err
			}

			blog, err := likeS.blogRepository.GetBlogByID(ctx, nil, blogId)
			if err != nil {
				return err.Error(), err
			}

			err = likeS.likeRepository.SetBlogLikeCount(ctx, nil, blog)
			if err != nil {
				return err.Error(), err
			}

			return "successfully unliked blog", nil
		}
	default:
		return "failed posting like to blog", err
	}

	if err != nil {
		return err.Error(), err
	}

	blog, err := likeS.blogRepository.GetBlogByID(ctx, nil, blogId)
	if err != nil {
		return err.Error(), err
	}

	err = likeS.likeRepository.SetBlogLikeCount(ctx, nil, blog)
	if err != nil {
		return err.Error(), err
	}

	return "successfully liked blog", nil
}

func (likeS *likeService) GetAllCommentLikes(ctx context.Context) ([]entity.CommentLike, error) {
	clikes, err := likeS.likeRepository.GetAllCommentLikes(ctx, nil)
	if err != nil {
		return []entity.CommentLike{}, err
	}
	return clikes, nil
}
