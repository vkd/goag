package test

const SpecFile string = `paths:
  /pets/{pet_id}/names:
    parameters:
      - name: pet_id
        in: path
        schema:
          type: string
    get:
      responses:
        '200': {}
  /pets/{pet_id}/shops:
    parameters:
      - name: pet_id
        in: path
        schema:
          type: string
    get:
      responses:
        '200': {}
`
