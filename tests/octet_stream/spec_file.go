package test

const SpecFile string = `openapi: 3.1.0
info:
  version: 0.0.1
  title: any

paths:
  /pets:
    get:
      responses:
        200:
          description: "OK"
          content:
            application/octet-stream:
              schema:
                # a binary file of any type
                type: string
                format: binary

    post:
      requestBody:
        content:
          application/octet-stream:
            schema:
              # a binary file of any type
              type: string
              format: binary
      responses:
        200:
          description: "OK"
`
