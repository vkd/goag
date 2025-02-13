package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: custom_type

paths:
  /shops/{shop}:
    get:
      parameters:
        - name: shop
          in: path
          required: true
          schema:
            type: string
            x-goag-go-type: pkg.Shop
        - name: page_schema_ref_query
          in: query
          schema:
            $ref: "#/components/schemas/PageCustom"
        - name: page_custom_type_query
          in: query
          schema:
            type: string
            x-goag-go-type: pkg.PageCustomTypeQuery
      requestBody:
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/GetShop"
      responses:
        200:
          description: Shop response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Shop"

        default:
          description: Error response

components:
  schemas:
    PageCustom:
      type: string
      x-goag-go-type: pkg.Page

    Shop:
      $ref: "#/components/schemas/ShopName"

    ShopName:
      type: integer
      format: int64
      x-goag-go-type: pkg.Page

    GetShop:
      type: object
      required:
        - metadata
      properties:
        metadata:
          $ref: "#/components/schemas/Metadata"
        settings:
          $ref: "#/components/schemas/Settings"
        environments:
          $ref: "#/components/schemas/Environments"
        additionals:
          type: object
          nullable: true
          additionalProperties: true
          x-goag-go-type: pkg.Settings


    Metadata:
      type: object
      properties:
        inner_id:
          type: string
      x-goag-go-type: pkg.Metadata

    Settings:
      type: object
      nullable: true
      properties:
        theme:
          type: string
      x-goag-go-type: pkg.Settings

    Environments:
      type: array
      nullable: true
      items:
        $ref: "#/components/schemas/Environment"
      x-goag-go-type: pkg.Environments

    EnvironmentCreate:
      type: object
      required:
      - name
      properties:
        name:
          type: string
      x-goag-go-type: pkg.EnvironmentCreate

    Environment:
      allOf:
        - $ref: "#/components/schemas/EnvironmentCreate"
        - type: object
          required:
          - value
          properties:
            value:
              type: string
      x-goag-go-type: pkg.Environment
`
