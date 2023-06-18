package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	campaigns, err := h.CampaignService.GetById(id)

	return campaigns, http.StatusOK, err
}
