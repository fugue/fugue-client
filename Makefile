
GO=GO111MODULE=on go
BINARY=fugue
SWAGGER=swagger.yaml
SWAGGER_URL=https://api.riskmanager.fugue.co/v0/swagger
SOURCES=$(shell find . -name '*.go')
GOPATH?=$(shell go env GOPATH)
UPDATE_ENV_SRC=$(shell find models -name "update_environment_input.go")

$(BINARY): $(SOURCES)
	$(GO) build -v -o fugue

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
	swagger generate client -f $(SWAGGER)
	# Workaround for deficiencies in generated swagger types
	sed -i "" "s/BaselineID string/BaselineID *string/g" $(UPDATE_ENV_SRC)
	sed -i "" "s/Remediation bool/Remediation *bool/g" $(UPDATE_ENV_SRC)
	sed -i "" "s/ScanScheduleEnabled bool/ScanScheduleEnabled *bool/g" $(UPDATE_ENV_SRC)

.PHONY: clean
clean:
	rm -f $(BINARY)
