package test

const SpecFile string = `servers:
# - url: "https://demo.example.com:8443/api/v1"
- url: https://{username}.example.com:{port}/{basePath}
  description: The production API server
  variables:
    username:
      default: demo
    port:
      default: '8443'
    basePath:
      default: api/v1
paths:
  # some ` + "`" + `comment` + "`" + `
  /: {get: {responses: {default: {}}}}
  /shops: {get: {responses: {default: {}}}}
  /shops/: {get: {responses: {default: {}}}}
  /shops/{shop}: {get: {parameters: [{in: path, name: shop, required: true, schema: {type: string}}], responses: {default: {}}}}
  /shops/{shop}/: {get: {parameters: [{in: path, name: shop, required: true, schema: {type: string}}], responses: {default: {}}}}
  /shops/{shop}/pets: {get: {parameters: [{in: path, name: shop, required: true, schema: {type: string}}], responses: {default: {}}}}
  /shops/{shop}/pets/mike/paws: {get: {parameters: [{in: path, name: shop, required: true, schema: {type: string}}], responses: {default: {}}}}
  /shops/activate: {get: {responses: {default: {}}}}
`
