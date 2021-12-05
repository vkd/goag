package test

const SpecFile string = `paths:
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
