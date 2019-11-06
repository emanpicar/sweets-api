package routes

import (
	"encoding/json"
	"net/http"

	"github.com/emanpicar/sweets-api/logger"

	"github.com/emanpicar/sweets-api/sweets"

	"github.com/gorilla/mux"
)

type (
	Router interface {
		ServeHTTP(http.ResponseWriter, *http.Request)
	}

	routeHandler struct {
		sweetsManager sweets.Manager
		router        *mux.Router
	}

	JsonMessage struct {
		Message string `json:"message"`
	}
)

func NewRouter(sweetsManager sweets.Manager) Router {
	routeHandler := &routeHandler{sweetsManager: sweetsManager}

	return routeHandler.newRouter()
}

func (rh *routeHandler) newRouter() *mux.Router {
	router := mux.NewRouter()
	rh.registerRoutes(router)

	return router
}

func (rh *routeHandler) registerRoutes(router *mux.Router) {
	router.HandleFunc("/api/sweets", rh.getAllSweets).Methods("GET")
	router.HandleFunc("/api/sweets/{productId}", rh.getSweetsByID).Methods("GET")
	router.HandleFunc("/api/sweets", rh.createSweets).Methods("POST")
	router.HandleFunc("/api/sweets/{productId}", rh.updateSweet).Methods("PUT")
	router.HandleFunc("/api/sweets/{productId}", rh.deleteSweets).Methods("DELETE")

	rh.router = router
}

func (rh *routeHandler) getAllSweets(w http.ResponseWriter, r *http.Request) {
	logger.Log.Infoln("Getting all sweets")

	w.Header().Set("Content-Type", "application/json")
	data := rh.sweetsManager.GetAllSweets()

	rh.encodeError(json.NewEncoder(w).Encode(data), w)
}

func (rh *routeHandler) getSweetsByID(w http.ResponseWriter, r *http.Request) {
	logger.Log.Infof("Getting sweets by id:%v", mux.Vars(r)["productId"])

	w.Header().Set("Content-Type", "application/json")
	data, err := rh.sweetsManager.GetSweetsByID(mux.Vars(r))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		rh.encodeError(json.NewEncoder(w).Encode(&JsonMessage{err.Error()}), w)
		return
	}

	rh.encodeError(json.NewEncoder(w).Encode(data), w)
}

func (rh *routeHandler) createSweets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqData sweets.SweetsCollection
	rh.encodeError(json.NewDecoder(r.Body).Decode(&reqData), w)

	logger.Log.Infof("Creating sweets in id:%v", reqData.ProductID)
	data, err := rh.sweetsManager.CreateSweets(&reqData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		rh.encodeError(json.NewEncoder(w).Encode(&JsonMessage{err.Error()}), w)
		return
	}

	rh.encodeError(json.NewEncoder(w).Encode(data), w)
}

func (rh *routeHandler) updateSweet(w http.ResponseWriter, r *http.Request) {
	logger.Log.Infof("Updating sweets in id:%v", mux.Vars(r)["productId"])

	w.Header().Set("Content-Type", "application/json")

	var reqData sweets.SweetsCollection
	json.NewDecoder(r.Body).Decode(&reqData)

	data, err := rh.sweetsManager.UpdateSweet(mux.Vars(r), &reqData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		rh.encodeError(json.NewEncoder(w).Encode(&JsonMessage{err.Error()}), w)
		return
	}

	rh.encodeError(json.NewEncoder(w).Encode(data), w)
}

func (rh *routeHandler) deleteSweets(w http.ResponseWriter, r *http.Request) {
	logger.Log.Infof("Deleting sweets id:%v", mux.Vars(r)["productId"])

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	successMsg, err := rh.sweetsManager.DeleteSweet(params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		rh.encodeError(json.NewEncoder(w).Encode(&JsonMessage{err.Error()}), w)
		return
	}

	rh.encodeError(json.NewEncoder(w).Encode(&JsonMessage{successMsg}), w)
}

func (rh *routeHandler) encodeError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
