package test

const SpecFile string = `openapi: 3.1.0
info:
  version: 0.0.1
  title: any

paths:
  /pets:
    get:
      responses:
        200:
          description: "OK"
          content:
            application/json:
              schema:
                type: object
                required:
                - array
                - nullable
                properties:
                  array:
                    type: array
                    items:
                      type: string
                  nullable:
                    type: array
                    nullable: true
                    items:
                      type: string
`
