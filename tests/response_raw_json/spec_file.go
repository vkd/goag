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
                x-goag-go-type: json.RawMessage
`
