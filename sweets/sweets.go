package sweets

import (
	"encoding/json"
	"errors"
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
		ProductID             string   `json:"productId,omitempty"`
	}

	Manager interface {
		GetAllSweets() *[]SweetsCollection
		GetSweetsByID(params map[string]string) (*SweetsCollection, error)
		CreateSweets(reqData *SweetsCollection) (*SweetsCollection, error)
		UpdateSweet(params map[string]string, reqData *SweetsCollection) (*SweetsCollection, error)
		DeleteSweet(params map[string]string) (string, error)
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

func (sw *sweetHandler) GetSweetsByID(params map[string]string) (*SweetsCollection, error) {
	sweetData, err := sw.dbManager.GetSweetByID(params["productId"])
	if err != nil {
		return nil, err
	}

	jsonReadyData := sw.modelSingleSweetToJSON(sweetData)

	return &jsonReadyData, nil
}

func (sw *sweetHandler) CreateSweets(reqData *SweetsCollection) (*SweetsCollection, error) {
	if reqData.ProductID == "" {
		return nil, errors.New("ProductID cannot be empty")
	}

	sweetData := sw.jsonSigleSweetToModel(reqData)
	if err := sw.dbManager.Insert(&sweetData); err != nil {
		return nil, err
	}

	return reqData, nil
}

func (sw *sweetHandler) UpdateSweet(params map[string]string, reqData *SweetsCollection) (*SweetsCollection, error) {
	if reqData.ProductID != "" {
		return nil, errors.New("ProductID cannot be updated")
	}

	sweetData := sw.jsonSigleSweetToModel(reqData)
	if err := sw.dbManager.UpdateByProductID(params["productId"], &sweetData); err != nil {
		return nil, err
	}

	updatedData := sw.modelSingleSweetToJSON(&sweetData)

	return &updatedData, nil
}

func (sw *sweetHandler) DeleteSweet(params map[string]string) (string, error) {
	successMsg, err := sw.dbManager.DeleteByProductID(params["productId"])
	if err != nil {
		return "", err
	}

	return successMsg, nil
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

func (sw *sweetHandler) jsonSigleSweetToModel(sweet *SweetsCollection) entities.SweetsCollection {
	return entities.SweetsCollection{
		Name:                  sweet.Name,
		ImageClosed:           sweet.ImageClosed,
		ImageOpen:             sweet.ImageOpen,
		Description:           sweet.Description,
		Story:                 sweet.Story,
		AllergyInfo:           sweet.AllergyInfo,
		DietaryCertifications: sweet.DietaryCertifications,
		ProductID:             sweet.ProductID,
		SourcingValues:        sw.jsonSourcingValuesToModel(sweet.SourcingValues),
		Ingredients:           sw.jsonIngredientsToModel(sweet.Ingredients),
	}
}

func (sw *sweetHandler) jsonSweetsCollectionToModel(jsonStructSweets *[]SweetsCollection) *[]entities.SweetsCollection {
	var dbEntity []entities.SweetsCollection

	for _, swData := range *jsonStructSweets {
		dbEntity = append(dbEntity, sw.jsonSigleSweetToModel(&swData))
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

func (sw *sweetHandler) modelSingleSweetToJSON(sweet *entities.SweetsCollection) SweetsCollection {
	return SweetsCollection{
		Name:                  sweet.Name,
		ImageClosed:           sweet.ImageClosed,
		ImageOpen:             sweet.ImageOpen,
		Description:           sweet.Description,
		Story:                 sweet.Story,
		AllergyInfo:           sweet.AllergyInfo,
		DietaryCertifications: sweet.DietaryCertifications,
		ProductID:             sweet.ProductID,
		SourcingValues:        sw.modelSourcingValuesToJSON(sweet.SourcingValues),
		Ingredients:           sw.modelIngredientsToJSON(sweet.Ingredients),
	}
}

func (sw *sweetHandler) modelSweetsCollectionToJSON(jsonStructSweets *[]entities.SweetsCollection) *[]SweetsCollection {
	var dbEntity []SweetsCollection

	for _, swData := range *jsonStructSweets {
		dbEntity = append(dbEntity, sw.modelSingleSweetToJSON(&swData))
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
