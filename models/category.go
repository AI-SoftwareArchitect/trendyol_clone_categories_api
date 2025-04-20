package models

type Category struct {
	ID           uint   `gorm:"column:id;primaryKey" json:"id"`
	Name         string `gorm:"column:name" json:"name"`
	Slug         string `gorm:"column:slug" json:"slug"`
	ParentID     uint   `gorm:"column:parent_id" json:"parent_id"`
	Description  string `gorm:"column:description" json:"description"`
	DisplayOrder uint   `gorm:"column:display_order" json:"display_order"`
}

type CategoryTree struct {
	ID           uint   `gorm:"column:id;primaryKey" json:"id"`
	Name         string `gorm:"column:name" json:"name"`
	Slug         string `gorm:"column:slug" json:"slug"`
	ParentID     uint   `gorm:"column:parent_id" json:"parent_id"`
	Description  string `gorm:"column:description" json:"description"`
	DisplayOrder uint   `gorm:"column:display_order" json:"display_order"`

	Subcategories []CategoryTree `gorm:"-" json:"subcategories,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}

func (CategoryTree) TableName() string {
	return "categories"
}
