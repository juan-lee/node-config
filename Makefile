SHELL := /bin/bash
PROJECT = github.com/juan-lee/node-config
APIS_DIR = ${PROJECT}/apis
CACHE_DIR = /tmp/go-cache

all: build

build: 
	$(MAKE) shell COMMAND="make binary"

shell:
	mkdir -p $(CACHE_DIR)/bin $(CACHE_DIR)/src $(CACHE_DIR)/cache bin/cache
	docker run -it \
		-v $(shell pwd):/go/src/github.com/juan-lee/node-config \
		-v $(CACHE_DIR)/bin:/go/bin \
		-v $(CACHE_DIR)/src:/go/src \
		-v $(CACHE_DIR)/cache:/.cache/go-build \
		-w /go/src/github.com/juan-lee/node-config \
		-u $(shell id -u):$(shell id -g) \
		golang:1.21 \
		$(COMMAND)

binary: autogen vendor
	go build -mod=vendor -o bin/node-config .

autogen: /go/bin/deepcopy-gen /go/bin/defaulter-gen /go/bin/conversion-gen
	# Let the boilerplate be empty
	touch /tmp/boilerplate
	go mod vendor

	/go/bin/deepcopy-gen \
		--input-dirs ${APIS_DIR}/config,${APIS_DIR}/config/v1,${APIS_DIR}/config/v1beta1 \
		--bounding-dirs ${APIS_DIR} \
		-O zz_generated.deepcopy \
		-h /tmp/boilerplate

	/go/bin/defaulter-gen \
		--input-dirs ${APIS_DIR}/config/v1,${APIS_DIR}/config/v1beta1 \
		-O zz_generated.defaults \
		-h /tmp/boilerplate

	/go/bin/conversion-gen \
		--input-dirs ${APIS_DIR}/config,${APIS_DIR}/config/v1,${APIS_DIR}/config/v1beta1 \
		-O zz_generated.conversion \
		-h /tmp/boilerplate

/go/bin/%: vendor
	go install k8s.io/code-generator/cmd/$*@v0.28.2

vendor:
	if [[ ! -f go.mod ]]; then go mod init github.com/juan-lee/node-config; fi
	go mod tidy
	go mod vendor
	go mod verify

only-build:
	$(MAKE) shell COMMAND="go build -mod=vendor -o bin/node-config ."

clean:
	rm -rf bin vendor go.sum go.mod
	find . -type f | grep zz_generated | grep -v vendor | xargs -r rm
