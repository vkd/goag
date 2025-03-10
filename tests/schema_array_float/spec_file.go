package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: default

paths:
  /pets/ids:
    get:
      responses:
        '200':
          content:
            application/json:
              schema:
                type: array
                items:
                  type: number
`
