package test

const SpecFile string = `openapi: 3.1.0
info:
  version: 0.0.1
  title: get_custom_params

paths:
  /shops/{shop}/pages/{page}:
    get:
      operationId: getShopsShop
      parameters:
        - name: shop
          in: path
          required: true
          schema:
            type: string
            x-goag-go-type: Shop
        - name: page
          in: path
          required: true
          schema:
            $ref: '#/components/schemas/PageCustom'
        - name: page
          in: query
          schema:
            type: integer
            format: int32
            x-goag-go-type: Page
        - name: page_req
          in: query
          required: true
          schema:
            type: integer
            format: int32
            x-goag-go-type: Page
        - name: pages
          in: query
          schema:
            type: array
            items:
              type: integer
              format: int32
              x-goag-go-type: Page
        - name: pages_array
          in: query
          schema:
            type: array
            items:
              type: integer
              format: int32
            x-goag-go-type: Pages
        - name: page_custom
          in: query
          schema:
            $ref: '#/components/schemas/PageCustom'
        - name: request-id
          in: header
          schema:
            type: string
            x-goag-go-type: RequestID
      responses:
        200:
          description: "OK"
        default:
          description: "Default"

components:
  schemas:
    PageCustom:
      type: string
      x-goag-go-type: github.com/vkd/goag/tests/get_custom_params/pkg.Page
`
