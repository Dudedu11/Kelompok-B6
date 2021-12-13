package router

import (
	"be/transport"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/like", transport.Melike).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/like/{feed_id}", transport.TmplknSemuaLike).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/jmllike/{feed_id}", transport.TmplknLike).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/like/{user_id}/{feed_id}", transport.UptLike).Methods("PUT", "OPTIONS")

	router.HandleFunc("/api/komen", transport.TmbhKomen).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/komen/{feed_id}", transport.TmplknSemuaKomen).Methods("GET", "OPTIONS")

	return router
}