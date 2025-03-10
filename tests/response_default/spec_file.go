package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: default

paths:
  /pets:
    get:
      responses:
        '200': {}
        default:
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Error"
components:
  schemas:
    Error:
      type: object
      required:
        - message
      properties:
        message:
          type: string
`
