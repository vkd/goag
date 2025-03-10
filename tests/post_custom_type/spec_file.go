package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: post_custom_type

paths:
  /shops/{shop}/pets:
    post:
      parameters:
        - name: shop
          in: path
          required: true
          schema:
            type: string
            x-goag-go-type: pkg.ShopType
        - name: filter
          in: query
          schema:
            type: string
            x-goag-go-type: pkg.ShopType
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/NewPet'
      responses:
        '201':
          description: OK response
        default:
          description: Default response
components:
  schemas:
    NewPet:
      type: object
      required:
        - name
      properties:
        name:
          type: string
        tag:
          type: string
          x-goag-go-type: pkg.PetTag
`
