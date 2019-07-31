
GO=GO111MODULE=on go
BINARY=fugue
SWAGGER=swagger.yaml
SWAGGER_URL=https://api.riskmanager.fugue.co/v0/swagger
SOURCES=$(shell find . -name '*.go')
GOPATH?=$(shell go env GOPATH)

$(BINARY): $(SOURCES)
	$(GO) build -v -o fugue

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
	sed -i "" "s/,omitempty//g" $(shell find models -name "*.go")

.PHONY: clean
clean:
	rm -f $(BINARY)
