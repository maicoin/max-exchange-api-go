PHONY += all test clean

CURDIR := $(shell pwd)
BUILD_DIR := $(CURDIR)/build
GOBIN := $(BUILD_DIR)/bin
HOST_OS := $(shell uname -s)
TARGETS := $(sort $(notdir $(wildcard ./cmd/*)))
PHONY += $(TARGETS)
CLIENT_PACKAGE_NAME := api

export GOBIN
export PATH := $(GOBIN):$(PATH)

all: $(TARGETS)

.SECONDEXPANSION:
$(TARGETS): $(addprefix $(GOBIN)/,$$@)
$(GOBIN):
	@mkdir -p $@

$(GOBIN)/%: $(GOBIN) FORCE
	@echo "Building $(notdir $@)"
	@go build -v -o $@ ./cmd/$(notdir $@)
	@echo "Done building."
	@echo "Run \"$(subst $(CURDIR),.,$@)\" to launch $(notdir $@)."

DOCS_DIR := $(CURDIR)/docs
$(DOCS_DIR):
	@mkdir -p $@

SWAGGER_SPEC_URL := https://max-api.maicoin.com/api/doc/swagger
SWAGGER_SPEC := $(DOCS_DIR)/swagger-spec.json
$(SWAGGER_SPEC): $(DOCS_DIR) FORCE
	@/bin/echo -n "Get API specification from $(SWAGGER_SPEC_URL)" && curl -s $(SWAGGER_SPEC_URL) > $(SWAGGER_SPEC) && /bin/echo " ... done"

PHONY += swagger-spec
swagger-spec: $(SWAGGER_SPEC)

TEMPLATES_DIR := $(CURDIR)/templates

PHONY += sdk
sdk:
	docker run --rm \
	-v $(CURDIR):/local \
	swaggerapi/swagger-codegen-cli:v2.3.1 generate \
    -i /local/docs/$(notdir $(SWAGGER_SPEC)) \
    -l go \
	-t /local/$(notdir $(TEMPLATES_DIR)) \
	-DpackageName=$(CLIENT_PACKAGE_NAME) \
    -o /local/$(CLIENT_PACKAGE_NAME)

clean:
	@rm -rf $(GOBIN)

test:
	@go test -v .

PHONY: help
help:
	@echo  'Generic targets:'
	@echo  '  all                           - Build all targets marked with [*]'
	@echo  '* sdk                           - Build MAX Go SDK'
	@echo  '* swagger-spec                  - Get latest OpenAPI specification of MAX'
	@echo  ''
	@echo  'Test targets:'
	@echo  '  test                          - Run all unit tests'
	@echo  ''
	@echo  'Cleaning targets:'
	@echo  '  clean                         - Remove built executables'
	@echo  ''
	@echo  'Execute "make" or "make all" to build all targets marked with [*] '
	@echo  'For further info see the ./README.md file'

.PHONY: $(PHONY)

.PHONY: FORCE
FORCE:
