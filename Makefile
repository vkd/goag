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
# go run cmd/goag/main.go --dir ./${TESTS_DIR} --package test
	go run cmd/goag/main.go $(if ${RUN_TEST},--file ./${TESTS_DIR}/${RUN_TEST}/openapi.yaml --out ./${TESTS_DIR}/${RUN_TEST}/,--dir ./${TESTS_DIR})  --package test $(if ${TEST_CLIENT},--client=true,)

EXAMPLES_DIR?=examples
examples-gen:
	go run cmd/goag/main.go $(if ${RUN_EXAMPLE},--file ./${EXAMPLES_DIR}/${RUN_EXAMPLE}/openapi.yaml --out ./${EXAMPLES_DIR}/${RUN_EXAMPLE}/,--dir ./${EXAMPLES_DIR})  --package test $(if ${CLIENT},--client=true,)

test-only:
	go test $(if ${RUN},-run=${RUN},) ./...

test: test-gen test-only

# clean:
# 	rm -rf ${INTEGRATION_DIR}

COVER_FILE?=cover.prof

cover-web: test-gen
	go test -coverpkg=$(shell go list ./... | grep -v ${PACKAGE_TESTS} | tr "\n" "," | sed 's/.$$//') -coverprofile=${COVER_FILE} $(shell go list ./... | grep -v ${PACKAGE_TESTS})
	go tool cover -html=${COVER_FILE}
	@ rm ${COVER_FILE}

build:
	go build ./cmd/goag/main.go
