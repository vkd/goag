package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: default

paths:
  /pet:
    get:
      responses:
        '200':
          description: OK response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/RawResponse"
components:
  schemas:
    RawResponse:
      type: object
      additionalProperties: true
`
