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
	GetBlogLikeByID(ctx context.Context, tx *gorm.DB, clID uint64) (entity.BlogLike, error)
	CreateNewBlogLike(ctx context.Context, tx *gorm.DB, bl entity.BlogLike) (entity.BlogLike, error)
	DeleteBlogLike(ctx context.Context, tx *gorm.DB, blID uint64) error
	RestoreBlogLike(ctx context.Context, tx *gorm.DB, bl entity.BlogLike) (entity.BlogLike, error)
	CheckBlogLike(ctx context.Context, tx *gorm.DB, bl entity.BlogLike, blogId uint64, userId uint64) (int, entity.BlogLike, error)
	SetBlogLikeCount(ctx context.Context, tx *gorm.DB, blog entity.Blog) error

	// CommentLike functional
	GetAllCommentLikes(ctx context.Context, tx *gorm.DB) ([]entity.CommentLike, error)
	GetCommentLikeByID(ctx context.Context, tx *gorm.DB, clID uint64) (entity.CommentLike, error)
	CreateNewCommentLike(ctx context.Context, tx *gorm.DB, cl entity.CommentLike) (entity.CommentLike, error)
	DeleteCommentLike(ctx context.Context, tx *gorm.DB, clID uint64) error
	RestoreCommentLike(ctx context.Context, tx *gorm.DB, cl entity.CommentLike) (entity.CommentLike, error)
	CheckCommentLike(ctx context.Context, tx *gorm.DB, cl entity.CommentLike, blogId uint64, userId uint64) (int, entity.CommentLike, error)
	SetCommentLikeCount(ctx context.Context, tx *gorm.DB, comment entity.Comment) error
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

func (likeR *likeRepository) GetCommentLikeByID(ctx context.Context, tx *gorm.DB, clID uint64) (entity.CommentLike, error) {
	var err error
	var cl entity.CommentLike
	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Where("id = $1", clID).Take(&cl)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Where("id = $1", clID).Take(&cl).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return cl, err
	}
	return cl, nil
}

func (likeR *likeRepository) CreateNewCommentLike(ctx context.Context, tx *gorm.DB, cl entity.CommentLike) (entity.CommentLike, error) {
	var err error
	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Create(&cl)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Create(&cl).Error
	}

	if err != nil {
		return entity.CommentLike{}, err
	}
	return cl, nil
}

func (likeR *likeRepository) DeleteCommentLike(ctx context.Context, tx *gorm.DB, clID uint64) error {
	var err error
	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Delete(&entity.CommentLike{}, clID)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Delete(&entity.CommentLike{}, clID).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return err
	}
	return nil
}

func (likeR *likeRepository) RestoreCommentLike(ctx context.Context, tx *gorm.DB, cl entity.CommentLike) (entity.CommentLike, error) {
	var err error
	clRestore := cl
	clRestore.Model.DeletedAt = gorm.DeletedAt{}
	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Unscoped().Save(&clRestore)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Unscoped().Save(&clRestore).Error
	}

	if err != nil {
		return clRestore, err
	}
	return clRestore, nil
}

func (likeR *likeRepository) CheckCommentLike(ctx context.Context, tx *gorm.DB, cl entity.CommentLike, commentId uint64, userId uint64) (int, entity.CommentLike, error) {
	var clike entity.CommentLike
	var err error

	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Unscoped().Where("comment_id = $1 AND user_id = $2", commentId, userId).Take(&clike)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Unscoped().Where("comment_id = $1 AND user_id = $2", commentId, userId).Take(&clike).Error
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 1, clike, nil
		} else {
			return -1, entity.CommentLike{}, err
		}
	} else {
		return 2, clike, nil
	}
}

func (likeR *likeRepository) SetCommentLikeCount(ctx context.Context, tx *gorm.DB, comment entity.Comment) error {
	var err error
	commentSet := comment
	commentSet.LikeCount = len(commentSet.Likes)
	if tx == nil {
		tx = likeR.db.WithContext(ctx).Debug().Save(&commentSet)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Save(&commentSet).Error
	}

	if err != nil {
		return err
	}
	return nil
}
