package repository

import (
	"context"
	"gorm.io/gorm"
	"products_graphql_api/models"
)

type CategoryRepository interface {
	GetWomenMainCategories(ctx context.Context) ([]models.CategoryTree, error)
	GetSubcategories(ctx context.Context, parentID uint) ([]models.CategoryTree, error)
	GetCategoryTree(ctx context.Context, parentID uint) ([]models.CategoryTree, error)
	GetCategoryBySlug(ctx context.Context, slug string) (*models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetWomenMainCategories(ctx context.Context) ([]models.CategoryTree, error) {
	var categories []models.CategoryTree
	err := r.db.WithContext(ctx).
		Where("parent_id = ?", 1). // Assuming 1 is the ID for "Women" category
		Order("display_order ASC").
		Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) GetSubcategories(ctx context.Context, parentID uint) ([]models.CategoryTree, error) {
	var categories []models.CategoryTree
	err := r.db.WithContext(ctx).
		Where("parent_id = ?", parentID).
		Order("display_order ASC").
		Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) GetCategoryTree(ctx context.Context, parentID uint) ([]models.CategoryTree, error) {
	var categories []models.CategoryTree
	err := r.db.WithContext(ctx).
		Where("parent_id = ?", parentID).
		Order("display_order ASC").
		Find(&categories).Error
	if err != nil {
		return nil, err
	}

	for i := range categories {
		subcategories, err := r.GetCategoryTree(ctx, categories[i].ID)
		if err != nil {
			return nil, err
		}
		categories[i].Subcategories = subcategories
	}

	return categories, nil
}

func (r *categoryRepository) GetCategoryBySlug(ctx context.Context, slug string) (*models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).
		Where("slug = ?", slug).
		First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}
