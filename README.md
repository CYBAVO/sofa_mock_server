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
	- [Query withdrawal transaction state](#query-withdrawal-transaction-state)
	- [Query withdrawal wallet balance](#query-withdrawal-wallet-balance)
- Query API
	- [Query API token status](#query-api-token-status)
	- [Query notification callback history](#query-notification-callback-history)
	- [Query notification callback by ID](#query-notification-callback-by-id)
	- [Query vault/batch wallet transaction history](#query-vault/batch-wallet-transaction-history)
	- [Query wallet block info](#query-wallet-block-info)
	- [Query invalid deposit addresses](#query-invalid-deposit-addresses)
	- [Query wallet basic info](#query-wallet-basic-info)
	- [Verify addresses](#verify-addresses)
	- [Query wallet transaction avarage blockchain fee](#query-wallet-transaction-autofee)
- Testing
	- [Mock Server](#mock-server)
	- [CURL Testing Commands](#curl-testing-commands)
- Appendix
	- [Callback Definition](#callback-definition)
	- [Notification Callback Type Definition
](#notification-callback-type-definition)
	- [Transaction State Filter Definition](#transaction-state-filter-definition)
	- [Currency Definition](#currency-definition)
	- [Memo Requirement](#memo-requirement)

<a name="get-started"></a>
# Get Started

- Get API code and API secret of the wallet on web console
- Refer to [mock server](#mock-server) sample code 


# Deposit Wallet API

<a name="create-deposit-wallet-addresses"></a>
## Create deposit wallet addresses

You can create new deposit address through this API, once the addresses were created, both of the deposit and withdraw behavior occurs on these addresses will be callback..

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

| Field | Type  | Required | Description |
| :---  | :---  | :--- | :---        |
| count | int | YES | Specify address count, max value is 1000 |
| memos | array | YES (if create BNB, XLM, XRP or EOS wallet) | Specify memos for BNB, XLM, XRP or EOS deposit wallet. Refer to [Memo Requirement](#memo-requirement) |

> NOTE: The length of `memos` must equal to `count` while creating addresses for BNB or EOS wallet

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

| Field | Type  | Requried | Description |
| :---  | :---  | :--- | :---        |
| start_index | int | NO | Specify address start index (default: 0) |
| request_number | int | NO | Request address count (default: 1000, max: 5000) |

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

> Refer to [Currency Definition](#currency-definition) or [here](https://github.com/satoshilabs/slips/blob/master/slip-0044.md) for more detailed currency definitions

##### [Back to top](#table-of-contents)

<a name="query-pool-address-of-deposit-wallet"></a>
## Query pool address of deposit wallet

Get the pool address of a deposit wallet. The pool address has different functionality in different cryptocurrency.
> In BTC/ETH/BCH/LTC, the cryptocurrency in the pool address will be used to pay for token transfer(Ex. ERC20, USDT-Omni).
> 
> In EOS/XRP/XLM/BNB, the pool address is the user’s deposit address, and all user will be distinguished by memo/tag field.

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

When the deposit callback sent to your server and not got the correct response, the SOFA system will resend the callback automatically (time interval: 1-3-5-15-45 mins), and if all resends were failed, for example your server is under maintenance or network is corrupted at that time. All the callbacks will be put in failed zone, and you can use this api to ask SOFA system resend again or you can resend it through SOFA UI resend button.

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

| Field | Type  | Required | Description |
| :---  | :---  | :---  | :---        |
| notification_id | int64 | YES | Specify callback ID to resend, 0 means all |

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

Withdraw assets from withdraw wallet. You have to provide a unique **order_id** for each request, we will send the callback with the unique **order_id** when the withdraw is success (from in pool to in chain). 

**POST** /v1/sofa/wallets/`WALLET_ID`/sender/transactions

> Wallet ID must be a withdraw wallet's ID
> 
> order\_id must be prefixed. The prefix is 888888_ in following sample request.
>
>> **Find prefix from corresponding wallet detail on web console UI.**
>
> If withdraw BNB or XRP, this API will check whether the destination addresse needs memo / destination tag or not. If the address does need memo, the withdraw API will fail without memo specified.

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
      "order_id": "888888_1",
      "address": "0x60589A749AAC632e9A830c8aBE042D1899d8Dd15",
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
  ]
}
```

The request includes the following parameters:

###### Post body

| Field | Type  | Required | Description |
| :---  | :---  | :--- | :---        |
| order_id | string | YES | Specify an unique ID, order ID must be prefixed (Up to 255 chars) |
| address | string | YES | Outgoing address |
| amount | string | YES | Withdrawal amount |
| memo | string | NO | Memo on blockchain (This memo will be sent to blockchain). Refer to [Memo Requirement](#memo-requirement) |
| user_id | string | NO | Specify certain user |
| message | string | NO | Message (This message only savced on CYBAVO, not sent to blockchain) |
| block\_average_fee | int | NO | Use avarage blockchain fee within latest N blocks (acceptable value 1~30) |
| manual_fee | int | NO | Specify blockchain fee in smallest unit of wallet currency (acceptable value 1~1000) |

> The order\_id must be prefixed. Find prefix from corresponding wallet detail on web console UI
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

##### [Back to top](#table-of-contents)

[Get withdrawal transaction state](#query-withdrawal-transaction-state)

<a name="query-withdrawal-transaction-state"></a>
## Query withdrawal transaction state

Used to check the withdrawal state.

**GET** /v1/sofa/wallets/`WALLET_ID`/sender/transactions/`ORDER_ID`

- [Sample curl command](#curl-query-withdrawal-transaction-state)

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

##### [Back to top](#table-of-contents)

<a name="query-withdrawal-wallet-balance"></a>
## Query withdrawal wallet balance

Use to get the withdrawal wallet balance.

**GET** /v1/sofa/wallets/`WALLET_ID`/sender/balance

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
| unconfirm_balance | string | Unconfirmed withdrawal wallet balance |
| unconfirm_token_balance | string | Unconfirmed withdrawal wallet token balance |
| err_reason | string | Error message if fail to get balance |

> The currencies that support the unconfirmed balance are BTC, LTC, ETH, BCH, BSV, DASH

##### [Back to top](#table-of-contents)

# Query API

<a name="query-api-token-status"></a>
## Query API token status

Used to check the api token status, you can see if the api token is expired or not. Every time you apply a new api token, you need to call at least one api to activate it, this query api will be useful in this case. 

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

> The un-enabled API-CODE will automatically take effect when it is used for the first time (fill in **X-API-CODE**), and invalidate the currently enabled API-CODE

> If you use the invalid API-CODE query, you will get 403 Forbidden

##### [Back to top](#table-of-contents)

<a name="query-notification-callback-history"></a>
## Query notification callback history

Used to query some kind of callback during a time interval.

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

| Field | Type  | Required | Description |
| :---  | :---  | :--- | :---        |
| from_time | int64 | NO | Start date (unix time in UTC) (default: 0) |
| to_time | int64 | NO | End date (unix time in UTC) (default: current time) |
| type | int | NO | Refer to [Notification Callback Type](#notification-callback-type-definition) (default: -1)  |

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

##### [Back to top](#table-of-contents)


<a name="query-notification-callback-by-id"></a>
## Query notification callback by ID

Used to query if the callback exist or not by id, you can use this api for double confirming if an deposit callback is really existed or not.

**POST** /v1/sofa/wallets/`WALLET_ID`/notifications/get\_by_id

- [Sample curl command](#curl-query-notification-callback-by-id)

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

| Field | Type | Required | Description |
| :---  | :--- | :--- | :---        |
| ids | array | YES | Specify the IDs for query |


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

##### [Back to top](#table-of-contents)



<a name="query-vault/batch-wallet-transaction-history"></a>
## Query vault/batch wallet transaction history

Use to get the wallet’s transaction history.

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

| Field | Type | Required | Description |
| :---  | :--- | :--- | :---        |
| from_item | int64 | NO | Start date (unix time in UTC) (default: 0) |
| to_item | int64 | NO | End date (unix time in UTC) (default: current time) |
| start_index | int | NO | Index of starting transaction record (default: 0) |
| request_number | int | NO | Count of returning transaction record (default: 1000, max: 5000) |
| state | int | NO | Refer to [Transaction State Filter](#transaction-state-filter-definition) (default: -1) |

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

##### [Back to top](#table-of-contents)

<a name="query-wallet-block-info"></a>
## Query wallet block info

Use to get the wallet’s syncing block information.

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

The fake deposit will make the address be added to invalid deposit address. Used this api to get the invalid list for further usage.

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

Use to get wallet basic information.

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

##### [Back to top](#table-of-contents)

<a name="verify-addresses"></a>
## Verify addresses

Check if the addresses are well-formatted in this cryptocurrency, Ex. ETH must have the prefix 0x, BTC should be started with 1 or 3 or bc1, etc.

> If the wallet's cryptocurrency is BNB or XRP, there will be a `must_need_memo` flag to indicate whether the address needs a memo / destination tag when transferring cryptocurrency to the address.

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

| Field | Type  | Requried | Description |
| :---  | :---  | :--- | :---        |
| addresses | array | YES | Specify the address for verification |

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
| must\_need\_memo | boolean | Indicate whether the address needs a memo / destination tag when transferring cryptocurrency to the address |

##### [Back to top](#table-of-contents)


<a name="query-wallet-transaction-autofee"></a>
## Query wallet transaction avarage blockchain fee

Query avarage blockchain fee within latest N blocks.

**POST** /v1/sofa/wallets/`WALLET_ID`/autofee

- [Sample curl command](#curl-query-wallet-transaction-autofee)

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

| Field | Type  | Requried | Description |
| :---  | :---  | :--- | :---        |
| block_num | int | NO | Get avarage blockchain fee within latest N blocks (acceptable value 1~30)(default: 1) |

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
>	Set the backend server URL to the following configuration in mockserver.app.conf

```
api_server_url="BACKEND_SERVER_URL"
```

### Put wallet API code/secret into mock server
-	Get API code/secret on web console
	-	API-CODE, API-SECRET, WALLET-ID
- 	Put API code/secret to mock server's database

```
curl -X POST -d '{"api_code":"API-CODE","api_secret":"API-SECRET"}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/apitoken
```

### Register mock server callback URL
>	Operate on web admin console

Notification Callback URL

```
http://localhost:8889/v1/mock/wallets/callback
```

Withdrawal Authentication Callback URL

```
http://localhost:8889/v1/mock/wallets/withdrawal/callback
```

> The withdrawal authentication callback URL once set, every withrawal request will callback this URL to get authentication to proceed withdrawal request.
> 
> Refer to *WithdrawalCallback()* function in mock server OuterController.go

##### [Back to top](#table-of-contents)

<a name="curl-testing-commands"></a>
## CURL Testing Commands

<a name="curl-create-deposit-wallet-addresses"></a>
### Create deposit wallet addresses

For BNB or EOS wallet:

```
curl -X POST -H "Content-Type: application/json" -d '{"count":2,"memos":["001","002"]}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/addresses
```

For wallet excepts BNB or EOS:

```
curl -X POST -H "Content-Type: application/json" -d '{"count":2}' \
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
curl -X POST -H "Content-Type: application/json" -d '{"notification_id":0}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/callback/resend
```
- [API definition](#resend-pending-or-failed-deposit-callbacks)

<a name="curl-withdraw"></a>
### Withdraw

```
curl -X POST -H "Content-Type: application/json" -d '{"requests":[{"order_id":"888888_1","address":"0x60589A749AAC632e9A830c8aBE042D1899d8Dd15","amount":"0.0001","memo":"memo-001","user_id":"USER01","message":"message-001"},{"order_id":"888888_2","address":"0xf16B7B8900F0d2f682e0FFe207a553F52B6C7015","amount":"0.0002","memo":"memo-002","user_id":"USER01","message":"message-002"}]}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/withdraw
```
- [API definition](#withdraw)

<a name="curl-query-withdrawal-transaction-state"></a>
### Query withdrawal transaction state

```
curl -X GET http://localhost:8889/v1/mock/wallets/{WALLET-ID}/sender/transactions/{ORDER_ID}
```
- [API definition](#query-withdrawal-transaction-state)

<a name="curl-query-withdrawal-wallet-balance"></a>
### curl-query-withdrawal-wallet-balance

```
curl -X GET http://localhost:8889/v1/mock/wallets/{WALLET-ID}/sender/balance
```
- [API definition](#query-withdrawal-wallet-balance)

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

<a name="curl-query-notification-callback-by-id"></a>
### Query notification callback by ID

```
curl -X POST -H "Content-Type: application/json" -d '{"ids":[90000000140,90000000139]}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/notifications/get_by_id
```
- [API definition](#query-notification-callback-by-id)

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

<a name="curl-verify-addresses"></a>
### Verify addresses

```
curl -X POST -H "Content-Type: application/json" -d '{"addresses":["0x635B4764D1939DfAcD3a8014726159abC277BecC","1CK6KHY6MHgYvmRQ4PAafKYDrg1ejbH1cE"]}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/addresses/verify
```

- [API definition](#verify-addresses)

<a name="curl-query-wallet-transaction-autofee"></a>
### Query wallet transaction avarage blockchain fee

```
curl -X POST -H "Content-Type: application/json" -d '{"block_num":1}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/autofee
```

- [API definition](#verify-addresses)


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
    <td>The extra information of this callback</td>
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
| 1 | Deposit Callback |
| 2 | Withdraw Callback |
| 3 | Collect Callback |
| 4 | Airdrop Callback |
| -1 | All callbacks (for inquiry) |

##### [Back to top](#table-of-contents)

<a name="transaction-state-filter-definition"></a>
## Transaction State Filter Definition

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

##### [Back to top](#table-of-contents)

<a name="currency-definition"></a>
## Currency Definition

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
| 714| BNB |
| 1815| ADA |
  
> Refer to [here](https://github.com/satoshilabs/slips/blob/master/slip-0044.md) for more detailed currency definitions
 
##### [Back to top](#table-of-contents)

<a name="memo-requirement"></a>
## Memo Requirement

| Currency | Description |
| :--- | :---        |
| XRP | Up to 20 digits |
| XLM | Up to 28 bytes of ASCII/UTF-8 |
| EOS | Up to 256 chars |
| BNB | Up to 128 chars |

##### [Back to top](#table-of-contents)

