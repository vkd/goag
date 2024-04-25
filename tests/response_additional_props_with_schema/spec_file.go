package test

const SpecFile string = `paths:
  /pet:
    get:
      responses:
        '200':
          content:
            application/json:
              schema:
                type: object
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
      properties:
        name:
          type: string
        custom:
          $ref: '#/components/schemas/PetCustom'
    PetCustom:
      type: object
      x-goag-go-type: json.RawMessage
      additionalProperties:
        type: object
      properties:
        field1:
          type: string
`
