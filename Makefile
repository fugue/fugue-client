
GO=GO111MODULE=on go
BINARY=fugue
DOCKER_IMAGE=fugue/fugue-client
VERSION=$(shell cat VERSION)
SHORT_COMMIT=$(shell git rev-parse HEAD | cut -c 1-8)
LD_FLAGS=-ldflags "-X main.version=$(VERSION) -X main.commit=$(SHORT_COMMIT) -extldflags '-static'"

SWAGGER=swagger.yaml
SWAGGER_URL=https://api.riskmanager.fugue.co/v0/swagger
SOURCES=$(shell find . -name '*.go')
GOPATH?=$(shell go env GOPATH)

UPDATE_ENV_SRC=models/update_environment_input.go
UPDATE_RULE_SRC=models/update_custom_rule_input.go
CREATE_ENV_SRC=models/create_environment_input.go
INVITE_SRC=models/invite.go
UPDATE_FAMILY_SRC=models/update_family_input.go
CREATE_FAMILY_SRC=models/create_family_input.go

GOSWAGGER=docker run --rm -it \
	--volume $(shell pwd):/fugue-client \
	--user $(shell id -u):$(shell id -g) \
	--workdir /fugue-client \
	quay.io/goswagger/swagger:v0.23.0

$(BINARY): $(SOURCES)
	$(GO) build $(LD_FLAGS) -v -o $@

$(BINARY)-linux-amd64: $(SOURCES)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build $(LD_FLAGS) -o $@

$(BINARY)-darwin-amd64: $(SOURCES)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build $(LD_FLAGS) -o $@

$(BINARY).exe: $(SOURCES)
	CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GO) build $(LD_FLAGS) -o $@

.PHONY: help
help: ## Show this help
	@egrep '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: release
release: ## Create release binaries
release: $(BINARY)-linux-amd64 $(BINARY)-darwin-amd64 $(BINARY).exe

.PHONY: docker
docker: ## Build and tag docker image
docker: $(BINARY)-linux-amd64
	docker build -t $(DOCKER_IMAGE):v$(VERSION) -t $(DOCKER_IMAGE):latest .

.PHONY: docker-publish
docker-publish: ## Publish docker image to Docker Hub
docker-publish: docker
	docker push $(DOCKER_IMAGE):v$(VERSION)
	docker push $(DOCKER_IMAGE):latest

.PHONY: build
build: ## Build native binary
build: $(BINARY)

.PHONY: install
install: ## Insall binary to $GOPATH/bin
install: $(BINARY)
	cp $(BINARY) $(GOPATH)/bin/

$(SWAGGER): ## Download Swagger definitions from web
	wget -q -O $(SWAGGER) $(SWAGGER_URL)

.PHONY: validate
validate: ## Validate Swagger definitions
validate: $(SWAGGER)
	swagger validate $(SWAGGER)

.PHONY: gen
gen: ## Generate Go Swagger interface
gen: $(SWAGGER)
	# go-swagger: https://goswagger.io/
	$(GOSWAGGER) generate client -f $(SWAGGER)
	# Workaround for deficiencies in generated swagger types
	sed -i".bak" "s/BaselineID string/BaselineID *string/g" $(UPDATE_ENV_SRC)
	sed -i".bak" "s/Remediation bool/Remediation *bool/g" $(UPDATE_ENV_SRC)
	sed -i".bak" "s/ScanScheduleEnabled bool/ScanScheduleEnabled *bool/g" $(UPDATE_ENV_SRC)
	sed -i".bak" "s/ScanScheduleEnabled bool/ScanScheduleEnabled *bool/g" $(UPDATE_RULE_SRC)
	sed -i".bak" "s/ScanInterval int64/ScanInterval *int64/g" $(CREATE_ENV_SRC)
	sed -i".bak" "s/ScanScheduleEnabled bool/ScanScheduleEnabled *bool/g" $(CREATE_ENV_SRC)
	sed -i".bak" "s/int64(m.ScanInterval)/int64(*m.ScanInterval)/g" $(CREATE_ENV_SRC)
	sed -i".bak" "s/float64/int64/g" $(INVITE_SRC)
	sed -i".bak" "s/Recommended bool/Recommended *bool/g" $(UPDATE_FAMILY_SRC)
	sed -i".bak" "s/Recommended bool/Recommended *bool/g" $(CREATE_FAMILY_SRC)
	sed -i".bak" "s/AlwaysEnabled bool/AlwaysEnabled *bool/g" $(UPDATE_FAMILY_SRC)
	sed -i".bak" "s/AlwaysEnabled bool/AlwaysEnabled *bool/g" $(CREATE_FAMILY_SRC)

.PHONY: test
test: ## Run tests
test:
	$(GO) test -test.v ./...

.PHONY: clean
clean: ## Remove generated executables
	rm -f $(BINARY)
	rm -f $(BINARY)-linux-amd64
	rm -f $(BINARY)-darwin-amd64
	rm -f $(BINARY).exe
	docker images $(DOCKER_IMAGE) --format '{{.Repository}}:{{.Tag}}' | xargs docker image rm
