package ingestor

import (
	openapi "github.com/lmsilva-oss/messages/gen/go/proto/v1/openapi/go"
)

type Ingestor struct {
	// TODO: this chan could also receive contexts,
	// to cancel ingestion if the context asks for.
	// check usefulness by cancelling on a big payload
	input  chan openapi.ProtoPayload
	fanOut []chan openapi.ProtoPayload
}

func StartIngestor(bus chan openapi.ProtoPayload, fanOut []chan openapi.ProtoPayload) {
	ingestor := Ingestor{
		input:  bus,
		fanOut: fanOut,
	}
	ingestor.ingest()
}

func (i *Ingestor) ingest() {
	for payload := range i.input {
		for _, outChan := range i.fanOut {
			outChan <- payload
		}
	}
	for _, outChan := range i.fanOut {
		close(outChan)
	}
}
