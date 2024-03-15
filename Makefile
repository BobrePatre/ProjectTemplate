all: generate-proto generate-typescript



# Base Paths
PROTO_DIR=api
PROTO_FILES=$(shell find $(PROTO_DIR) -name '*.proto')
OUT_DIR=pkg/api

# Generated Libraries Paths
GO_OUT_DIR=$(OUT_DIR)/golang
DART_OUT_DIR=$(OUT_DIR)/dart
TYPESCRIPT_OUT_DIR=$(OUT_DIR)/typescript
OPENAPI_OUT_DIR=api/openapiv2

# Vendoring Paths
VENDOR_DIR=pkg/vendor
GOOGLE_APIS_DIR=$(VENDOR_DIR)/googleapis
GRPC_GATEWAY_DIR=$(VENDOR_DIR)/grpc-gateway



# Rules To Create Directories For PROTOC
$(GO_OUT_DIR):
	mkdir -p $(GO_OUT_DIR)

$(DART_OUT_DIR):
	mkdir -p $(DART_OUT_DIR)

$(OPENAPI_OUT_DIR):
	mkdir -p $(OPENAPI_OUT_DIR)



# Installing Tools For Code Generation; Using Golang
setup-golang-tools:
	@echo "Setup golang tools"
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest


# Installing Tools For Code Generation; Using NPM
setup-npm-tools:
	@echo "Setup openapi generator"
	@npm install -g @openapitools/openapi-generator-cli

# Setup All Tools
setup-tools: validate-project setup-golang-tools setup-npm-tools vendor

# Vendoring Dependencies
vendor:
	@echo "Vendoring dependencies"
	@echo "After vendoring completes, if you use jetbrains ide, enhance IDE dependency recognition with this quick setup: Go to Settings > Languages & Frameworks > Protocol Buffers > Import Paths, and add the path to your new dependency folder."
	@if [ ! -d "$(VENDOR_DIR)" ]; then \
  		echo "Creating vendor directory"; \
		mkdir -p $(VENDOR_DIR); \
	fi
	@if [ ! -d "$(GOOGLE_APIS_DIR)" ] || [ -z "$$(ls -A $(GOOGLE_APIS_DIR) 2>/dev/null)"  ]; then \
  		echo "Vendoring googleapis"; \
		git clone https://github.com/googleapis/googleapis.git $(GOOGLE_APIS_DIR); \
		echo "Done vendoring googleapis"; \
	fi
	@if [ ! -d "$(GRPC_GATEWAY_DIR)" ] || [ -z "$$(ls -A $(GRPC_GATEWAY_DIR) 2>/dev/null)" ]; then \
  		echo "Vendoring grpc-gateway"; \
		git clone https://github.com/grpc-ecosystem/grpc-gateway.git $(GRPC_GATEWAY_DIR); \
	fi
	@echo "Dependencies vendored"

# Generating Libraries From Protobuf
generate-proto: $(GO_OUT_DIR) $(DART_OUT_DIR) $(OPENAPI_OUT_DIR)
	@echo "Generating Protobuf libraries..."
	@protoc -I$(PROTO_DIR) \
	-I$(GOOGLE_APIS_DIR) \
	-I$(GRPC_GATEWAY_DIR) \
	--go_out=$(GO_OUT_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(GO_OUT_DIR) \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=$(GO_OUT_DIR) \
	--grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=$(OPENAPI_OUT_DIR) \
	--openapiv2_opt=use_go_templates=true,preserve_rpc_order=true \
	--dart_out=grpc:$(DART_OUT_DIR) \
	$(PROTO_FILES)



# Generating Libraries From OpenAPI
generate-typescript:
	@echo "Generating TypeScript Libraries From OpenAPI...";
	@SWAGGER_FILES=$$(find $(OPENAPI_OUT_DIR) -name "*.swagger.json"); \
	for file in $$SWAGGER_FILES; do \
		output_path=$(TYPESCRIPT_OUT_DIR)/$$(basename $$file .swagger.json); \
		echo "Generating for $$file..."; \
		echo "Generating to path: $$output_path"; \
		openapi-generator-cli generate \
			-i $$file \
			-g typescript \
			-o $$output_path \
			--additional-properties=usePromises=true,useES6=true; \
	done


validate-project:
	@echo "Validating project..."
	@if [ ! -d $(PROTO_DIR) ]; then \
		echo "Protobuf directory not found !!!!!!"; \
		echo "Please, place your proto files in the $(PROTO_DIR) directory !!!!!!"; \
		exit 1; \
	fi
	@echo "Project Is Valid!"



