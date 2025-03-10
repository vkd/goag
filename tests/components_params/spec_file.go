package test

const SpecFile string = `openapi: 3.1.0
info:
  title: components_params
  version: 0.0.1

paths:
  /shops/{shop}:
    get:
      parameters:
        - $ref: '#/components/parameters/ShopPathRequired'
        - $ref: '#/components/parameters/PageQuery'
      responses:
        '200': {}
        default: {}
  /shops/{shop}/reviews:
    get:
      parameters:
        - $ref: '#/components/parameters/ShopPathRequired'
        - $ref: '#/components/parameters/PageQuery'
      responses:
        '200': {}
        default: {}
  /shops/new:
    post:
      parameters:
        - $ref: '#/components/parameters/PageQuery'
      responses:
        '200': {}
        default: {}

components:
  parameters:
    ShopPathRequired:
      name: shop
      in: path
      required: true
      schema:
        type: string

    PageQuery:
      name: page
      in: query
      schema:
        type: integer
        format: int32
`
