package ingestor

import (
	"fmt"
	"time"

	openapi "github.com/lmsilva-oss/messages/gen/go/proto/v1/openapi/go"
)

type Printer struct {
	input chan openapi.ProtoPayload
}

func NewPrinter(printerBufferSize int) (*Printer, chan openapi.ProtoPayload) {
	newChan := make(chan openapi.ProtoPayload, printerBufferSize)
	printer := &Printer{
		input: newChan,
	}
	printer.println("WARN", "Printer times are UTC")

	return printer, newChan
}

func (p *Printer) Print() {
	p.println("Starting printer")
	for payload := range p.input {
		p.println("source:", payload.Source)
		p.println("payload:", payload.Payload)
	}
}

func (p *Printer) println(a ...any) {
	val := []any{p.getTimestamp(), "Printer:"}
	val = append(val, a...)
	fmt.Println(val...)
}

func (p *Printer) getTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.9999")
}
