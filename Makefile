TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=$(shell hostname)
NAMESPACE=packetfabric
NAME=packetfabric
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
OS_ARCH=linux_amd64

SOURCES := $(shell find . -name "*.go" -or -name "go.mod" -or -name "go.sum" \
	-or -name "Makefile")

default: install

#
# Tools
#

# external tool
define tool # 1: binary-name, 2: go-import-path
TOOLS += $(TOOLDIR)/$(1)

$(TOOLDIR)/$(1): Makefile
	GOBIN="$(CURDIR)/$(TOOLDIR)" go install "$(2)"
endef

$(eval $(call tool,gofumports,mvdan.cc/gofumpt/gofumports@latest))
$(eval $(call tool,golangci-lint,github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44))
$(eval $(call tool,gomod,github.com/Helcaraxan/gomod@latest))
$(eval $(call tool,tfplugindocs,github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.7))
$(eval $(call tool,tfproviderlint,github.com/bflad/tfproviderlint/cmd/tfproviderlint@latest))

build:
	go build -o ${BINARY}

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

.PHONY: test
test: 
	go test -v ./...               

.PHONY: testacc
testacc: 
	TF_ACC=1 go test -v ./... -timeout 120m

#
# Coverage
#

.PHONY: cov
cov: coverage.out

.PHONY: cov-html
cov-html: coverage.out
	go tool cover -html=coverage.out

.PHONY: cov-func
cov-func: coverage.out
	go tool cover -func=coverage.out

coverage.out: $(SOURCES)
	VCR=replay go test $(V) -timeout=120m \
			-covermode=count -coverprofile=coverage.out \
			$(TESTARGS) $(TEST)

#
# Dependencies
#

.PHONY: deps
deps:
	$(info Downloading dependencies)
	go mod download

.PHONY: deps-update
deps-update:
	$(info Downloading dependencies)
	go get -u -t ./...

.PHONY: deps-analyze
deps-analyze: $(TOOLDIR)/gomod
	gomod analyze

.PHONY: tidy
tidy:
	go mod tidy $(V)

.PHONY: verify
verify:
	go mod verify

.SILENT: check-tidy
.PHONY: check-tidy
check-tidy:
	cp go.mod go.mod.tidy-check
	cp go.sum go.sum.tidy-check
	go mod tidy
	( \
		diff go.mod go.mod.tidy-check && \
		diff go.sum go.sum.tidy-check && \
		rm -f go.mod go.sum && \
		mv go.mod.tidy-check go.mod && \
		mv go.sum.tidy-check go.sum \
	) || ( \
		rm -f go.mod go.sum && \
		mv go.mod.tidy-check go.mod && \
		mv go.sum.tidy-check go.sum; \
		exit 1 \
	)

#
# Documentation
#

# Force set provider configuration environment variables, as required vars get
# listed as "Optional" if the corresponding var is not empty.
.PHONY: docs
docs: $(TOOLDIR)/tfplugindocs
		tfplugindocs generate

.PHONY: check-docs
check-docs: $(TOOLDIR)/tfplugindocs
		tfplugindocs validate
