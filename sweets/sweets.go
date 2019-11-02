package sweets

import (
	"github.com/emanpicar/sweets-api/db/entities"
	"github.com/emanpicar/sweets-api/logger"
)

type (
	Manager interface {
		GetAllSweets() *[]entities.SweetEntity
		CreateSweets() *[]entities.SweetEntity
		UpdateSweet(params map[string]string) *[]entities.SweetEntity
		DeleteSweet(params map[string]string) *[]entities.SweetEntity
	}

	sweetHandler struct {
	}
)

func NewManager() Manager {
	return &sweetHandler{}
}

func (sw *sweetHandler) GetAllSweets() *[]entities.SweetEntity {
	sweetsList := make([]entities.SweetEntity, 2)
	return &sweetsList
}

func (sw *sweetHandler) CreateSweets() *[]entities.SweetEntity {
	sweetsList := make([]entities.SweetEntity, 3)
	return &sweetsList
}

func (sw *sweetHandler) UpdateSweet(params map[string]string) *[]entities.SweetEntity {
	logger.Log.Infof("################# %v", params)
	sweetsList := make([]entities.SweetEntity, 3)
	return &sweetsList
}

func (sw *sweetHandler) DeleteSweet(params map[string]string) *[]entities.SweetEntity {
	logger.Log.Infof("################# %v", params)
	sweetsList := make([]entities.SweetEntity, 2)
	return &sweetsList
}
