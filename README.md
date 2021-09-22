<a name="table-of-contents"></a>
## Table of contents

- [Get Started](#get-started)
- [API Authentication](#api-authentication)
- [Callback Integration](#callback-integration)
- [Cryptocurrency Unit Conversion](#cryptocurrency-unit-conversion)
- REST API
	- Deposit Wallet API
		- [Create Deposit Addresses](#create-deposit-addresses)
		- [Query Deposit Addresses](#query-deposit-addresses)
		- [Query Deployed Contract Deposit Addresses](#query-deployed-contract-deposit-addresses)
		- [Query Pool Address](#query-pool-address)
		- [Query Pool Address Balance](#query-pool-address-balance)
		- [Query Invalid Deposit Addresses](#query-invalid-deposit-addresses)
		- [Query Deposit Callback Detail](#query-deposit-callback-detail)
		- [Resend Deposit/Collection Callbacks](#resend-deposit-callbacks)
		- [Query Deposit Wallet Balance](#query-deposit-wallet-balance)
		- [Update Deposit Address Label](#update-deposit-address-label)
		- [Query Deposit Address Label](#query-deposit-address-label)
	- Withdraw Wallet API
		- [Withdraw Assets](#withdraw-assets)
		- [Cancel Withdrawal Request](#cancel-withdrawal-request)
		- [Query Latest Withdrawal Transaction State](#query-latest-withdrawal-transaction-state)
		- [Query All Withdrawal Transaction States](#query-all-withdrawal-transaction-states)
		- [Query Withdrawal Wallet Balance](#query-withdrawal-wallet-balance)
		- [Query Withdrawal Callback Detail](#query-withdrawal-callback-detail)
		- [Set Withdrawal Request ACL](#set-withdrawal-request-acl)
		- [Resend Withdrawal Callbacks](#resend-withdrawal-callbacks)
		- [Query Withdrawal Whitelist Configuration](#query-withdrawal-whitelist-configuration)
		- [Add Withdrawal Whitelist Entry](#add-withdrawal-whitelist-entry)
		- [Remove Withdrawal Whitelist Entry](#remove-withdrawal-whitelist-entry)
		- [Check Withdrawal Whitelist](#check-withdrawal-whitelist)
		- [Query Withdrawal Whitelist](#query-withdrawal-whitelist)
		- [Query Withdrawal Wallet Transaction History](#query-withdrawal-wallet-transaction-history)
	- Deposit / Withdraw Wallet Common API
		- [Query Callback History](#query-callback-history)
		- [Query Callback Detail](#query-callback-detail)
		- [Query Wallet Synchronization Info](#query-wallet-synchronization-info)
		- [Query Transaction Average Fee](#query-transaction-average-fee)
		- [Batch Query Transaction Average Fees](#batch-query-transaction-average-fees)
	- Vault Wallet API
		- [Query Vault Wallet Transaction History](#query-vault-wallet-transaction-history)
		- [Query Vault Wallet Balance](#query-vault-wallet-balance)
	- Common API
		- [Activate API Code](#activate-api-code)
		- [Query API Code Status](#query-api-code-status)
		- [Refresh API Code](#refresh-api-code)
		- [Query Wallet Info](#query-wallet-info)
		- [Verify Addresses](#verify-addresses)
		- [Inspect Callback Endpoint](#inspect-callback-endpoint)
	- Read-only API code API
		- [List Wallets](#list-wallets)
- Testing
	- [Mock Server](#mock-server)
	- [cURL Testing Commands](#curl-testing-commands)
	- [Other Language Versions](#other-language-versions)
- Appendix
	- [Callback Definition](#callback-definition)
	- [Transaction State Definition](#transaction-state-definition)
	- [Callback Type Definition](#callback-type-definition)
	- [Currency Definition](#currency-definition)
	- [Memo Requirement](#memo-requirement)

<a name="get-started"></a>
# Get Started

### How to deposit?
- Setup a deposit wallet and configure it (via web control panel)
	- Refer to CYBAVO VAULT SOFA User Manual for detailed steps.
- Request an API code/secret (via web control panel)
- Create deposit addresses (via CYBAVO SOFA API)
	- Refer to [Create deposit addresses](#create-deposit-addresses) API
- Waiting for the CYBAVO SOFA system detecting transactions to those deposit addresses
- Handle the deposit callback
	- Use the callback data to update certain data on your system.
	- <b>Security Enhancement</b>: Use the Query Callback Detail API to confirm the callback is sent from the CYBAVO SOFA system.
	- Refer to [Callback Integration](#callback-integration) for detailed information.

### How to withdraw?
- Setup a withdrawal wallet and configure it (via web control panel)
	- Refer to CYBAVO VAULT SOFA User Manual for detailed steps.
- Request an API code/secret (via web control panel)
- Make withdraw request (via CYBAVI SOFA API)
	- Refer to [Withdraw Assets](#withdraw-assets) API
	- <b>Security Enhancement</b>: Also set the withdrawal authentication callback URL to authorize the withdrawal requests sent to the CYBAVO SOFA system.
- Waiting for the CYBAVO SOFA system broadcasting transactions to blockchain
- Handle the withdrawal callback
	- Use the callback data to update certain data on your system.
	- Refer to [Callback Integration](#callback-integration) for detailed information.

### Try it now
- Use [mock server](#mock-server) to test CYBAVO SOFA API right away.

### Start integration
- To make a correct API call, refer to [API Authentication](#api-authentication).
- To handle callback correctly, refer to [Callback Integration](#callback-integration).

<a name="api-authentication"></a>
# API Authentication

- The CYBAVO SOFA system verifies all incoming requests. All requests must include X-API-CODE, X-CHECKSUM headers otherwise caller will get a 403 Forbidden error.

### How to acquire and refresh API code and secret
- Request the API code/secret from the **Wallet Details** page on the web control panel for the first time.
- A paired refresh code can be used in the [refresh API](#refresh-api-code) to acquire the new inactive API code/secret of the wallet.
	- Before the inactive API code is activated, the currently activated API code is still valid.
	- Once the paired API code becomes invalid, the paired refresh code will also become invalid.

### How to make a correct request?
- Put the API code in the X-API-CODE header.
	- Use the inactivated API code in any request will activate it automatically. Once activated, the currently activated API code will immediately become invalid.
	- Or you can explicitly call the [activation API](#activate-api-code) to activate the API code before use
- Calculate the checksum with the corresponding API secret and put the checksum in the X-CHECKSUM header.
  - The checksum calculation will use all the query parameters, the current timestamp, user-defined random string and the post body (if any).
- Please refer to the code snippet on the github project to know how to calculate the checksum.
	- [Go](https://github.com/CYBAVO/API_CHECKSUM_CALC/blob/main/go/checksum.go#L40)
	- [Java](https://github.com/CYBAVO/API_CHECKSUM_CALC/blob/main/java/checksum.java#L49)
	- [Javascript](https://github.com/CYBAVO/API_CHECKSUM_CALC/blob/main/javascript/checksum.js#L27)
	- [PHP](https://github.com/CYBAVO/API_CHECKSUM_CALC/blob/main/php/checksum.php#L27)
	- [C#](https://github.com/CYBAVO/API_CHECKSUM_CALC/blob/main/c%23/checksum.cs#L55)
	- [Python](https://github.com/CYBAVO/API_CHECKSUM_CALC/blob/main/python/checksum.py#L29)

<a name="readonly-api-code"></a>
# Read-only API Code

- A read-only API code can be used to call all the read functions of wallets.
	- All the read functions will be labeled `VIEW` in front of the API definition.
- Use [activation API](#activate-api-code) with the `WALLET_ID` set as `readonly` to activate a read-only API code.
	- The full API path is `/v1/sofa/wallets/readonly/apisecret/activate`
	- After activation, the API code will remain valid until it is replaced by a newly activated read-only API code.
- Use [listing API](#list-wallets) to list all wallets that can be accessed through a read-only API code.

<a name="callback-integration"></a>
# Callback Integration

- Please note that the wallet must have an activated API code, otherwise no callback will be sent.
	- Use the [activation API](#activate-api-code) to activate an API code.
- How to distinguish between deposit and withdrawal callbacks?
	- Deposit Callback (callback type 1)
	  - The combination of **txid** and **vout_index** of the callback is unique, use this combined ID to identify the deposit request, not to use only the transaction ID (txid field). Because multiple deposit callbacks may have the same transaction ID, for example, BTC many-to-many transactions.
	- Withdrawal Callback (callback type 2)
	  - The **order_id** of the callback is unique, use this ID to identify the withdrawal request.

<div class="alert alert-warning">
It is important to distinguish between unique callbacks to avoid improper handling of deposit / withdrawal requests.
</div>

- To ensure that the callbacks have processed by callback handler, the CYBAVO SOFA system will continue to send the callbacks to the callback URL until a callback confirmation (HTTP/1.1 200 OK) is received or exceeds the number of retries (retry time interval: 1-3-5-15-45 mins).
	- If all attempts fail, the callback will be set to a failed state, the callback handler can call the [resend deposit callback](#resend-deposit-callbacks) or [resend withdrawal callback](#resend-withdrawal-callbacks) API to request CYBAVO SOFA system to resend such kind of callback(s) or through the web control panel.

- Refer to [Callback Definition](#callback-definition), [Callback Type Definition](#callback-type-definition) for detailed definition.
- Please refer to the code snippet on the github project to know how to validate the callback payload.
	- [Go](https://github.com/CYBAVO/SOFA_MOCK_SERVER/blob/master/controllers/OuterController.go#L197)
	- [Java](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVA/blob/master/src/main/java/com/cybavo/sofa/mock/MockController.java#L93)
	- [Javascript](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVASCRIPT/blob/master/routes/wallets.js#L399)
	- [PHP](https://github.com/CYBAVO/SOFA_MOCK_SERVER_PHP/blob/master/index.php#L207)
	- [C#](https://github.com/CYBAVO/API_CHECKSUM_CALC/blob/main/c%23/checksum.cs#L89)
	- [Python](https://github.com/CYBAVO/API_CHECKSUM_CALC/blob/main/python/checksum.py#L64)

- Best practice:
	- While processing a deposit callback, in addition to verifying the checksum of the callback, use [Query Callback Detail](#query-callback-detail) API with the serial ID of the callback to perform an additional confirmation.


<a name="callback-state-change"></a>
# Callback State Change

#### The state of a successful withdrawal request is changed as follows:

processing state(1) -> transaction in pool state(2) -> transaction in chain state(3) -> repeats state 3 until the confirmation count is met

#### The state of a successful deposit request is changed as follows:

transaction in chain state(3) -> repeats state 3 until the confirmation count is met

> Refer to [Transaction State Definition](#transaction-state-definition) for all transaction states definition.


<a name="cryptocurrency-unit-conversion"></a>
# Cryptocurrency Unit Conversion

#### For callback

- The amount and fees fields in the callback are in the smallest cryptocurrency unit, use `decimal` and `fee_decimal`(in the addon field) fields of callback data to convert the unit.

#### For API

- Refer to decimals of [Currency Definition](#currency-definition) to convert main cryptocurrency unit.
- For the cryptocurrency token, use the token_decimals field of the [Wallet Info](#query-wallet-info) API to convert cryptocurrency token unit.


# REST API

# Deposit Wallet API

<a name="create-deposit-addresses"></a>
### Create Deposit Addresses

Create deposit addresses on certain wallet. Once addresses are created, the CYBAVO SOFA system will callback when transactions are detected on these addresses.

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/addresses

> `WALLET_ID` must be a deposit wallet ID

- [Sample curl command](#curl-create-deposit-addresses)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/107950/addresses
```

###### Post body

For BNB, XLM, XRP or EOS wallet:

```json
{
  "count": 2,
  "memos": [
    "10001",
    "10002"
  ],
  "labels": [
  	"note-for-001",
  	"note-for-002"
  ]
}
```

For wallet excepts BNB, XLM, XRP and EOS:

```json
{
  "count": 2,
   "labels": [
  	"note-for-address-1",
  	"note-for-address-2"
  ]
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| count | int | required, max `1000` | Specify address count |
| memos | array | required (creating BNB, XLM, XRP or EOS wallet) | Specify memos for BNB, XLM, XRP or EOS deposit wallet. Refer to [Memo Requirement](#memo-requirement) |
| labels | array | optional | Specify the labels of the generated addresses or memos |

> NOTE: The length of `memos` must equal to `count` while creating addresses for BNB, XLM, XRP or EOS wallet.
> 
> NOTE: The memos(or called destination tags) of XRP must be strings that can be converted to numbers.
> 
> If use the `labels` to assign labels, the array length of the labels must equal to `count`.
> The label will be automatically synced to the child wallet.

##### Response Format

An example of a successful response:

For BNB, XLM, XRP or EOS wallet:
	
```json
{
  "addresses": [
    "10001",
    "10002"
  ]
}
```
	
For wallet excepts BNB, XLM, XRP or EOS:
	
```json
{
  "addresses": [
    "0x2E7248BBCD61Ad7C33EA183A85B1856bc02C40b6",
    "0x4EB990D527c96c64eC5Bfb0D1e304840052d4975",
    "0x86434604FF857702fbE11cBFf5aC7689Af19c4Ed"
  ]
}
```

For the ETH wallet that uses contract collection:

```json
{
  "txids": [
    "0xe6dfe0d283690f636df5ea4b9df25552e6b576b88887bfb5837016cdd696e754",
    "0xdb18fd33c9a6809bfc341a1c0b2c092be5a360f394c85367f9cf316579281ab4",
    "0x18075ff1693026f93722f8b2cc0e29bf148ded5bce4dc173c8118951eceabe60",
    "0x7c6acb506ef033c09f781cc5ad6b2d0a216346758d7f955e720d6bc7a52731a5",
    "0x7da19f8c0d82cde16636da3307a6bef46eb9f398af3eb2362d230ce300509d63"
  ]
}
```

Use [Query Deployed Contract Deposit Addresses](#query-deployed-contract-deposit-addresses) API to query deployed contract addresses.


The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| addresses | array | Array of just created deposit addresses |
| txids | array | Array of transaction IDs used to deploy collection contract |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 112 | Invalid parameter | - | The count and the count of memos mismatched |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 403 | 706 | Exceed max allow wallet limitation, Upgrade your SKU to get more wallets | - | Reached the limit of the total number of deposit addresses |
| 400 | 421 | Mapped(Token) wallet not allow to create deposit addresses, please create the deposit wallet in parent wallet, the address will be synced to mapped wallet automatically | - | Only the parent wallet can create deposit addresses |
| 400 | 500 | insufficient fund | - | Insufficient balance to deploy collection contract |
| 400 | 703 | Operation failed | Error message returned by JSON parser | Malformatted post body |
| 400 | 818 | Destination Tag must be integer | - | Wrong XRP destination tag format |
| 400 | 945 | The max length of BNB memo is 256 chars | - | Reached the limit of the length of BNB memo |
| 400 | 946 | The max length of EOS memo is 128 chars | - | Reached the limit of the length of EOS memo |
| 400 | 947 | The max length of XRP destination tag is 20 chars | - | Reached the limit of the length of XRP destination tag |
| 400 | 948 | The max length of XLM memo is 20 chars | - | Reached the limit of the length of XLM memo |
| 404 | 304 | Wallet ID invalid | archived wallet or wrong wallet type | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-deposit-addresses"></a>
### Query Deposit Addresses

Query the deposit addresses created by the [Create Deposit Addresses](#create-deposit-addresses) API.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/addresses?start\_index=`from`&request\_number=`count`

> `WALLET_ID` must be a deposit wallet ID

- [Sample curl command](#curl-query-deposit-addresses)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/179654/addresses?start_index=0&request_number=3

--- THEN ---

/v1/sofa/wallets/179654/addresses?start_index=3&request_number=3
```

The request includes the following parameters:

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| start_index | int | optional, default `0` | Specify address start index |
| request_number | int | optional, default `1000`, max `5000` | Request address count |

##### Response Format

An example of a successful response:

```json
{
  "wallet_id": 179654,
  "wallet_count": 6,
  "wallet_address": [
    {
      "currency": 60,
      "token_address": "",
      "address": "0x8c42fD03A5cfba7C3Cd97AB8a09e1a3137Ef33C3",
      "memo": ""
    },
    {
      "currency": 60,
      "token_address": "",
      "address": "0x4d3EB54b602BF4985CE457089F9fB084Af597A2C",
      "memo": ""
    },
    {
      "currency": 60,
      "token_address": "",
      "address": "0x74dc3fB523295C87C0b93E48744Ce94fe3a8Ef5e",
      "memo": ""
    }
  ]
}

--- THEN ---

{
  "wallet_id": 179654,
  "wallet_count": 6,
  "wallet_address": [
    {
      "currency": 60,
      "token_address": "",
      "address": "0x6d68443D6564cF257A48c1b16aa6d0EF13c5A719",
      "memo": ""
    },
    {
      "currency": 60,
      "token_address": "",
      "address": "0x26F103322B6f0ed2D35B85F1611589c92F023986",
      "memo": ""
    },
    {
      "currency": 60,
      "token_address": "",
      "address": "0x2b91918Bee4411DaD6293EA5d6D38251E72723Ca",
      "memo": ""
    }
  ]
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| wallet_id | int64 | ID of request wallet |
| wallet_address | array | Array of wallet addresses |
| wallet_count | int64 | Total count of deposit addresses |

> Refer to [Currency Definition](#currency-definition) or [here](https://github.com/satoshilabs/slips/blob/master/slip-0044.md) for more detailed currency definitions

> If this is an ETH contract collection deposit wallet, only the deployed address will be returned.

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-deployed-contract-deposit-addresses"></a>
### Query Deployed Contract Deposit Addresses

Query deployed contract deposit addresses created by the [Create Deposit Addresses](#create-deposit-addresses) API.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/addresses/contract_txid?txids=`txid1,txid2`

> `WALLET_ID` must be an ETH contract collection deposit wallet ID
> 
> Only deployed addresses will be returned

- [Sample curl command](#curl-query-deployed-contract-deposit-addresses)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/179654/addresses/contract_txid?txids=0xe6dfe0d283690f636df5ea4b9df25552e6b576b88887bfb5837016cdd696e754,0xdb18fd33c9a6809bfc341a1c0b2c092be5a360f394c85367f9cf316579281ab4
```

The request includes the following parameters:

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| txids | string | requried, max `10` transaction IDs | Transaction ID used to deploy collection contract |

##### Response Format

An example of a successful response:

```json
{
  "addresses": {
    "0xdb18fd33c9a6809bfc341a1c0b2c092be5a360f394c85367f9cf316579281ab4": {
      "address": "0x00926cE2BbF56317c72234a0Fb8A65A1A15F7103",
      "currency": 60,
      "memo": "",
      "token_address": ""
    },
    "0xe6dfe0d283690f636df5ea4b9df25552e6b576b88887bfb5837016cdd696e754": {
      "address": "0xf3747e3edbd8B8414718dd51330415c171e79208",
      "currency": 60,
      "memo": "",
      "token_address": ""
    }
  }
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| addresses | map object | The map KEY is Transaction ID used to deploy collection contract and the map VALUE is the address information |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-pool-address"></a>
### Query Pool Address

Get the pool address of a deposit wallet. The pool address has different functionality in different cryptocurrencies.

> In BTC, ETH, BCH or LTC, the cryptocurrency in the pool address will be used to pay for token transfer(ex. ERC20, USDT-Omni).
> 
> In EOS, XRP, XLM or BNB, the pool address is the user's deposit address. All user deposits will be distinguished by memo / tag field.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/pooladdress

> `WALLET_ID` must be a deposit wallet ID

- [Sample curl command](#curl-get-pool-address)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/179654/pooladdress
```

##### Response Format

An example of a successful response:

```json
{
  "address": "0x36099775afa8d6363aC8e5d0fC698306C021a858"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| address  | string | Pool address of wallet |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-pool-address-balance"></a>
### Query Pool Address Balance

Get the pool address balance of a deposit wallet.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/pooladdress/balance

> `WALLET_ID` must be a deposit wallet ID

- [Sample curl command](#curl-query-pool-address-balance)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/179654/pooladdress/balance
```

##### Response Format

An example of a successful response:

```json
{
  "balance": "0.515",
  "currency": 60,
  "unconfirm_balance": "0",
  "wallet_address": "0xb6ad80c96D093EA584AfcB9443927812d3e4Bd94"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| balance | string | Pool address balance |
| unconfirm\_balance | string | Unconfirmed pool address balance |
| currency | int64 | Cryptocurrency of the wallet |
| wallet_address  | string | Pool address of the wallet |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)



<a name="query-invalid-deposit-addresses"></a>
### Query Invalid Deposit Addresses

When an abnormal deposit is detected, the CYBAVO SOFA system will set the deposit address to invalid. Use this API to obtain the all invalid deposit addresses for further usage.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/addresses/invalid-deposit

> `WALLET_ID` must be a deposit wallet ID

- [Sample curl command](#curl-query-invalid-deposit-addresses)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/179654/addresses/invalid-deposit
```

##### Response Format

An example of a successful response:

```json
{
  "addresses": ["0x5dB3d8C70dAa9C919F9962221c2fDDbe8EBAa5F2"]
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| addresses | array | Array of invalid deposit address |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-deposit-callback-detail"></a>
### Query Deposit Callback Detail

Query the detailed information of the deposit callback by the tx ID and the vout index.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/receiver/notifications/txid/`TX_ID`/`VOUT_INDEX`

- [Sample curl command](#curl-query-deposit-callback-detail)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/179654/receiver/notifications/txid/0xb72a81976f780445decd13a35c24974c4e32665cb57d79b3f7a601c775f6a7d8/0
```

##### Response Format

An example of a successful response:

```json
{
  "notification": {
    "addon": {},
    "amount": "2000000000000000000",
    "block_height": 7757485,
    "broadcast_at": 1587441501,
    "chain_at": 1587441501,
    "confirm_blocks": 166027,
    "currency": "ETH",
    "fees": "126000000000000",
    "from_address": "0x8382Cc1B05649AfBe179e341179fa869C2A9862b",
    "memo": "",
    "order_id": "",
    "processing_state": 2,
    "serial": 90000000547,
    "state": 3,
    "tindex": 27,
    "to_address": "0x32d638773cB85965422b3B98e9312Fc9392307BC",
    "txid": "0xb72a81976f780445decd13a35c24974c4e32665cb57d79b3f7a601c775f6a7d8",
    "type": 1,
    "vout_index": 0,
    "wallet_id": 179654
  }
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| notification | object | Refer to [Callback Definition](#callback-definition) |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request or the callback (txid+vout_index) not found |

##### [Back to top](#table-of-contents)


<a name="resend-deposit-callbacks"></a>
### Resend Deposit/Collection Callbacks

The callback handler can call this API to resend pending, risk-controlled or failed deposit/collection callbacks.

Refer to [Callback Integration](#callback-integration) for callback rules.

> The resend operation could be requested on the web control panel as well.

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/collection/notifications/manual

> `WALLET_ID` must be a deposit wallet ID

- [Sample curl command](#curl-resend-deposit-callbacks)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/179654/collection/notifications/manual
```

###### Post body

```json
{
  "notification_id": 0
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :---  | :---        |
| notification_id | int64 | required, 0 means all | Specify callback ID to resend |

> This ID equal to the serial field of callback data.

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
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | Malformatted post body |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-deposit-wallet-balance"></a>
### Query Deposit Wallet Balance

Get the deposit wallet balance.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/receiver/balance

> `WALLET_ID` must be a deposit wallet ID

- [Sample curl command](#curl-query-deposit-wallet-balance)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/959272/receiver/balance
```

##### Response Format

An example of a successful response:

```json
{
  "currency": 60,
  "token_address": "",
  "balance": "0.619673333517576",
  "token_balance": ""
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| currency | int64 | Registered coin types. Refer to [Currency Definition](#currency-definition) |
| token_address | string | Token contract address |
| balance | string | Deposit wallet balance |
| token_balance | string | Deposit wallet token balance |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="update-deposit-address-label"></a>
### Update Deposit Address Label

Update the label of the deposit address.

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/addresses/label

> `WALLET_ID` must be a deposit wallet ID
> 
> The label will be automatically synced between the parent and child wallet.

- [Sample curl command](#curl-update-deposit-address-label)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/98675/addresses/label
```

###### Post body

```json
{
  "address": "0x2B974a3De0b491bB26e0bF146E6cDaC36EFD574a",
  "label": "take-some-notes"
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :---  | :---        |
| address | string | required | Specify address to update the label |
| label | string | optional, set empty to clear the label | Specify the label of the address |

##### Response Format

Status code 200 represnts a successful operation

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | Malformatted post body |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |
| 404 | 112 | Invalid parameter | - | The address can not be found |

##### [Back to top](#table-of-contents)


<a name="query-deposit-address-label"></a>
### Query Deposit Address Label

Query the labels of the deposit addresses.

##### Request

`VIEW` **POST** /v1/sofa/wallets/`WALLET_ID`/addresses/get_labels

> `WALLET_ID` must be a deposit wallet ID

- [Sample curl command](#curl-query-deposit-address-label)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/98675/addresses/label
```

###### Post body

```json
{
  "addresses": [
    "0x2B974a3De0b491bB26e0bF146E6cDaC36EFD574a",
    "0xF401AC94D9672e79c68e56A6f822b666E5A7d644"
  ]
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :---  | :---        |
| addresses | array | required | Specify the addresses to query labels |

##### Response Format

An example of a successful response:

```json
{
  "labels": {
    "0x2B974a3De0b491bB26e0bF146E6cDaC36EFD574a": "take-some-notes",
    "0xF401AC94D9672e79c68e56A6f822b666E5A7d644": ""
  }
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| labels | key-value pairs | The address-label pairs |

> If the address can not be found, it will not be listed in the response.

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | Malformatted post body |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


# Withdraw Wallet API

<a name="withdraw-assets"></a>
### Withdraw Assets

To withdraw assets from an withdrawal wallet, the caller must to provide an unique **order_id** for each request, the CYBAVO SOFA system will send the callback with the unique **order_id** when the withdrawal is success (from `in pool` state to `in chain` state). 

By default, the withdraw API will perform the address check to verify that the outgoing address is good or not. If the address in the request is marked as a problematic address, the request will be aborted. The error message will identify the problematic addresses. Set the `ignore_black_list` to true to skip the address check.

The withdrawal API can also interact with the contracts (ERC/BEP 721/1155) deployed in the SOFA system.

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/sender/transactions

> `WALLET_ID` must be a withdrawal wallet ID
> 
> The order\_id must be prefixed. **Find prefix from corresponding wallet detail on web control panel.**
>
> If withdraw BNB or XRP, this API will check whether the destination addresse needs memo / destination tag or not. If the address does need memo / destination tag, the API will fail without memo / destination tag specified.

- [Sample curl command](#curl-withdraw-assets)

##### Request Format

An example of the request:

> The prefix is 888888_ in following sample request.

###### API

```
/v1/sofa/wallets/68451/sender/transactions
```

###### Post body

```json
{
  "requests": [
    {
      "order_id": "888888_1",
      "address": "0x83eE561B2aBD000FF00d6ca22f38b29d4a760d4D",
      "amount": "0.0001",
      "memo": "memo-001",
      "user_id": "USER01",
      "message": "message-001",
      "block_average_fee": 5
    },
    {
      "order_id": "888888_2",
      "address": "0xf16B7B8900F0d2f682e0FFe207a553F52B6C7015",
      "amount": "0.0002",
      "manual_fee": 50
    },
    {
      "order_id": "888888_3",
      "address": "0x9638fa816ccd35389a9a98a997ee08b5321f3eb9",
      "amount": "0.0002",
      "message": "message-003"
    },
    {
      "order_id": "888888_4",
      "address": "0x2386b18e76184367b844a402332703dd2eec2a90",
      "amount": "0",
      "contract_abi":"create:0x000000000000000000000000000000000000000000000000000000000000138800000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"
      "user_id": "USER04"
    },
    {
      "order_id": "888888_5",
      "address": "0x2386b18e76184367b844a402332703dd2eec2a90",
      "amount": "1",
      "token_id": "985552421"
    }
  ],
  "ignore_black_list": false
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| order_id | string | required, max `255` chars | Specify an unique ID, order ID must be prefixed |
| address | string | required | Outgoing address (`address` must be a contract address, if the contract_abi is not empty) |
| amount | string | required | Withdrawal amount |
| contract_abi | string | required, if calls contract ABI | Specify the ABI method and the parameters, in the format `ABI_method:parameters` |
| memo | string | optional | Memo on blockchain (This memo will be sent to blockchain). Refer to [Memo Requirement](#memo-requirement) |
| user_id | string | optional | Specify certain user |
| message | string | optional | Message (This message only saved on CYBAVO, not sent to blockchain) |
| block\_average_fee | int | optional, range `1~30` | Use average blockchain fee within latest N blocks |
| manual_fee | int | optional, range `1~1000` | Specify blockchain fee in smallest unit of wallet currency |
| token_id | string | optional | Specify the token ID to be transferred |
| ignore\_black_list| boolean | optional, default `false` | After setting, the address check will not be performed. |

> The order\_id must be prefixed. Find prefix from corresponding wallet detail on web control panel
>
> block\_average\_fee and manual_fee are mutually exclusive configurations. If neither of these fields is set, the fee will refer to corresponding withdrawal policy of the withdrawal wallet.
> 
> The format of the `contract_abi` is `ABI_method:hex_parameters`, for example: create:0x000000000000000000000000000000000000000000000000000000000000138800000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000. The parameters must be encoded by [web3.eth.abi.encodeParameters() of web3.js](https://web3js.readthedocs.io/en/v1.3.4/web3-eth-abi.html#encodeparameters).
> 
> Only ERC721/1155 wallet can use `token_id` to transfer token. For ERC721 wallets, if `token_id` is specified, the amount will be ignored.

##### Response Format

An example of a successful response:

```json
{
  "results": {
    "1": 20000000001,
    "2": 20000000002
  }
}
```	

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| results | array | Array of withdraw result (order ID/withdraw transaction ID pair), if succeeds |

An example response of the request contains problematic addresses:

```json
{
    "error_code": 827,
    "error": "Outgoing address in black list, abort transaction",
    "blacklist": {
        "0x83eE561B2aBD000FF00d6ca22f38b29d4a760d4D": [
            "Involve phishing activity",
            "Involve cybercrime related"
        ]
    }
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| error_code | int | The error code |
| error | string | The error message |
| blacklist | object | The object describes all problematic addresses and their causes. |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 403 | 827 | Outgoing address in black list, abort transaction | - | Some outgoing addresses are blacklisted, examine the response 'blacklist' field for detailed information |
| 400 | 112 | Invalid parameter | - | Malformatted post body |
| 400 | 955 | There is no content in your withdrawal request, please check your input | - | The post body of request doesn't conform the API request specification |
| 400 | 703 | Operation failed | order_id must start with {ORDERID\_PREFIX} | The prefix of order_id is incorrect |
| 400 | 703 | Operation failed | order_id: {ORDER\_ID} - the character \\ or / is prohibited | {ORDER\_ID} is invalid |
| 400 | 703 | Operation failed | order_id: {ORDER\_ID} is invalid | {ORDER\_ID} is invalid |
| 400 | 703 | Operation failed | order_id: {ORDER\_ID} - memo is required | The outgoing address of {ORDER\_ID} needs memo specified |
| 400 | 703 | Operation failed | order_id: {ORDER\_ID} - destination tag is required | The outgoing address of {ORDER\_ID} needs destination tag specified |
| 400 | 703 | Operation failed | order_id: {ORDER\_ID} - invalid block\_average\_fee | The block\_average\_fee is out of range |
| 400 | 703 | Operation failed | order_id: {ORDER\_ID} - invalid manual\_fee | The manual\_fee is out of range |
| 400 | 399 | Duplicated entry: {ORDER\_ID} | - | The {ORDER\_ID} is duplicated |
| 400 | 945 | The max length of BNB memo is 256 chars | - | Reached the limit of the length of BNB memo |
| 400 | 946 | The max length of EOS memo is 128 chars | - | Reached the limit of the length of EOS memo |
| 400 | 947 | The max length of XRP destination tag is 20 chars | - | Reached the limit of the length of XRP destination tag |
| 400 | 948 | The max length of XLM memo is 20 chars | - | Reached the limit of the length of XLM memo |
| 400 | 818 | Destination Tag must be integer | - | Wrong XRP destination tag format |
| 400 | 944 | The max length of order id is 255 chars | - | Reached the limit of the length of order_id |
| 400 | 703 | Operation failed | Detailed error message | Failed to connect to authentication callback URL |
| 400 | 703 | Operation failed | The withdrawal request has been rejected, {RESPONSE_BODY} | The withdrawal request has been rejected by the authentication callback |
| 400 | 703 | Operation failed | The withdrawal request has been rejected, unexpected response {HTTP\_CODE}: {RESPONSE_BODY} | The authentication callback URL returned status code other than 200 or 400 |
| 400 | 703 | Operation failed | Unrecognized response: {RESPONSE_BODY}, 'OK' expected | The returned status code is 200 but the body is not **OK** |
| 400 | 703 | Operation failed | request IP ({IPv4}) not in ACL | The request IP not in the withdrawal ACL |
| 400 | 703 | Operation failed | invalid amount {AMOUNT} | The requested amount is not a valid number |
| 400 | 703 | Operation failed | invalid amount decimals {AMOUNT} | The decimals of the requested amount exceeds the valid range |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |
| 404 | 312 | Policy not found | no active withdrawal policy found | No active withdrawal policy found |


##### [Back to top](#table-of-contents)


<a name="cancel-withdrawal-request"></a>
### Cancel Withdrawal Request

To cancel the withdrawal request which state is `Init` or `Failed`. The request state can be checked on web control panel or query through this [API](#query-withdrawal-callback-detail) (represents `state` = 0 or 5 ).

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/sender/transactions/`ORDER_ID`/cancel

> `WALLET_ID` must be a withdrawal wallet ID

- [Sample curl command](#curl-cancel-withdrawal-request)

##### Request Format

An example of the request:

> The prefix is 888888_ in following sample request.

###### API

```
/v1/sofa/wallets/68451/sender/transactions/88888_1/cancel
```

##### Response Format

The HTTP 200 means the withdrawal request has been cancelled successfully.

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 177 | Illegal state | - | The {ORDER\_ID} withdrawal request is not in `Init` or `Failed` state |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The {ORDER\_ID} not found |

##### [Back to top](#table-of-contents)


<a name="query-latest-withdrawal-transaction-state"></a>
### Query Latest Withdrawal Transaction State

Check the latest withdrawal transaction state of certain order ID.

> The order ID is used in the [withdraw assets](#withdraw-assets) API.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/sender/transactions/`ORDER_ID`

> `WALLET_ID` must be a withdrawal wallet ID

- [Sample curl command](#curl-query-latest-withdrawal-transaction-state)

##### Request Format

An example of the request:

> The prefix is 888888_ in following sample request.

###### API

```
/v1/sofa/wallets/68451/sender/transactions/888888_1
```

##### Response Format

An example of a successful response:

```json
{
  "order_id": "888888_1",
  "address": "0xaa0cA2f9bA3A33a915a27e289C9719adB2ad7d73",
  "amount": "1.11",
  "memo": "",
  "in_chain_block": 1016603,
  "txid": "db0f3a27de564a411aeff1d2cb3234c54817de1ecc2258a510a50c5a1063d41c",
  "create_time": "2020-03-16T10:27:57Z"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| order_id | string | The unique ID specified in `sender/transactions` API |
| address | string | Outgoing address |
| amount | string | Withdrawal amount |
| memo | string | Memo on blockchain |
| in\_chain\_block | int64 | The block that contains this transaction |
| txid | string | Transaction ID |
| create_time | string | The withdrawal time in UTC |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The {ORDER\_ID} not found |

##### [Back to top](#table-of-contents)


<a name="query-all-withdrawal-transaction-states"></a>
### Query All Withdrawal Transaction States

Check the all withdrawal transaction states of certain order ID.

> The order ID is used in the [withdraw assets](#withdraw-assets) API.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/sender/transactions/`ORDER_ID`/all

> `WALLET_ID` must be a withdrawal wallet ID

- [Sample curl command](#curl-query-all-withdrawal-transaction-states)

##### Request Format

An example of the request:

> The prefix is 888888_ in following sample request.

###### API

```
/v1/sofa/wallets/68451/sender/transactions/888888_1/all
```

##### Response Format

An example of a successful response:

> The sample shows the states of a resent transaction

```json
{
  "transactions": [
    {
      "address": "0x36a49c68EF1e3f39CDbaE2f5636C74BA10815cea",
      "amount": "0.105",
      "create_time": "2020-09-24T03:43:17Z",
      "in_chain_block": 0,
      "memo": "",
      "order_id": "888888_1",
      "state": 6,
      "txid": "0x2a8a44f1cfed9cd7b86d86170e2418566765f88c5186246f571374df218fd1a1"
    },
    {
      "address": "0x36a49c68EF1e3f39CDbaE2f5636C74BA10815cea",
      "amount": "0.105",
      "create_time": "2020-09-24T03:44:35Z",
      "in_chain_block": 8742982,
      "memo": "",
      "order_id": "888888_1",
      "state": 4,
      "txid": "0xfbeaae4b87f977bcce8ef44672e035d287b96be24e779757c1a7f598501881ef"
    }
  ]
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| order_id | string | The unique ID specified in `sender/transactions` API |
| address | string | Outgoing address |
| amount | string | Withdrawal amount |
| memo | string | Memo on blockchain |
| in\_chain\_block | int64 | The block that contains this transaction |
| txid | string | Transaction ID |
| create_time | string | The withdrawal time in UTC |
| state | int | Refer to [Transaction State Definition](#transaction-state-definition) |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The {ORDER\_ID} not found |

##### [Back to top](#table-of-contents)


<a name="query-withdrawal-wallet-balance"></a>
### Query Withdrawal Wallet Balance

Get the withdrawal wallet balance. Facilitate to establish a real-time balance monitoring mechanism.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/sender/balance

> `WALLET_ID` must be a withdrawal wallet ID

- [Sample curl command](#curl-query-withdrawal-wallet-balance)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/632543/sender/balance
```

##### Response Format

An example of a successful response:

```json
{
  "currency": 60,
  "wallet_address": "0xaa0cA2f9bA3A33a915a27e289C9719adB2ad7d73",
  "token_address": "",
  "balance": "0.619673333517576",
  "token_balance": "",
  "unconfirm_balance": "0",
  "unconfirm_token_balance": ""
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| currency | int64 | Registered coin types. Refer to [Currency Definition](#currency-definition) |
| wallet_address | string | Wallet address |
| token_address | string | Token contract address |
| balance | string | Withdrawal wallet balance |
| token_balance | string | Withdrawal wallet token balance |
| unconfirm\_balance | string | Unconfirmed withdrawal wallet balance |
| unconfirm\_token_balance | string | Unconfirmed withdrawal wallet token balance |
| err_reason | string | Error message if fail to get balance |

> The currencies that support the unconfirmed balance are BTC, LTC, ETH, BCH, BSV, DASH

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-withdrawal-callback-detail"></a>
### Query Withdrawal Callback Detail

Query the detailed information of the withdrawal callback by the order ID.

> This API only provides in-chain transactions query, for those not in-chain transactions use [Query All Withdrawal Transaction States](#query-all-withdrawal-transaction-states) API instead.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/sender/notifications/order_id/`ORDER_ID`

- [Sample curl command](#curl-query-withdrawal-callback-detail)

##### Request Format

An example of the request:

> The prefix is 888888_ in following sample request.

###### API

```
/v1/sofa/wallets/68451/sender/notifications/order_id/888888_1
```

##### Response Format

An example of a successful response:

```json
{
  "notification": {
    "addon": {},
    "amount": "100000000000000",
    "block_height": 7813953,
    "broadcast_at": 1588211126,
    "chain_at": 1588211126,
    "confirm_blocks": 109490,
    "currency": "ETH",
    "fees": "21000000000000",
    "from_address": "0xaa0cA2f9bA3A33a915a27e289C9719adB2ad7d73",
    "memo": "",
    "order_id": "888888_1",
    "processing_state": 2,
    "serial": 90000000554,
    "state": 3,
    "tindex": 30,
    "to_address": "0x60589A749AAC632e9A830c8aBE042D1899d8Dd15",
    "txid": "0x471c11f139ce1a7e0627a05cea50d64e47e797c94fd72025f1978cc919e07aa9",
    "type": 2,
    "vout_index": 0,
    "wallet_id": 68451
  }
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| notification | object | Refer to [Callback Definition](#callback-definition) |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request or the {ORDER\_ID} is not found |

##### [Back to top](#table-of-contents)


<a name="set-withdrawal-request-acl"></a>
### Set Withdrawal Request ACL

Set an authorized IP to the withdrawal request ACL dynamically.

> If a static ACL has been set in web control panel, the API will fail.

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/sender/transactions/acl

- [Sample curl command](#curl-set-withdrawal-request-acl)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/632543/sender/transactions/acl
```

###### Post body

```json
{
  "acl": "192.168.101.55"
}
```

The request includes the following parameters:

###### Post body

| Field | Type | Note | Description |
| :---  | :--- | :--- | :---        |
| acl | string | requried | Specify an authorized IP in IPv4 format |


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
| result | int | Specify a successful API call |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The wallet is invalid to perform this API call |
| 400 | 180 | Invalid format | - | The acl field is empty or does not conform to the IPv4 format |
| 400 | 180 | Operation failed | ACL has been set via web | The static ACL is not empty |

##### [Back to top](#table-of-contents)


<a name="resend-withdrawal-callbacks"></a>
### Resend Withdrawal Callbacks

The callback handler can call this API to resend pending, risk-controlled or failed withdrawal callbacks.

Refer to [Callback Integration](#callback-integration) for callback rules.

> The resend operation could be requested on the web control panel as well.

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/sender/notifications/manual

> `WALLET_ID` must be a withdrawal wallet ID

- [Sample curl command](#curl-resend-withdrawal-callbacks)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/55743/sender/notifications/manual
```

###### Post body

```json
{
  "notification_id": 0
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :---  | :---        |
| notification_id | int64 | required, 0 means all | Specify callback ID to resend |

> This ID equal to the serial field of callback data.

##### Response Format

An example of a successful response:

```json
{
  "count": 3
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| count | int | Count of callbacks just resent |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | Malformatted post body |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-withdrawal-whitelist-configuration"></a>
### Query Withdrawal Whitelist Configuration

Query the whitelist configuration of the withdrawal wallet.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/sender/whitelist/config

> `WALLET_ID` must be a withdrawal wallet ID

- [Sample curl command](#curl-query-withdrawal-whitelist-configuration)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/55743/sender/whitelist/config
```

##### Response Format

An example of a successful response:

```json
{
  "effective_latency": 0,
  "whitelist_check": true
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| effective_latency | int64 | The effective latency of the whitelist entry, 0 means the whitelist entry will take effect immediately. |
| whitelist_check | boolean | Indicate whether the withdrawal wallet has enabled whitelist checking. |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="add-withdrawal-whitelist-entry"></a>
### Add Withdrawal Whitelist Entry

Add an outgoing address to the withdrawal wallet's whitelist.

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/sender/whitelist

> `WALLET_ID` must be a withdrawal wallet ID

- [Sample curl command](#curl-add-withdrawal-whitelist-entry)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/55743/sender/whitelist
```

###### Post body

```json
{
  "items": [
    {
      "address": "GCIFMEYIEWSX3K6EOPMEJ3FHW5AAPD6NW76J7LPBRAKD4JZKTISKUPHJ",
      "memo": "865314",
      "user_id": "USER001"
    }
  ]
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| items | array | required | Specify the whitelist entries |
| address | string | required | The outgoing address |
| memo | string | optional | The memo of the outgoing address |
| user_id | string | optional, max length `255` | The custom user ID of the outgoing address |

##### Response Format

An example of a successful response:

```json
{
  "added_items": [
    {
      "address": "GCIFMEYIEWSX3K6EOPMEJ3FHW5AAPD6NW76J7LPBRAKD4JZKTISKUPHJ",
      "memo": "865314",
      "user_id": "USER001"
    }
  ]
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| added_items | array | Array of added whitelist entries |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | Malformatted post body |
| 400 | 703 | Operation failed | invalid address: {INVALID_ADDRESS} | The address format does not comply with the cryptocurrency specification |
| 400 | 703 | Operation failed | invalid user id: {INVALID_USER_ID} | The length of the user ID exceeds 255 characters |
| 400 | 703 | Operation failed | this wallet does not support memo | The cryptocurrency does not support memo |
| 400 | 945 | The max length of BNB memo is 256 chars | - | Reached the limit of the length of BNB memo |
| 400 | 946 | The max length of EOS memo is 128 chars | - | Reached the limit of the length of EOS memo |
| 400 | 947 | The max length of XRP destination tag is 20 chars | - | Reached the limit of the length of XRP destination tag |
| 400 | 948 | The max length of XLM memo is 20 chars | - | Reached the limit of the length of XLM memo |
| 400 | 818 | Destination Tag must be integer | - | Wrong XRP destination tag format |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="remove-withdrawal-whitelist-entry"></a>
### Remove Withdrawal Whitelist Entry

Remove an outgoing address from the withdrawal wallet's whitelist.

##### Request

**DELETE** /v1/sofa/wallets/`WALLET_ID`/sender/whitelist

> `WALLET_ID` must be a withdrawal wallet ID

> Only the entries exactly matches all the fields will be removed.

- [Sample curl command](#curl-remove-withdrawal-whitelist-entry)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/55743/sender/whitelist
```

###### Post body

```json
{
  "items": [
    {
      "address": "GCIFMEYIEWSX3K6EOPMEJ3FHW5AAPD6NW76J7LPBRAKD4JZKTISKUPHJ",
      "memo": "865314",
      "user_id": "USER001"
    }
  ]
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| items | array | required | Specify the whitelist entries |
| address | string | required | The outgoing address |
| memo | string | optional | The memo of the outgoing address |
| user_id | string | optional | The custom user ID of the outgoing address |

##### Response Format

An example of a successful response:

```json
{
  "removed_items": [
    {
      "address": "GCIFMEYIEWSX3K6EOPMEJ3FHW5AAPD6NW76J7LPBRAKD4JZKTISKUPHJ",
      "memo": "865314",
      "user_id": "USER001"
    }
  ]
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| removed_items | array | Array of removed whitelist entries |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | Malformatted post body |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="check-withdrawal-whitelist"></a>
### Check Withdrawal Whitelist

Check the withdrawal whitelist entry status in the withdrawal whitelist.

##### Request

`VIEW` **POST** /v1/sofa/wallets/`WALLET_ID`/sender/whitelist/check

> `WALLET_ID` must be a withdrawal wallet ID

- [Sample curl command](#curl-check-withdrawal-whitelist)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/557432/sender/whitelist/check
```

###### Post body

```json
{
  "address": "GCIFMEYIEWSX3K6EOPMEJ3FHW5AAPD6NW76J7LPBRAKD4JZKTISKUPHJ",
  "memo": "865314",
  "user_id": "USER001"
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| address | string | required | The inquiry whitelist entry address |
| memo | string | optional | The memo of the whitelist entry |
| user_id | string | optional | The custom user ID of the whitelist entry |

##### Response Format

An example of a successful response:

```json
{
  "address": "0x79D6660b2aB1d37AD5D11C2ca2B3EBba7Efd13F6",
  "create_time": "2020-12-30T13:09:39Z",
  "effective": true,
  "effective_time": "2020-12-30T13:09:39Z",
  "memo": "",
  "state": 1,
  "update_time": "2020-12-30T13:09:39Z",
  "user_id": "USER001"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| address | string | The inquiry whitelist entry address |
| create_time | string | The creation time in UTC |
| effective | boolean | Indicate whether the whitelist entry has taken effect |
| effective_time | string | The effective time in UTC |
| memo | string | The memo of the whitelist entry |
| state | int | `1` means the entry is active, `2` means the entry is removed |
| update_time | string | Last modification time in UTC |
| user_id | string | The custom user ID of the whitelist entry |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | Malformatted post body |
| 400 | 703 | Operation failed | not found | Cannot find the inquiry whitelist entry |
| 400 | 703 | Operation failed | invalid address: {INVALID_ADDRESS} | The address format does not comply with the cryptocurrency specification |
| 400 | 703 | Operation failed | invalid user id: {INVALID_USER_ID} | The length of the user ID exceeds 255 characters |
| 400 | 703 | Operation failed | this wallet does not support memo | The cryptocurrency does not support memo |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-withdrawal-whitelist"></a>
### Query Withdrawal Whitelist

Used to query some kind of callbacks within a time interval.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/sender/whitelist?from\_time=`from`&to\_time=`to`&start\_index=`offset`&request_number=`count`&state=`state`

- [Sample curl command](#curl-query-withdrawal-whitelist)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/557432/sender/whitelist
```

The request includes the following parameters:

###### Query Parameters

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| from_time | int64 | optional, default `0` | Start date (unix time in UTC) |
| to_time | int64 | optional, default `current time` | End date (unix time in UTC) |
| start_index | int64 | optional, default `0` | The offset to the first entry |
| request_number | int64 | optional, default `1000`, max `2000` | The count to request |
| state | int | optional, default `-1` | Use `1` to query the active entries and `2` to query the removed entries, `-1` means all entries |

##### Response Format

An example of a successful response:

```json
{
  "items": [
    {
      "address": "GCIFMEYIEWSX3K6EOPMEJ3FHW5AAPD6NW76J7LPBRAKD4JZKTISKUPHJ",
      "create_time": "2020-12-30T06:02:25Z",
      "effective": false,
      "effective_time": "2020-12-30T13:27:37Z",
      "memo": "",
      "state": 1,
      "update_time": "2020-12-30T06:02:25Z",
      "user_id": "USER001"
    },
  ],
  "total_count": 1
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| items | array | Arrary of the whitelist entries |
| address | string | The whitelist entry address |
| create_time | string | The creation time in UTC |
| effective | boolean | Indicate whether the whitelist entry has taken effect |
| effective_time | string | The effective time in UTC |
| memo | string | The memo of the whitelist entry |
| state | int | `1` means the entry is active, `2` means the entry is removed |
| update_time | string | Last modification time in UTC |
| user_id | string | The custom user ID of the whitelist entry |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-withdrawal-wallet-transaction-history"></a>
### Query Withdrawal Wallet Transaction History

Get transaction history of withdrawal wallets.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/sender/transactions?from\_time=`from`&to\_time=`to`&start\_index=`start`&request_number=`count`

> `WALLET_ID` should be a withdrawal wallet ID

- [Sample curl command](#curl-query-withdrawal-wallet-transaction-history)

##### Request Format

An example of the request:

###### API with parameters

```
/v1/sofa/wallets/345312/sender/transactions?from_time=1559664000&to_time=1562255999&start_index=0&request_number=10
```

The request includes the following parameters:

###### Query Parameters

| Field | Type | Note | Description |
| :---  | :--- | :--- | :---        |
| from_time | int64 | optional, default `0` | Start date (unix time in UTC) |
| to_time | int64 | optional, default `current time` | End date (unix time in UTC) |
| start_index | int | optional, default `0` | Index of starting transaction record |
| request_number | int | optional, default `10`, max `500` | Count of returning transaction record |

##### Response Format

An example of a successful response:

```json
{
  "total_count": 169,
  "transactions": [
    {
      "amount": "0.1",
      "block_height": 10813730,
      "block_time": "2021-08-11T06:13:01Z",
      "blocklist_tags": [],
      "fee": "0.000693",
      "from_address": "0xaa0cA2f9bA3A33a915a27e289C9719adB2ad7d73",
      "memo": "",
      "out": true,
      "source": "",
      "state": 1,
      "to_address": "0x79D6660b2aB1d37AD5D11C2ca2B3EBba7Efd13F6",
      "txid": "0xe3607325e3b7c0190089d1fb41ce9fa059858c6b2e5dd220e55ba46707fc38f0"
    },
    {
      "amount": "1",
      "block_height": 10811102,
      "block_time": "2021-08-10T17:24:21Z",
      "blocklist_tags": [],
      "fee": "0.000021",
      "from_address": "0xaa0cA2f9bA3A33a915a27e289C9719adB2ad7d73",
      "memo": "",
      "out": true,
      "source": "withdraw-api",
      "state": 1,
      "to_address": "0x8382Cc1B05649AfBe179e341179fa869C2A9862b",
      "txid": "0x19657382aa16520c32eef0dacc0f16d78e9105e83d37d126b4f6687c0d651859"
    },
  ]
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| total_count | int | Total transactions in specified date duration |
| transactions | array | Array of transaction record |
| txid | string | Transaction ID |
| from_address | string | Sender address of the transaction |
| to_address | string | Recipient address of the transaction |
| out | boolean | True means outgoing transaction |
| amount | string | Transaction amount |
| blocklist_tags | array | The tags of CYBAVO AML detection. If `out` is true, the `to_address` is tagged. Otherwise, the `from_address` is tagged |
| block_height | int64 | The block height |
| block_time | time | When was the transaction packed into block (in chain) in UTC time |
| fee | string | Transaction blockchain fee |
| memo | string | Memo of the transaction |
| source | string | `withdraw-api` means that the transaction was triggered by the withdrawal API, otherwise it was triggered from the web withdrawal UI |
| state | int | Refer to [Transaction State Definition](#api-transaction-state-filter) bellow |

<a name="api-transaction-state-filter"></a>
###### Transaction State Definition

| ID   | Description |
| :--- | :---        |
| 1 | Success, the transaction status is successful |
| 2 | Failed, the transaction status is failed |
| 3 | Invalid, the transaction status is successful but is identified as invalid by the SOFA system |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | `from_time` or `to_time` is invalid |

##### [Back to top](#table-of-contents)


# Deposit / Withdrawal Wallet Common API

<a name="query-callback-history"></a>
### Query Callback History

Used to query some kind of callbacks within a time interval.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/notifications?from\_time=`from`&to\_time=`to`&type=`type`

- [Sample curl command](#curl-query-callback-history)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/677414/notifications?from_time=1561651200&to_time=1562255999&type=2
```

The request includes the following parameters:

###### Query Parameters

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| from_time | int64 | optional, default `0` | Start date (unix time in UTC) |
| to_time | int64 | optional, default `current time` | End date (unix time in UTC) |
| type | int | optional, default `-1` | Refer to [Callback Type](#callback-type-definition) |

##### Response Format

An example of a successful response:

```json
{
  "notifications": [
    {
      "type": 2,
      "serial": 90000000003,
      "order_id": "a206",
      "currency": "BNB",
      "txid": "76B8B2B1E25472FFE7B8FCE85742E0964FEDB1B679DE963FA19F430E8B287F93",
      "block_height": 25844472,
      "tindex": 2,
      "vout_index": 0,
      "amount": "15000000",
      "fees": "37500",
      "memo": "CC",
      "broadcast_at": 0,
      "chain_at": 1562234190,
      "from_address": "tbnb1f805kv6z8nq2whrcnkagjte3jjss2sxf2rfls0",
      "to_address": "tbnb1655kasahedvaeudaeq6jggr7kal8qgwygu9xqk",
      "wallet_id": 677414,
      "state": 3,
      "confirm_blocks": 1,
      "processing_state": 1,
      "addon": {}
    }
  ]
}

```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| notifications | array | Arrary of callbacks, refer to [Callback Definition](#callback-definition) |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |

##### [Back to top](#table-of-contents)


<a name="query-callback-detail"></a>
### Query Callback Detail

Query the detailed information of the callback by its serial ID. It can be used to reconfirm whether a deposit callback exists.

##### Request

`VIEW` **POST** /v1/sofa/wallets/`WALLET_ID`/notifications/get\_by_id

- [Sample curl command](#curl-query-callback-detail)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/677414/notifications/get_by_id
```

###### Post body

```json
{
  "ids": [
    90000000140,
    90000000139    
  ]
}
```

The request includes the following parameters:

###### Post body

| Field | Type | Note | Description |
| :---  | :--- | :--- | :---        |
| ids | array | requried | Specify the IDs for query |


##### Response Format

An example of a successful response:

```json
{
  "notifications": [
    {
      "type": 3,
      "serial": 90000000139,
      "order_id": "",
      "currency": "ADA",
      "txid": "35c283a6f13f5886240fe2e815bc149154ec066cd2061202318dd4e4bf8af35e",
      "block_height": 1003304,
      "tindex": 0,
      "vout_index": 0,
      "amount": "24447",
      "fees": "0",
      "memo": "",
      "broadcast_at": 1584088556,
      "chain_at": 1584088556,
      "from_address": "",
      "to_address": "37btjrVyb4KG8gKeZjJguinwdsbcRV65ngHhBUaJWf36QxiakTV3UHiNUP9arReXMZQnpRBVVdkcBB4GyiWzPRSTmg41mTzMpxgfhtfRHtaBCKJNbX",
      "wallet_id": 677414,
      "state": 3,
      "confirm_blocks": 2,
      "processing_state": 1,
      "addon": {}
    },
    {
      "type": 3,
      "serial": 90000000140,
      "order_id": "",
      "currency": "ADA",
      "txid": "fa120b6283509f0ab2b136a3ac8b613aa3ca2f36ce7c2744e122668d013cfdb5",
      "block_height": 1003305,
      "tindex": 0,
      "vout_index": 0,
      "amount": "55497180",
      "fees": "0",
      "memo": "",
      "broadcast_at": 1584088576,
      "chain_at": 1584088576,
      "from_address": "",
      "to_address": "37btjrVyb4KDKCyAPRUPxpGiUPWunpBAkGRX8U3h7LYzS2UrHUnEQozcCyqR2GfBVnM3frTaUNEb8DoNGo9JakrskAtaWt6vED6R6ohkmaJ2qr4oCg",
      "wallet_id": 677414,
      "state": 3,
      "confirm_blocks": 1,
      "processing_state": 1,
      "addon": {}
    }
  ]
}

```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| notifications | array | Arrary of callbacks, refer to [Callback Definition](#callback-definition) |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |

##### [Back to top](#table-of-contents)


<a name="query-wallet-synchronization-info"></a>
### Query Wallet Synchronization Info

Get the blockchain synchronization status of a wallet.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/blocks

- [Sample curl command](#curl-query-wallet-synchronization-info)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/677414/blocks
```

##### Response Format

An example of a successful response:

```json
{
  "latest_block_height": 29317651,
  "synced_block_height": 28529203
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| latest\_block_height | int64 | The latest block height on blockchain |
| synced\_block_height | int64 | The current synced block height |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |

##### [Back to top](#table-of-contents)


<a name="query-transaction-average-fee"></a>
### Query Transaction Average Fee

Query average blockchain fee within latest N blocks.

##### Request

`VIEW` **POST** /v1/sofa/wallets/`WALLET_ID`/autofee

- [Sample curl command](#curl-query-transaction-average-fee)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/635718/autofee
```

###### Post body

```json
{
  "block_num": 1
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :---     | :---        |
| block_num | int | optional, default `1`, range `1~30` | Query the average blockchain fee in the last N blocks |

##### Response Format

An example of a successful response:
	
```json
{
	"auto_fee": "1"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| auto_fee | string | Mining fee denominated in the smallest cryptocurrency unit |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | The `block_num` is out of range |

##### [Back to top](#table-of-contents)


<a name="batch-query-transaction-average-fees"></a>
### Batch Query Transaction Average Fees

Batch Query average blockchain fee within latest N blocks.

##### Request

`VIEW` **POST** /v1/sofa/wallets/`WALLET_ID`/autofees

- [Sample curl command](#curl-batch-query-transaction-average-fees)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/854714/autofees
```

###### Post body

```json
{
  "block_nums": [
    1,
    5,
    10
  ]
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :---     | :---        |
| block_nums | array | required, max 5 entries, each entry is range `1~30` | Batch query the average blockchain fee in the last N blocks |

##### Response Format

An example of a successful response:
	
```json
{
  "auto_fees": [
    {
      "auto_fee": "49000000000",
      "block_num": 1
    },
    {
      "auto_fee": "49000000000",
      "block_num": 5
    },
    {
      "auto_fee": "38000000000",
      "block_num": 10
    }
  ]
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| auto_fees | array | Result of the inquiry |
| block_num | int | Inquiry block number |
| auto_fee | string | Mining fee denominated in the smallest cryptocurrency unit |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | Exceeds 5 entries or each entry not in range |

##### [Back to top](#table-of-contents)


# Vault Wallet API

<a name="query-vault-wallet-transaction-history"></a>
### Query Vault Wallet Transaction History

Get transaction history of vault wallets.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/transactions?from\_time=`from`&to\_time=`to`&start\_index=`start`&request_number=`count`&state=`state`

> `WALLET_ID` should be a vault wallet ID

- [Sample curl command](#curl-query-vault-wallet-transaction-history)

##### Request Format

An example of the request:

###### API with parameters

```
/v1/sofa/wallets/488471/transactions?from_time=1559664000&to_time=1562255999&start_index=0&request_number=1
```

The request includes the following parameters:

###### Query Parameters

| Field | Type | Note | Description |
| :---  | :--- | :--- | :---        |
| from_time | int64 | optional, default `0` | Start date (unix time in UTC) |
| to_time | int64 | optional, default `current time` | End date (unix time in UTC) |
| start_index | int | optional, default `0` | Index of starting transaction record |
| request_number | int | optional, default `1000`, max `5000` | Count of returning transaction record |
| state | int | optional, default `-1` | Refer to [Transaction State Filter Definition](#transaction-state-filter) bellow |

<a name="transaction-state-filter"></a>
###### Transaction State Filter Definition

| ID   | Description |
| :--- | :---        |
| 0 | WaitApproval |
| 1 | Rejected |
| 2 | Approved |
| 3 | Failed |
| 4 | NextLevel |
| 5 | Cancelled |
| 6 | BatchDone |
| -1 | All states (for inquiry) |

##### Response Format

An example of a successful response:

```json
{
  "transaction_count": 3,
  "transaction_item": [
    {
      "issue_user_id": 3,
      "issue_user_name": "wallet owner (user@gmail.com)",
      "description": "TO SND",
      "wallet_id": 488471,
      "wallet_name": "BNB I",
      "wallet_address": "tbnb1655kasahedvaeudaeq6jggr7kal8qgwygu9xqk",
      "token_address": "",
      "txid": "3E6D61D1D39FA5DD3B86C2C28FFAD3D98CD7AFDB62346468D3C4FFE710CAAF85",
      "currency": 714,
      "currency_name": "BNB",
      "outgoing_address": "tbnb1f805kv6z8nq2whrcnkagjte3jjss2sxf2rfls0",
      "outgoing_address_name": "BNB SND",
      "amount": "2",
      "fee": "0",
      "txno": 100087,
      "approval_item": [
        {
          "approval_id": 3,
          "approval_user": "wallet owner (user@gmail.com)",
          "approval_time": 1562210142,
          "user_message": "",
          "level": 0,
          "owner": 1,
          "confirm": 1,
          "state": 2,
          "error_code": 0
        }
      ],
      "state": 2,
      "create_time": 1562210129,
      "transaction_time": 1562210142,
      "scheduled_name": "",
      "transaction_type": 0,
      "eos_transaction_type": 0,
      "real_amount": "2",
      "chain_fee": "0.000375",
      "platform_fee": "0",
      "tx_category": "",
      "memo": "TO SND"
    }
  ]
}

```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| transaction_count | int | Total transactions in specified date duration |
| transaction_item | array | Array of transaction record |

> Refer to [Currency Definition](#currency-definition) or [here](https://github.com/satoshilabs/slips/blob/master/slip-0044.md) for more detailed currency definitions

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |

##### [Back to top](#table-of-contents)


<a name="query-vault-wallet-balance"></a>
### Query Vault Wallet Balance

Get the vault wallet balance. Facilitate to establish a real-time balance monitoring mechanism.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/vault/balance

> `WALLET_ID` must be a vault wallet ID

- [Sample curl command](#curl-query-vault-wallet-balance)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/488471/vault/balance
```

##### Response Format

An example of a successful response:

BTC vault wallet

```json
{
  "balance": "0.00009798",
  "currency": 0,
  "token_address": "",
  "token_balance": "",
  "unconfirm_balance": "0",
  "unconfirm_token_balance": "",
  "wallet_address": "2Mw1iJnQvAt3hNEvEZKdHkij8TNtzjaF3LH"
}
```

USDT-Omni vault wallet that mapping to above BTC vault wallet

```json
{
  "balance": "0.00009798",
  "currency": 0,
  "token_address": "31",
  "token_balance": "0.1",
  "unconfirm_balance": "0",
  "unconfirm_token_balance": "",
  "wallet_address": "2Mw1iJnQvAt3hNEvEZKdHkij8TNtzjaF3LH"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| currency | int64 | Registered coin types. Refer to [Currency Definition](#currency-definition) |
| wallet_address | string | Wallet address |
| token_address | string | Token contract address |
| balance | string | Withdrawal wallet balance |
| token_balance | string | Withdrawal wallet token balance |
| unconfirm\_balance | string | Unconfirmed withdrawal wallet balance |
| unconfirm\_token_balance | string | Unconfirmed withdrawal wallet token balance |

> The currencies that support the unconfirmed balance are BTC, LTC, ETH, BCH, BSV, DASH

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |

##### [Back to top](#table-of-contents)


# Common API

<a name="activate-api-code"></a>
### Activate API Code

Activate the API code of a certain wallet. Once activated, the currently activated API code will immediately become invalid.

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/apisecret/activate

- [Sample curl command](#curl-activate-api-code)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/488471/apisecret/activate
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
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |

##### [Back to top](#table-of-contents)


<a name="query-api-code-status"></a>
### Query API Code Status

Query the API code info of a certain wallet. Use the `inactivated` API code in any request will activate it. Once activated, the currently activated API code will immediately become invalid.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/apisecret

- [Sample curl command](#curl-query-api-code-status)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/488471/apisecret
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
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |

##### [Back to top](#table-of-contents)


<a name="refresh-api-code"></a>
### Refresh API Code

Use paired refresh code to acquire the new inactive API code/secret of the wallet.

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/refreshsecret

- [Sample curl command](#curl-refresh-api-code)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/357818/refreshsecret
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
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | Malformatted post body or the refresh code is invalid |


##### [Back to top](#table-of-contents)


<a name="query-wallet-info"></a>
### Query Wallet Info

Get wallet basic information.

##### Request

`VIEW` **GET** /v1/sofa/wallets/`WALLET_ID`/info

- [Sample curl command](#curl-query-wallet-info)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/488471/info
```

##### Response Format

An example of a successful response:

```json
{
  "currency": 60,
  "currency_name": "ETH",
  "address": "0xd11Bd6E308b8DC1c5243D54cf41A427Ca0F46943",
  "token_name": "TTF TOKEN",
  "token_symbol": "TTF",
  "token_contract_address": "0xd0ee17a4e1866c1ac53a54cc2cd4dd64b503cf40",
  "token_decimals": "18"
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| currency | int64 | Registered coin types. Refer to [Currency Definition](#currency-definition) |
| currency_name | string | Name of currency |
| address | string | Wallet address |

> Refer to [here](https://github.com/satoshilabs/slips/blob/master/slip-0044.md) for more detailed currency definitions

If `WALLET_ID` is a token wallet, the following fields present:

| Field | Type  | Description |
| :---  | :---  | :---        |
| token_name | string | Token name |
| token_symbol | string | Token symbol |
| token\_contract_address | string | Token contract address |
| token_decimals | string | Token decimals |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |

##### [Back to top](#table-of-contents)


<a name="verify-addresses"></a>
### Verify Addresses

Check if the address conforms to the wallet cryptocurrency address format (for example, ETH must have the prefix 0x, BTC should start with 1, 3 or bc1, etc).

> If the wallet's cryptocurrency is BNB or XRP, there will be a `must_need_memo` flag to indicate whether the address needs a memo / destination tag when transferring cryptocurrency to that address.

##### Request

`VIEW` **POST** /v1/sofa/wallets/`WALLET_ID`/addresses/verify

> Wallet ID must be a deposit or withdraw wallet's ID

- [Sample curl command](#curl-verify-addresses)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/488471/addresses/verify
```

###### Post body

```json
{
  "addresses": [
    "0x635B4764D1939DfAcD3a8014726159abC277BecC",
    "1CK6KHY6MHgYvmRQ4PAafKYDrg1ejbH1cE"
  ]
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| addresses | array | requried | Specify the address for verification |

##### Response Format

An example of a successful response:
	
```json
{
  "result": [
    {
      "address": "0x635B4764D1939DfAcD3a8014726159abC277BecC",
      "valid": true,
      "must_need_memo": false
    },
    {
      "address": "1CK6KHY6MHgYvmRQ4PAafKYDrg1ejbH1cE",
      "valid": false,
      "must_need_memo": false
    }
  ]
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| result | array | Array of addresses' verification result |
| must\_need\_memo | boolean | Indicate whether the address needs a memo / destination tag when transferring cryptocurrency to that address |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | Malformatted post body |

##### [Back to top](#table-of-contents)


<a name="inspect-callback-endpoint"></a>
### Inspect Callback Endpoint

Use to inspect the notification and withdrawal authentication endpoint.

##### Request

`VIEW` **POST** /v1/sofa/wallets/`WALLET_ID`/notifications/inspect

- [Sample curl command](#curl-inspect-callback-endpoint)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/896541/notifications/inspect
```

###### Post body

```json
{
  "test_number": 102999
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| test_number | int64 | requried | The test number will be sent to the notification endpoint in the format `{"msg":"CONNECTION TEST","server_time":1622195270,"client_ip":"xxx.xxx.xxx.xxx","test_number":102999}`. |

##### Response Format

An example of a successful response:
	
```json
{
  "server_time": 1622195253,
  "client_ip": "::1",
  "notification_endpoint": {
    "url": "http%3A%2F%2Flocalhost%3A8889%2Fv1%2Fmock%2Fwallets%2Fcallback",
    "status_code": 400,
    "response": "NDAw"
  },
  "withdrawal_authentication_endpoint": {
    "error": "no endpoint found"
  }
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| server_time | int64 | Current server unix time in UTC |
| client_ip | string | The request client IP |
| notification\_endpoint | struct | Specify the test result of notification endpoint |
| withdrawal\_authentication_endpoint | struct | Specify the test result of withdrawal authentication endpoint|
| url | string | The escaped endpoint URL |
| status_code | int | The HTTP response status code from endpoint |
| response | string | The base64 encoded response from endpoint |
| error | string | Specify the connection error if any |


##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |
| 400 | 112 | Invalid parameter | - | Malformatted post body |

##### [Back to top](#table-of-contents)


# Read-only API code API

<a name="list-wallets"></a>
### List Wallets

List all wallets can be accessed by the inquiry read-only API code.

##### Request

`VIEW` **GET** /v1/sofa/wallets/readonly/walletlist

> The API code must be a read-only API code.

- [Sample curl command](#curl-list-wallets)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/readonly/walletlist
```

##### Response Format

An example of a successful response:

```json
{
  "total": 2,
  "wallets": [
    {
      "address": "2NAnSkEp6SpUPLsdP3ChvN6K5qPMZyoB3RF",
      "currency": 0,
      "currency_name": "BTC",
      "decimals": "8",
      "type": 2,
      "wallet_id": 101645
    },
    {
      "address": "0x85AfD8F88C0347aFF89AFc6C0749322719396616",
      "currency": 60,
      "currency_name": "ETH",
      "decimals": "18",
      "token_contract_address": "0xdf2ce4af00b10644d00316b3d99e029d82d5d2f3",
      "token_decimals": "18",
      "token_name": "JGB2",
      "token_symbol": "JGB2",
      "type": 0,
      "wallet_id": 118970
    }
  ]
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| address | string | Wallet address |
| currency | int64 | Registered coin types. Refer to [Currency Definition](#currency-definition) |
| currency_name | string | Name of currency |
| decimals | string | Decimals of currency |
| type | int | Wallet Type. Refer to [Wallet Type Definition](#wallet-type-definition)|
| wallet_id | int64 | Wallet ID |
| token_name | string | Token name |
| token_symbol | string | Token symbol |
| token\_contract_address | string | Token contract address |
| token_decimals | string | Token decimals |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 385   | API Secret not valid | - | Invalid API code permission |

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

### Put wallet API code/secret into mock server
-	Get API code/secret on web control panel
	-	API_CODE, API\_SECRET, WALLET\_ID
- 	Put API code/secret to mock server's databaså

```
curl -X POST -H "Content-Type: application/json" -d '{"api_code":"API_CODE","api_secret":"API_SECRET"}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/apitoken
```

### Register mock server callback URL
>	Operate on web control panel

Notification Callback URL

```
http://localhost:8889/v1/mock/wallets/callback
```

Withdrawal Authentication Callback URL

```
http://localhost:8889/v1/mock/wallets/withdrawal/callback
```

> The withdrawal authentication callback URL once set, every withdrawal request will callback this URL to get authentication to proceed withdrawal request.
> 
> Refer to [WithdrawalCallback()](https://github.com/CYBAVO/SOFA_MOCK_SERVER/blob/master/controllers/OuterController.go#L183) function in mock server OuterController.go

##### [Back to top](#table-of-contents)

<a name="curl-testing-commands"></a>
# cURL Testing Commands

<a name="curl-create-deposit-addresses"></a>
### Create Deposit Addresses

For BNB, XLM, XRP or EOS wallet:

```
curl -X POST -H "Content-Type: application/json" -d '{"count":2,"memos":["10001","10002"]}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/addresses
```

For wallet excepts BNB, XLM, XRP and EOS:

```
curl -X POST -H "Content-Type: application/json" -d '{"count":2}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/addresses
```
- [API definition](#create-deposit-addresses)

<a name="curl-query-deposit-addresses"></a>
### Query Deposit Addresses

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/addresses?start_index=0&request_number=1000
```
- [API definition](#query-deposit-addresses)

<a name="curl-query-deployed-contract-deposit-addresses"></a>
### Query Deployed Contract Deposit Addresses

```
curl 'http://localhost:8889/v1/mock/wallets/{WALLET_ID}/addresses/contract_txid?txids={TXID1},{TXID2}'
```
- [API definition](#query-deployed-contract-deposit-addresses)

<a name="curl-query-pool-address"></a>
### Query Pool Address

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/pooladdress
```
- [API definition](#query-pool-address)


<a name="curl-query-pool-address-balance"></a>
### Query Pool Address Balance

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/pooladdress/balance
```
- [API definition](#query-pool-address-balance)


<a name="curl-query-invalid-deposit-addresses"></a>
### Query Invalid Deposit Addresses

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/addresses/invalid-deposit
```
- [API definition](#query-invalid-deposit-addresses)


<a name="curl-query-deposit-callback-detail"></a>
### Query Deposit Callback Detail

```
curl 'http://localhost:8889/v1/mock/wallets/{WALLET_ID}/receiver/notifications/txid/{TX_ID}/{VOUT_INDEX}'
```
- [API definition](#query-deposit-callback-detail)


<a name="curl-resend-deposit-callbacks"></a>
### Resend Deposit/Collection Callbacks

```
curl -X POST -H "Content-Type: application/json" -d '{"notification_id":0}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/collection/notifications/manual
```
- [API definition](#resend-deposit-callbacks)


<a name="curl-query-deposit-wallet-balance"></a>
### Query Deposit Wallet Balance

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/receiver/balance
```
- [API definition](#query-deposit-wallet-balance)


<a name="curl-update-deposit-address-label"></a>
### Update Deposit Address Label

```
curl -X POST -H "Content-Type: application/json" -d '{"address":"0x2B974a3e0b491bB26e0bF146E6cDaC36EFD574a","label":"take-some-notes"}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/addresses/label
```
- [API definition](#update-deposit-address-label)


<a name="curl-query-deposit-address-label"></a>
### Query Deposit Address Label

```
curl -X POST -H "Content-Type: application/json" -d '{"addresses":["0x2B974a3De0b491bB26e0bF146E6cDaC36EFD574a","0xF401AC94D9672e79c68e56A6f822b666E5A7d644"]}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/addresses/get_labels
```
- [API definition](#query-deposit-address-label)


<a name="curl-withdraw-assets"></a>
### Withdraw Assets

```
curl -X POST -H "Content-Type: application/json" -d '{"requests":[{"order_id":"888888_1","address":"0x60589A749AAC632e9A830c8aBE042D1899d8Dd15","amount":"0.0001","memo":"memo-001","user_id":"USER01","message":"message-001"},{"order_id":"888888_2","address":"0xf16B7B8900F0d2f682e0FFe207a553F52B6C7015","amount":"0.0002","memo":"memo-002","user_id":"USER01","message":"message-002"}]}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/transactions
```
- [API definition](#withdraw-assets)


<a name="curl-cancel-withdrawal-request"></a>
### Cancel Withdrawal Request

```
curl -X POST http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/transactions/{ORDER_ID}/cancel
```
- [API definition](#cancel-withdrawal-request)


<a name="curl-query-latest-withdrawal-transaction-state"></a>
### Query Latest Withdrawal Transaction State

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/transactions/{ORDER_ID}
```
- [API definition](#query-latest-withdrawal-transaction-state)


<a name="curl-query-all-withdrawal-transaction-states"></a>
### Query All Withdrawal Transaction States

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/transactions/{ORDER_ID}/all
```
- [API definition](#query-all-withdrawal-transaction-states)


<a name="curl-query-withdrawal-wallet-balance"></a>
### Query Withdrawal Wallet Balance

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/balance
```
- [API definition](#query-withdrawal-wallet-balance)


<a name="curl-query-withdrawal-callback-detail"></a>
### Query Withdrawal Callback Detail

```
curl 'http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/notifications/order_id/{ORDER_ID}'
```
- [API definition](#query-withdrawal-callback-detail)


<a name="curl-set-withdrawal-request-acl"></a>
### Set Withdrawal Request ACL

```
curl -X POST -H "Content-Type: application/json" -d '{"acl":"192.168.101.55"}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/transactions/acl
```
- [API definition](#set-withdrawal-request-acl)


<a name="curl-query-withdrawal-whitelist-configuration"></a>
### Query Withdrawal Whitelist Configuration

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/whitelist/config
```
- [API definition](#query-withdrawal-whitelist-configuration)


<a name="curl-query-withdrawal-wallet-transaction-history"></a>
### Query Withdrawal Wallet Transaction History

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/transactions
```
- [API definition](#query-withdrawal-wallet-transaction-history)


<a name="curl-add-withdrawal-whitelist-entry"></a>
### Add Withdrawal Whitelist Entry

```
curl -X POST -H "Content-Type: application/json" -d '{"items":[{"address":"GCIFMEYIEWSX3K6EOPMEJ3FHW5AAPD6NW76J7LPBRAKD4JZKTISKUPHJ","memo":"85666","user_id":"USER002"}]}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/whitelist
```
- [API definition](#add-withdrawal-whitelist-entry)


<a name="curl-remove-withdrawal-whitelist-entry"></a>
### Remove Withdrawal Whitelist Entry

```
curl -X DELETE -H "Content-Type: application/json" -d '{"items":[{"address":"GCIFMEYIEWSX3K6EOPMEJ3FHW5AAPD6NW76J7LPBRAKD4JZKTISKUPHJ","memo":"85666","user_id":"USER002"}]}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/whitelist
```
- [API definition](#remove-withdrawal-whitelist-entry)


<a name="curl-check-withdrawal-whitelist"></a>
### Check Withdrawal Whitelist

```
curl -X POST -H "Content-Type: application/json" -d '{"address":"GCIFMEYIEWSX3K6EOPMEJ3FHW5AAPD6NW76J7LPBRAKD4JZKTISKUPHJ","memo":"85666","user_id":"USER002"}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/whitelist/check
```
- [API definition](#check-withdrawal-whitelist)


<a name="curl-query-withdrawal-whitelist"></a>
### Query Withdrawal Whitelist

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/whitelist
```
- [API definition](#query-withdrawal-whitelist)


<a name="curl-resend-withdrawal-callbacks"></a>
### Resend Withdrawal Callbacks

```
curl -X POST -H "Content-Type: application/json" -d '{"notification_id":0}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/sender/notifications/manual
```
- [API definition](#resend-withdrawal-callbacks)


<a name="curl-query-callback-history"></a>
### Query Callback History

```
curl 'http://localhost:8889/v1/mock/wallets/{WALLET_ID}/notifications?from_time=1561651200&to_time=1562255999&type=2'
```
- [API definition](#query-callback-history)


<a name="curl-query-callback-detail"></a>
### Query Callback Detail

```
curl -X POST -H "Content-Type: application/json" -d '{"ids":[90000000140,90000000139]}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/notifications/get_by_id
```
- [API definition](#query-callback-detail)


<a name="curl-query-wallet-synchronization-info"></a>
### Query Wallet Synchronization Info

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/blocks
```
- [API definition](#query-wallet-synchronization-info)


<a name="curl-query-transaction-average-fee"></a>
### Query Transaction Average Fee

```
curl -X POST -H "Content-Type: application/json" -d '{"block_num":1}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/autofee
```
- [API definition](#query-transaction-average-fee)


<a name="curl-batch-query-transaction-average-fees"></a>
### Batch Query Transaction Average Fees

```
curl -X POST -H "Content-Type: application/json" -d '{"block_nums":[1,5,10]}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/autofees
```
- [API definition](#batch-query-transaction-average-fees)


<a name="curl-query-vault-wallet-transaction-history"></a>
### Query Vault Wallet Transaction History

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/transactions?start_index=0&from_time=1559664000&to_time=1562255999&request_number=8
```
- [API definition](#query-vault-wallet-transaction-history)

<a name="curl-query-vault-wallet-balance"></a>
### Query Vault Wallet Balance

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/vault/balance
```
- [API definition](#query-vault-wallet-balance)


<a name="curl-activate-api-code"></a>
### Activate API Code

```
curl -X POST http://localhost:8889/v1/mock/wallets/{WALLET_ID}/apisecret/activate
```
- [API definition](#activate-api-code)


<a name="curl-query-api-code-status"></a>
### Query API Code Status

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/apisecret
```
- [API definition](#query-api-code-status)


<a name="curl-refresh-api-code"></a>
### Refresh API Code

```
curl -X POST -H "Content-Type: application/json" -d '{"refresh_code":"3EbaSPUpKzHJ9wYgYZqy6W4g43NT365bm9vtTfYhMPra"}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/refreshsecret
```
- [API definition](#query-api-code-status)


<a name="curl-query-wallet-info"></a>
### Query Wallet Info

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/info
```
- [API definition](#query-wallet-info)


<a name="curl-verify-addresses"></a>
### Verify Addresses

```
curl -X POST -H "Content-Type: application/json" -d '{"addresses":["0x635B4764D1939DfAcD3a8014726159abC277BecC","1CK6KHY6MHgYvmRQ4PAafKYDrg1ejbH1cE"]}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/addresses/verify
```
- [API definition](#verify-addresses)


<a name="curl-inspect-callback-endpoint"></a>
### Inspect Callback Endpoint

```
curl -X POST -H "Content-Type: application/json" -d '{"test_number":102999}' \
http://localhost:8889/v1/mock/wallets/896541/notifications/inspect
```
- [API definition](#inspect-callback-endpoint)


<a name="curl-list-wallets"></a>
### List Wallets

```
curl http://localhost:8889/v1/mock/wallets/readonly/walletlist
```
- [API definition](#list-wallets)

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
    <td>type</td>
    <td>int</td>
    <td>
   	 	<table>
   	 	  <thead><tr><td>ID</td><td>Type</td></tr></thead>
   	 	  <tbody>
		    <tr><td>1</td><td>Deposit</td></tr>
  		    <tr><td>2</td><td>Withdraw</td></tr>
  		    <tr><td>3</td><td>Collect</td></tr>
  		    <tr><td>4</td><td>Airdrop</td></tr>
   	 	  </tbody>
		</table>
    </td>
  </tr>
  <tr>
    <td>serial</td>
    <td>int</td>
    <td>The unique serial of callback</td>
  </tr>
  <tr>
    <td>order_id</td>
    <td>string</td>
    <td>The unique order ID of withdrawal request</td>
  </tr>
  <tr>
    <td>currency</td>
    <td>string</td>
    <td>Cryptocurrency of the callback<br>This field is for human reading only and may change in the future. Do not use this string as currency definition, use the fields <b>currency_bip44</b> and <b>token_address</b> as currency definition.</td>
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
    <td>tindex</td>
    <td>int</td>
    <td>The index of transaction in its block</td>
  </tr>
  <tr>
    <td>vout_index</td>
    <td>int</td>
    <td>The index of vout in its transaction</td>
  </tr>
  <tr>
    <td>amount</td>
    <td>string</td>
    <td>Transaction amount denominated in the smallest cryptocurrency unit</td>
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
    <td>chain_at</td>
    <td>int64</td>
    <td>When was the transaction packed into block (in chain) in UTC time</td>
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
    <td>wallet_id</td>
    <td>int64</td>
    <td>The wallet ID of the callback</td>
  </tr>
  <tr>
    <td>state</td>
    <td>int</td>
    <td>
      Possible states (listed in the Transaction State Definition table)
	 	<table>
	 	  <thead><tr><td>ID</td><td>Description</td></tr></thead>
	 	  <tbody>
		    <tr><td>1</td><td>Processing</td></tr>
		    <tr><td>2</td><td>TXID in pool</td></tr>
		    <tr><td>3</td><td>TXID in chain</td></tr>
		    <tr><td>5</td><td>Failed (the err_reason of addon field will contain detailed error reason)</td></tr>
		    <tr><td>8</td><td>Cancelled</td></tr>
		    <tr><td>10</td><td>Dropped</td></tr>
	 	  </tbody>
		</table>
    </td>
  </tr>
  <tr>
    <td>confirm_blocks</td>
    <td>int64</td>
    <td>Number of confirmations</td>
  </tr>
  <tr>
    <td>processing_state</td>
    <td>int</td>
    <td>
	 	<table>
	 	  <thead><tr><td>ID</td><td>Description</td></tr></thead>
	 	  <tbody>
		    <tr><td>-1</td><td>If the state is 5(failed), 8(cacelled), 10(dropped) or 11(transaction failed)</td></tr>
		    <tr><td>0</td><td>In fullnode mempool</td></tr>
		    <tr><td>1</td><td>In chain (the transaction is already on the blockchain but the confirmations have not been met)</td></tr>
		    <tr><td>2</td><td>Done (the transaction is already on the blockchain and satisfy confirmations)</td></tr>
	 	  </tbody>
		</table>
    </td>
  </tr>
  <tr>
    <td>addon</td>
    <td>key-value pairs</td>
    <td>
    The extra information of this callback
	 	<table>
	 	  <thead><tr><td>Key</td><td>Value (Description)</td></tr></thead>
	 	  <tbody>
		    <tr><td>err_reason</td><td>Will contain detail error reason if state is 5(Failed)</td></tr>
		    <tr><td>fee_decimal</td><td>The decimal of cryptocurrency miner fee</td></tr>
		    <tr><td>blocklist_tags</td><td>The tags of CYBAVO AML detection</td></tr>
		    <tr><td>address_label</td><td>The label of the deposit address</td></tr>
		    <tr><td>contract_abi</td><td>The contract ABI of the withdrawal request</td></tr>
		    <tr><td>token_id</td><td>Transferred token ID</td></tr>
		    <tr><td>aml_tags</td><td>Detailed CYBAVO AML detection information includes score, tags and blocked flag</td></tr>
		    <tr><td>aml_screen_pass</td><td>Pass or fail CYBAVO AML screening</td></tr>
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
    <td>currency_bip44</td>
    <td>int64</td>
    <td>
   	 	<table>
   	 	  <thead><tr><td>ID</td><td>Currency Symbol</td><td>Decimals</td></tr></thead>
   	 	  <tbody>
		    <tr><td>0</td><td>BTC</td><td>8</td></tr>
  		    <tr><td>2</td><td>LTC</td><td>8</td></tr>
  		    <tr><td>3</td><td>DOGE</td><td>8</td></tr>
  		    <tr><td>5</td><td>DASH</td><td>8</td></tr>
  		    <tr><td>60</td><td>ETH</td><td>18</td></tr>
  		    <tr><td>144</td><td>XRP</td><td>6</td></tr>
  		    <tr><td>145</td><td>BCH (BCHN)</td><td>8</td></tr>
  		    <tr><td>148</td><td>XLM</td><td>7</td></tr>
  		    <tr><td>194</td><td>EOS</td><td>4</td></tr>
   		    <tr><td>195</td><td>TRX</td><td>6</td></tr>
   		    <tr><td>236</td><td>BSV</td><td>8</td></tr>
   		    <tr><td>354</td><td>DOT</td><td>10</td></tr>
   		    <tr><td>461</td><td>FIL</td><td>18</td></tr>
   		    <tr><td>472</td><td>AR</td><td>12</td></tr>
   		    <tr><td>501</td><td>SOL</td><td>18</td></tr>
   		    <tr><td>539</td><td>FLOW</td><td>8</td></tr>
   		    <tr><td>700</td><td>XDAI</td><td>8</td></tr>
   		    <tr><td>714</td><td>BNB</td><td>8</td></tr>
   		    <tr><td>966</td><td>MATIC</td><td>8</td></tr>
   		    <tr><td>1815</td><td>ADA</td><td>6</td></tr>
   		    <tr><td>5353</td><td>HNS</td><td>6</td></tr>
   		    <tr><td>52752</td><td>CELO</td><td>18</td></tr>
   		    <tr><td>99999999989</td><td>PALM*</td><td>18</td></tr>
   		    <tr><td>99999999990</td><td>FTM*</td><td>18</td></tr>
   		    <tr><td>99999999991</td><td>OKT*</td><td>12</td></tr>
   		    <tr><td>99999999992</td><td>OPTIMISM*</td><td>12</td></tr>
   		    <tr><td>99999999993</td><td>ARBITRUM*</td><td>12</td></tr>
   		    <tr><td>99999999994</td><td>HECO*</td><td>12</td></tr>
   		    <tr><td>99999999996</td><td>WND*</td><td>12</td></tr>
   		    <tr><td>99999999997</td><td>BSC*</td><td>18</td></tr>
  	 	  </tbody>
		</table>
		*pseudo cryptocurrency definition in the CYBAVO SOFA system
    </td>
  </tr>
  <tr>
    <td>token_address</td>
    <td>string</td>
    <td>The contract address of cryptocurrency</td>
  </tr>
  <tr>
    <td>memo</td>
    <td>string</td>
    <td>The memo/destination tag of the transaction</td>
  </tr>
</table>

<a name="transaction-state-definition"></a>
### Transaction State Definition

| ID   | State | Description | Callback | Callback Type |
| :-:  | :---  | :--- | :-: | :-- |
| 0    | Init        | The withdrawal request has been enqueued in the CYBAVO SOFA system | - | - |
| 1    | Processing  | The withdrawal request is processing in the CYBAVO SOFA system | O | Withdrawal(2) |
| 2    | TXID in pool | The withdrawal transaction is pending in the fullnode mempool | O | Withdrawal(2) |
| 3    | TXID in chain | The transaction is already on the blockchain | O | Deposit(1), Withdrawal(2), Collect(3) |
| 4    | TXID confirmed in N blocks | The withdrawal transaction is already on the blockchain and satisfy confirmations | - | - |
| 5    | Failed | Fail to create the withdrawal transaction | O | Withdrawal(2) |
| 6    | Resent | The transaction has been successfully resent | - | - |
| 7    | Blocked due to risk controlled | The deposit or withdrawal transaction was temporarily blocked due to a violation of the risk control rules | - | - |
| 8    | Cancelled | The withdrawal request is cancelled via web console | O | Withdrawal(2) |
| 9    | UTXO temporarily not available | The withdrawal request has been set as pending due to no available UTXO | - | - |
| 10   | Dropped | A long-awaited withdrawal transaction that does not appear in the memory pool of the fullnode will be regarded as dropped  | O | Withdrawal(2) |
| 11   | Transaction Failed | The withdrawal transaction is regarded as a failed transaction by the fullnode | O | Withdrawal(2) |
| 12   | Paused | The withdrawal request has been paused | - | - |

Callback sample:

> If the `state` of callback is 5 (Failed), the detailed failure reason will put in `addon` field (key is `err_reason`). See the callback sample bellow.

```json
{
  "type": 1,
  "serial": 90000000619,
  "order_id": "",
  "currency": "ETH",
  "txid": "0xc99a4941f87364c9679fe834f99bc12cbacfc577dedf4f34c4fd8833a68a0b00",
  "block_height": 8336269,
  "tindex": 43,
  "vout_index": 0,
  "amount": "500000000000000000",
  "fees": "945000000000000",
  "memo": "",
  "broadcast_at": 1595296751,
  "chain_at": 1595296751,
  "from_address": "0x8382Cc1B05649AfBe179e341179fa869C2A9862b",
  "to_address": "0x32d638773cB85965422b3B98e9312Fc9392307BC",
  "wallet_id": 5,
  "state": 3,
  "confirm_blocks": 2,
  "processing_state": 2,
  "addon": {
    "fee_decimal": 18
  },
  "decimal": 18,
  "currency_bip44": 60,
  "token_address": ""
}
```

Callback with state 5 (Failed) sample:

```json
{
  "type": 2,
  "serial": 20000000155,
  "order_id": "1_69",
  "currency": "ETH",
  "txid": "",
  "block_height": 0,
  "tindex": 0,
  "vout_index": 0,
  "amount": "1000000000000000",
  "fees": "",
  "memo": "",
  "broadcast_at": 0,
  "chain_at": 0,
  "from_address": "",
  "to_address": "0x60589A749AAC632e9A830c8aBE041899d8Dd15",
  "wallet_id": 2,
  "state": 5,
  "confirm_blocks": 0,
  "processing_state": 0,
  "addon": {
    "err_reason": "Illegal Transaction Format: To 0x60589A749AAC632e9A830c8aBE041899d8Dd15"
  },
  "decimal": 18,
  "currency_bip44": 60,
  "token_address": ""
}
```

Deposit callback with blocklist_tags sample:

```json
{
  "type": 1,
  "serial": 90000002797,
  "order_id": "",
  "currency": "ETH",
  "txid": "0xbb38e22c33cbc33ad5a58a88bfee0905968062fe34e33eb6e28861771686cf45",
  "block_height": 11075566,
  "tindex": 7,
  "vout_index": 0,
  "amount": "10000000000000000",
  "fees": "31500000315000",
  "memo": "",
  "broadcast_at": 1632195931,
  "chain_at": 1632195931,
  "from_address": "0xD5909BacFc1faD78e4e45E9e2feF8b4e61c8Fd0d",
  "to_address": "0x319b269ef02AB7e6660f7e6cb181D0CD06E2E4a0",
  "wallet_id": 854512,
  "processing_state": 2,
  "addon": {
    "address_label": "",
    "aml_screen_pass": false,
    "aml_tags": {
      "cybavo": {
        "score": 100,
        "tags": [
          "TEST"
        ],
        "blocked": true
      }
    },
    "blocklist_tags": [
      "cybavo(100): TEST"
    ],
    "fee_decimal": 18
  },
  "decimal": 18,
  "currency_bip44": 60,
  "token_address": ""
}
```


##### [Back to top](#table-of-contents)

<a name="callback-type-definition"></a>
### Callback Type Definition

| ID   | Description |
| :--- | :---        |
| 1 | Deposit Callback |
| 2 | Withdraw Callback |
| 3 | Collect Callback |
| 4 | Airdrop Callback |
| -1 | All callbacks (for inquiry) |

##### [Back to top](#table-of-contents)

<a name="currency-definition"></a>
### Currency Definition

| ID   | Currency Symbol | Decimals |
| :--- | :---            | :---     |
| 0    | BTC             | 8 |
| 2    | LTC             | 8 |
| 3    | DOGE            | 8 |
| 5    | DASH            | 8 |
| 60   | ETH             | 18 |
| 144  | XRP             | 6 |
| 145  | BCH (BCHN)      | 8 |
| 148  | XLM             | 7 |
| 194  | EOS             | 4 |
| 195  | TRX             | 6 |
| 236  | BSV             | 8 |
| 354  | DOT             | 10 |
| 461  | FIL             | 18 |
| 472  | AR              | 12 |
| 501  | SOL             | 9 |
| 539  | FLOW            | 8 |
| 700  | XDAI            | 18 |
| 714  | BNB             | 8 |
| 966  | MATIC           | 18 |
| 1815 | ADA             | 6 |
| 5353 | HNS             | 6 |
| 52752 | CELO           | 18 |
| 99999999989 | PALM*    | 18 |      
| 99999999990 | FTM*     | 18 |
| 99999999991 | OKT*     | 18 |
| 99999999992 | OPTIMISM* | 18 |
| 99999999993 | ARBITRUM* | 18 |
| 99999999994 | HECO*    | 18 |
| 99999999996 | WND*     | 12 |
| 99999999997 | BSC*     | 18 |
  
> Refer to [here](https://github.com/satoshilabs/slips/blob/master/slip-0044.md) for more detailed currency definitions
> 
> The * mark represents the definition of pseudo-cryptocurrency in the CYBAVO SOFA system
 
##### [Back to top](#table-of-contents)

<a name="memo-requirement"></a>
### Memo Requirement

| Currency | Description |
| :--- | :---        |
| XRP | Up to 32-bit unsigned integer (max 4294967295) |
| XLM | Up to 20 chars |
| EOS | Up to 256 chars |
| BNB | Up to 128 chars |

##### [Back to top](#table-of-contents)

<a name="wallet-type-definition"></a>
### Wallet Type Definition

| Type | Description |
| :--- | :---        |
| 0 | Vault wallet |
| 1 | Batch wallet |
| 2 | Deposit wallet |
| 3 | Withdrawal wallet |
| 5 | Deposit-withdrawal hybrid wallet |

##### [Back to top](#table-of-contents)

