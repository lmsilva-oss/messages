package main

import (
	"github.com/lmsilva-oss/messages/api"
	openapi "github.com/lmsilva-oss/messages/gen/go/proto/v1/openapi/go"
	"github.com/lmsilva-oss/messages/ingestor"
)

func main() {
	const HTTP_BUS_SIZE = 200
	const printerBufferSize = 10
	httpBus := make(chan openapi.ProtoPayload, HTTP_BUS_SIZE)

	fanOut := make([]chan openapi.ProtoPayload, 1)
	fanOut[0] = make(chan openapi.ProtoPayload, printerBufferSize)

	printer := ingestor.NewPrinter(fanOut[0])
	go printer.Print()
	go ingestor.StartIngestor(httpBus, fanOut)
	api.StartAPI(httpBus)
}
