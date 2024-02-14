package test

const SpecFile string = `security:
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
