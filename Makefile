IMAGE_NAME_PREFIX ?= hiot-platform

SERVICES = auth
CONTAINERS = $(addprefix $(SERVICES),_container)

GO111MODULE ?= on
CGO_ENABLED ?= 0
GOARCH ?= amd64
PACKAGE ?= github.com/reecerose/hiot-platform
VERSION ?= $(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')
COMMIT_HASH ?= $(shell git rev-parse --short HEAD)
BUILD_TIMESTAMP ?= $(shell date '+%Y-%m-%dT%H:%M:%S')

.PHONY: all $(SERVICES) dc-up dc-down clean

define build_service
	GO111MODULE=$(GO111MODULE) CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) \
	go build -mod=vendor -ldflags "-s -w \
	-X 'main.Version=$(VERSION)' \
	-X 'main.CommitHash=$(COMMIT_HASH)' \
	-X 'main.BuildTimestamp=$(BUILD_TIMESTAMP)'" \
	-o build/$(1) cmd/$(1)/main.go
	upx -4 build/$(1)
endef

define build_container
	$(eval service=$(subst _container,,$(1)))

	docker build \
		--no-cache \
		--build-arg SERVICE=$(service) \
		--tag=$(IMAGE_NAME_PREFIX)/$(service) \
		-f docker/Dockerfile .
endef

dc-up:
	docker-compose -f docker/docker-compose.yml up

dc-down:
	docker-compose -f docker/docker-compose.yml down

clean:
	docker-compose -f docker/docker-compose.yml down --rmi all -v --remove-orphans

$(SERVICES):
	$(call build_service,$(@))

$(CONTAINERS):
	$(call build_container,$(@))