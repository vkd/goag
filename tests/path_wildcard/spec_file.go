package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: default

paths:
  /pets/{pet_id}:
    parameters:
      - name: pet_id
        in: path
        schema:
          type: integer
          format: int32
    get:
      responses:
        '200':
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Pet"
components:
  schemas:
    Pet:
      type: object
      required:
        - id
        - name
      properties:
        id:
          type: integer
          format: int32
        name:
          type: string
`
