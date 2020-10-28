# api-golang

This api has been made for a test at Vade Secure

## Run the api

`make run` run the api server

`make test` execute all the tests

`make mux-package` install the package mux to use the router

`make assert-package` install the package assert to use assert clauses for the different tests

## Authorization

No authorization needed. All the datas are saved locally.

## Use Cases

### - Create a new document

`POST /create`

|Paraemter | type | Description|
|----------|------|------------|
|id|integer|the id of document|
|name|string|the name of the document|
|desc|string|the description of the document|

#### Response

`
    response : string
`

### - Remove a document

`POST /remove/{id}`

|Paraemter | type | Description|
|----------|------|------------|
|id|integer|the id of document|

#### Response

`
    response : string
`

### - Get a document information

`POST /document/{id}`

|Paraemter | type | Description|
|----------|------|------------|
|id|integer|the id of document|

#### Response

`
    response : string
`


## Maintainers

This project is maintened by :
- [Abdoulaye Chérif Touré](https://github.com/cherift)