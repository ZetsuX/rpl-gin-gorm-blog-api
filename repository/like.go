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
	GetBlogLikeByID(ctx context.Context, tx *gorm.DB, blID uint64) (entity.BlogLike, error)
	CreateNewBlogLike(ctx context.Context, tx *gorm.DB, bl entity.BlogLike) (entity.BlogLike, error)
	DeleteBlogLike(ctx context.Context, tx *gorm.DB, blID uint64) error
	RestoreBlogLike(ctx context.Context, tx *gorm.DB, bl entity.BlogLike) (entity.BlogLike, error)
	CheckBlogLike(ctx context.Context, tx *gorm.DB, bl entity.BlogLike, blogId uint64, userId uint64) (int, entity.BlogLike, error)
	SetBlogLikeCount(ctx context.Context, tx *gorm.DB, blog entity.Blog) error

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
		tx = likeR.db.WithContext(ctx).Debug().Find(&blikes)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Find(&blikes).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return blikes, err
	}
	return blikes, nil
}

func (likeR *likeRepository) GetBlogLikeByID(ctx context.Context, tx *gorm.DB, blID uint64) (entity.BlogLike, error) {
	var err error
	var bl entity.BlogLike
	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Where("id = $1", blID).Take(&bl)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Where("id = $1", blID).Take(&bl).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return bl, err
	}
	return bl, nil
}

func (likeR *likeRepository) CreateNewBlogLike(ctx context.Context, tx *gorm.DB, bl entity.BlogLike) (entity.BlogLike, error) {
	var err error
	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Create(&bl)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Create(&bl).Error
	}

	if err != nil {
		return entity.BlogLike{}, err
	}
	return bl, nil
}

func (likeR *likeRepository) DeleteBlogLike(ctx context.Context, tx *gorm.DB, blID uint64) error {
	var err error
	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Delete(&entity.BlogLike{}, blID)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Delete(&entity.BlogLike{}, blID).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return err
	}
	return nil
}

func (likeR *likeRepository) RestoreBlogLike(ctx context.Context, tx *gorm.DB, bl entity.BlogLike) (entity.BlogLike, error) {
	var err error
	blRestore := bl
	blRestore.Model.DeletedAt = gorm.DeletedAt{}
	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Unscoped().Save(&blRestore)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Unscoped().Save(&blRestore).Error
	}

	if err != nil {
		return blRestore, err
	}
	return blRestore, nil
}

func (likeR *likeRepository) CheckBlogLike(ctx context.Context, tx *gorm.DB, bl entity.BlogLike, blogId uint64, userId uint64) (int, entity.BlogLike, error) {
	var blike entity.BlogLike
	var err error

	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Unscoped().Where("blog_id = $1 AND user_id = $2", blogId, userId).Take(&blike)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Unscoped().Where("blog_id = $1 AND user_id = $2", blogId, userId).Take(&blike).Error
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 1, blike, nil
		} else {
			return -1, entity.BlogLike{}, err
		}
	} else {
		return 2, blike, nil
	}
}

func (likeR *likeRepository) SetBlogLikeCount(ctx context.Context, tx *gorm.DB, blog entity.Blog) error {
	var err error
	blogSet := blog
	blogSet.LikeCount = len(blogSet.Likes)
	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Save(&blogSet)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Save(&blogSet).Error
	}

	if err != nil {
		return err
	}
	return nil
}

func (likeR *likeRepository) GetAllCommentLikes(ctx context.Context, tx *gorm.DB) ([]entity.CommentLike, error) {
	var err error
	var clikes []entity.CommentLike

	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Find(&clikes)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Find(&clikes).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return clikes, err
	}
	return clikes, nil
}
