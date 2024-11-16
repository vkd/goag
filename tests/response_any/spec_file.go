package test

const SpecFile string = `paths:
  /pet:
    get:
      responses:
        '200':
          description: OK response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/RawResponse"
components:
  schemas:
    RawResponse:
      type: object
      additionalProperties: true
`
