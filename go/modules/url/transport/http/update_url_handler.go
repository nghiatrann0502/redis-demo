package httphandler

import (
	"github.com/nghiatrann0502/demo-redis/go/common"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/model"
	"net/http"
)

func (h *handler) UpdateUrlHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("id")
	var input struct {
		Url string `json:"url"`
	}

	if err := common.ReadJSON(r, &input); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
	}

	data, err := h.biz.UpdateById(r.Context(), model.UrlUpdate{URL: input.Url, ShortURL: shortUrl})
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	common.SimpleSuccess(w, http.StatusOK, data)
}
