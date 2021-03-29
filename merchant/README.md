<a name="table-of-contents"></a>
## Table of contents

- [Get Started](#get-started)
- [API Authentication](#api-authentication)
- [Callback Integration](#callback-integration)
- REST API
	- [Create Payment Order](#create-payment-order)
	- [Query Payment Order Status](#query-payment-order-status)
	- [Cancel Payment Order](#cancel-payment-order)
	- [Update Payment Order Expiration Duration](#update-payment-order-expiration-duration)
	- [Activate Merchant API Code](#activate-merchant-api-code)
	- [Query Merchant API Code Status](#query-merchant-api-code-status)
	- [Refresh Merchant API Code](#refresh-merchant-api-code)
- Testing
	- [Mock Server](#mock-server)
	- [cURL Testing Commands](#curl-testing-commands)
	- [Other Language Versions](#other-language-versions)
- Appendix
	- [Callback Definition](#callback-definition)
	- [Currency Definition](#currency-definition)

<a name="get-started"></a>
# Get Started

CYBAVO Merchant Service(CMS) is a comprehensive cryptocurrency payment servcie, based on the deposit wallet of CYBAVO VAULT.

### Try it now
- Use [mock server](#mock-server) to test CYBAVO Merchant API right away.

### Start integration
- To make a correct API call, refer to [API Authentication](#api-authentication).
- To handle callback correctly, refer to [Callback Integration](#callback-integration).

<a name="api-authentication"></a>
# API Authentication

- The CYBAVO Merchant Service verifies all incoming requests. All requests must include X-API-CODE, X-CHECKSUM headers otherwise caller will get a 403 Forbidden error.

### How to acquire and refresh API code and secret
- Request the merchant API code/secret from the **Merchant Details** page on the web control panel for the first time.
- A paired refresh code can be used in the [refresh API](#refresh-merchant-api-code) to acquire the new inactive API code/secret of the wallet.
	- Before the inactive API code is activated, the currently activated API code is still valid.
	- Once the paired API code becomes invalid, the paired refresh code will also become invalid.

### How to make a correct request?
- Put the merchant API code in the X-API-CODE header.
	- Use the inactivated API code in any request will activate it automatically. Once activated, the currently activated API code will immediately become invalid.
	- Or you can explicitly call the [activation API](#activate-merchant-api-code) to activate the API code before use
- Calculate the checksum with the corresponding API secret and put the checksum in the X-CHECKSUM header.
  - The checksum calculation will use all the query parameters, the current timestamp, user-defined random string and the post body (if any).
- Please refer to the code snippet on the github project to know how to calculate the checksum.
	- [Go](https://github.com/CYBAVO/SOFA_MOCK_SERVER/blob/master/api/apicaller.go#L40)
	- [Java](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVA/blob/master/src/main/java/com/cybavo/sofa/mock/Api.java#L71)
	- [Javascript](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVASCRIPT/blob/master/helper/apicaller.js#L58)
	- [PHP](https://github.com/CYBAVO/SOFA_MOCK_SERVER_PHP/blob/master/helper/apicaller.php#L26)


<a name="callback-integration"></a>
# Callback Integration

- Please note that the merchant must have an activated API code, otherwise no callback will be sent.
	- Use the [activation API](#activate-merchant-api-code) to activate an API code.

- To ensure that the callbacks have processed by callback handler, the CYBAVO Merchant Service will continue to send the callbacks to the callback URL until a callback confirmation (HTTP/1.1 200 OK) is received or exceeds the number of retries (retry time interval: 1-3-5-15-45 mins).
	- If all attempts fail, the callback will be set to a failed state, the callback handler can call the [resend failed callbacks](#resend-failed-merchant-callbacks) API to request CYBAVO Merchant Serviceto resend such kind of callback(s) or through the web control panel.

- Refer to [Callback Definition](#callback-definition), [Callback Type Definition](#callback-type-definition) for detailed definition.
- Please refer to the code snippet on the github project to know how to validate the callback payload.
	- [Go](https://github.com/CYBAVO/SOFA_MOCK_SERVER/blob/master/controllers/MerchantController.go#L248)
	- [Java](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVA/blob/master/src/main/java/com/cybavo/sofa/mock/MerchantController.java#L165)
	- [Javascript](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVASCRIPT/blob/master/routes/merchant.js#L163)
	- [PHP](https://github.com/CYBAVO/SOFA_MOCK_SERVER_PHP/blob/master/index.php#L418)

# REST API

<a name="create-payment-order"></a>
### Create Payment Order

Create a payment order.

##### Request

**POST** /v1/merchant/`MERCHANT_ID`/order

- [Sample curl command](#curl-create-payment-order)

> The order\_id must be prefixed. **Find prefix from corresponding merchant detail on web control panel.**
> 
> The prefix is `N520335069_` is the following example.

##### Request Format

An example of the request:

###### API

```
/v1/merchant/520335069/order
```

###### Post body

```json
{
  "currency": 60,
  "token_address": "",
  "amount": "0.01",
  "duration": 50,
  "description": "ETH payment",
  "order_id": "N520335069_1000022",
  "redirect_url": "https%3A%2F%2Fmyredirect.example.com%3Fk%3Dv%26k1%3Dv1"
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :---     | :---        |
| currency | int64 | required | The cryptocurrency used to pay the order |
| token_address | string | required | The token contract address of cryptocurrency used to pay the order |
| amount | string | required | The required amount of the payment order |
| duration | int64 | required | The expiration duration (in minutes) of the payment order |
| description | string | optional, max `255` chars | The description of the payment order |
| order_id | string | required | The user defined order ID (must be prefixed) |
| redirect_url | string | optional | User defined redirect URL (must be encoded) |

> The `redirect_url` must be encoded. Please refer to the code snippet on the github project to know how to encode the URL. [Go](https://github.com/CYBAVO/SOFA_MOCK_SERVER/blob/master/controllers/MerchantController.go#L92), [Java](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVA/blob/master/src/main/java/com/cybavo/sofa/mock/MerchantController.java#L72), [Javascript](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVASCRIPT/blob/master/routes/merchant.js#L45), [PHP](https://github.com/CYBAVO/SOFA_MOCK_SERVER_PHP/blob/master/index.php#L370)

##### Response Format

An example of a successful response:

```json
{
  "access_token": "ybJWKM_CT1yXxzLO2z1Y5fg1EzHuMyRA14ubzR8i-RE",
  "address": "0xed965D0A23eC4583f55Fb5d4109C0fE069B396fC",
  "expired_time": 1615975467,
  "order_id": "N520335069_1000022"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| access_token | string | The access token used to query, update or cancel the payment order |
| address | string | The address to accept the payment |
| expired_time | int64 | The due date of the payment order (unix time in UTC) |
| order_id | string | The order ID of the payment order |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No merchant ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replayed request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 400 | 112 | Invalid parameter | - | Malformatted post body |
| 403 | 113   | Permission denied | - | Merchant API not allowed |
| 400 | 703 | Operation failed | Invalid redirect URL, {DETAILED_ERROR} | Fail to parse the redirect URL |
| 400 | 20003 | Wallet address not available | - | There is no available deposit address |
| 400 | 20007 | Wallet not found | - | No linked deposit wallet found |
| 400 | 20009 | Duplicated Order ID| - | The order ID has been used |
| 400 | 20010 | Wrong Order prefix | - | The order prefix is wrong |
| 400 | 20011 | Wrong Order format | - | The order contains invalid characters |

##### [Back to top](#table-of-contents)


<a name="query-payment-order-status"></a>
### Query Payment Order Status

Query current status of a payment order.

##### Request

`VIEW` **GET** /v1/merchant/`MERCHANT_ID`/order?token=`ACCESS_TOKEN`&order=`ORDER_ID`

- [Sample curl command](#curl-query-payment-order-status)

##### Request Format

An example of the request:

###### API

```
/v1/merchant/520335069/order?token=ybJWKM_CT1yXxzLO2z1Y5fg1EzHuMyRA14ubzR8i-RE&order=N520335069_1000022
```

##### Response Format

An example of a successful response:

```json
{
  "address": "0xed965D0A23eC4583f55Fb5d4109C0fE069B396fC",
  "expired_time": 1615975468,
  "redirect_url": "https%3A%2F%2Fmyredirect.example.com%3Fk%3Dv%26k1%3Dv1",
  "state": 1,
  "tx_id": ""
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| address | string | The deposit address to accept payment |
| expired_time | int64 | The due date of the payment (unix time in UTC) |
| redirect_url | string | User defined redirect URL (encoded) |
| state | int | Refer to [Order State Definition](#order-state-definition) |
| tx_id | string | The TX ID of the payment if state is 0, 2 or 3 |

> The `redirect_url` is encoded. Please refer to the code snippet on the github project to know how to decode the URL. [Go](https://github.com/CYBAVO/SOFA_MOCK_SERVER/blob/master/controllers/MerchantController.go#L129), [Java](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVA/blob/master/src/main/java/com/cybavo/sofa/mock/MerchantController.java#L97), [Javascript](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVASCRIPT/blob/master/routes/merchant.js#L66), [PHP](https://github.com/CYBAVO/SOFA_MOCK_SERVER_PHP/blob/master/index.php#L381)

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No merchant ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replayed request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 112   | Invalid parameter | - | The invalid ACCESS\_TOKEN/ORDER\_ID |
| 403 | 113   | Permission denied | - | Merchant API not allowed |

##### [Back to top](#table-of-contents)
		

<a name="cancel-payment-order"></a>
### Cancel Payment Order

Cancel a payment order. Only pending payment order can be cancelled.

##### Request

**DELETE** /v1/merchant/`MERCHANT_ID`/order?token=`ACCESS_TOKEN`&order=`ORDER_ID`

- [Sample curl command](#curl-cancel-payment-order)

##### Request Format

An example of the request:

###### API

```
/v1/merchant/520335069/order?token=ybJWKM_CT1yXxzLO2z1Y5fg1EzHuMyRA14ubzR8i-RE&order=N520335069_1000022
```

##### Response Format

An example of a successful response:

```json
{
  "result": 1
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| result | int | Always be 1, means the payment order has been successfully cancelled |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No merchant ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replayed request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 112   | Invalid parameter | - | The invalid ACCESS\_TOKEN/ORDER\_ID or the payment can't be cancelled |
| 403 | 113   | Permission denied | - | Merchant API not allowed |


##### [Back to top](#table-of-contents)


<a name="update-payment-order-expiration-duration"></a>
### Update Payment Order Expiration Duration

Update a payment order expiration duration. Only pending payment order can be updated.

**POST** /v1/merchant/`MERCHANT_ID`/order/duration

- [Sample curl command](#curl-update-payment-order-expiration-duration)

##### Request Format

An example of the request:

###### API

```
/v1/merchant/520335069/order
```

###### Post body

```json
{
  "access_token": "ybJWKM_CT1yXxzLO2z1Y5fg1EzHuMyRA14ubzR8i-RE",
  "order_id": "N520335069_1000022",
  "duration": 50
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :---     | :---        |
| access_token | string | required | The access token of the payment order returned with [Create order](#create-payment-order) API |
| order_id | string | required | The order ID of the payment order |
| duration | int64 | required | The expiration duration (in minutes) of the payment order |

> The new due date is calculated based on the submission time of the payment order.

##### Response Format

An example of a successful response:

```json
{
  "result": 1
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| result | int | Always be 1, means the payment order has been successfully updated |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No merchant ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replayed request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 112   | Invalid parameter | - | The invalid ACCESS\_TOKEN/ORDER\_ID or the payment can't be updated |
| 403 | 113   | Permission denied | - | Merchant API not allowed |

##### [Back to top](#table-of-contents)


<a name="resend-failed-merchant-callbacks"></a>
### Resend Failed Merchant Callbacks

The callback handler can call this API to resend pending or failed merchant callbacks.

Refer to [Callback Integration](#callback-integration) for callback rules.

> The resend operation could be requested on the web control panel as well.

##### Request

**POST** /v1/merchant/`MERCHANT_ID`/notifications/manual

- [Sample curl command](#curl-resend-failed-merchant-callbacks)

##### Request Format

An example of the request:

###### API

```
/v1/merchant/520335069/notifications/manual
```

##### Response Format

An example of a successful response:

```json
{
  "count": 0
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| count | int | Count of callbacks just resent |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No merchant ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replayed request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 113   | Permission denied | - | Merchant API not allowed |

##### [Back to top](#table-of-contents)


<a name="activate-merchant-api-code"></a>
### Activate Merchant API Code

Activate the API code of a certain merchant. Once activated, the currently activated API code will immediately become invalid.

##### Request

**POST** /v1/merchant/`MERCHANT_ID`/apisecret/activate

- [Sample curl command](#curl-activate-merchant-api-code)

##### Request Format

An example of the request:

###### API

```
/v1/merchant/520335069/apisecret/activate
```

##### Response Format

An example of a successful response:

```json
{
  "api_code": "4PcdE9VjXfrk7WjC1",
  "exp": 1609646716
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| api_code | string | The activated API code |
| exp | int64 | The API code expiration unix time in UTC |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No merchant ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replayed request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 113   | Permission denied | - | Merchant API not allowed |

##### [Back to top](#table-of-contents)


<a name="query-merchant-api-code-status"></a>
### Query Merchant API Code Status

Query the API code info of a certain merchant. Use the `inactivated` API code in any request will activate it. Once activated, the currently activated API code will immediately become invalid.

##### Request

`VIEW` **GET** /v1/merchant/`MERCHANT_ID`/apisecret

- [Sample curl command](#curl-query-merchant-api-code-status)

##### Request Format

An example of the request:

###### API

```
/v1/merchant/520335069/apisecret
```

##### Response Format

An example of a successful response:

```json
{
  "valid": {
    "api_code": "H4Q6xFZgiTZb37GN",
    "exp": 1583144863
  },
  "inactivated": {
    "api_code": "32PmGCjNzXda4mNHX"
  }
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| valid | object | The activated API code |
| inactivated | object | Not active API code |
| api_code | string | The API code for querying wallet |
| exp | int64 | The API code expiration unix time in UTC |

> Use an invalid API-CODE, the caller will get a 403 Forbidden error.

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No merchant ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replayed request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 113   | Permission denied | - | Merchant API not allowed |

##### [Back to top](#table-of-contents)


<a name="refresh-merchant-api-code"></a>
### Refresh Merchant API Code

Use paired refresh code to acquire the new inactive API code/secret of the merchant.

##### Request

**POST** /v1/merchant/`MERCHANT_ID`/refreshsecret

- [Sample curl command](#curl-refresh-api-code)

##### Request Format

An example of the request:

###### API

```
/v1/merchant/520335069/refreshsecret
```

###### Post body

```json
{
  "refresh_code":"3EbaSPUpKzHJ9wYgYZqy6W4g43NT365bm9vtTfYhMPra"
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :---     | :---        |
| refresh_code | string | required | The corresponding refresh code of the API code specified in the X-API-CODE header |

##### Response Format

An example of a successful response:

```json
{
  "api_code": "4QjbY3qES4tEh19PU",
  "api_secret": "3jC1qjr4mrKxfoXkxoN27Uhmbm1E",
  "refresh_code": "HcN17gxZ3ojrBYSXnjKsU9Pun8krP6J9Pn678k4rZ13m"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| api_code | string | The new inactive API code |
| api_secret | string | The API secret |
| refresh_code | string | The paired refresh code |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No merchant ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replayed request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 400 | 112 | Invalid parameter | - | Malformatted post body or the refresh code is invalid |
| 403 | 113   | Permission denied | - | Merchant API not allowed |

##### [Back to top](#table-of-contents)



<a name="mock-server"></a>
# Mock Server

### How to compile
- Put sample code to {YOUR\_GO\_PATH}/github.com/cybavo/SOFA\_MOCK\_SERVER
- Execute
	- glide install
	- go build ./mockserver.go
	- ./mockserver

### Setup configuration

>	Configure CYBAVO API server URL in mockserver.app.conf

```
api_server_url="BACKEND_SERVER_URL"
```

### Put merchant API code/secret into mock server
-	Get API code/secret on web control panel
	-	API_CODE, API\_SECRET, WALLET\_ID
- 	Put API code/secret to mock server's database

```
curl -X POST -H "Content-Type: application/json" -d '{"api_code":"API_CODE","api_secret":"API_SECRET"}' \
http://localhost:8889/v1/mock/merchant/{MERCHANT_ID}/apitoken
```

### Register mock server callback URL
>	Operate on web control panel

Notification Callback URL

```
http://localhost:8889/v1/mock/merchant/callback
```

##### [Back to top](#table-of-contents)

<a name="curl-testing-commands"></a>
# cURL Testing Commands

<a name="curl-create-payment-order"></a>
### Create Payment Order

```
curl -X POST -H "Content-Type: application/json" -d '{"currency":60,"token_address":"","amount":"0.01","duration":50,"description":"TEST Order","redirect_url":"","order_id":"N520335069_10001"}' \
http://localhost:8889/v1/mock/merchant/{MERCHANT_ID}/order
```

- [API definition](#create-payment-order)


<a name="curl-query-payment-order-status"></a>
### Query Payment Order Status

```
curl http://localhost:8889/v1/mock/merchant/{MERCHANT_ID}/order?token={ACCESS_TOKEN}&order={ORDER_ID}
```

- [API definition](#query-payment-order-status)


<a name="curl-cancel-payment-order"></a>
### Cancel Payment Order

```
curl -X DELETE http://localhost:8889/v1/mock/merchant/{MERCHANT_ID}/order?token={ACCESS_TOKEN}&order={ORDER_ID}
```

- [API definition](#cancel-payment-order)


<a name="curl-update-payment-order-expiration-duration"></a>
### Update Payment Order Expiration Duration

```
curl -X POST -H "Content-Type: application/json" -d '{"access_token":"IUxyiWsdIBp_FS6Tu2afaecH4F_dqpYOLj4oXxn02AA","order_id":"N520335069_10001","duration":100}' \
http://localhost:8889/v1/mock/merchant/{MERCHANT_ID}/order/duration
```

- [API definition](#update-payment-order-expiration-duration)

<a name="curl-resend-failed-merchant-callbacks"></a>
### Resend Failed Merchant Callbacks

```
curl -X POST http://localhost:8889/v1/mock/merchant/{MERCHANT_ID}/notifications/manual
```
- [API definition](#resend-failed-merchant-callbacks)


<a name="curl-activate-merchant-api-code"></a>
### Activate Merchant API Code

```
curl -X POST http://localhost:8889/v1/mock/merchant/{MERCHANT_ID}/apisecret/activate
```
- [API definition](#activate-merchant-api-code)


<a name="curl-query-merchant-api-code-status"></a>
### Query Merchant API Code Status

```
curl http://localhost:8889/v1/mock/merchant/{MERCHANT_ID}/apisecret
```
- [API definition](#query-merchant-api-code-status)


<a name="curl-refresh-merchant-api-code"></a>
### Refresh Merchant API Code

```
curl -X POST -H "Content-Type: application/json" -d '{"refresh_code":"3EbaSPUpKzHJ9wYgYZqy6W4g43NT365bm9vtTfYhMPra"}' \
http://localhost:8889/v1/mock/merchant/{MERCHANT_ID}/apisecret/refreshsecret
```
- [API definition](#query-merchant-api-code-status)


##### [Back to top](#table-of-contents)

<a name="other-language-versions"></a>
# Other Language Versions
- [Java](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVA)
- [Javascript](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVASCRIPT)
- [PHP](https://github.com/CYBAVO/SOFA_MOCK_SERVER_PHP)

##### [Back to top](#table-of-contents)

# Appendix

<a name="callback-definition"></a>
### Callback Definition

<table>
  <tr>
    <td>Field</td>
    <td>Type</td>
    <td>Description</td>
  </tr>
  <tr>
    <td>merchant_id</td>
    <td>int64</td>
    <td>The merchant ID of the callback</td>
  </tr>
  <tr>
    <td>order_id</td>
    <td>string</td>
    <td>The unique order ID of the payment order</td>
  </tr>
  <tr>
    <td>currency</td>
    <td>string</td>
    <td>Cryptocurrency of the callback</td>
  </tr>
  <tr>
    <td>txid</td>
    <td>string</td>
    <td>Transaction identifier</td>
  </tr>
  <tr>
    <td>block_height</td>
    <td>int64</td>
    <td>The block height show the transaction was packed in which block</td>
  </tr>
  <tr>
    <td>recv_amount</td>
    <td>string</td>
    <td>Received amount denominated in the smallest cryptocurrency unit</td>
  </tr>
  <tr>
    <td>fees</td>
    <td>string</td>
    <td>Mining fee denominated in the smallest cryptocurrency unit</td>
  </tr>
  <tr>
    <td>broadcast_at</td>
    <td>int64</td>
    <td>When to broadcast the transaction in UTC time</td>
  </tr>
  <tr>
    <td>from_address</td>
    <td>string</td>
    <td>The source address of the transaction</td>
  </tr>
  <tr>
    <td>to_address</td>
    <td>string</td>
    <td>The destination address of the transaction</td>
  </tr>
  <tr>
    <td>state</td>
    <td>int</td>
    <td>
      Possible states (listed in the Order State Definition table)
	 	<table>
	 	  <thead><tr><td>ID</td><td>Description</td></tr></thead>
	 	  <tbody>
		    <tr><td>0</td><td>The order has been successfully paid</td></tr>
		    <tr><td>1</td><td>The order has expired</td></tr>
		    <tr><td>2</td><td>The amount received is less than the amount requested</td></tr>
		    <tr><td>3</td><td>The amount received is greater than the amount requested</td></tr>
		    <tr><td>4</td><td>The order cancelled</td></tr>
	 	  </tbody>
		</table>
    </td>
  </tr>
  <tr>
    <td>decimal</td>
    <td>int</td>
    <td>The decimal of cryptocurrency</td>
  </tr>
  <tr>
    <td>addon</td>
    <td>key-value pairs</td>
    <td>
    The extra information of this callback
	 	<table>
	 	  <thead><tr><td>Key</td><td>Value (Description)</td></tr></thead>
	 	  <tbody>
		    <tr><td>fee_decimal</td><td>The decimal of cryptocurrency miner fee</td></tr>
	 	  </tbody>
		</table>
    </td>
  </tr>
  <tr>
    <td>currency_bip44</td>
    <td>int64</td>
    <td>
   	 	<table>
   	 	  <thead><tr><td>ID</td><td>Currency Symbol</td><td>Decimals</td></tr></thead>
   	 	  <tbody>
		    <tr><td>0</td><td>BTC</td><td>8</td></tr>
  		    <tr><td>2</td><td>LTC</td><td>8</td></tr>
  		    <tr><td>5</td><td>DASH</td><td>8</td></tr>
  		    <tr><td>60</td><td>ETH</td><td>18</td></tr>
  		    <tr><td>144</td><td>XRP</td><td>6</td></tr>
  		    <tr><td>145</td><td>BCH</td><td>8</td></tr>
  		    <tr><td>148</td><td>XLM</td><td>7</td></tr>
  		    <tr><td>194</td><td>EOS</td><td>4</td></tr>
   		    <tr><td>195</td><td>TRX</td><td>6</td></tr>
   		    <tr><td>236</td><td>BSV</td><td>8</td></tr>
   		    <tr><td>354</td><td>DOT</td><td>10</td></tr>
   		    <tr><td>461</td><td>FIL</td><td>18</td></tr>
   		    <tr><td>714</td><td>BNB</td><td>8</td></tr>
   		    <tr><td>1815</td><td>ADA</td><td>6</td></tr>
   		    <tr><td>99999999997</td><td>BSC</td><td>18</td></tr>
   	 	  </tbody>
		</table>
    </td>
  </tr>
  <tr>
    <td>token_address</td>
    <td>string</td>
    <td>The contract address of cryptocurrency</td>
  </tr>
</table>

##### [Back to top](#table-of-contents)


<a name="order-state-definition"></a>
### Order State Definition

| ID   | Description |
| :--- | :---        |
| -1 | Waiting for payment |
| 0 | The order has been successfully paid |
| 1 | The order has expired |
| 2 | The amount received is less than the amount requested |
| 3 | The amount received is greater than the amount requested |
| 4 | The order cancelled |
 
##### [Back to top](#table-of-contents)

