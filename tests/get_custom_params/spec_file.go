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
            x-goag-go-type: Shop
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
              type: string
              x-goag-go-type: Page
        - name: request-id
          in: header
          schema:
            type: string
            x-goag-go-type: RequestID
      responses:
        '200': {}
        default: {}
`
