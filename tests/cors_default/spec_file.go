package test

const SpecFile string = `openapi: 3.1.0
info:
  title: cors_default
  version: 0.0.1

paths:
  /shops:
    get:
      parameters:
        - name: page
          in: query
          schema:
            type: integer
            format: int32
        - name: access-key
          in: header
          schema:
            type: string
      responses:
        '200': {}
        default: {}
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
    post:
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
        - name: query-id
          in: header
          schema:
            type: string
      responses:
        '200': {}
        default: {}
`
