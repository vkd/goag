package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: schema_all_of

paths:
  /pet:
    get:
      responses:
        200:
          description: "Ok"
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Pet"
components:
  schemas:
    Pet:
      allOf:
        - $ref: '#/components/schemas/NewPet'
        - type: object
          required:
          - id
          properties:
            id:
              type: integer
              format: int64

    NewPet:
      type: object
      required:
        - name
      properties:
        name:
          type: string
        tag:
          type: string
`
