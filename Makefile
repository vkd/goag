PACKAGE_NAME?=github.com/vkd/goag
TESTS_DIR?=tests
PACKAGE_TESTS?=${PACKAGE_NAME}/${TESTS_DIR}/

simple-example:
	go run cmd/goag/main.go --file ../goag-examples/openapi.yaml --out ../goag-examples/simple


# INTEGRATION_DIR=./tests/integration
# integration-tests:
# 	@ for DIR in $(wildcard tests/testdata/*); do \
# 		rm -rf ${INTEGRATION_DIR}; mkdir ${INTEGRATION_DIR}; cp $$DIR/h_test.go ${INTEGRATION_DIR}; go run cmd/goag/main.go --file $$DIR/openapi.yml --out ${INTEGRATION_DIR} --package tests; go test -v ${INTEGRATION_DIR}; \
# 	done

test-gen:
	@ go run cmd/goag/main.go --dir ./${TESTS_DIR} --package test

test: test-gen
	go test ./...

# clean:
# 	rm -rf ${INTEGRATION_DIR}

COVER_FILE?=cover.prof

cover-web: test-gen
	go test -coverpkg=$(shell go list ./... | grep -v ${PACKAGE_TESTS} | tr "\n" "," | sed 's/.$$//') -coverprofile=${COVER_FILE} $(shell go list ./... | grep -v ${PACKAGE_TESTS})
	go tool cover -html=${COVER_FILE}
	@ rm ${COVER_FILE}