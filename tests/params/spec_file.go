package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: default

paths:
  /shops/{shop}/reviews:
    get:
      operationId: getReviews
      parameters:
        # Path
        - $ref: '#/components/parameters/ShopPathRequired'

        # Query
        - $ref: '#/components/parameters/IntRequired'
        - $ref: '#/components/parameters/Int'
        - $ref: '#/components/parameters/Int32Required'
        - $ref: '#/components/parameters/Int32'
        - $ref: '#/components/parameters/Int64Required'
        - $ref: '#/components/parameters/Int64'
        - $ref: '#/components/parameters/Float32Required'
        - $ref: '#/components/parameters/Float32'
        - $ref: '#/components/parameters/Float64Required'
        - $ref: '#/components/parameters/Float64'

        - $ref: '#/components/parameters/StringRequired'
        - $ref: '#/components/parameters/String'

        - name: tag
          in: query
          schema:
            type: array
            items:
              type: string
        - name: filter
          in: query
          schema:
            type: array
            items:
              type: integer
              format: int32

        # Headers
        - name: request-id
          in: header
          schema:
            type: string
        - name: user-id
          in: header
          required: true
          schema:
            type: string

      responses:
        default: {}

components:
  parameters:
    ShopPathRequired:
      name: shop
      in: path
      required: true
      schema:
        type: integer
        format: int32

    IntRequired:
      name: int_req
      in: query
      required: true
      schema:
        type: integer
    Int:
      name: int
      in: query
      schema:
        type: integer

    Int32Required:
      name: int32_req
      in: query
      required: true
      schema:
        type: integer
        format: int32
    Int32:
      name: int32
      in: query
      schema:
        type: integer
        format: int32

    Int64Required:
      name: int64_req
      in: query
      required: true
      schema:
        type: integer
        format: int64
    Int64:
      name: int64
      in: query
      schema:
        type: integer
        format: int64

    Float32Required:
      name: float32_req
      in: query
      required: true
      schema:
        type: number
        format: float
    Float32:
      name: float32
      in: query
      schema:
        type: number
        format: float

    Float64Required:
      name: float64_req
      in: query
      required: true
      schema:
        type: number
        format: double
    Float64:
      name: float64
      in: query
      schema:
        type: number
        format: double

    StringRequired:
      name: string_req
      in: query
      required: true
      schema:
        type: string
    String:
      name: string
      in: query
      schema:
        type: string
`
