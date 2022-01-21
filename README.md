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


## Data types mapping
