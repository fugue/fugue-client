
GO=GO111MODULE=on go
BINARY=fugue
WINDOWS_BINARY=fugue.exe
VERSION=$(shell cat VERSION)
SHORT_COMMIT=$(shell git rev-parse HEAD | cut -c 1-8)
LD_FLAGS=-ldflags "-X main.version=$(VERSION) -X main.commit=$(SHORT_COMMIT)"

SWAGGER=swagger.yaml
SWAGGER_URL=https://api.riskmanager.fugue.co/v0/swagger
SOURCES=$(shell find . -name '*.go')
GOPATH?=$(shell go env GOPATH)

UPDATE_ENV_SRC=models/update_environment_input.go
UPDATE_RULE_SRC=models/update_custom_rule_input.go
CREATE_ENV_SRC=models/create_environment_input.go

GOSWAGGER=docker run --rm -it \
	--volume $(shell pwd):/fugue-client \
	--user $(shell id -u):$(shell id -g) \
	--workdir /fugue-client \
	quay.io/goswagger/swagger:v0.23.0

$(BINARY): $(SOURCES)
	$(GO) build $(LD_FLAGS) -v -o $@

$(WINDOWS_BINARY): $(SOURCES)
	GOOS=windows GOARCH=386 $(GO) build $(LD_FLAGS) -v -o $@

.PHONY: build
build: $(BINARY)

.PHONY: install
install: $(BINARY)
	cp $(BINARY) $(GOPATH)/bin/

$(SWAGGER):
	wget -q -O $(SWAGGER) $(SWAGGER_URL)

.PHONY: validate
validate: $(SWAGGER)
	swagger validate $(SWAGGER)

.PHONY: gen
gen: $(SWAGGER)
	# go-swagger: https://goswagger.io/
	$(GOSWAGGER) generate client -f $(SWAGGER)
	# Workaround for deficiencies in generated swagger types
	sed -i "" "s/BaselineID string/BaselineID *string/g" $(UPDATE_ENV_SRC)
	sed -i "" "s/Remediation bool/Remediation *bool/g" $(UPDATE_ENV_SRC)
	sed -i "" "s/ScanScheduleEnabled bool/ScanScheduleEnabled *bool/g" $(UPDATE_ENV_SRC)
	sed -i "" "s/ScanScheduleEnabled bool/ScanScheduleEnabled *bool/g" $(UPDATE_RULE_SRC)
	sed -i "" "s/ScanInterval int64/ScanInterval *int64/g" $(CREATE_ENV_SRC)
	sed -i "" "s/ScanScheduleEnabled bool/ScanScheduleEnabled *bool/g" $(CREATE_ENV_SRC)
	sed -i "" "s/int64(m.ScanInterval)/int64(*m.ScanInterval)/g" $(CREATE_ENV_SRC)

.PHONY: test
test:
	$(GO) test -test.v ./...

.PHONY: clean
clean:
	rm -f $(BINARY)
	rm -f $(WINDOWS_BINARY)
