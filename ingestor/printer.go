package ingestor

import (
	"fmt"

	openapi "github.com/lmsilva-oss/messages/gen/go/proto/v1/openapi/go"
)

type Printer struct {
	input chan openapi.ProtoPayload
}

func NewPrinter(input chan openapi.ProtoPayload) *Printer {
	return &Printer{
		input: input,
	}
}

func (p *Printer) Print() {
	fmt.Println("Starting printer")
	for payload := range p.input {
		fmt.Println(payload.Source)
		fmt.Println(payload.Payload)
	}
}
