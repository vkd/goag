package test

const SpecFile string = `openapi: 3.0.3
info:
  title: security_jwt_apikey
  version: 0.0.0

security:
  - bearerAuth: []
  - apiKey: []

paths:
  /login:
    post:
      security: []
      responses:
        200:
          description: "OK"
        401:
          description: "Unauthorized"

  /shops:
    post:
      responses:
        200:
          description: "OK"
        401:
          description: "Unauthorized"

components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT
    apiKey:
      type: apiKey
      name: Access-Token
      in: header
`
