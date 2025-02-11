package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: json

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
        - birthday
      properties:
        name:
          type: string
        tag:
          nullable: true
          type: string
        tago:
          nullable: true
          type: string
        birthday:
          type: string
          format: "date-time"
        metadata:
          $ref: '#/components/schemas/Metadata'

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

    Metadata:
      type: object
      required:
        - owner
      properties:
        owner:
          type: string
        tags:
          $ref: '#/components/schemas/Tags'

    Tags:
      type: array
      items:
        $ref: '#/components/schemas/Tag'

    Tag:
      type: object
      required:
      - name
      - value
      properties:
        name:
          type: string
        value:
          type: string

  responses:
    Pet:
      description: "Pet output response"
      content:
        application/json:
          schema:
            $ref: "#/components/schemas/Pet"
`
