package api

import (
	"context"

	openapi "github.com/lmsilva-oss/messages/gen/go/proto/v1/openapi/go"
)

type GreeterService struct {
	input chan openapi.ProtoPayload
}

// NewGreeterService creates a default api service
func NewGreeterService(input chan openapi.ProtoPayload) openapi.GreeterServiceApiServicer {
	return &GreeterService{
		input: input,
	}
}

// GreeterServiceSayHello - Sends a greeting
func (s *GreeterService) GreeterServiceSayHello(ctx context.Context, body openapi.ProtoSayHelloRequest) (openapi.ImplResponse, error) {
	// TODO - update GreeterServiceSayHello with the required logic for this service method.
	// Add api_greeter_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	for _, payload := range body.Data {
		s.input <- payload
	}

	//TODO: Uncomment the next line to return response Response(200, map[string]interface{}{}) or use other options such as http.Ok ...
	return openapi.Response(202, map[string]interface{}{}), nil
}
