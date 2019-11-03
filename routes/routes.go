package routes

import (
	"encoding/json"
	"net/http"

	"github.com/emanpicar/sweets-api/sweets"

	"github.com/gorilla/mux"
)

type (
	Router interface {
		ServeHTTP(http.ResponseWriter, *http.Request)
	}

	routeHandler struct {
		sweetsManager sweets.Manager
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
	router.HandleFunc("/api/sweets", rh.createSweets).Methods("POST")
	router.HandleFunc("/api/sweets/{productId}", rh.updateSweet).Methods("PUT")
	router.HandleFunc("/api/sweets/{productId}", rh.deleteSweets).Methods("DELETE")
}

func (rh *routeHandler) getAllSweets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := rh.sweetsManager.GetAllSweets()

	rh.encodeError(json.NewEncoder(w).Encode(data), w)
}

func (rh *routeHandler) createSweets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqData sweets.SweetsCollection
	json.NewDecoder(r.Body).Decode(&reqData)

	data, err := rh.sweetsManager.CreateSweets(&reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rh.encodeError(json.NewEncoder(w).Encode(data), w)
}

func (rh *routeHandler) updateSweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqData sweets.SweetsCollection
	json.NewDecoder(r.Body).Decode(&reqData)

	data, err := rh.sweetsManager.UpdateSweet(mux.Vars(r), &reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rh.encodeError(json.NewEncoder(w).Encode(data), w)
}

func (rh *routeHandler) deleteSweets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	data := rh.sweetsManager.DeleteSweet(params)

	rh.encodeError(json.NewEncoder(w).Encode(data), w)
}

func (rh *routeHandler) encodeError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
