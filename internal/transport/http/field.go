package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/arturfil/gorilla_soccer/internal/field"
)

type FieldService interface {
	GetFields(ctx context.Context) ([]field.Field, error)
	CreateField(ctx context.Context, field field.Field) (field.Field, error)
}

type Message struct {
	Key  string `json:"key"`
	Data string `json:"data"`
}

func (h *Handler) GetFields(w http.ResponseWriter, r *http.Request) {
	all, err := h.FieldService.GetFields(r.Context())
	if err != nil {
		log.Print(err)
		return
	}
	log.Println(all)
	w.Header().Set("Content-Type", "application/json")
	encjson, _ := json.Marshal(all)
	w.Write(encjson)
}

func (h *Handler) CreateField(w http.ResponseWriter, r *http.Request) {
	var fieldBody field.Field
	if err := json.NewDecoder(r.Body).Decode(&fieldBody); err != nil {
		return
	}
	field, err := h.FieldService.CreateField(r.Context(), fieldBody)
	if err != nil {
		log.Print(err)
		return
	}
	if err := json.NewEncoder(w).Encode(field); err != nil {
		panic(err)
	}
}
