package test

const SpecFile string = `paths:
  /pet:
    get:
      responses:
        '200':
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/RawResponse"
components:
  schemas:
    RawResponse:
      type: object
      x-goag-go-type: json.RawMessage
`
