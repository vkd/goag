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
                type: object
                required:
                  - length
                properties:
                  length:
                    type: integer
                additionalProperties:
                  $ref: '#/components/schemas/Pets'

components:
  schemas:
    Pets:
      type: array
      items:
        $ref: '#/components/schemas/Pet'
    Pet:
      type: object
      additionalProperties: true
      required:
        - name
        - custom
      properties:
        name:
          type: string
        custom:
          $ref: '#/components/schemas/PetCustom'
    PetCustom:
      type: object
      additionalProperties:
        type: object
      properties:
        field1:
          type: string
`
