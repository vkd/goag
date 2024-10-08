package test

const SpecFile string = `openapi: 3.1.0
info:
  title: components
  version: 0.0.1


paths:
  /shops:
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
        500: {$ref: "#/components/responses/ErrorResponse"}

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
        500: {$ref: "#/components/responses/ErrorResponse"}


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
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string

    Error:
      type: object
      required:
        - detail
      properties:
        detail:
          type: string

  responses:
    ErrorResponse:
      description: Error response
      headers:
        X-Error-Code:
          $ref: '#/components/headers/ErrorCode'
      content:
        application/json:
          schema:
            $ref: "#/components/schemas/Error"

  requestBodies:
    CreatePet:
      content:
        application/json:
          schema:
            type: object
            required:
            - name
            properties:
              name:
                type: string

  headers:
    ErrorCode:
      description: Error code
      schema:
        type: integer
`
