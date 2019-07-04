# Deposit Wallet API

#### Create deposit wallet addresses
- **POST** /v1/sofa/wallets/`WALLET_ID`/addresses
	- Request
		-	Params
			- 	`count`: Specify address count, max value is 1000.
		-  Sample:
	
		```
		{ "count": 3 }
		```

	- Response
		-	Params
			-	`addresses`: Array of just created deposit addresses
		-	Sample:

		```
		{
		  "addresses": [
		    "0x2E7248BBCD61Ad7C33EA183A85B1856bc02C40b6",
		    "0x4EB990D527c96c64eC5Bfb0D1e304840052d4975",
		    "0x86434604FF857702fbE11cBFf5aC7689Af19c4Ed"
		  ]
		}
		```

#### Query address of deposit wallet

- **GET** /v1/sofa/wallets/`WALLET_ID`/addresses?start\_index=`from`&request\_number=`count`
	- Request
		-	Params
			- 	`from`: Specify address start index
			-  `count`: Request address count
		-  Sample:
		
		```
		/v1/sofa/wallets/17/addresses??start_index=0&request_number=3
		--- AND ---
		/v1/sofa/wallets/17/addresses??start_index=3&request_number=3
		```

	- Response
		-	Params
			-	`wallet_id `: ID of request wallet
			- 	`wallet_address`: Array of wallet addresses
		-	Sample:

		```
		{
		  "wallet_id": 17,
		  "wallet_address": [
		    {
		      "currency": 60,
		      "token_address": "",
		      "address": "0x8c42fD03A5cfba7C3Cd97AB8a09e1a3137Ef33C3"
		    },
		    {
		      "currency": 60,
		      "token_address": "",
		      "address": "0x4d3EB54b602BF4985CE457089F9fB084Af597A2C"
		    },
		    {
		      "currency": 60,
		      "token_address": "",
		      "address": "0x74dc3fB523295C87C0b93E48744Ce94fe3a8Ef5e"
		    }
		  ]
		}
		--- AND ---
		{
		  "wallet_id": 17,
		  "wallet_address": [
		    {
		      "currency": 60,
		      "token_address": "",
		      "address": "0x6d68443D6564cF257A48c1b16aa6d0EF13c5A719"
		    },
		    {
		      "currency": 60,
		      "token_address": "",
		      "address": "0x26F103322B6f0ed2D35B85F1611589c92F023986"
		    },
		    {
		      "currency": 60,
		      "token_address": "",
		      "address": "0x2b91918Bee4411DaD6293EA5d6D38251E72723Ca"
		    }
		  ]
		}
		```

#### Resend pending or failed deposit callbacks

- **POST** /v1/sofa/wallets/`WALLET_ID`/collection/notifications/manual
	-	Request
		-	Params
			- 	`notification_id `: Specify callback ID to resend, 0 means all
				-  This ID equal to callback data's serial/order_id
		-  Sample:

		```
		{ "notification_id": 0 }
		```
	
	- Response
		-	Params
			-	`count `: Count of callbacks just resent
		-	Sample:

		```
		{ "count": 0 }
		```

# Withdraw Wallet API

#### Withdraw

- **POST** /v1/sofa/wallets/`WALLET_ID`/sender/transactions
	-	Request
		-	Params
			- 	`order_id `: Specify an unique ID
			-  `address`: Outgoing address
			-  `amount`: Withdrawal amount
			-  `memo`: Memo on blockchain
		-  Sample:

		```
		{
		  "requests": [
		    {
		      "order_id": "1",
		      "address": "0x60589A749AAC632e9A830c8aBE042D1899d8Dd15",
		      "amount": "0.0001",
		      "memo": "WHATEVER_YOU_WANT"
		    },
		    {
		      "order_id": "2",
		      "address": "0xf16B7B8900F0d2f682e0FFe207a553F52B6C7015",
		      "amount": "0.0002",
		      "memo": "WHATEVER_YOU_WANT"
		    }
		  ]
		}
		```
	
	- Response
		-	Params
			-	`results `: Array of withdraw result (order ID/withdraw transaction ID pair), if succeeds
		-	Sample:

		```
		{
		  "results": {
		    "1": 20000000001,
		    "2": 20000000002
		  }
		}
		```
		
# Query API

#### Query API token status

- **GET** /v1/sofa/wallets/`WALLET_ID`/apisecret

	- Response
		-	Params
			-	`api_code `: API token for querying wallet
			-	`exp`: API token expiration unix time in UTC
		-	Sample:

		```
		{
		  "api_code": "2cVhTyZrkctoEiUQa",
		  "exp": 1564801781
		}
		```

#### Query notification callback history

- **GET** /v1/sofa/wallets/`WALLET_ID`/notifications?from\_time=`from`&to\_time=`to`&type=`type`
	- Request
		-	Params
			- 	`from_time`: Specify address start index
			-  `to_time`: Request address count
			-  `type`: Notification callback type [1|2|3]
				-	1: Deposit Callback (入金回調)
				-	2: Withdraw Callback (出金回調)
				- 	3: Collect Callback (歸帳回調)
		-  Sample:
		
		```
		/v1/sofa/wallets/67/notifications?from_time=1561651200&to_time=1562255999&type=2
		```

	- Response
		-	Params
			-	`notifications `: Arrary of callbacks
		-	Sample:

		```
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



# Mock Server

### Set API server URL
-	mockserver.app.conf

```
api_server_url=""
```

### Register mock server callback URL
-	Operate on web console

```
http://localhost:8889/v1/mock/wallets/callback
```

-	callback structure
	-	Params
		-	`type `: callback type [1|2|3]
			-	1: Deposit Callback (入金回調)
			-	2: Withdraw Callback (出金回調)
			- 	3: Collect Callback (歸帳回調)
	-	Sample:

```
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
  "addon": {}
}
```

### Query API token

-	Operate on web console
	-	API-CODE, API-SECRET, WALLET-ID
- 	Save API token to database

```
curl -X POST -d '{"api_code":"API-CODE","api_secret":"API-SECRET"}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/apitoken
```

### Create deposit wallet addresses

```
curl -X POST -d '{"count":10}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/addresses
```

### Get deposit wallet addresses

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/addresses?start_index=0&request_number=1000'
```

### Get deposit wallet pool address, only for USDT-omni 

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/pooladdress'
```

### Resend all pending or failed deposit callbacks

```
curl -X POST -d '{"notification_id":0}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/callback/resend
```

### Withdraw

```
curl -X POST -d '{"requests":[{"order_id":"1","address":"0x60589A749AAC632e9A830c8aBE042D1899d8Dd15","amount":"0.0001","memo":"WHATEVER_YOU_WANT"},{"order_id":"2","address":"0xf16B7B8900F0d2f682e0FFe207a553F52B6C7015","amount":"0.0002","memo":"WHATEVER_YOU_WANT"}]}' \
http://localhost:8889/v1/mock/wallets/{WALLET-ID}/withdraw
```

### Query API token status

```
curl -X GET http://localhost:8889/v1/mock/wallets/{WALLET-ID}/apisecret
```

### Query notification callback history

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/notifications?from_time=1561651200&to_time=1562255999&type=2'
```