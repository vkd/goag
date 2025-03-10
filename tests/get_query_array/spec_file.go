package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: default

paths:
  /pets:
    get:
      parameters:
        - name: tag
          in: query
          schema:
            type: array
            items:
              type: string
        - name: page
          in: query
          schema:
            type: array
            items:
              type: integer
              format: int64
      responses:
        '200': {}
        default: {}
`
