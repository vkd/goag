package test

const SpecFile string = `paths:
  /pet:
    get:
      responses:
        200:
          $ref: "#/components/responses/Pet"
  /v2/pet:
    get:
      responses:
        201:
          $ref: "#/components/responses/Pet2"
        default:
          $ref: "#/components/responses/Error"
  /v3/pet:
    get:
      responses:
        202:
          $ref: "#/components/responses/Pet3"

components:
  schemas:
    Error:
      type: object
      properties:
        message:
          type: string
    Pet:
      type: object
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string

  responses:
    Error:
      description: "Error output response"
      content:
        application/json:
          schema:
            $ref: "#/components/schemas/Error"
    Pet:
      description: "Pet output response"
      content:
        application/json:
          schema:
            $ref: "#/components/schemas/Pet"
    Pet2:
      $ref: "#/components/responses/Pet"
    Pet3:
      $ref: "#/components/responses/Pet2"
`
