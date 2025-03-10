package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: default

paths:
  /shops/{shop}:
    get:
      parameters:
        - name: shop
          in: path
          required: true
          schema:
            type: string
        - name: page
          in: query
          schema:
            type: integer
            format: int32
        - name: request-id
          in: header
          schema:
            type: string
      responses:
        '200': {}
        default: {}
`
