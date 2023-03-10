package repository

import (
	"context"
	"errors"
	"go-blogrpl/entity"

	"gorm.io/gorm"
)

type likeRepository struct {
	db *gorm.DB
}

type LikeRepository interface {
	// db transaction
	BeginTx(ctx context.Context) (*gorm.DB, error)
	CommitTx(ctx context.Context, tx *gorm.DB) error
	RollbackTx(ctx context.Context, tx *gorm.DB)

	// BlogLike functional
	GetAllBlogLikes(ctx context.Context, tx *gorm.DB) ([]entity.BlogLike, error)

	// CommentLike functional
	GetAllCommentLikes(ctx context.Context, tx *gorm.DB) ([]entity.CommentLike, error)
}

func NewLikeRepository(db *gorm.DB) *likeRepository {
	return &likeRepository{db: db}
}

func (likeR *likeRepository) BeginTx(ctx context.Context) (*gorm.DB, error) {
	tx := likeR.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (likeR *likeRepository) CommitTx(ctx context.Context, tx *gorm.DB) error {
	err := tx.WithContext(ctx).Commit().Error
	if err == nil {
		return err
	}
	return nil
}

func (likeR *likeRepository) RollbackTx(ctx context.Context, tx *gorm.DB) {
	tx.WithContext(ctx).Debug().Rollback()
}

func (likeR *likeRepository) GetAllBlogLikes(ctx context.Context, tx *gorm.DB) ([]entity.BlogLike, error) {
	var err error
	var blikes []entity.BlogLike

	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Preload("Users").Find(&blikes)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Preload("Users").Find(&blikes).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return blikes, err
	}
	return blikes, nil
}

func (likeR *likeRepository) GetAllCommentLikes(ctx context.Context, tx *gorm.DB) ([]entity.CommentLike, error) {
	var err error
	var clikes []entity.CommentLike

	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Preload("Users").Find(&clikes)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Preload("Users").Find(&clikes).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return clikes, err
	}
	return clikes, nil
}
