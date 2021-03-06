package main

import (
	"fmt"

	"github.com/emanpicar/sweets-api/db"
	"github.com/emanpicar/sweets-api/logger"
	"github.com/emanpicar/sweets-api/routes"
	"github.com/emanpicar/sweets-api/settings"
	"github.com/emanpicar/sweets-api/sweets"

	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	logger.Init(settings.GetLogLevel())
	logger.Log.Infoln("Initializing Sweets API")

	dbManager := db.NewDBManager()
	sweetsManager := sweets.NewManager(dbManager)
	sweetsManager.PopulateDefaultData()

	logger.Log.Fatal(http.ListenAndServeTLS(
		fmt.Sprintf("%v:%v", settings.GetServerHost(), settings.GetServerPort()),
		settings.GetServerPublicKey(),
		settings.GetServerPrivateKey(),
		routes.NewRouter(sweetsManager),
	))
}
