package test

const SpecFile string = `paths:
  /pets:
    post:
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/NewPet'
      responses:
        '201': {}
        default: {}
  /pets2:
    post:
      requestBody:
        $ref: '#/components/requestBodies/Pets2'
      responses:
        '201': {}
        default: {}
components:
  schemas:
    NewPet:
      type: object
      required:
        - id
        - name
      properties:
        name:
          type: string
        tag:
          type: string
  requestBodies:
    Pets2:
      $ref: '#/components/requestBodies/NewPet'
    NewPet:
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/NewPet'
`
