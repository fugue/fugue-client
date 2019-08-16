
GO=GO111MODULE=on go
BINARY=fugue
VERSION=$(shell cat VERSION)
SHORT_COMMIT=$(shell git rev-parse HEAD | cut -c 1-8)
LD_FLAGS=-ldflags "-X main.version=$(VERSION) -X main.commit=$(SHORT_COMMIT)"

SWAGGER=swagger.yaml
SWAGGER_URL=https://api.riskmanager.fugue.co/v0/swagger
SOURCES=$(shell find . -name '*.go')
GOPATH?=$(shell go env GOPATH)

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
	swagger generate client -f $(SWAGGER)
	# sed -i "" "s/,omitempty//g" $(shell find models -name "*.go")

.PHONY: clean
clean:
	rm -f $(BINARY)
