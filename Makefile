ARCH=$(shell arch)
UNAME=$(shell uname)
SHORTHASH=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)-$(SHORTHASH)
FILEVERSION=$(VERSION)-$(UNAME)-$(ARCH)

all: gourfdash

.PHONY: version
version:
	echo $(VERSION)

gourfdash:
	cd backend/cmd/dashboard && GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) go build -o ../../../gourfdash -ldflags "-X main.dver=$(VERSION)"

.PHONY: package
package: gourfdash gourfdash-$(FILEVERSION).tar.gz

.PHONY: dev-frontend
dev-frontend:
	cd frontend/m17-dashboard-quasar && quasar dev

.PHONY: compile-frontend
compile-frontend:
	cd frontend/m17-dashboard-quasar && quasar build
	-rm -rf frontend/spa
	cp -r frontend/m17-dashboard-quasar/dist/spa frontend/spa

gourfdash-$(FILEVERSION).tar.gz: gourfdash
	tar cfz gourfdash-$(FILEVERSION).tar.gz gourfdash templates static .env.sample

.PHONY: docker
docker:
	docker build -t gourfdash .

.PHONY: docker-clean
docker-clean:
	-docker image rm gourfdash

.PHONY: docker-compose-up
docker-compose-up:
	docker-compose up -d
	docker-compose logs -f

.PHONY: docker-compose-down
docker-compose-down:
	docker-compose down

.PHONY: docker-compose-logs
	docker-compose logs -f

.PHONY: run
run:
	go run -ldflags "-X main.dver=$(VERSION)" backend/cmd/dashboard/*.go

.PHONY: clean
clean:
	-rm gourfdash
	-rm -f gourfdash-$(FILEVERSION).tar.gz