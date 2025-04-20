package models

type CategoryAttributes struct {
	CategoryAttributeID           uint                      `gorm:"primaryKey" json:"CategoryAttributeID"`
	CategoryID                    uint                      `json:"CategoryID"` // Ekledim
	CategoryAttributeName         string                    `json:"CategoryAttributeName"`
	CategoryAttributeSlug         string                    `json:"CategoryAttributeSlug"`
	CategoryAttributeFilterable   bool                      `json:"CategoryAttributeFilterable"`
	CategoryAttributeSearchable   bool                      `json:"CategoryAttributeSearchable"`
	CategoryAttributeDisplayOrder uint                      `json:"CategoryAttributeDisplayOrder"`
	CategoryAttributeIsRequired   bool                      `json:"CategoryAttributeIsRequired"`
	CategoryAttributeIsVariant    bool                      `json:"CategoryAttributeIsVariant"`
	CategoryAttributeDataType     CategoryAttributeDataType `json:"CategoryAttributeDataType"`
}
