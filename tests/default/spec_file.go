package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: default

paths:
  /:
    get:
      responses:
        default:
          description: "Default"
  /shops:
    get:
      responses:
        default:
          description: "Default"
  /shops/:
    get:
      responses:
        default:
          description: "Default"
  /shops/{shop}:
    get:
      parameters:
        - $ref: '#/components/parameters/ShopPathRequired'
      responses:
        default:
          description: "Default"
  /shops/{shop}/:
    get:
      parameters:
        - $ref: '#/components/parameters/ShopPathRequired'
      responses:
        default:
          description: "Default"
  /shops/activate:
    get:
      responses:
        default:
          description: "Default"
  /shops/activate/:
    get:
      responses:
        default:
          description: "Default"
  /shops/activate/tag:
    get:
      responses:
        default:
          description: "Default"

  /shops/{shop}/pets:
    get:
      parameters:
        - $ref: '#/components/parameters/ShopPathRequired'
        - $ref: '#/components/parameters/PageQuery'
        - $ref: '#/components/parameters/PageSizeQuery'
      responses:
        '200':
          description: List of pets
          headers:
            x-next:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
                required:
                - groups
                properties:
                  groups:
                    type: object
                    additionalProperties:
                      $ref: '#/components/schemas/Pets'
        default:
          description: "Default"
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Error"
  /shops/{shop}/review:
    post:
      operationId: reviewShop
      description: |
        Review shop.
        Returns a current pet.
      parameters:
        - $ref: '#/components/parameters/ShopPathRequired'
        - $ref: '#/components/parameters/PageQuery'
        - $ref: '#/components/parameters/PageSizeQuery'
        - name: request-id
          in: header
          schema:
            type: string
        - name: user-id
          in: header
          required: true
          schema:
            type: string
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
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/NewPet'
      responses:
        200:
          description: "OK"
          headers:
            x-next:
              schema:
                type: string
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Pet"
        default:
          description: "Default"
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Error"

components:
  parameters:
    ShopPathRequired:
      name: shop
      in: path
      required: true
      schema:
        type: integer
        format: int32

    PageQuery:
      name: page
      in: query
      schema:
        type: integer
        format: int32

    PageSizeQuery:
      name: page_size
      in: query
      required: true
      schema:
        type: integer
        format: int32

  schemas:
    Error:
      type: object
      required:
        - message
      properties:
        message:
          type: string

    NewPet:
      type: object
      required:
        - name
        - tag
      properties:
        name:
          type: string
        tag:
          type: string

    Pets:
      type: array
      items:
        $ref: '#/components/schemas/Pet'
    Pet:
      allOf:
        - $ref: '#/components/schemas/NewPet'
        - type: object
          required:
          - id
          properties:
            id:
              type: integer
              format: int64
`
