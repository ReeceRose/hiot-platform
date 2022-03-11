IMAGE_NAME_PREFIX ?= hiot-platform

SERVICES = auth
CONTAINERS = $(addprefix $(SERVICES),_container)

GO111MODULE ?= on
CGO_ENABLED ?= 0
GOOS ?= linux
GOARCH ?= arm64

.PHONY: all $(SERVICES) dc-up dc-down clean

define build_service
	GO111MODULE=$(GO111MODULE) CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) \
	go build -mod=vendor -ldflags "-s -w" -o build/$(1) cmd/$(1)/main.go
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