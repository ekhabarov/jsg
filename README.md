# JSON Schema Generator (jsg)
The generator produces structure(s) and validation method(s) for them based on
JSON schema.

## The flow

* Parse JSON schema file.
* Build internal tree representation. AST.
* Generate Go code based on AST.


## Modules

* `ast`: reads the JSON schema and builds Abstract Syntax Tree (AST)
* `generator`: produces Go code out of AST.


## What's supported

### Features

| Feature            | Parse | Generate | Validation | Notes |
|:-------------------|:-----:|:--------:|:----------:|:-----:|
| `string`           | x     |          |            |       |
| `string`:`length`  | x     |          |            |       |
| `string`:`pattern` | x     |          |            |       |
| `string`:`format`  | x     |          |            |       |
| `number`           | x     |          |            |       |
| `integer`          | x     |          |            |       |
| `object`           | x     |          |            |       |
| `array`            | x     |          |            |       |
| `boolean`          | x     |          |            |       |
| `null`             | x     |          |            |       |
| `multi types`      | x     |          |            |       |

* `Parse`: library recognizes the feature inside a JSON schema and converts it’s
  into AST.
* `Generate`: library reads AST and generated Go code for the feature.
* `Validation`: library is able to validate a data againt feature.
