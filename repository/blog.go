package repository

import (
	"context"
	"errors"
	"go-blogrpl/entity"

	"gorm.io/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

type BlogRepository interface {
	// db transaction
	BeginTx(ctx context.Context) (*gorm.DB, error)
	CommitTx(ctx context.Context, tx *gorm.DB) error
	RollbackTx(ctx context.Context, tx *gorm.DB)

	// functional
	CreateNewBlog(ctx context.Context, tx *gorm.DB, blog entity.Blog) (entity.Blog, error)
	GetAllBlogs(ctx context.Context, tx *gorm.DB) ([]entity.Blog, error)
	GetBlogBySlug(ctx context.Context, tx *gorm.DB, slug string) (entity.Blog, error)
}

func NewBlogRepository(db *gorm.DB) *blogRepository {
	return &blogRepository{db: db}
}

func (blogR *blogRepository) BeginTx(ctx context.Context) (*gorm.DB, error) {
	tx := blogR.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (blogR *blogRepository) CommitTx(ctx context.Context, tx *gorm.DB) error {
	err := tx.WithContext(ctx).Commit().Error
	if err == nil {
		return err
	}
	return nil
}

func (blogR *blogRepository) RollbackTx(ctx context.Context, tx *gorm.DB) {
	tx.WithContext(ctx).Debug().Rollback()
}

func (blogR *blogRepository) CreateNewBlog(ctx context.Context, tx *gorm.DB, blog entity.Blog) (entity.Blog, error) {
	var err error
	if tx == nil {
		tx = blogR.db.WithContext(ctx).Debug().Create(&blog)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Create(&blog).Error
	}

	if err != nil {
		return entity.Blog{}, err
	}
	return blog, nil
}

func (blogR *blogRepository) GetAllBlogs(ctx context.Context, tx *gorm.DB) ([]entity.Blog, error) {
	var err error
	var blogs []entity.Blog

	if tx == nil {
		tx = blogR.db.WithContext(ctx).Debug().Preload("Likes").Find(&blogs)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Preload("Likes").Find(&blogs).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return blogs, err
	}
	return blogs, nil
}

func (blogR *blogRepository) GetBlogBySlug(ctx context.Context, tx *gorm.DB, slug string) (entity.Blog, error) {
	var err error
	var blog entity.Blog
	if tx == nil {
		tx = blogR.db.WithContext(ctx).Debug().Where("slug = $1", slug).Preload("Comments").Preload("Likes").Take(&blog)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Where("slug = $1", slug).Preload("Comments").Preload("Likes").Take(&blog).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return blog, err
	}
	return blog, nil
}
