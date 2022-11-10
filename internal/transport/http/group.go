package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/arturfil/gorilla_soccer/internal/group"
	"github.com/gorilla/mux"
)

type GroupService interface {
	GetGroups(ctx context.Context) ([]group.Group, error)
	GetGroupById(ctx context.Context, id string) (group.Group, error)
	CreateGroup(context.Context, group.Group) (group.Group, error)
	UpdateGroup(ctx context.Context, id string, updatedGroup group.Group) (group.Group, error)
	DeleteGroup(ctx context.Context, id string) error
}

func (h *Handler) GetGroups(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetGroupById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	group, err := h.GroupService.GetGroupById(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(group); err != nil {
		panic(err)
	}
}

func (h *Handler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group group.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		return
	}
	group, err := h.GroupService.CreateGroup(r.Context(), group)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(group); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateGroup(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteGroup(w http.ResponseWriter, r *http.Request) {

}
