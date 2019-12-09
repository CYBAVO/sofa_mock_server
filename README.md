<a name="table-of-contents"></a>
## Table of contents

- [Get Started](#get-started)
- Deposit Wallet API
	- [Create deposit wallet addresses](#create-deposit-wallet-addresses)
	- [Query address of deposit wallet](#query-address-of-deposit-wallet)
	- [Query pool address of deposit wallet](#query-pool-address-of-deposit-wallet)
	- [Resend pending or failed deposit callbacks](#resend-pending-or-failed-deposit-callbacks)
- Withdraw Wallet API
	- [Withdraw](#withdraw)
- Query API
	- [Query API token status](#query-api-token-status)
	- [Query notification callback history](#query-notification-callback-history)
	- [Query vault/batch wallet transaction history](#query-vault/batch-wallet-transaction-history)
	- [Query wallet block info](#query-wallet-block-info)
	- [Query invalid deposit addresses](#query-invalid-deposit-addresses)
	- [Query wallet basic info](#query-wallet-basic-info)
- Testing
	- [Mock Server](#mock-server)
	- [CURL Testing Commands](#curl-testing-commands)
- Appendix
	- [Callback Definition](#callback-definition)
	- [Notification Callback Type Definition
](#notification-callback-type-definition)
	- [Transaction State Filter Definition](#transaction-state-filter-definition)

<a name="get-started"></a>
# Get Started

- Get API code and API secret of the wallet on web console
- Refer to [mock server](#mock-server) sample code 


# Deposit Wallet API

<a name="create-deposit-wallet-addresses"></a>
## Create deposit wallet addresses

Create deposit addresses in a deposit wallet.

**POST** /v1/sofa/wallets/`WALLET_ID`/addresses

> Wallet ID must be a deposit wallet's ID

- [Sample curl command](#curl-create-deposit-wallet-addresses)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/1/addresses
```

###### Post body

For BNB or EOS wallet:

```json
{
  "count": 2,
  "memos": [
    "001",
    "002"
  ]
}
```

For wallet excepts BNB and EOS:

```json
{
  "count": 2
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Description |
| :---  | :---  | :---        |
| count | int | Specify address count, max value is 1000 |
| memos | array | Specify memos for BNB or EOS deposit wallet |

> **NOTE: The length of `memos` must equal to `count` while creating addresses for BNB or EOS wallet**

##### Response Format

An example of a successful response:

For BNB or EOS wallet:
	
```json
{
  "addresses": [
    "002",
    "001"
  ]
}
```
	
For wallet excepts BNB and EOS:
	
```json
{
  "addresses": [
    "0x2E7248BBCD61Ad7C33EA183A85B1856bc02C40b6",
    "0x4EB990D527c96c64eC5Bfb0D1e304840052d4975",
    "0x86434604FF857702fbE11cBFf5aC7689Af19c4Ed"
  ]
}
```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| addresses | array | Array of just created deposit addresses |

##### [Back to top](#table-of-contents)


<a name="query-address-of-deposit-wallet"></a>
## Query address of deposit wallet

Query deposit addresses those created by address creation API in a deposit wallet.

**GET** /v1/sofa/wallets/`WALLET_ID`/addresses?start\_index=`from`&request\_number=`count`

> Wallet ID must be a deposit wallet's ID

- [Sample curl command](#curl-get-deposit-wallet-addresses)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/17/addresses?start_index=0&request_number=3

--- THEN ---

/v1/sofa/wallets/17/addresses?start_index=3&request_number=3
```

The request includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| start_index | int | Specify address start index |
| request_number | int | Request address count |

##### Response Format

An example of a successful response:

```json
{
  "wallet_id": 17,
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

##### [Back to top](#table-of-contents)

<a name="query-pool-address-of-deposit-wallet"></a>
## Query pool address of deposit wallet

Get pool address of a deposit wallet.

**GET** /v1/sofa/wallets/`WALLET_ID`/pooladdress

> Wallet ID must be a deposit wallet's ID

- [Sample curl command](#curl-get-deposit-wallet-pool-address)

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

##### [Back to top](#table-of-contents)

<a name="resend-pending-or-failed-deposit-callbacks"></a>
## Resend pending or failed deposit callbacks

Trigger pending/failed deposit callback resending process manually.

**POST** /v1/sofa/wallets/`WALLET_ID`/collection/notifications/manual

- [Sample curl command](#curl-resend-all-pending-or-failed-deposit-callbacks)

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

| Field | Type  | Description |
| :---  | :---  | :---        |
| notification_id | int64 | Specify callback ID to resend, 0 means all |

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

##### [Back to top](#table-of-contents)


# Withdraw Wallet API

<a name="withdraw"></a>
## Withdraw

Withdraw assets from withdraw wallet.

**POST** /v1/sofa/wallets/`WALLET_ID`/sender/transactions

> Wallet ID must be a withdraw wallet's ID

- [Sample curl command](#curl-withdraw)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/1/sender/transactions
```

###### Post body

```json
{
  "requests": [
    {
      "order_id": "1",
      "address": "0x60589A749AAC632e9A830c8aBE042D1899d8Dd15",
      "amount": "0.0001",
      "memo": "memo-001",
      "user_id": "USER01",
      "message": "message-001",
    },
    {
      "order_id": "2",
      "address": "0xf16B7B8900F0d2f682e0FFe207a553F52B6C7015",
      "amount": "0.0002",
      "memo": "memo-002",
      "user_id": "USER01",
      "message": "message-002",
    }
  ]
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Description |
| :---  | :---  | :---        |
| order_id | string | Specify an unique ID |
| address | string | Outgoing address |
| amount | string | Withdrawal amount |
| memo | string | Memo on blockchain (This memo will be sent to blockchain) |
| user_id | string | Specify certain user (optional) |
| message | string | Message (This message only savced on CYBAVO, not sent to blockchain) |

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
| results | Array | Array of withdraw result (order ID/withdraw transaction ID pair), if succeeds |

##### [Back to top](#table-of-contents)

# Query API

<a name="query-api-token-status"></a>
## Query API token status

**GET** /v1/sofa/wallets/`WALLET_ID`/apisecret

- [Sample curl command](#curl-query-api-token-status)

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
| valid | object | Activated API token |
| inactivated | object | Not active API token |
| api_code | string | API code for querying wallet |
| exp | int64 | API code expiration unix time in UTC |

> 未啟用的API-CODE會在第一次使用(填在**X-API-CODE**)時自動生效，並讓當前已啟用的API-CODE失效

> 若使用已失效的API-CODE查詢會得到403 Forbidden

##### [Back to top](#table-of-contents)

<a name="query-notification-callback-history"></a>
## Query notification callback history

**GET** /v1/sofa/wallets/`WALLET_ID`/notifications?from\_time=`from`&to\_time=`to`&type=`type`

- [Sample curl command](#curl-query-notification-callback-history)

##### Request Format

An example of the request:

###### API

```
/v1/sofa/wallets/67/notifications?from_time=1561651200&to_time=1562255999&type=2
```

The request includes the following parameters:

###### API Parameters

| Field | Type  | Description |
| :---  | :---  | :---        |
| from_time | int64 | Start date (unix time in UTC) |
| to_time | int64 | End date (unix time in UTC) |
| type | int | [Notification Callback Type](#notification-callback-type-definition)  |

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
      "addon": {}
    }
  ]
}

```

The response includes the following parameters:

| Field | Type  | Description |
| :---  | :---  | :---        |
| notifications | array | Arrary of callbacks, refer to [Callback Definition](#callback-definition) |

##### [Back to top](#table-of-contents)

<a name="query-vault/batch-wallet-transaction-history"></a>
## Query vault/batch wallet transaction history

**GET** /v1/sofa/wallets/`WALLET_ID`/transactions?from\_time=`from`&to\_time=`to`&start\_index=`start`&request_number=`count`&state=`state`

- [Sample curl command](#curl-query-vault/batch-wallet-transaction-history)

##### Request Format

An example of the request:

###### API with parameters

```
/v1/sofa/wallets/48/transactions?from_time=1559664000&to_time=1562255999&start_index=0&request_number=1
```

The request includes the following parameters:

###### API Parameters

| Field | Type  | Description |
| :---  | :---  | :---        |
| from_item | int64 | Start date (unix time in UTC) |
| to_item | int64 | End date (unix time in UTC) |
| start_index | int | Index of starting transaction record |
| request_number | int | Count of returning transaction record |
| state | int | Refer to [Transaction State Filter](#transaction-state-filter-definition) (optional, default: -1) |

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

##### [Back to top](#table-of-contents)

<a name="query-wallet-block-info"></a>
## Query wallet block info

**GET** /v1/sofa/wallets/`WALLET_ID`/blocks

- [Sample curl command](#curl-query-deposit/withdraw-wallet-block-info)

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

##### [Back to top](#table-of-contents)

<a name="query-invalid-deposit-addresses"></a>
## Query invalid deposit addresses

**GET** /v1/sofa/wallets/`WALLET_ID`/addresses/invalid-deposit

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

##### [Back to top](#table-of-contents)

<a name="query-wallet-basic-info"></a>
## Query wallet basic info

**GET** /v1/sofa/wallets/`WALLET_ID`/info

- [Sample curl command](#curl-query-invalid-deposit-addresses)

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
| currency | int64 | Registered coin types |
| currency_name | string | Name of currency |
| address | string | Wallet address |

> Refer to https://github.com/satoshilabs/slips/blob/master/slip-0044.md for currency definitions

If `WALLET_ID` is a token wallet, the following fields present:

| Field | Type  | Description |
| :---  | :---  | :---        |
| token_name | string | Token name |
| token_symbol | string | Token symbol |
| token\_contract_address | string | Token contract address |
| token_decimals | string | Token decimals |

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
>	Set following configuration in mockserver.app.conf

```
api_server_url=""
```

### Put wallet API code/secret into mock server
-	Get API code/secret on web console
	-	API-CODE, API-SECRET, WALLET-ID
- 	Put API code/secret to mock server's database

```
curl -X POST -d '{"api_code":"API-CODE","api_secret":"API-SECRET"}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/apitoken
```

### Register mock server URL
>	Operate on web admin console

Callback URL

```
http://localhost:8889/v1/mock/wallets/callback
```

##### [Back to top](#table-of-contents)

<a name="curl-testing-commands"></a>
## CURL Testing Commands

<a name="curl-create-deposit-wallet-addresses"></a>
### Create deposit wallet addresses

For BNB or EOS wallet:

```
curl -X POST -d '{"count":2,"memos":["001","002"]}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/addresses
```

For wallet excepts BNB or EOS:

```
curl -X POST -d '{"count":2}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/addresses
```
- [API definition](#create-deposit-wallet-addresses)

<a name="curl-get-deposit-wallet-addresses"></a>
### Get deposit wallet addresses

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/addresses?start_index=0&request_number=1000'
```
- [API definition](#query-address-of-deposit-wallet)

<a name="curl-get-deposit-wallet-pool-address"></a>
### Get deposit wallet pool address

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/pooladdress'
```
- [API definition](#query-pool-address-of-deposit-wallet)

<a name="curl-resend-all-pending-or-failed-deposit-callbacks"></a>
### Resend all pending or failed deposit callbacks

```
curl -X POST -d '{"notification_id":0}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/callback/resend
```
- [API definition](#resend-pending-or-failed-deposit-callbacks)

<a name="curl-withdraw"></a>
### Withdraw

```
curl -X POST -d '{"requests":[{"order_id":"1","address":"0x60589A749AAC632e9A830c8aBE042D1899d8Dd15","amount":"0.0001","memo":"memo-001","user_id":"USER01","message":"message-001"},{"order_id":"2","address":"0xf16B7B8900F0d2f682e0FFe207a553F52B6C7015","amount":"0.0002","memo":"memo-002","user_id":"USER01","message":"message-002"}]}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/withdraw
```
- [API definition](#withdraw)

<a name="curl-query-api-token-status"></a>
### Query API token status

```
curl -X GET http://localhost:8889/v1/mock/wallets/{WALLET-ID}/apisecret
```
- [API definition](#query-api-token-status)

<a name="curl-query-notification-callback-history"></a>
### Query notification callback history

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/notifications?from_time=1561651200&to_time=1562255999&type=2'
```
- [API definition](#query-notification-callback-history)

<a name="curl-query-vault/batch-wallet-transaction-history"></a>
### Query vault/batch wallet transaction history

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/transactions?start_index=0&from_time=1559664000&to_time=1562255999&request_number=8'
```
- [API definition](#query-vault/batch-wallet-transaction-history)

<a name="curl-query-deposit/withdraw-wallet-block-info"></a>
### Query deposit/withdraw wallet block info

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/blocks'
```
- [API definition](#query-wallet-block-info)

<a name="curl-query-invalid-deposit-addresses"></a>
#### Query invalid deposit addresses

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/addresses/invalid-deposit'
```
- [API definition](#query-invalid-deposit-addresses)

<a name="curl-query-wallet-basic-info"></a>
### Query wallet basic info

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/info'
```
- [API definition](#query-wallet-basic-info)

##### [Back to top](#table-of-contents)

# Appendix

<a name="callback-definition"></a>
## Callback Definition

<table>
  <tr>
    <td>Field</td>
    <td>Type</td>
    <td>Description</td>
  </tr>
  <tr>
    <td>type</td>
    <td>int</td>
    <td rowspan="3">
      <b>1</b> - Deposit Callback (入金回調)<br>
      <b>2</b> - Withdraw Callback (出金回調)<br>
      <b>3</b> - Collect Callback (歸帳回調)<br>
    </td>
  </tr>
  <tr></tr>
  <tr></tr>
  <tr>
    <td>state</td>
    <td>int</td>
    <td rowspan="9">
      <b>0</b> - Enqueue<br>
      <b>1</b> - Processing batch in KMS<br>
      <b>2</b> - TXID in pool<br>
      <b>3</b> - TXID in chain<br>
      <span style="text-decoration:line-through"><b>4 (DEPRECATED)</b> - TXID confirmed in N blocks</span><br>
      <b>5</b> - Failed<br>
      <b>6</b> - Resent<br>
      <b>7</b> - Blocked due to risk controlled<br>
      <b>8</b> - Cancelled<br>
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
  <tr>
    <td>processing_state</td>
    <td>int</td>
    <td rowspan="3">
      <b>0</b> - inpool (mempool中)<br>
      <b>1</b> - in chain (已經上鏈還沒滿confirm block)<br>
      <b>2</b> - done (已經上鍊也滿足confirm block)<br>
    </td>
  </tr>
</table>

Callback sample:

```json
{
  "type": 1,
  "serial": 90000000001,
  "order_id": "",
  "currency": "ETH",
  "txid": "0x100e84230923124582b7feb5daf638df79616fb3dea37fc2ea80659f5de3472e",
  "block_height": 5905092,
  "tindex": 11,
  "vout_index": 0,
  "amount": "500000000000000000",
  "fees": "210000000000000",
  "broadcast_at": 1562057483,
  "chain_at": 1562057483,
  "from_address": "0x8382Cc1B05649AfBe179e341179fa869C2A9862b",
  "to_address": "0x87F907C868D92a5d97E59CD1E9383a2E51dC4778",
  "wallet_id": 21,
  "state": 3,
  "confirm_blocks": 1,
  "processing_state": 1,
  "addon": {}
}
```


##### [Back to top](#table-of-contents)

<a name="notification-callback-type-definition"></a>
## Notification Callback Type Definition

| ID   | Description |
| :--- | :---        |
| 1 | Deposit Callback (入金回調) |
| 2 | Withdraw Callback (出金回調) |
| 3 | Collect Callback (歸帳回調) |

##### [Back to top](#table-of-contents)

<a name="transaction-state-filter-definition"></a>
## Transaction State Filter Definition

| ID   | Description |
| :--- | :---        |
| -1 | All states |
| 0 | WaitApproval |
| 1 | Rejected |
| 2 | Approved |
| 3 | Failed |
| 4 | NextLevel |
| 5 | Cancelled |
| 6 | BatchDone |

##### [Back to top](#table-of-contents)

