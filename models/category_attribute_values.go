package models

type CategoryAttributeValues struct {
	CategoryAttributeValueID uint   `gorm:"primaryKey" json:"CategoryAttributeValueID"`
	CategoryAttributeID      uint   `json:"CategoryAttributeID"` // Bağlantı
	Value                    string `json:"Value"`
}
