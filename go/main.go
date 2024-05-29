package main

import (
	"github.com/nghiatrann0502/demo-redis/go/modules/url/business"
	"github.com/nghiatrann0502/demo-redis/go/modules/url/repository/mysql"
	redis2 "github.com/nghiatrann0502/demo-redis/go/modules/url/repository/redis"
	httphandler "github.com/nghiatrann0502/demo-redis/go/modules/url/transport/http"
	"log"
	"net/http"
)

var (
	httpAddr = ":8080"
)

func main() {
	mux := http.NewServeMux()

	repository, err := mysql.NewMySqlStorage()
	if err != nil {
		log.Panic(err)
	}

	redis, err := redis2.NewRedisStorage()
	if err != nil {
		log.Panic(err)
	}

	biz := business.NewBusiness(repository, redis)
	handler := httphandler.NewHandler(biz)
	handler.RegisterRouter(mux)

	log.Println("starting url server on", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
