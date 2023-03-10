package service

import (
	"context"
	"go-blogrpl/dto"
	"go-blogrpl/entity"
	"go-blogrpl/repository"

	"github.com/jinzhu/copier"
)

type commentService struct {
	commentRepository repository.CommentRepository
}

type CommentService interface {
	// // Comments
	// GetAllComments(ctx context.Context) ([]entity.Comment, error)

	// BlogComments
	CreateNewBlogComment(ctx context.Context, bcDTO dto.CommentRequest, blogId uint64, userId uint64) (entity.Comment, error)
}

func NewCommentService(commentR repository.CommentRepository) CommentService {
	return &commentService{commentRepository: commentR}
}

func (commentS *commentService) CreateNewBlogComment(ctx context.Context, bcDTO dto.CommentRequest, blogId uint64, userId uint64) (entity.Comment, error) {
	// Copy CommentDTO to empty newly created comment var
	var bc entity.Comment
	bcDTO.BlogID = blogId
	bcDTO.UserID = userId

	copier.Copy(&bc, &bcDTO)

	// create new comment
	newBlogComment, err := commentS.commentRepository.CreateNewBlogComment(ctx, nil, bc)
	if err != nil {
		return entity.Comment{}, err
	}
	return newBlogComment, nil
}
