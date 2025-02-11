package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: schema_one_of

paths:
  /pet:
    get:
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Pet'
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Resp"
  /pet2:
    get:
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Pet2'
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Resp2"
components:
  schemas:
    Pet:
      oneOf:
        - $ref: '#/components/schemas/Cat'
        - $ref: '#/components/schemas/Dog'
        - type: object
          required:
          - id
          properties:
            id:
              type: integer
              format: int64
        - type: integer
          format: int64
        - type: string

    Cat:
      type: object
      required:
        - label
      properties:
        label:
          type: string

    Dog:
      type: object
      required:
        - name
      properties:
        name:
          type: string
        tag:
          type: string
      x-goag-go-type: github.com/vkd/goag/tests/schema_one_of/pkg.Dog

    Resp:
      type: object
      required:
        - pet
      properties:
        pet:
          $ref: "#/components/schemas/Pet"

    Pet2:
      oneOf:
        - $ref: '#/components/schemas/Cat2'
        - $ref: '#/components/schemas/Dog2'
      discriminator:
        propertyName: petType
        mapping:
          cati: '#/components/schemas/Cat2'
          dogi: '#/components/schemas/Dog2'

    Cat2:
      type: object
      required:
        - petType
      properties:
        name:
          type: string
        petType:
          type: string

    Dog2:
      type: object
      required:
        - petType
      properties:
        name:
          type: string
        tag:
          type: string
        petType:
          type: string
      x-goag-go-type: github.com/vkd/goag/tests/schema_one_of/pkg.Dog2

    Resp2:
      type: object
      required:
        - pet
      properties:
        pet:
          $ref: "#/components/schemas/Pet2"
`
