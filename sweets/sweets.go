package sweets

import (
	"encoding/json"
	"io/ioutil"

	"github.com/emanpicar/sweets-api/db"
	"github.com/emanpicar/sweets-api/db/entities"
	"github.com/emanpicar/sweets-api/logger"
)

type (
	SweetsCollection struct {
		Name                  string   `json:"name"`
		ImageClosed           string   `json:"image_closed"`
		ImageOpen             string   `json:"image_open"`
		Description           string   `json:"description"`
		Story                 string   `json:"story"`
		SourcingValues        []string `json:"sourcing_values"`
		Ingredients           []string `json:"ingredients"`
		AllergyInfo           string   `json:"allergy_info"`
		DietaryCertifications string   `json:"dietary_certifications"`
		ProductID             string   `json:"productId"`
	}

	Manager interface {
		GetAllSweets() *[]SweetsCollection
		CreateSweets() *[]entities.SweetsCollection
		UpdateSweet(params map[string]string) *[]entities.SweetsCollection
		DeleteSweet(params map[string]string) *[]entities.SweetsCollection
		PopulateDefaultData()
	}

	sweetHandler struct {
		dbManager db.DBManager
	}
)

func NewManager(dbManager db.DBManager) Manager {
	return &sweetHandler{dbManager}
}

func (sw *sweetHandler) GetAllSweets() *[]SweetsCollection {
	sweetsList := sw.dbManager.GetSweetCollections()
	jsonReadyList := sw.modelSweetsCollectionToJSON(sweetsList)

	return jsonReadyList
}

func (sw *sweetHandler) CreateSweets() *[]entities.SweetsCollection {
	sweetsList := []entities.SweetsCollection{
		entities.SweetsCollection{SourcingValues: []entities.SourcingValues{}, Ingredients: []entities.Ingredients{}},
		entities.SweetsCollection{SourcingValues: []entities.SourcingValues{}, Ingredients: []entities.Ingredients{}},
	}

	return &sweetsList
}

func (sw *sweetHandler) UpdateSweet(params map[string]string) *[]entities.SweetsCollection {
	logger.Log.Infof("################# %v", params)
	sweetsList := []entities.SweetsCollection{
		entities.SweetsCollection{SourcingValues: []entities.SourcingValues{}, Ingredients: []entities.Ingredients{}},
		entities.SweetsCollection{SourcingValues: []entities.SourcingValues{}, Ingredients: []entities.Ingredients{}},
	}

	return &sweetsList
}

func (sw *sweetHandler) DeleteSweet(params map[string]string) *[]entities.SweetsCollection {
	logger.Log.Infof("################# %v", params)
	sweetsList := []entities.SweetsCollection{
		entities.SweetsCollection{SourcingValues: []entities.SourcingValues{}, Ingredients: []entities.Ingredients{}},
		entities.SweetsCollection{SourcingValues: []entities.SourcingValues{}, Ingredients: []entities.Ingredients{}},
	}

	return &sweetsList
}

func (sw *sweetHandler) PopulateDefaultData() {
	var jsonStructSweets []SweetsCollection

	bytesData, err := ioutil.ReadFile("./jsondata/sweets.json")
	if err != nil {
		logger.Log.Errorf("Unable to create default data due to: %v", err)
	}

	if err = json.Unmarshal(bytesData, &jsonStructSweets); err != nil {
		logger.Log.Errorf("Unable to create default data due to: %v", err)
	}

	sweetsCollectionModel := sw.jsonSweetsCollectionToModel(&jsonStructSweets)

	sw.dbManager.BatchFirstOrCreate(sweetsCollectionModel)
}

func (sw *sweetHandler) jsonSweetsCollectionToModel(jsonStructSweets *[]SweetsCollection) *[]entities.SweetsCollection {
	var dbEntity []entities.SweetsCollection

	for _, swData := range *jsonStructSweets {
		dbEntity = append(dbEntity, entities.SweetsCollection{
			Name:                  swData.Name,
			ImageClosed:           swData.ImageClosed,
			ImageOpen:             swData.ImageOpen,
			Description:           swData.Description,
			Story:                 swData.Story,
			AllergyInfo:           swData.AllergyInfo,
			DietaryCertifications: swData.DietaryCertifications,
			ProductID:             swData.ProductID,
			SourcingValues:        sw.jsonSourcingValuesToModel(swData.SourcingValues),
			Ingredients:           sw.jsonIngredientsToModel(swData.Ingredients),
		})
	}

	return &dbEntity
}

func (sw *sweetHandler) jsonSourcingValuesToModel(sourcingValues []string) []entities.SourcingValues {
	var sourcingEntity []entities.SourcingValues

	for _, value := range sourcingValues {
		sourcingEntity = append(sourcingEntity, entities.SourcingValues{
			Value: value,
		})
	}

	return sourcingEntity
}

func (sw *sweetHandler) jsonIngredientsToModel(ingredients []string) []entities.Ingredients {
	var ingredientsEntity []entities.Ingredients

	for _, value := range ingredients {
		ingredientsEntity = append(ingredientsEntity, entities.Ingredients{
			Value: value,
		})
	}

	return ingredientsEntity
}

func (sw *sweetHandler) modelSweetsCollectionToJSON(jsonStructSweets *[]entities.SweetsCollection) *[]SweetsCollection {
	var dbEntity []SweetsCollection

	for _, swData := range *jsonStructSweets {
		dbEntity = append(dbEntity, SweetsCollection{
			Name:                  swData.Name,
			ImageClosed:           swData.ImageClosed,
			ImageOpen:             swData.ImageOpen,
			Description:           swData.Description,
			Story:                 swData.Story,
			AllergyInfo:           swData.AllergyInfo,
			DietaryCertifications: swData.DietaryCertifications,
			ProductID:             swData.ProductID,
			SourcingValues:        sw.modelSourcingValuesToJSON(swData.SourcingValues),
			Ingredients:           sw.modelIngredientsToJSON(swData.Ingredients),
		})
	}

	return &dbEntity
}

func (sw *sweetHandler) modelSourcingValuesToJSON(sourcingValues []entities.SourcingValues) []string {
	var sourcingEntity []string

	for _, sv := range sourcingValues {
		sourcingEntity = append(sourcingEntity, sv.Value)
	}

	return sourcingEntity
}

func (sw *sweetHandler) modelIngredientsToJSON(ingredients []entities.Ingredients) []string {
	var ingredientsEntity []string

	for _, ingredient := range ingredients {
		ingredientsEntity = append(ingredientsEntity, ingredient.Value)
	}

	return ingredientsEntity
}
