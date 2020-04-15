
GO=GO111MODULE=on go
BINARY=fugue
VERSION=$(shell cat VERSION)
SHORT_COMMIT=$(shell git rev-parse HEAD | cut -c 1-8)
LD_FLAGS=-ldflags "-X main.version=$(VERSION) -X main.commit=$(SHORT_COMMIT)"

SWAGGER=swagger.yaml
SWAGGER_URL=https://api.riskmanager.fugue.co/v0/swagger
SOURCES=$(shell find . -name '*.go')
GOPATH?=$(shell go env GOPATH)
UPDATE_ENV_SRC=$(shell find models -name "update_environment_input.go")
UPDATE_RULE_SRC=$(shell find models -name "update_custom_rule_input.go")

$(BINARY): $(SOURCES)
	$(GO) build $(LD_FLAGS) -v -o fugue

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
	swagger generate client -f $(SWAGGER)
	# Workaround for deficiencies in generated swagger types
	sed -i "" "s/BaselineID string/BaselineID *string/g" $(UPDATE_ENV_SRC)
	sed -i "" "s/Remediation bool/Remediation *bool/g" $(UPDATE_ENV_SRC)
	sed -i "" "s/ScanScheduleEnabled bool/ScanScheduleEnabled *bool/g" $(UPDATE_ENV_SRC)
	sed -i "" "s/ScanScheduleEnabled bool/ScanScheduleEnabled *bool/g" $(UPDATE_RULE_SRC)

.PHONY: test
test:
	$(GO) test -test.v ./...

.PHONY: clean
clean:
	rm -f $(BINARY)
