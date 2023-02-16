
setup:
	brew install protobuf
	brew install openapi-generator

generate: clean
	protoc -I . \
	  --csharp_out ./gen/csharp/ \
		--go_out ./gen/go/ \
		--go_opt paths=source_relative \
		--go-grpc_out ./gen/go/ \
		--go-grpc_opt paths=source_relative \
		--openapiv2_out ./gen/go \
  	--openapiv2_opt grpc_api_configuration=proto/v1/service.yaml \
		proto/v1/service.proto
	openapi-generator generate -g go-server \
    -o gen/go/proto/v1/openapi/ -i gen/go/proto/v1/service.swagger.json
	rm gen/go/proto/v1/openapi/go.mod \
	  gen/go/proto/v1/openapi/main.go \
		gen/go/proto/v1/openapi/README.md \
		gen/go/proto/v1/openapi/Dockerfile \
		gen/go/proto/v1/openapi/.openapi-generator-ignore \
		gen/go/proto/v1/openapi/go/api_greeter_service_service.go
	rm -rf gen/go/proto/v1/openapi/.openapi-generator \
	  gen/go/proto/v1/openapi/api
	cp ./gen/csharp/Service.cs ./example/client/messages/

	
clean:
	find . -name "*.pb.*" -exec rm {} +;
	find . -name "*.swagger.*" -exec rm {} +;