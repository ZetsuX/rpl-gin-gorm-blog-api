package service

import (
	"context"
	"go-blogrpl/entity"
	"go-blogrpl/repository"
	"reflect"
)

type likeService struct {
	likeRepository    repository.LikeRepository
	blogRepository    repository.BlogRepository
	commentRepository repository.CommentRepository
}

type LikeService interface {
	// BlogLikes
	GetAllBlogLikes(ctx context.Context) ([]entity.BlogLike, error)
	ChangeLikeForBlog(ctx context.Context, blogId uint64, userId uint64) (string, error)

	// CommentLikes
	GetAllCommentLikes(ctx context.Context) ([]entity.CommentLike, error)
	ChangeLikeForComment(ctx context.Context, commentId uint64, userId uint64) (string, error)
}

func NewLikeService(likeR repository.LikeRepository, blogR repository.BlogRepository, commentR repository.CommentRepository) LikeService {
	return &likeService{likeRepository: likeR, blogRepository: blogR, commentRepository: commentR}
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

func (likeS *likeService) ChangeLikeForComment(ctx context.Context, commentId uint64, userId uint64) (string, error) {
	bl := entity.CommentLike{CommentID: commentId, UserID: userId}

	status, blRes, err := likeS.likeRepository.CheckCommentLike(ctx, nil, bl, commentId, userId)
	if err != nil {
		return "failed posting like to comment", err
	}

	switch status {
	case 1:
		_, err = likeS.likeRepository.CreateNewCommentLike(ctx, nil, bl)
	case 2:
		check, _ := likeS.likeRepository.GetCommentLikeByID(ctx, nil, blRes.ID)

		if reflect.DeepEqual(check, entity.CommentLike{}) {
			_, err = likeS.likeRepository.RestoreCommentLike(ctx, nil, blRes)
		} else {
			err = likeS.likeRepository.DeleteCommentLike(ctx, nil, blRes.ID)
			if err != nil {
				return err.Error(), err
			}

			comment, err := likeS.commentRepository.GetCommentByID(ctx, nil, commentId)
			if err != nil {
				return err.Error(), err
			}

			err = likeS.likeRepository.SetCommentLikeCount(ctx, nil, comment)
			if err != nil {
				return err.Error(), err
			}

			return "successfully unliked comment", nil
		}
	default:
		return "failed posting like to comment", err
	}

	if err != nil {
		return err.Error(), err
	}

	comment, err := likeS.commentRepository.GetCommentByID(ctx, nil, commentId)
	if err != nil {
		return err.Error(), err
	}

	err = likeS.likeRepository.SetCommentLikeCount(ctx, nil, comment)
	if err != nil {
		return err.Error(), err
	}

	return "successfully liked comment", nil
}
