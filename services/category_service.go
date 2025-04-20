package services

import (
	"context"
	"products_graphql_api/models"
	"products_graphql_api/repository"
)

type CategoryService interface {
	GetWomenMainCategories(ctx context.Context) ([]models.CategoryTree, error)
	GetSubcategories(ctx context.Context, parentID uint) ([]models.CategoryTree, error)
	GetFullCategoryTree(ctx context.Context) ([]models.CategoryTree, error)
	GetCategoryBySlug(ctx context.Context, slug string) (*models.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetWomenMainCategories(ctx context.Context) ([]models.CategoryTree, error) {
	return s.repo.GetWomenMainCategories(ctx)
}

func (s *categoryService) GetSubcategories(ctx context.Context, parentID uint) ([]models.CategoryTree, error) {
	return s.repo.GetSubcategories(ctx, parentID)
}

func (s *categoryService) GetFullCategoryTree(ctx context.Context) ([]models.CategoryTree, error) {
	// Start with parentID = 1 (Women category) to get the full tree
	return s.repo.GetCategoryTree(ctx, 1)
}

func (s *categoryService) GetCategoryBySlug(ctx context.Context, slug string) (*models.Category, error) {
	return s.repo.GetCategoryBySlug(ctx, slug)
}
