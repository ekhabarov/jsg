# JSON Schema to Go Generator (jsg) | WORK IN PROGRESS

![Tests](https://github.com/ekhabarov/jsg/actions/workflows/tests.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/ekhabarov/jsg.svg)](https://pkg.go.dev/github.com/ekhabarov/jsg)

`jsg` generates Golang code out of [JSON schema](https://json-schema.org/).

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
| `string`           | x     |x          |            |       |
| `string`:`length`  | x     |          |            |       |
| `string`:`pattern` | x     |          |            |       |
| `string`:`format`  | x     |x          |            |       |
| `number`           | x     |x         |            |       |
| `integer`          | x     |x         |            |       |
| `object`           | x     |          |            |       |
| `array`            | x     |x         |            |       |
| `boolean`          | x     |x         |            |       |
| `null`             | x     |x         |            |       |
| `multi types`      | x     |          |            |       |

* `Parse`: library recognizes the feature inside a JSON schema and converts it’s
  into AST.
* `Generate`: library reads AST and generated Go code for the feature.
* `Validation`: library is able to validate a data againt feature.

## Specification support

[Draft 2020-12](https://json-schema.org/draft/2020-12/json-schema-validation.html)

| Specification section                              | Parse   | Generate   | Validation   | Notes   |
|:---------------------------------------------------|:-------:|:----------:|:------------:|:-------:|
| `6.1. Validation Keywords for Any Instance Type`   |         |            |              |         |
| `6.1.1. type`                                      | x       |            |              |         |
| `6.1.2. enum`                                      |         |            |              |         |
| `6.1.3. const`                                     |         |            |              |         |
| `6.2. Validation Keywords for Numeric Instances`   |         |            |              |         |
| `6.2.1. multipleOf`                                |x        |            |              |         |
| `6.2.2. maximum`                                   |x        |            |              |         |
| `6.2.3. exclusiveMaximum`                          |x        |            |              |         |
| `6.2.4. minimum`                                   |x        |            |              |         |
| `6.2.5. exclusiveMinimum`                          |x        |            |              |         |
| `6.3. Validation Keywords for Strings`             |         |            |              |         |
| `6.3.1. maxLength`                                 | x       |            |              |         |
| `6.3.2. minLength`                                 | x       |            |              |         |
| `6.3.3. pattern`                                   | x       |            |              |         |
| `6.4. Validation Keywords for Arrays`              |         |            |              |         |
| `6.4.1. maxItems`                                  |         |            |              |         |
| `6.4.2. minItems`                                  |         |            |              |         |
| `6.4.3. uniqueItems`                               |         |            |              |         |
| `6.4.4. maxContains`                               |         |            |              |         |
| `6.4.5. minContains`                               |         |            |              |         |
| `6.5. Validation Keywords for Objects`             |         |            |              |         |
| `6.5.1. maxProperties`                             |         |            |              |         |
| `6.5.2. minProperties`                             |         |            |              |         |
| `6.5.3. required`                                  |         |            |              |         |
| `6.5.4. dependentRequired`                         |         |            |              |         |
| :------------------------------------------------- | :-----: | :--------: | :----------: | :-----: |
| `7.3. Defined Formats`                             | x       |            |              |         |
| `7.3.1. Dates, Times, and Duration`                | x       |            |              |         |
| `date-time`                                        | x       |x           |              |         |
| `date`                                             | x       |x           |              |         |
| `time`                                             | x       |x           |              |         |
| `duration`                                         | x       |x           |              |         |
| `7.3.2. Email Addresses`                           | x       |            |              |         |
| `email`                                            | x       |x           |              |         |
| `idn-email`                                        | x       |x           |              |         |
| `7.3.3. Hostnames`                                 | x       |            |              |         |
| `hostname`                                         | x       |x           |              |         |
| `idn-hostname`                                     | x       |x           |              |         |
| `7.3.4. IP Addresses`                              | x       |            |              |         |
| `ipv4`                                             | x       |x           |              |         |
| `ipv6`                                             | x       |x           |              |         |
| `7.3.5. Resource Identifiers`                      | x       |            |              |         |
| `uri`                                              | x       |x           |              |         |
| `uri-reference`                                    | x       |x           |              |         |
| `iri`                                              | x       |x           |              |         |
| `iri-reference`                                    | x       |x           |              |         |
| `uuid`                                             | x       |x           |              |         |
| `7.3.6. uri-template`                              | x       |x           |              |         |
| `7.3.7. JSON Pointers`                             | x       |x           |              |         |
| `json-pointer`                                     | x       |x           |              |         |
| `relative-json-pointer`                            | x       |x           |              |         |
| `7.3.8. regex`                                     | x       |x           |              |         |

