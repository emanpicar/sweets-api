package main

import (
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

	sweetsManager := sweets.NewManager()

	logger.Log.Fatal(http.ListenAndServeTLS(
		"127.0.0.1:9988",
		"./certs/cert.pem",
		"./certs/key.pem",
		routes.NewRouter(sweetsManager),
	))
}
