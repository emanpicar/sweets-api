package db

import (
	"errors"

	"github.com/emanpicar/sweets-api/db/entities"
	"github.com/emanpicar/sweets-api/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type (
	DBManager interface {
		BatchFirstOrCreate(data *[]entities.SweetsCollection)
		Insert(data *entities.SweetsCollection) error
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
	dbHandler.database, err = gorm.Open("postgres", "host=localhost port=54320 user=secretdbuser dbname=sweetscollection password=secretdbpass sslmode=disable")

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
	if dbHandler.database.NewRecord(data) {
		return errors.New("Already existing in database")
	}

	dbHandler.database.Create(data)
	return nil
}

func (dbHandler *dbHandler) GetSweetCollections() *[]entities.SweetsCollection {
	var swData []entities.SweetsCollection
	dbHandler.database.Set("gorm:auto_preload", true).Find(&swData)

	return &swData
}
