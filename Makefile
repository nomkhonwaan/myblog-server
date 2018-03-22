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
PACKAGE := github.com/nomkhonwaan/myblog-server

# Docker 
DOCKER := $(shell which docker)
DOCKER_IMAGE_REPOSITORY := nomkhonwaan/myblog-server
DOCKER_IMAGE_TAG := latest

ifeq ($(CURRENT_BRANCH), master)
	$(eval DOCKER_IMAGE_TAG := $(VERSION))
endif

.PHONY: default
default: generate-bindata
	$(GO) run cmd/myblog/main.go

.PHONY: install
install:
ifeq ($(DEP),)
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	$(eval DEP := $(shell which dep))
endif
	$(DEP) ensure

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
	# $(MOCKGEN) -source=pkg/post/post.go -package post_test Repositorier > pkg/post/post_mock.go

.PHONY: test
test: generate-bindata generate-mock
ifeq ($(GINKGO),)
	$(GO) get -v -u github.com/onsi/ginkgo/ginkgo
	$(GO) get -v -u github.com/onsi/gomega/...
	$(eval GINKGO := $(shell which ginkgo))
endif
	# $(GINKGO) ./cmd/...
	$(GINKGO) pkg/...

.PHONY: clean
clean:
	rm -f pkg/generated/bindata.go

.PHONY: build
build: clean generate-bindata
	$(GO) build \
		-o $(GOPATH)/bin/myblog-server \
		-ldflags " \
			-X $(PACKAGE)/cmd/myblog/app.version=$(VERSION) \
			-X $(PACKAGE)/cmd/myblog/app.revision=$(REVISION) \
		" \
		cmd/myblog/main.go

.PHONY: build-docker
build-docker:
	$(DOCKER) build \
		--build-arg="VERSION=$(VERSION)" \
		--build-arg="REVISION=$(REVISION)" \
		--tag $(DOCKER_IMAGE_REPOSITORY):$(DOCKER_IMAGE_TAG) \
		.

.PHONY: publish-to-registry
publish-to-registry:
	$(DOCKER) push $(DOCKER_IMAGE_REPOSITORY):$(DOCKER_IMAGE_TAG)