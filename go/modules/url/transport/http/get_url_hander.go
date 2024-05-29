package httphandler

import (
	"github.com/nghiatrann0502/demo-redis/go/common"
	"net/http"
)

func (h *handler) GetUrlHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("id")

	data, err := h.biz.GetUrlShorten(r.Context(), shortUrl)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	common.SimpleSuccess(w, http.StatusOK, data)
}
