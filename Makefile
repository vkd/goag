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
	go run cmd/goag/main.go $(if ${RUN_TEST},--file ./${TESTS_DIR}/${RUN_TEST}/openapi.yaml --out ./${TESTS_DIR}/${RUN_TEST}/,--dir ./${TESTS_DIR})  --package test $(if ${TEST_CLIENT},--client=true,) --donotedit=false

EXAMPLES_DIR?=examples
examples-gen:
	go run cmd/goag/main.go $(if ${RUN_EXAMPLE},--file ./${EXAMPLES_DIR}/${RUN_EXAMPLE}/openapi.yaml --out ./${EXAMPLES_DIR}/${RUN_EXAMPLE}/,--dir ./${EXAMPLES_DIR})  --package test $(if ${CLIENT},--client=true,)

test-only:
	go test $(if ${RUN},-run=${RUN},) ./...

test-local: test-gen test-only
test: test-gen examples-gen update-readme
	go test $(if ${RUN},-run=${RUN},) ./...

# clean:
# 	rm -rf ${INTEGRATION_DIR}

COVER_FILE?=cover.prof

cover-web: test-gen
	go test -coverpkg=$(shell go list ./... | grep -v ${PACKAGE_TESTS} | tr "\n" "," | sed 's/.$$//') -coverprofile=${COVER_FILE} $(shell go list ./... | grep -v ${PACKAGE_TESTS})
	go tool cover -html=${COVER_FILE}
	@ rm ${COVER_FILE}

build:
	go build ./cmd/goag/main.go


README_ANCHOR?=petstore-example_test
README_LANG?=golang

update-readme: TMP_FILE=readme_new.md
update-readme: START_LINE=\[${README_ANCHOR}\]: \# \(PRINT START\)
update-readme: END_LINE=\[${README_ANCHOR}\]: \# \(END\)
update-readme: START_LINE_TEXT=$(shell echo "${START_LINE}" | sed -r 's/\\//g')
update-readme: END_LINE_TEXT=$(shell echo "${END_LINE}" | sed -r 's/\\//g')
update-readme:
	@ grep -F -q "${START_LINE_TEXT}" README.md || (echo "README.md should contain line: ${START_LINE_TEXT}" && exit 1)
	@ grep -F -q "${END_LINE_TEXT}" README.md || (echo "README.md should contain line: ${END_LINE_TEXT}" && exit 1)
	cat README.md | sed -r '/${START_LINE}$$/q' > ${TMP_FILE}
	@ echo "\`\`\`${README_LANG}" >> ${TMP_FILE}
	cat examples/petstore/example_test.go | sed 1,8d | sed 's/\t/    /g' >> ${TMP_FILE}
	@ echo "\`\`\`" >> ${TMP_FILE}
	cat README.md | sed -r -n '/${END_LINE}$$/,$$ p' >> ${TMP_FILE}
	@ cp -f ${TMP_FILE} README.md
	@ rm ${TMP_FILE}
