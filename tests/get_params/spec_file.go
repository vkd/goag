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
        - name: page
          in: query
          schema:
            type: integer
            format: int32
        - name: request-id
          in: header
          schema:
            type: string
      responses:
        '200': {}
        default: {}
`
