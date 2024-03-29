# Go API client for cnsbackend

The responses of the open api in swagger focus on the data field rather than the code and the message fields

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import cnsbackend "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), cnsbackend.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), cnsbackend.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), cnsbackend.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), cnsbackend.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://101.42.88.184/cns-backend/v0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CensorApi* | [**CensorText**](docs/CensorApi.md#censortext) | **Post** /censor/text | censor text if compliance
*CommitsApi* | [**GetCommit**](docs/CommitsApi.md#getcommit) | **Get** /commits/{commit_hash} | get commit
*CommitsApi* | [**MakeCommits**](docs/CommitsApi.md#makecommits) | **Post** /commits | make commit
*CommitsApi* | [**QueryCommit**](docs/CommitsApi.md#querycommit) | **Get** /commits | query commit
*RegistersApi* | [**GetRegister**](docs/RegistersApi.md#getregister) | **Get** /registers/{commit_hash} | get register info by admin
*RegistersApi* | [**GetRegisterOrder**](docs/RegistersApi.md#getregisterorder) | **Get** /registers/order/{commit_hash} | get register order
*RegistersApi* | [**MakeRegisterOrder**](docs/RegistersApi.md#makeregisterorder) | **Post** /registers/order/{commit_hash} | make register order
*RegistersApi* | [**RefreshRegisterOrderUrl**](docs/RegistersApi.md#refreshregisterorderurl) | **Put** /registers/order/refresh-url/{commit_hash} | refresh register order url
*RegistersApi* | [**Register**](docs/RegistersApi.md#register) | **Post** /registers | register by admin
*RenewsApi* | [**GetRenew**](docs/RenewsApi.md#getrenew) | **Get** /renews/{id} | get renew order by admin
*RenewsApi* | [**GetRenewOrder**](docs/RenewsApi.md#getreneworder) | **Get** /renews/order/{id} | get renew order
*RenewsApi* | [**MakeRenewOrder**](docs/RenewsApi.md#makereneworder) | **Post** /renews/order | make renew order
*RenewsApi* | [**RefreshRenewOrderUrl**](docs/RenewsApi.md#refreshreneworderurl) | **Put** /renews/order/refresh-url/{id} | refresh renew order url
*RenewsApi* | [**Renew**](docs/RenewsApi.md#renew) | **Post** /renews | renew by admin


## Documentation For Models

 - [CnsErrorsCnsErrorResp](docs/CnsErrorsCnsErrorResp.md)
 - [ModelsCommitCore](docs/ModelsCommitCore.md)
 - [ModelsRegisterCore](docs/ModelsRegisterCore.md)
 - [ModelsRenewCore](docs/ModelsRenewCore.md)
 - [ModelsRenewOrderArgs](docs/ModelsRenewOrderArgs.md)
 - [ServicesCensorTextReq](docs/ServicesCensorTextReq.md)
 - [ServicesCensorTextResp](docs/ServicesCensorTextResp.md)
 - [ServicesData](docs/ServicesData.md)
 - [ServicesHits](docs/ServicesHits.md)
 - [ServicesMakeCommitResp](docs/ServicesMakeCommitResp.md)
 - [ServicesMakeRegisterOrderReq](docs/ServicesMakeRegisterOrderReq.md)
 - [ServicesMakeRegisterOrderResp](docs/ServicesMakeRegisterOrderResp.md)
 - [ServicesMakeRenewOrderReq](docs/ServicesMakeRenewOrderReq.md)
 - [ServicesMakeRenewOrderResp](docs/ServicesMakeRenewOrderResp.md)
 - [ServicesRegisterByAdminResp](docs/ServicesRegisterByAdminResp.md)
 - [ServicesRnewByAdminResp](docs/ServicesRnewByAdminResp.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



