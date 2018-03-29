# Version Control
GIT := git
CURRENT_BRANCH := $(shell $(GIT) rev-parse --abbrev-ref HEAD)
VERSION := $(shell $(GIT) describe --match 'v[0-9]*' --dirty='.m' --always)
REVISION := $(shell $(GIT) rev-parse HEAD)$(shell if ! $(GIT) diff --no-ext-diff --quiet --exit-code; then echo .m; fi)

# Golang
GO := go
BINDATA := go-bindata
DEP := dep
MOCKGEN := mockgen
GINKGO := ginkgo
PACKAGE := github.com/nomkhonwaan/myblog-server

# Docker 
DOCKER := docker
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
ifeq ($(shell which $(DEP)),)
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
	$(DEP) ensure

.PHONY: generate-bindata
generate-bindata:
ifeq ($(shell which $(BINDATA)),)
	$(GO) get -v -u github.com/jteeuwen/go-bindata
	$(GO) build -o $(GOPATH)/bin/go-bindata $(GOPATH)/src/github.com/jteeuwen/go-bindata/go-bindata/*.go
endif
	$(BINDATA) -o pkg/generated/bindata.go -pkg generated pkg/graphql/schema/... pkg/graphql/graphiql/...

.PHONY: generate-mock
generate-mock:
ifeq ($(shell which $(MOCKGEN)),)
	$(GO) get -v -u github.com/golang/mock/mockgen
endif
	$(MOCKGEN) -source vendor/github.com/nicksrandall/dataloader/dataloader.go -package dataloader_mock Interface > pkg/dataloader/mock/dataloader_mock.go
	make fix-dataloader-mock
	$(MOCKGEN) -source pkg/mongodb/collection.go -package mongodb_mock Collection > pkg/mongodb/mock/collection_mock.go
	$(MOCKGEN) -source pkg/mongodb/database.go -package mongodb_mock Database > pkg/mongodb/mock/database_mock.go
	$(MOCKGEN) -source pkg/mongodb/iter.go -package mongodb_mock Iter > pkg/mongodb/mock/iter_mock.go
	$(MOCKGEN) -source pkg/mongodb/query.go -package mongodb_mock Query > pkg/mongodb/mock/query_mock.go
	$(MOCKGEN) -source pkg/mongodb/session.go -package mongodb_mock Session > pkg/mongodb/mock/session_mock.go
	make fix-mongodb-mock

.PHONY: fix-dataloader-mock
fix-dataloader-mock:
	sed -i -e 's/gomock "github.com\/golang\/mock\/gomock"/gomock "github.com\/golang\/mock\/gomock"\'$$'\n        dataloader "github.com\/nicksrandall\/dataloader"/' pkg/dataloader/mock/dataloader_mock.go
	sed -i -e 's/arg1 Key/arg1 dataloader.Key/g; s/key Key/key dataloader.Key/g' pkg/dataloader/mock/dataloader_mock.go
	sed -i -e 's/arg1 Keys/arg1 dataloader.Keys/g' pkg/dataloader/mock/dataloader_mock.go
	sed -i -e 's/) Thunk/) dataloader.Thunk/g; s/(Thunk)/(dataloader.Thunk)/g' pkg/dataloader/mock/dataloader_mock.go
	sed -i -e 's/) ThunkMany/) dataloader.ThunkMany/g; s/(ThunkMany)/\(dataloader.ThunkMany)/g' pkg/dataloader/mock/dataloader_mock.go
	sed -i -e 's/) Interface/) dataloader.Interface/g; s/(Interface)/\(dataloader.Interface)/g' pkg/dataloader/mock/dataloader_mock.go

.PHONY: fix-mongodb-mock
fix-mongodb-mock:
	sed -i -e 's/gomock "github.com\/golang\/mock\/gomock"/gomock "github.com\/golang\/mock\/gomock"\'$$'\n        mongodb "github.com\/nomkhonwaan\/myblog-server\/pkg\/mongodb"/' pkg/mongodb/mock/collection_mock.go
	sed -i -e 's/ Query / mongodb.Query /g; s/(Query)/(mongodb.Query)/g' pkg/mongodb/mock/collection_mock.go
	sed -i -e 's/gomock "github.com\/golang\/mock\/gomock"/gomock "github.com\/golang\/mock\/gomock"\'$$'\n        mongodb "github.com\/nomkhonwaan\/myblog-server\/pkg\/mongodb"/' pkg/mongodb/mock/database_mock.go
	sed -i -e 's/ Collection / mongodb.Collection /g; s/(Collection)/(mongodb.Collection)/g' pkg/mongodb/mock/database_mock.go
	sed -i -e 's/gomock "github.com\/golang\/mock\/gomock"/gomock "github.com\/golang\/mock\/gomock"\'$$'\n        mongodb "github.com\/nomkhonwaan\/myblog-server\/pkg\/mongodb"/' pkg/mongodb/mock/query_mock.go
	sed -i -e 's/ Iter / mongodb.Iter /g; s/(Iter)/(mongodb.Iter)/g;' pkg/mongodb/mock/query_mock.go
	sed -i -e 's/ Query / mongodb.Query /g; s/(Query)/(mongodb.Query)/g' pkg/mongodb/mock/query_mock.go
	sed -i -e 's/gomock "github.com\/golang\/mock\/gomock"/gomock "github.com\/golang\/mock\/gomock"\'$$'\n        mongodb "github.com\/nomkhonwaan\/myblog-server\/pkg\/mongodb"/' pkg/mongodb/mock/session_mock.go
	sed -i -e 's/ Database / mongodb.Database /g; s/(Database)/(mongodb.Database)/g' pkg/mongodb/mock/session_mock.go
	sed -i -e 's/ Session / mongodb.Session /g; s/(Session)/(mongodb.Session)/g' pkg/mongodb/mock/session_mock.go

.PHONY: test
test:
ifeq (,$(wildcard vendor/github.com/golang/mock/gomock))
	$(GO) get -v -u github.com/golang/mock/gomock
endif
ifeq ($(shell which $(GINKGO)),)
	$(GO) get -v -u github.com/onsi/ginkgo/ginkgo
	$(GO) get -v -u github.com/onsi/gomega/...
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
