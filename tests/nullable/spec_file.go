package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: nullable

paths:
  /pets:
    post:
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/NewPet'
      responses:
        200:
          $ref: "#/components/responses/Pet"
components:
  schemas:
    NewPet:
      type: object
      required:
        - name
        - tag
      properties:
        name:
          type: string
        tag:
          nullable: true
          type: string
        tago:
          nullable: true
          type: string

    Pet:
      type: object
      required:
        - id
        - name
        - tag
        - parents
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string
        tag:
          nullable: true
          type: string
        tago:
          nullable: true
          type: string
        parents:
          nullable: true
          type: object
          additionalProperties: true

  responses:
    Pet:
      description: "Pet output response"
      content:
        application/json:
          schema:
            $ref: "#/components/schemas/Pet"
`
