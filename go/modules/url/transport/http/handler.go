package httphandler

import (
	"github.com/nghiatrann0502/demo-redis/go/common"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/business"
	"log"
	"net/http"
)

type handler struct {
	biz business.URLBusiness
}

func NewHandler(biz business.URLBusiness) *handler {
	return &handler{biz}
}

func (h *handler) RegisterRouter(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/ping", h.PingHandler)
	mux.HandleFunc("POST /api/v1/shorten", h.CreateHandler)
	mux.HandleFunc("GET /api/v1/shorten/{id}", h.GetUrlHandler)
	mux.HandleFunc("PATCH /api/v1/shorten/{id}", h.UpdateUrlHandler)
}

func (h *handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Pong")
	common.SimpleSuccess(w, http.StatusOK, "pong")
}
