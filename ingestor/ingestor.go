package ingestor

import (
	"fmt"

	openapi "github.com/lmsilva-oss/messages/gen/go/proto/v1/openapi/go"
)

type Ingestor struct {
	input chan openapi.ProtoPayload
}

func StartIngestor(bus chan openapi.ProtoPayload) {
	ingestor := Ingestor{
		input: bus,
	}
	ingestor.Ingest()
}

func (i *Ingestor) Ingest() {
	for payload := range i.input {
		fmt.Println(payload.Source)
		fmt.Println(payload.Payload)
	}
}
