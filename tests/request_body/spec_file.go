package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: default

paths:
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
        - name
        - tag
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
