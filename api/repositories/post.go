package repositories

import (
	"blog/database"
	"blog/models"
)

type PostRepository struct {
	db database.Database
}

func NewPostRepository(db database.Database) PostRepository {
	return PostRepository{
		db: db,
	}
}

func (p PostRepository) Save(post models.Post) error {
	return p.db.DB.Create(&post).Error
}

func (p PostRepository) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	var posts []models.Post
	var totalRows int64 = 0

	queryBuilder := p.db.DB.Model(&models.Post{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			p.db.DB.Where("post.title LIKE ? ", queryKeyword))
	}

	// Count total rows first before applying order and pagination
	err := queryBuilder.Where(post).Count(&totalRows).Error
	if err != nil {
		return &posts, 0, err
	}

	// Then find the actual records with ordering
	err = queryBuilder.Order("created_at desc").Find(&posts).Error
	return &posts, totalRows, err
}

func (p PostRepository) Update(post models.Post) error {
	return p.db.DB.Save(&post).Error
}

func (p PostRepository) Find(post models.Post) (models.Post, error) {
	var posts models.Post
	err := p.db.DB.
		Model(&models.Post{}).
		Where(&post).
		Take(&posts).Error
	return posts, err
}

func (p PostRepository) Delete(post models.Post) error {
	return p.db.DB.Delete(&post).Error
}
