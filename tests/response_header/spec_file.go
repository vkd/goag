package test

const SpecFile string = `paths:
  /pets:
    get:
      responses:
        '200':
          headers:
            x-next:
              schema:
                type: string
`
