package entities

import (
	"github.com/jinzhu/gorm"
)

type (
	SweetsCollection struct {
		Name                  string           `gorm:"type:varchar(100)"`
		ImageClosed           string           `gorm:"type:varchar(200)"`
		ImageOpen             string           `gorm:"type:varchar(200)"`
		Description           string           `gorm:"type:varchar(500)"`
		Story                 string           `gorm:"type:text"`
		SourcingValues        []SourcingValues `gorm:"foreignkey:ProductID"`
		Ingredients           []Ingredients    `gorm:"foreignkey:ProductID"`
		AllergyInfo           string           `gorm:"type:varchar(200)"`
		DietaryCertifications string           `gorm:"type:varchar(200)"`
		ProductID             string           `gorm:"unique;primary_key"`
	}

	SourcingValues struct {
		gorm.Model
		Value     string `gorm:"type:varchar(100)"`
		ProductID string
	}

	Ingredients struct {
		gorm.Model
		Value     string `gorm:"type:varchar(100)"`
		ProductID string
	}
)

func (SweetsCollection) TableName() string {
	return "sweets_collection"
}

func (SourcingValues) TableName() string {
	return "sweets_sourcing_values"
}

func (Ingredients) TableName() string {
	return "sweets_ingredients"
}
