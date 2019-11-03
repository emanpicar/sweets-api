package db

import (
	"fmt"

	"github.com/emanpicar/sweets-api/settings"

	"github.com/emanpicar/sweets-api/db/entities"
	"github.com/emanpicar/sweets-api/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type (
	DBManager interface {
		BatchFirstOrCreate(data *[]entities.SweetsCollection)
		Insert(data *entities.SweetsCollection) error
		UpdateByProductID(pID string, data *entities.SweetsCollection) error
		GetSweetCollections() *[]entities.SweetsCollection
	}

	dbHandler struct {
		database *gorm.DB
	}
)

func NewDBManager() DBManager {
	dbHandler := &dbHandler{}
	dbHandler.connect()
	dbHandler.migrateTables()

	return dbHandler
}

func (dbHandler *dbHandler) connect() {
	logger.Log.Infoln("Establishing connection to DB")

	var err error
	dbHandler.database, err = gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=sweetscollection password=%v sslmode=disable",
		settings.GetDBHost(), settings.GetDBPort(), settings.GetDBUser(), settings.GetDBPass(),
	))

	if err != nil {
		logger.Log.Fatalln(err)
	}

	logger.Log.Infoln("Successfully connected to DB")
}

func (dbHandler *dbHandler) migrateTables() {
	dbHandler.database.AutoMigrate(&entities.SweetsCollection{})
	dbHandler.database.AutoMigrate(&entities.SourcingValues{}).AddForeignKey("product_id", "sweets_collection(product_id)", "CASCADE", "CASCADE")
	dbHandler.database.AutoMigrate(&entities.Ingredients{}).AddForeignKey("product_id", "sweets_collection(product_id)", "CASCADE", "CASCADE")
}

func (dbHandler *dbHandler) BatchFirstOrCreate(swCollection *[]entities.SweetsCollection) {
	for _, sweet := range *swCollection {
		dbHandler.database.FirstOrCreate(&sweet, entities.SweetsCollection{})
	}
}

func (dbHandler *dbHandler) Insert(data *entities.SweetsCollection) error {
	err := dbHandler.database.Where(&entities.SweetsCollection{ProductID: data.ProductID}).First(data).Error
	if err == nil {
		return fmt.Errorf("Sweets with productID:%v already exist", data.ProductID)
	}

	dbHandler.database.FirstOrCreate(data)

	return nil
}

func (dbHandler *dbHandler) UpdateByProductID(pID string, data *entities.SweetsCollection) error {
	err := dbHandler.database.Where(&entities.SweetsCollection{ProductID: pID}).First(data).Error
	if err != nil {
		return fmt.Errorf("Sweets with productID:%v does not exist", pID)
	}

	logger.Log.Infof("########## %+v", data)
	if err := dbHandler.database.Save(data).Error; err != nil {
		return err
	}
	logger.Log.Infof("########## ssss %+v", data)

	return nil
}

func (dbHandler *dbHandler) GetSweetCollections() *[]entities.SweetsCollection {
	var swData []entities.SweetsCollection
	dbHandler.database.Set("gorm:auto_preload", true).Find(&swData)

	return &swData
}
