# Version Control
GIT := $(shell which git)
CURRENT_BRANCH := $(shell $(GIT) rev-parse --abbrev-ref HEAD)
VERSION := $(shell $(GIT) describe --match 'v[0-9]*' --dirty='.m' --always)
REVISION := $(shell $(GIT) rev-parse HEAD)$(shell if ! $(GIT) diff --no-ext-diff --quiet --exit-code; then echo .m; fi)

# Golang
GO := $(shell which go)
BINDATA := $(shell which go-bindata)
DEP := $(shell which dep)
MOCKGEN := $(shell which mockgen)
GINKGO := $(shell which ginkgo)

# Docker 
DOCKER := $(shell which docker)
DOCKER_IMAGE_REPOSITORY := nomkhonwaan/myblog-server
DOCKER_IMAGE_TAG := latest

if [ "$(CURRENT_BRANCH)" == "master" ]; \
then \
	$(eval DOCKER_IMAGE_TAG := $(VERSION)); \
fi

.PHONY: default
default: generate-bindata
	$(GO) run cmd/myblog/main.go

.PHONY: generate-bindata
generate-bindata:
ifeq ($(BINDATA),)
	$(GO) get -v -u github.com/jteeuwen/go-bindata
	$(eval BINDATA := $(shell which go-bindata))
endif
	$(BINDATA) -o pkg/generated/bindata.go -pkg generated pkg/graphql/schema/... pkg/graphql/graphiql/...

.PHONY: generate-mock
generate-mock:
ifeq ($(MOCKGEN),)
	$(GO) get -v -u github.com/golang/mock/mockgen
	$(eval MOCKGEN := $(shell which mockgen))
endif
	# $(MOCKGEN) -source=pkg/post/repository.go -package post Repositorier > pkg/post/repository_mock.go

.PHONY: test
test: generate-bindata generate-mock
ifeq ($(GINKGO),)
	$(GO) get -v -u github.com/onsi/ginkgo/ginkgo
	$(GO) get -v -u github.com/onsi/gomega/...
	$(eval GINKGO := $(shell which ginkgo))
endif
	$(GINKGO) cmd/...
	$(GINKGO) pkg/...

.PHONY: clean
clean:
	rm -f pkg/generated/bindata.go

.PHONY: build
build: clean generate-bindata
	$(GO) build \
		-o $(GOPATH)/bin/myblog-server \
		-ldflags " \
			-X main.version=$(VERSION) \
			-X main.revision=$(REVISION) \
		" \
		cmd/myblog/main.go

.PHONY: build-docker
build-docker:
	$(DOCKER) build \
		--build-arg="VERSION=$(VERSION)" \
		--build-arg="REVISION=$(REVISION)" \
		--tag $(DOCKER_IMAGE_REPOSITORY):$(DOCKER_IMAGE_TAG) \
		.