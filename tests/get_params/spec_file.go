package test

const SpecFile string = `paths:
  /shops/{shop}/pets/{petId}:
    get:
      parameters:
        - name: shop
          in: path
          required: true
          schema:
            type: string
        - name: petId
          in: path
          required: true
          schema:
            type: integer
            format: int64
        - name: color
          in: query
          required: true
          schema:
            type: string
        - name: page
          in: query
          schema:
            type: integer
            format: int32
      responses:
        '200': {}
        default: {}
`
