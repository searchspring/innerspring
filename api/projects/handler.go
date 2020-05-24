package projects

import (
	"net/http"

	"searchspring.com/innerspring/model"
)

// Handler handles a request
type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request, u *model.User) {
	w.Write([]byte(`[{"name":"project 1", "id":"1"},{"name":"project 2","id":"2"}]`))
}

func (h *Handler) GetProject(w http.ResponseWriter, r *http.Request, u *model.User) {
	w.Write([]byte(`{"name":"project 1", "id":"1"}`))
}
