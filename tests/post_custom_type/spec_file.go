package test

const SpecFile string = `paths:
  /shops/{shop}/pets:
    post:
      parameters:
        - name: shop
          in: path
          required: true
          schema:
            type: string
            x-goag-go-type: github.com/vkd/goag/tests/post_custom_type/pkg.ShopType
        - name: filter
          in: query
          schema:
            type: string
            x-goag-go-type: github.com/vkd/goag/tests/post_custom_type/pkg.ShopType
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
          type: string
          x-goag-go-type: github.com/vkd/goag/tests/post_custom_type/pkg.PetTag
`
