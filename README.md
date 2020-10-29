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

### - Get the document list

`GET /`

#### Response

|Paraemter | type | Description|
|----------|------|------------|
|Message|string|the message|
|Code|integer|the code of response (0 if faillure and 200 if success)|
|Result|list|the list of document returned (null if no document founded)|

### - Create a new document

`POST /create`

|Paraemter | type | Description|
|----------|------|------------|
|id|integer|the id of document|
|name|string|the name of the document|
|desc|string|the description of the document|

#### Response

|Paraemter | type | Description|
|----------|------|------------|
|Message|string|the message|
|Code|integer|the code of response (0 if faillure and 200 if success)|
|Result|list|the list of document returned (null if no document founded)|

### - Remove a document

`POST /remove/{id}`

|Paraemter | type | Description|
|----------|------|------------|
|id|integer|the id of document|

#### Response

|Paraemter | type | Description|
|----------|------|------------|
|Message|string|the message|
|Code|integer|the code of response (0 if faillure and 200 if success)|
|Result|list|the list of document returned (null if no document founded)|

### - Get a document information

`POST /document/{id}`

|Paraemter | type | Description|
|----------|------|------------|
|id|integer|the id of document|

#### Response

|Paraemter | type | Description|
|----------|------|------------|
|Message|string|the message|
|Code|integer|the code of response (0 if faillure and 200 if success)|
|Result|list|the list of document returned (null if no document founded)|

## Maintainers

This project is maintened by :
- [Abdoulaye Chérif Touré](https://github.com/cherift)