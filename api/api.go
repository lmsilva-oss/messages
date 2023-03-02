package api

import (
	"log"
	"net/http"

	openapi "github.com/lmsilva-oss/messages/gen/go/proto/v1/openapi/go"
)

func StartAPI(bus chan openapi.ProtoPayload, healthcheck *Healthcheck) {
	GreeterService := NewGreeterService(bus)
	GreeterServiceApiController := openapi.NewGreeterServiceApiController(GreeterService)

	router := openapi.NewRouter(GreeterServiceApiController)

	router.Handle("/healthz", http.HandlerFunc(healthcheck.handlerFunc()))

	log.Fatal(http.ListenAndServe(":8080", router))
}
