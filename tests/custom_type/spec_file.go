package test

const SpecFile string = `paths:
  /shops/{shop}:
    get:
      parameters:
        - name: shop
          in: path
          required: true
          schema:
            type: string
            x-goag-go-type: github.com/vkd/goag/tests/custom_type/pkg.Shop
        - name: page_schema_ref_query
          in: query
          schema:
            $ref: '#/components/schemas/PageCustom'
        - name: page_custom_type_query
          in: query
          schema:
            type: string
            x-goag-go-type: github.com/vkd/goag/tests/custom_type/pkg.PageCustomTypeQuery
      responses:
        200:
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Shop'


        default: {}

components:
  schemas:
    PageCustom:
      type: string
      x-goag-go-type: github.com/vkd/goag/tests/custom_type/pkg.Page

    Shop:
      $ref: '#/components/schemas/ShopName'

    ShopName:
      type: integer
      format: int64
      x-goag-go-type: github.com/vkd/goag/tests/custom_type/pkg.Page
`
