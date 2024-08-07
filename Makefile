BUILD_DIR=./target

.PHONY: build-app
build-app: mk-build-dir
	go build \
		-o "${BUILD_DIR}" \
		-tags osusergo,netgo \
		-ldflags="-s -w -X 'main.Version=${BUILD_VERSION}'" \
		./cmd

.PHONY: clean
clean:
	rm -rf "${BUILD_DIR}"
	docker compose rm -f

.PHONY: mk-build-dir
mk-build-dir:
	mkdir -p "${BUILD_DIR}"
