<a name="table-of-contents"></a>
## Table of contents

- [Get Started](#get-started)
- [API Authentication](#api-authentication)
- [Callback Integration](#callback-integration)
- REST API
	- Deposit Wallet API
		- [Create Deposit Addresses](#create-deposit-addresses)
		- [Query Deposit Addresses](#query-deposit-addresses)
		- [Query Deployed Contract Deposit Addresses](#query-deployed-contract-deposit-addresses)
		- [Query Pool Address](#query-pool-address)
		- [Query Pool Address Balance](#query-pool-address-balance)
		- [Query Invalid Deposit Addresses](#query-invalid-deposit-addresses)
		- [Query Deposit Callback Detail](#query-deposit-callback-detail)
		- [Resend Deposit Callbacks](#resend-deposit-callbacks)
		- [Query Deposit Wallet Balance](#query-deposit-wallet-balance)
	- Withdraw Wallet API
		- [Withdraw Assets](#withdraw-assets)
		- [Cancel Withdrawal Request](#cancel-withdrawal-request)
		- [Query Latest Withdrawal Transaction State](#query-latest-withdrawal-transaction-state)
		- [Query All Withdrawal Transaction States](#query-all-withdrawal-transaction-states)
		- [Query Withdrawal Wallet Balance](#query-withdrawal-wallet-balance)
		- [Query Withdrawal Callback Detail](#query-withdrawal-callback-detail)
	- Deposit / Withdraw Wallet Common API
		- [Query Callback History](#query-callback-history)
		- [Query Callback Detail](#query-callback-detail)
		- [Query Wallet Synchronization Info](#query-wallet-synchronization-info)
		- [Query Transaction Average Fee](#query-transaction-average-fee)
	- Vault Wallet API
		- [Query Vault Wallet Transaction History](#query-vault-wallet-transaction-history)
		- [Query Vault Wallet Balance](#query-vault-wallet-balance)
	- Common API
		- [Activate API Code](#activate-api-code)
		- [Query API Code Status](#query-api-code-status)
		- [Query Wallet Info](#query-wallet-info)
		- [Verify Addresses](#verify-addresses)
- Testing
	- [Mock Server](#mock-server)
	- [cURL Testing Commands](#curl-testing-commands)
	- [Other Language Versions](#other-language-versions)
- Appendix
	- [Callback Definition](#callback-definition)
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

### How to make a correct request?
- Put the API code in the X-API-CODE header.
	- Use the inactivated API code in any request will activate it automatically. Once activated, the currently activated API code will immediately become invalid.
	- Or you can explicitly call the [activation API](#activate-api-code) to activate the API code before use
- Calculate the checksum with the corresponding API secret and put the checksum in the X-CHECKSUM header.
  - The checksum calculation will use all the query parameters, the current timestamp, user-defined random string and the post body (if any).
- Please refer to the code snippet on the github project to know how to calculate the checksum.
	- [Go](https://github.com/CYBAVO/SOFA_MOCK_SERVER/blob/master/api/apicaller.go#L40)
	- [Java](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVA/blob/master/src/main/java/com/cybavo/sofa/mock/Api.java#L71)
	- [Javascript](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVASCRIPT/blob/master/helper/apicaller.js#L58)
	- [PHP](https://github.com/CYBAVO/SOFA_MOCK_SERVER_PHP/blob/master/helper/apicaller.php#L26)

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
	- If all attempts fail, the callback will be set to a failed state, for deposit callbacks the callback handler can call the [resend](#resend-deposit-callbacks) API to request CYBAVO SOFA system to resend such kind of callback(s) or through the web control panel. For withdrawal callbacks, the resend operation must be completed on the web control panel.

- Refer to [Callback Definition](#callback-definition), [Callback Type Definition](#callback-type-definition) for detailed definition.
- Please refer to the code snippet on the github project to know how to validate the callback payload.
	- [Go](https://github.com/CYBAVO/SOFA_MOCK_SERVER/blob/master/controllers/OuterController.go#L197)
	- [Java](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVA/blob/master/src/main/java/com/cybavo/sofa/mock/MockController.java#L82)
	- [Javascript](https://github.com/CYBAVO/SOFA_MOCK_SERVER_JAVASCRIPT/blob/master/routes/wallets.js#L385)
	- [PHP](https://github.com/CYBAVO/SOFA_MOCK_SERVER_PHP/blob/master/index.php#L203)



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
    "001",
    "002"
  ]
}
```

For wallet excepts BNB, XLM, XRP and EOS:

```json
{
  "count": 2
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Note | Description |
| :---  | :---  | :--- | :---        |
| count | int | required, max `1000` | Specify address count |
| memos | array | required (creating BNB, XLM, XRP or EOS wallet) | Specify memos for BNB, XLM, XRP or EOS deposit wallet. Refer to [Memo Requirement](#memo-requirement) |

> NOTE: The length of `memos` must equal to `count` while creating addresses for BNB, XLM, XRP or EOS wallet.

##### Response Format

An example of a successful response:

For BNB, XLM, XRP or EOS wallet:
	
```json
{
  "addresses": [
    "002",
    "001"
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 400 | 703 | Operation failed | Error message returned by JSON parser | Malformatted post body |
| 400 | 945 | The max length of BNB memo is 256 chars | - | Reached the limit of the length of BNB memo |
| 400 | 946 | The max length of EOS memo is 128 chars | - | Reached the limit of the length of EOS memo |
| 400 | 947 | The max length of XRP destination tag is 20 chars | - | Reached the limit of the length of XRP destination tag |
| 400 | 948 | The max length of XLM memo is 20 chars | - | Reached the limit of the length of XLM memo |
| 400 | 818 | Destination Tag must be integer | - | Wrong XRP destination tag format |
| 400 | 500 | insufficient fund | - | Insufficient balance to deploy collection contract |
| 403 | 706 | Exceed max allow wallet limitation, Upgrade your SKU to get more wallets | - | Reached the limit of the total number of deposit addresses |
| 403 | 112 | Invalid parameter | - | The count and the count of memos mismatched |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-deposit-addresses"></a>
### Query Deposit Addresses

Query the deposit addresses created by the [Create Deposit Addresses](#create-deposit-addresses) API.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/addresses?start\_index=`from`&request\_number=`count`

> `WALLET_ID` must be a deposit wallet ID

- [Sample curl command](#curl-query-deposit-addresses)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/17/addresses?start_index=0&request_number=3

--- THEN ---

/v1/sofa/wallets/17/addresses?start_index=3&request_number=3
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
  "wallet_id": 17,
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
  "wallet_id": 17,
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-deployed-contract-deposit-addresses"></a>
### Query Deployed Contract Deposit Addresses

Query deployed contract deposit addresses created by the [Create Deposit Addresses](#create-deposit-addresses) API.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/addresses/contract_txid?txids=`txid1,txid2`

> `WALLET_ID` must be an ETH contract collection deposit wallet ID
> 
> Only deployed addresses will be returned

- [Sample curl command](#curl-query-deployed-contract-deposit-addresses)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/17/addresses/contract_txid?txids=0xe6dfe0d283690f636df5ea4b9df25552e6b576b88887bfb5837016cdd696e754,0xdb18fd33c9a6809bfc341a1c0b2c092be5a360f394c85367f9cf316579281ab4
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-pool-address"></a>
### Query Pool Address

Get the pool address of a deposit wallet. The pool address has different functionality in different cryptocurrencies.

> In BTC, ETH, BCH or LTC, the cryptocurrency in the pool address will be used to pay for token transfer(ex. ERC20, USDT-Omni).
> 
> In EOS, XRP, XLM or BNB, the pool address is the user's deposit address. All user deposits will be distinguished by memo / tag field.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/pooladdress

> `WALLET_ID` must be a deposit wallet ID

- [Sample curl command](#curl-get-pool-address)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/17/pooladdress
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-pool-address-balance"></a>
### Query Pool Address Balance

Get the pool address balance of a deposit wallet.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/pooladdress/balance

> `WALLET_ID` must be a deposit wallet ID

- [Sample curl command](#curl-query-pool-address-balance)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/17/pooladdress/balance
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)



<a name="query-invalid-deposit-addresses"></a>
### Query Invalid Deposit Addresses

When an abnormal deposit is detected, the CYBAVO SOFA system will set the deposit address to invalid. Use this API to obtain the all invalid deposit addresses for further usage.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/addresses/invalid-deposit

> `WALLET_ID` must be a deposit wallet ID

- [Sample curl command](#curl-query-invalid-deposit-addresses)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/1/addresses/invalid-deposit
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-deposit-callback-detail"></a>
### Query Deposit Callback Detail

Query the detailed information of the deposit callback by the tx ID and the vout index.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/receiver/notifications/txid/`TX_ID`/`VOUT_INDEX`

- [Sample curl command](#curl-query-deposit-callback-detail)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/5/receiver/notifications/txid/0xb72a81976f780445decd13a35c24974c4e32665cb57d79b3f7a601c775f6a7d8/0
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
    "wallet_id": 5
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently || 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request or the callback (txid+vout_index) not found |

##### [Back to top](#table-of-contents)


<a name="resend-deposit-callbacks"></a>
### Resend Deposit Callbacks

The callback handler can call this API to resend pending or failed deposit callbacks.

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
/v1/sofa/wallets/1/collection/notifications/manual
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

> This ID equal to callback data's serial/order_id

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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 400 | 112 | Invalid parameter | - | Malformatted post body |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-deposit-wallet-balance"></a>
### Query Deposit Wallet Balance

Get the deposit wallet balance.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/receiver/balance

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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


# Withdraw Wallet API

<a name="withdraw-assets"></a>
### Withdraw Assets

To withdraw assets from an withdrawal wallet, the caller must to provide an unique **order_id** for each request, the CYBAVO SOFA system will send the callback with the unique **order_id** when the withdrawal is success (from `in pool` state to `in chain` state). 

By default, the withdraw API will perform the address check to verify that the outgoing address is good or not. If the address in the request is marked as a problematic address, the request will be aborted. The error message will identify the problematic addresses. Set the `ignore_black_list` to true to skip the address check. 

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
/v1/sofa/wallets/1/sender/transactions
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
      "memo": "memo-002",
      "user_id": "USER02",
      "message": "message-002",
      "manual_fee": 50
    },
    {
      "order_id": "888888_3",
      "address": "0x9638fa816ccd35389a9a98a997ee08b5321f3eb9",
      "amount": "0.0002",
      "memo": "memo-003",
      "user_id": "USER03",
      "message": "message-003"
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
| address | string | required | Outgoing address |
| amount | string | required | Withdrawal amount |
| memo | string | optional | Memo on blockchain (This memo will be sent to blockchain). Refer to [Memo Requirement](#memo-requirement) |
| user_id | string | optional | Specify certain user |
| message | string | optional | Message (This message only saved on CYBAVO, not sent to blockchain) |
| block\_average_fee | int | optional, range `1~30` | Use average blockchain fee within latest N blocks |
| manual_fee | int | optional, range `1~1000` | Specify blockchain fee in smallest unit of wallet currency |
| ignore\_black_list| boolean | optional, default `false` | After setting, the address check will not be performed. |

> The order\_id must be prefixed. Find prefix from corresponding wallet detail on web control panel
>
> block\_average\_fee and manual_fee are mutually exclusive configurations. If neither of these fields is set, the fee will refer to corresponding withdrawal policy of the withdrawal wallet.

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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
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
| 400 | 703 | Operation failed | HTTP resp failed {HTTP\_CODE}, body: {RESPONSE_BODY} | The authentication callback URL returned status code other than 200 |
| 403 | 827 | Outgoing address in black list, abort transaction | - | Some outgoing addresses are blacklisted, examine the response 'blacklist' field for detailed information |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |
| 404 | 312 | Policy not found | - | No active withdrawal policy found |
| 404 | 703 | Operation failed | Unrecognized response: {RESPONSE_BODY}, 'OK' expected | The withdrawal request is not allowed by authentication callback URL |

##### [Back to top](#table-of-contents)


<a name="cancel-withdrawal-request"></a>
### Cancel Withdrawal Request

To cancel the withdrawal request which state is `Init`. The request state can be checked on web control panel or query through this [API](#query-withdrawal-callback-detail) (represents `state` = 0).

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/sender/transactions/`ORDER_ID`/cancel

> `WALLET_ID` must be a withdrawal wallet ID

- [Sample curl command](#curl-cancel-withdrawal-request)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/1/sender/transactions/94531/cancel
```

##### Response Format

The HTTP 200 means the withdrawal request has been cancelled successfully.

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 403 | 177 | Illegal state | - | The {ORDER\_ID} withdrawal request is not in `Init` state |
| 404 | 304 | Wallet ID invalid | - | The {ORDER\_ID} not found |

##### [Back to top](#table-of-contents)


<a name="query-latest-withdrawal-transaction-state"></a>
### Query Latest Withdrawal Transaction State

Check the latest withdrawal transaction state of certain order ID.

> The order ID is used in the [withdraw assets](#withdraw-assets) API.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/sender/transactions/`ORDER_ID`

> `WALLET_ID` must be a withdrawal wallet ID

- [Sample curl command](#curl-query-latest-withdrawal-transaction-state)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/1/sender/transactions/1
```

##### Response Format

An example of a successful response:

```json
{
  "order_id": "1",
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
| create_time | string | The withdrawal unix time in UTC |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 404 | 304 | Wallet ID invalid | - | The {ORDER\_ID} not found |

##### [Back to top](#table-of-contents)


<a name="query-all-withdrawal-transaction-states"></a>
### Query All Withdrawal Transaction States

Check the all withdrawal transaction states of certain order ID.

> The order ID is used in the [withdraw assets](#withdraw-assets) API.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/sender/transactions/`ORDER_ID`/all

> `WALLET_ID` must be a withdrawal wallet ID

- [Sample curl command](#curl-query-all-withdrawal-transaction-states)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/1/sender/transactions/100005/all
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
      "order_id": "100005",
      "state": 6,
      "txid": "0x2a8a44f1cfed9cd7b86d86170e2418566765f88c5186246f571374df218fd1a1"
    },
    {
      "address": "0x36a49c68EF1e3f39CDbaE2f5636C74BA10815cea",
      "amount": "0.105",
      "create_time": "2020-09-24T03:44:35Z",
      "in_chain_block": 8742982,
      "memo": "",
      "order_id": "100005",
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
| create_time | string | The withdrawal unix time in UTC |
| state | int | Refer to `state` field in [Callback Definition](#callback-definition) |

##### Error Code

| HTTP Code | Error Code | Error | Message | Description |
| :---      | :---       | :---  | :---    | :---        |
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 404 | 304 | Wallet ID invalid | - | The {ORDER\_ID} not found |

##### [Back to top](#table-of-contents)


<a name="query-withdrawal-wallet-balance"></a>
### Query Withdrawal Wallet Balance

Get the withdrawal wallet balance. Facilitate to establish a real-time balance monitoring mechanism.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/sender/balance

> `WALLET_ID` must be a withdrawal wallet ID

- [Sample curl command](#curl-query-withdrawal-wallet-balance)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/1/sender/balance
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

##### [Back to top](#table-of-contents)


<a name="query-withdrawal-callback-detail"></a>
### Query Withdrawal Callback Detail

Query the detailed information of the withdrawal callback by the order ID.

> This API only provides in-chain transactions query, for those in-pool transactions use [Query All Withdrawal Transaction States](#query-all-withdrawal-transaction-states) API instead.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/sender/notifications/order_id/`ORDER_ID`

- [Sample curl command](#curl-query-withdrawal-callback-detail)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/2/sender/notifications/order_id/94531
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
    "order_id": "94531",
    "processing_state": 2,
    "serial": 90000000554,
    "state": 3,
    "tindex": 30,
    "to_address": "0x60589A749AAC632e9A830c8aBE042D1899d8Dd15",
    "txid": "0x471c11f139ce1a7e0627a05cea50d64e47e797c94fd72025f1978cc919e07aa9",
    "type": 2,
    "vout_index": 0,
    "wallet_id": 2
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request or the {ORDER\_ID} is not found |

##### [Back to top](#table-of-contents)


# Deposit / Withdraw Wallet Common API

<a name="query-callback-history"></a>
### Query Callback History

Used to query some kind of callbacks within a time interval.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/notifications?from\_time=`from`&to\_time=`to`&type=`type`

- [Sample curl command](#curl-query-callback-history)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/67/notifications?from_time=1561651200&to_time=1562255999&type=2
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
      "wallet_id": 67,
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |

##### [Back to top](#table-of-contents)


<a name="query-callback-detail"></a>
### Query Callback Detail

Query the detailed information of the callback by its serial ID. It can be used to reconfirm whether a deposit callback exists.

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/notifications/get\_by_id

- [Sample curl command](#curl-query-callback-detail)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/120/notifications/get_by_id
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
      "wallet_id": 120,
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
      "wallet_id": 120,
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |

##### [Back to top](#table-of-contents)


<a name="query-wallet-synchronization-info"></a>
### Query Wallet Synchronization Info

Get the blockchain synchronization status of a wallet.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/blocks

- [Sample curl command](#curl-query-wallet-synchronization-info)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/1/blocks
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |

##### [Back to top](#table-of-contents)


<a name="query-transaction-average-fee"></a>
### Query Transaction Average Fee

Query average blockchain fee within latest N blocks.

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/autofee

- [Sample curl command](#curl-query-transaction-average-fee)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/1/autofee
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
| block_num | int | optional, default `1`, range `1~30` | Get average blockchain fee within latest N blocks |

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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 400 | 112 | Invalid parameter | - | The `block_num` is out of range |

##### [Back to top](#table-of-contents)


# Vault Wallet API

<a name="query-vault-wallet-transaction-history"></a>
### Query Vault Wallet Transaction History

Get transaction history of vault wallets.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/transactions?from\_time=`from`&to\_time=`to`&start\_index=`start`&request_number=`count`&state=`state`

> `WALLET_ID` must be a vault wallet ID

- [Sample curl command](#curl-query-vault-wallet-transaction-history)

##### Request Format

An example of the request:

###### API with parameters

```
/v1/sofa/wallets/48/transactions?from_time=1559664000&to_time=1562255999&start_index=0&request_number=1
```

The request includes the following parameters:

###### Query Parameters

| Field | Type | Note | Description |
| :---  | :--- | :--- | :---        |
| from_item | int64 | optional, default `0` | Start date (unix time in UTC) |
| to_item | int64 | optional, default `current time` | End date (unix time in UTC) |
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
      "wallet_id": 48,
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |

##### [Back to top](#table-of-contents)


<a name="query-vault-wallet-balance"></a>
### Query Vault Wallet Balance

Get the vault wallet balance. Facilitate to establish a real-time balance monitoring mechanism.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/vault/balance

> `WALLET_ID` must be a vault wallet ID

- [Sample curl command](#curl-query-vault-wallet-balance)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/7/vault/balance
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 404 | 304 | Wallet ID invalid | - | The wallet is not allowed to perform this request |

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
/v1/sofa/wallets/1/apisecret/activate
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |

##### [Back to top](#table-of-contents)


<a name="query-api-code-status"></a>
### Query API Code Status

Query the API code info of a certain wallet. Use the `inactivated` API code in any request will activate it. Once activated, the currently activated API code will immediately become invalid.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/apisecret

- [Sample curl command](#curl-query-api-code-status)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/1/apisecret
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |

##### [Back to top](#table-of-contents)


<a name="query-wallet-info"></a>
### Query Wallet Info

Get wallet basic information.

##### Request

**GET** /v1/sofa/wallets/`WALLET_ID`/info

- [Sample curl command](#curl-query-wallet-info)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/1/info
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |

##### [Back to top](#table-of-contents)


<a name="verify-addresses"></a>
### Verify Addresses

Check if the address conforms to the wallet cryptocurrency address format (for example, ETH must have the prefix 0x, BTC should start with 1, 3 or bc1, etc).

> If the wallet's cryptocurrency is BNB or XRP, there will be a `must_need_memo` flag to indicate whether the address needs a memo / destination tag when transferring cryptocurrency to that address.

##### Request

**POST** /v1/sofa/wallets/`WALLET_ID`/addresses/verify

> Wallet ID must be a deposit or withdraw wallet's ID

- [Sample curl command](#curl-verify-addresses)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/1/addresses/verify
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
| 403 | -   | Forbidden. Invalid wallet ID | - | No wallet ID found |
| 403 | -   | Forbidden. Header not found | - | Missing `X-API-CODE`, `X-CHECKSUM` header or query param `t` |
| 403 | -   | Forbidden. Invalid timestamp | - | The timestamp `t` is not in the valid time range |
| 403 | -   | Forbidden. Invalid checksum | - | The request is considered a replay request |
| 403 | -   | Forbidden. Invalid API code | - | `X-API-CODE` header contains invalid API code |
| 403 | -   | Invalid API code for wallet {WALLET_ID} | - | The API code mismatched |
| 403 | -   | Forbidden. Checksum unmatch | - | `X-CHECKSUM` header contains wrong checksum |
| 403 | -   | Forbidden. Call too frequently ({THROTTLING_COUNT} calls/minute) | - | Send requests too frequently |
| 400 | 112 | Invalid parameter | - | Malformatted post body |

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
curl -X POST -H "Content-Type: application/json" -d '{"count":2,"memos":["001","002"]}' \
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
### Resend Deposit Callbacks

```
curl -X POST -H "Content-Type: application/json" -d '{"notification_id":0}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/callback/resend
```
- [API definition](#resend-deposit-callbacks)


<a name="curl-query-deposit-wallet-balance"></a>
### Query Deposit Wallet Balance

```
curl http://localhost:8889/v1/mock/wallets/{WALLET_ID}/receiver/balance
```
- [API definition](#query-deposit-wallet-balance)


<a name="curl-withdraw-assets"></a>
### Withdraw Assets

```
curl -X POST -H "Content-Type: application/json" -d '{"requests":[{"order_id":"888888_1","address":"0x60589A749AAC632e9A830c8aBE042D1899d8Dd15","amount":"0.0001","memo":"memo-001","user_id":"USER01","message":"message-001"},{"order_id":"888888_2","address":"0xf16B7B8900F0d2f682e0FFe207a553F52B6C7015","amount":"0.0002","memo":"memo-002","user_id":"USER01","message":"message-002"}]}' \
http://localhost:8889/v1/mock/wallets/{WALLET_ID}/withdraw
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
    <td rowspan="4">
      <b>1</b> - Deposit Callback<br>
      <b>2</b> - Withdraw Callback<br>
      <b>3</b> - Collect Callback<br>
      <b>4</b> - Airdrop Callback<br>
    </td>
  </tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
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
    <td rowspan="13">
      <b>0</b> - Init<br>
      <b>1</b> - Processing batch in KMS<br>
      <b>2</b> - TXID in pool<br>
      <b>3</b> - TXID in chain<br>
      <b>4</b> - TXID confirmed in N blocks<br>
      <b>5</b> - Failed (addon field of callback will contain detailed error reason)<br>
      <b>6</b> - Resent<br>
      <b>7</b> - Blocked due to risk controlled<br>
      <b>8</b> - Cancelled<br>
      <b>9</b> - Retry for UTXO Temporarily Not Available<br>
      <b>10</b> - Dropped<br>
      <b>11</b> - Transaction Failed<br>
      <b>12</b> - Paused<br>
    </td>
  </tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr></tr>
  <tr>
    <td>confirm_blocks</td>
    <td>int64</td>
    <td>Number of confirmations</td>
  </tr>
  <tr>
    <td>processing_state</td>
    <td>int</td>
    <td rowspan="3">
      <b>0</b> - in pool (in fullnode mempool)<br>
      <b>1</b> - in chain (the transaction is already on the blockchain but the confirmations have not been met)<br>
      <b>2</b> - done (the transaction is already on the blockchain and satisfy confirmations)<br>
    </td>
  </tr>
  <tr></tr>
  <tr></tr>
  <tr>
    <td>addon</td>
    <td>key-value pairs</td>
    <td>
    The extra information of this callback<br>
    <b>err_reason</b> - will contain detail error reason if state is 5(Failed)<br>
    <b>fee_decimal</b> - the decimal of cryptocurrency miner fee<br>
    <b>blocklist_tags</b> - the tags of CYBAVO AML detection
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
    <td rowspan="">
   	 	The coin type definition of cryptocurrency. <br>
		0 - BTC<br>
		2 - LTC<br>
		5 - DASH<br>
		60 - ETH<br>
		144 - XRP<br>
		145 - BCH<br>
		148 - XLM<br>
		194 - EOS<br>
		195 - TRX<br>
		236 - BSV<br>
		461 - FIL<br>
		714 - BNB<br>
	   Refer to Currency Definition table below.
    </td>
  </tr>
  <tr>
    <td>token_address</td>
    <td>string</td>
    <td>The contract address of cryptocurrency</td>
  </tr>
</table>

> If the `state` of callback is 5 (Failed), the detailed failure reason will put in `addon` field (key is `err_reason`). See the callback sample bellow.

Callback sample:

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
  "type": 4,
  "serial": 90000009949,
  "order_id": "",
  "currency": "ETH",
  "txid": "0xb74bb5cccf7c024b1b86ded449a7f40e01a2a81dc2298af1eee452de64a7bd15",
  "block_height": 8771833,
  "tindex": 1,
  "vout_index": 0,
  "amount": "88800000000000",
  "fees": "2142000000000000",
  "memo": "",
  "broadcast_at": 1601287122,
  "chain_at": 1601287122,
  "from_address": "0xf6DabB290FCE73f5617ED381ca90dBb7af0E8295",
  "to_address": "0x42214E5c0eb0960f14AEd3623c172986d5c24A0D",
  "wallet_id": 2,
  "state": 3,
  "confirm_blocks": 2,
  "processing_state": 2,
  "addon": {
    "blocklist_tags": [
      "Involve phishing activity",
      "Involve cybercrime related"
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

| ID   | Description |
| :--- | :---        |
| 0 | BTC |
| 2 | LTC |
| 5 | DASH |
| 60 | ETH |
| 144 | XRP |
| 145 | BCH |
| 148 | XLM |
| 194 | EOS |
| 195 | TRX |
| 236 | BSV |
| 461 | FIL |
| 714| BNB |
| 1815| ADA |
  
> Refer to [here](https://github.com/satoshilabs/slips/blob/master/slip-0044.md) for more detailed currency definitions
 
##### [Back to top](#table-of-contents)

<a name="memo-requirement"></a>
### Memo Requirement

| Currency | Description |
| :--- | :---        |
| XRP | Up to 20 digits |
| XLM | Up to 28 bytes of ASCII/UTF-8 |
| EOS | Up to 256 chars |
| BNB | Up to 128 chars |

##### [Back to top](#table-of-contents)
