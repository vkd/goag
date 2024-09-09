package test

const SpecFile string = `paths:
  /pet:
    get:
      responses:
        '200':
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Pet"
components:
  schemas:
    Pet:
      type: object
      required:
        - id
        - name
        - created_at
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string
        created_at:
          type: string
          format: "date-time"
`
