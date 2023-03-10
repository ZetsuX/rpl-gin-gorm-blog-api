package repository

import (
	"context"
	"go-blogrpl/entity"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

type CommentRepository interface {
	// db transaction
	BeginTx(ctx context.Context) (*gorm.DB, error)
	CommitTx(ctx context.Context, tx *gorm.DB) error
	RollbackTx(ctx context.Context, tx *gorm.DB)

	// // Comments functional
	// GetAllComments(ctx context.Context, tx *gorm.DB) ([]entity.Comment, error)

	// BlogComments functional
	CreateNewBlogComment(ctx context.Context, tx *gorm.DB, comment entity.Comment) (entity.Comment, error)
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db: db}
}

func (commentR *commentRepository) BeginTx(ctx context.Context) (*gorm.DB, error) {
	tx := commentR.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (commentR *commentRepository) CommitTx(ctx context.Context, tx *gorm.DB) error {
	err := tx.WithContext(ctx).Commit().Error
	if err == nil {
		return err
	}
	return nil
}

func (commentR *commentRepository) RollbackTx(ctx context.Context, tx *gorm.DB) {
	tx.WithContext(ctx).Debug().Rollback()
}

func (commentR *commentRepository) CreateNewBlogComment(ctx context.Context, tx *gorm.DB, comment entity.Comment) (entity.Comment, error) {
	var err error
	if tx == nil {
		tx = commentR.db.WithContext(ctx).Debug().Create(&comment)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Create(&comment).Error
	}

	if err != nil {
		return entity.Comment{}, err
	}
	return comment, nil
}
