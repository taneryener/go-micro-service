
build-protos:
	docker run --volume "./:/workspace" --workdir /workspace taneryener/buf-go format -w
	docker run --volume "./:/workspace" --workdir /workspace taneryener/buf-go generate proto

build-swagger:
	buf generate
	docker run --rm -v "./:/local" swaggerapi/swagger-codegen-cli generate -i /local/api/swagger_output/example_service/v1/example_service.swagger.json -l swagger -o /local/api/swagger_output/example_service/v1

clear-protos:
	@rm -rf ./pkg/*