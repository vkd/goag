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
                required:
                  - groups
                properties:
                  groups:
                    type: object
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
      required:
        - name
      properties:
        name:
          type: string
`
