package main

import (
	"github.com/lmsilva-oss/messages/api"
	openapi "github.com/lmsilva-oss/messages/gen/go/proto/v1/openapi/go"
	"github.com/lmsilva-oss/messages/ingestor"
)

func main() {
	const HTTP_BUS_SIZE = 200
	httpBus := make(chan openapi.ProtoPayload, HTTP_BUS_SIZE)
	go ingestor.StartIngestor(httpBus)
	api.StartAPI(httpBus)
}
