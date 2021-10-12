package test

const SpecFile string = `servers:
  - url: "https://development.example.com/api/v1"
paths:
  /: {get: {responses: {default: {}}}}
  /shops: {get: {responses: {default: {}}}}
  /shops/: {get: {responses: {default: {}}}}
  /shops/{shop}: {get: {parameters: [{in: path, name: shop, required: true, schema: {type: string}}], responses: {default: {}}}}
  /shops/{shop}/: {get: {parameters: [{in: path, name: shop, required: true, schema: {type: string}}], responses: {default: {}}}}
  /shops/{shop}/pets: {get: {parameters: [{in: path, name: shop, required: true, schema: {type: string}}], responses: {default: {}}}}
  /shops/activate: {get: {responses: {default: {}}}}
`
