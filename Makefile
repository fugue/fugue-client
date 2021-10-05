
GO=GO111MODULE=on go
BINARY=fugue
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
INVITE_SRC=models/invite.go
UPDATE_FAMILY_SRC=models/update_family_input.go
CREATE_FAMILY_SRC=models/create_family_input.go
NOTIFICATION_SRC=models/notification.go

GOSWAGGER=docker run --rm -it \
	--volume $(shell pwd):/fugue-client \
	--user $(shell id -u):$(shell id -g) \
	--workdir /fugue-client \
	quay.io/goswagger/swagger:v0.23.0

$(BINARY): $(SOURCES)
	$(GO) build $(LD_FLAGS) -v -o $@

$(BINARY)-linux-amd64: $(SOURCES)
	GOOS=linux GOARCH=amd64 $(GO) build $(LD_FLAGS) -o $@

$(BINARY)-darwin-amd64: $(SOURCES)
	GOOS=darwin GOARCH=amd64 $(GO) build $(LD_FLAGS) -o $@

$(BINARY).exe: $(SOURCES)
	GOOS=windows GOARCH=386 $(GO) build $(LD_FLAGS) -o $@

.PHONY: release
release: $(BINARY)-linux-amd64 $(BINARY)-darwin-amd64 $(BINARY).exe

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
	sed -i "" "s/float64/int64/g" $(INVITE_SRC)
	sed -i "" "s/Recommended bool/Recommended *bool/g" $(UPDATE_FAMILY_SRC)
	sed -i "" "s/Recommended bool/Recommended *bool/g" $(CREATE_FAMILY_SRC)
	sed -i "" "s/AlwaysEnabled bool/AlwaysEnabled *bool/g" $(UPDATE_FAMILY_SRC)
	sed -i "" "s/AlwaysEnabled bool/AlwaysEnabled *bool/g" $(CREATE_FAMILY_SRC)
	sed -i "" 's|CreatedAt.*|// CreatedAt|g' $(NOTIFICATION_SRC)
	sed -i "" 's|Environments.*|Environments map\[string\]string `json:"environments"`|g' $(NOTIFICATION_SRC)

.PHONY: test
test:
	$(GO) test -test.v ./...

.PHONY: clean
clean:
	rm -f $(BINARY)
	rm -f $(BINARY)-linux-amd64
	rm -f $(BINARY)-darwin-amd64
	rm -f $(BINARY).exe
