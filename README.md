Goag - Golang OpenAPIv3 Generator
=

This tool generates boilerplate code for the http handling. It handles the parsing of queries, headers or json bodies.

It makes impossible to return undocumented type/response by endpoint.

Example
-

[petstore-example_test]: # (PRINT START)
```golang
var db interface {
    GetPet(_ context.Context, id string) (Pet, error)
}

func ExampleAPI_petsStore() {
    api := &API{
        ShowPetByIDHandler: func(ctx context.Context, r ShowPetByIDRequest) ShowPetByIDResponse {
            req, err := r.Parse()
            if err != nil {
                return NewShowPetByIDResponseDefaultJSON(http.StatusBadRequest, Error{
                    Code:    400,
                    Message: fmt.Sprintf("Bad request: %v", err),
                })
            }

            out, err := db.GetPet(r.HTTP().Context(), req.Path.PetID)
            if err != nil {
                return NewShowPetByIDResponseDefaultJSON(http.StatusInternalServerError, Error{
                    Code:    500,
                    Message: fmt.Sprintf("Internal server error: %v", err),
                })
            }

            return NewShowPetByIDResponse200JSON(out)
        },
        // ...
    }

    _ = http.ListenAndServe(":8080", api)
}
```
[petstore-example_test]: # (END)

```yaml
openapi: "3.0.0"
info:
  version: 1.0.0
  title: Swagger Petstore
  license:
    name: MIT
servers:
  - url: http://petstore.swagger.io/v1
paths:
  /pets/{petId}:
    get:
      summary: Info for a specific pet
      operationId: showPetById
      tags:
        - pets
      parameters:
        - name: petId
          in: path
          required: true
          description: The id of the pet to retrieve
          schema:
            type: string
      responses:
        '200':
          description: Expected response to a valid request
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Pet"
        default:
          description: unexpected error
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Error"
components:
  schemas:
    Pet:
      type: object
      required:
        - id
        - name
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string
        tag:
          type: string
    Error:
      type: object
      required:
        - code
        - message
      properties:
        code:
          type: integer
          format: int32
        message:
          type: string
```
