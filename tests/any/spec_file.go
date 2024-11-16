package test

const SpecFile string = `openapi: 3.1.0
info:
  version: 0.0.1
  title: any


paths:
  /pets:
    post:
      requestBody:
        $ref: "#/components/requestBodies/CreatePet"
      responses:
        200:
          description: Pets response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Pets"


components:
  schemas:
    Pets:
      type: array
      items:
        $ref: "#/components/schemas/Pet"

    Pet:
      type: object
      required:
      - id
      - name
      - payload
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string
        payload: {}

  requestBodies:
    CreatePet:
      content:
        application/json:
          schema:
            type: object
            required:
            - name
            - payload
            properties:
              name:
                type: string
              payload: {}
`
