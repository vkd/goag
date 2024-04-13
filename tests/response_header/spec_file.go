package test

const SpecFile string = `paths:
  /pets:
    get:
      responses:
        '200':
          headers:
            x-next:
              schema:
                type: string
            x-next-two:
              required: true
              schema:
                type: array
                items:
                  type: integer
`
