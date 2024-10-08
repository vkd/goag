package test

const SpecFile string = `paths:
  /:
    get:
      responses:
        default: {}
  /shops:
    get:
      responses:
        default: {}
  /shops/:
    get:
      responses:
        default: {}
  /shops/{shop}:
    get:
      parameters:
        - $ref: '#/components/parameters/ShopPathRequired'
      responses:
        default: {}
  /shops/{shop}/:
    get:
      parameters:
        - $ref: '#/components/parameters/ShopPathRequired'
      responses:
        default: {}
  /shops/activate:
    get:
      responses:
        default: {}
  /shops/activate/:
    get:
      responses:
        default: {}
  /shops/activate/tag:
    get:
      responses:
        default: {}

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
        '200':
          headers:
            x-next:
              schema:
                type: string
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Pet"
        default:
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
