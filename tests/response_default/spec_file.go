package test

const SpecFile string = `paths:
  /pets:
    get:
      responses:
        '200': {}
        default:
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Error"
components:
  schemas:
    Error:
      type: object
      properties:
        message:
          type: string
`
