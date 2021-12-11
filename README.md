Goag - Golang OpenAPIv3 Generator
=

This tool generates boilerplate code for the http handling. It handles the parsing of queries, headers or json bodies.

It makes impossible to return undocumented type/response by endpoint.

Example
-

Instead of the unmarshaling of parameters we can just write the logic itself of endpoint:

```golang
var db interface {
  GetPet(context.Context, int64) (Pet, error)
}

api := &API{
  GetPetsHandler: ...,
  PostPetsHandler: ...,
  GetPetsIDHandler: func(r GetPetsIDRequester) GetPetsIDResponder {
    req, err := r.Parse()
    if err != nil {
      return GetPetsIDResponseDefaultJSON(http.StatusBadRequest, Error{
        Code:    400,
        Message: fmt.Sprintf("Bad request: %v", err),
      })
    }

    out, err := db.GetPet(req.HTTPRequest.Context(), req.ID)
    if err != nil {
      return GetPetsIDResponseDefaultJSON(http.StatusInternalServerError, Error{
        Code:    500,
        Message: fmt.Sprintf("Internal server error: %v", err),
      })
    }

    return GetPetsIDResponse200JSON(out)
  },
}

_ = http.ListenAndServe(":8080", api)
```

```yaml
openapi: "3.0.0"
info:
  title: Swagger Petstore
...
servers:
  - url: http://petstore.swagger.io/api
paths:
  ...
  /pets/{id}:
    get:
      parameters:
        - name: id
          in: path
          description: ID of pet to fetch
          required: true
          schema:
            type: integer
            format: int64
      responses:
        '200':
          description: pet response
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Pet'
        default:
          description: unexpected error
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
components:
  schemas:
    Pet:
      type: object
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string

    Error:
      type: object
      properties:
        code:
          type: integer
          format: int32
        message:
          type: string
```

---

Limitations:

* JSON only
* only one spec file
* allOf only for objects
* (?) component cannot has reference item
* no webhooks
* no file uploads
* no multipart
* no callbacks
* no links
* no discriminators
* no XML object
* no security scheme objects
* (?) not full JSON Schema supported
* schema.required is not applied
* no externalDocs
* no cookies
