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

	printer, printerOutChan := ingestor.NewPrinter(printerBufferSize)
	fanOut[0] = printerOutChan
	go printer.Print()
	go ingestor.StartIngestor(httpBus, fanOut)

	healthcheck := api.NewHealthcheck()

	healthcheck.RegisterService("printer", func() api.ServiceHealthCheckStatus {
		return api.ServiceHealthy
	})

	api.StartAPI(httpBus, healthcheck)
}
