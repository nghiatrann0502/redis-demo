package httphandler

import (
	"github.com/nghiatrann0502/demo-redis/go/common"
	"net/http"
)

func (h *handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Url string `json:"url"`
	}

	err := common.ReadJSON(r, &body)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	shortUrl, err := h.biz.CreateUrlShorten(r.Context(), body.Url)

	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	common.SimpleSuccess(w, http.StatusOK, shortUrl)
}
