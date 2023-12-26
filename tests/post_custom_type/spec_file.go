package test

const SpecFile string = `paths:
  /shops/{shop}/pets:
    post:
      parameters:
        - name: shop
          in: path
          required: true
          schema:
            x-go-type: ShopType
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/NewPet'
      responses:
        '201': {}
        default: {}
components:
  schemas:
    NewPet:
      type: object
      required:
        - id
        - name
      properties:
        name:
          type: string
        tag:
          x-go-type: PetTag
`
