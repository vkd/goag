package test

const SpecFile string = `openapi: "3.0.3"
info:
  version: 0.0.1
  title: default

security:
  - bearerAuth: []

paths:
  /login:
    post:
      security: []
      responses:
        200: {}
        401: {}

  /shops:
    post:
      responses:
        200: {}
        401: {}

components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT
`
