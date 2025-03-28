package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: default

paths:
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
