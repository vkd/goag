package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: default

paths:
  /pets:
    get:
      responses:
        '200':
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/Pet"
  /pets/names:
    get:
      responses:
        '200':
          content:
            application/json:
              schema:
                type: array
                items:
                  type: string
components:
  schemas:
    Pet:
      type: object
      required:
        - id
        - name
        - tag
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string
        tag:
          type: string
`
