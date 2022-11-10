package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router       *mux.Router
	GroupService GroupService
	FieldService FieldService
	Server       *http.Server
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken, X-CSRF-Token, Authorization, Token")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("content-type", "application/json;charset=UTF-8")
		next.ServeHTTP(w, r)
	})
}

func NewHandler(grp_service GroupService, fld_service FieldService) *Handler {
	h := &Handler{
		GroupService: grp_service,
		FieldService: fld_service,
	}

	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Router.Use(corsMiddleware)
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LogginMiddleware)
	h.Router.Use(TimeOutMiddleware)
	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/api/v1/groups", h.GetGroups).Methods("GET")
	h.Router.HandleFunc("/api/v1/groups/group/{id}", h.GetGroupById).Methods("GET")
	h.Router.HandleFunc("/api/v1/groups/group", h.CreateGroup).Methods("POST")
	h.Router.HandleFunc("/api/v1/groups/group/{id}", h.UpdateGroup).Methods("PUT")
	h.Router.HandleFunc("/api/v1/groups/group/{id}", h.DeleteGroup).Methods("DELETE")

	h.Router.HandleFunc("/api/v1/fields", h.GetFields).Methods("GET")
	h.Router.HandleFunc("/api/v1/fields/field", h.CreateField).Methods("POST")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
	log.Println("Shut down gracefully")

	return nil
}
