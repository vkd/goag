package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: double_wildcard

paths:
  /pets/{pet_id}/names:
    parameters:
      - name: pet_id
        in: path
        schema:
          type: string
    get:
      responses:
        '200': {}
  /pets/{pet_id}/shops:
    parameters:
      - name: pet_id
        in: path
        schema:
          type: string
    get:
      responses:
        '200': {}
`
