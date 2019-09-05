# Deposit Wallet API

#### Create deposit wallet addresses
- **POST** /v1/sofa/wallets/`WALLET_ID`/addresses
	- Request
		-	Params
			- 	`count`: Specify address count, max value is 1000.
			-  `memos`: Specify memos for BNB or EOS deposit wallet
				-	**NOTE: The length of `memos` must equal to `count` while creating addresses for BNB or EOS wallet**
		-  Sample:
	
			For BNB or EOS wallet:
			
			```
			{
			  "count": 2,
			  "memos": [
			    "001",
			    "002"
			  ]
			}
			```
			
			For wallet excepts BNB and EOS:
			
			```
			{ "count": 2 }
			```

	- Response
		-	Params
			-	`addresses`: Array of just created deposit addresses
		-	Sample:

		For BNB or EOS wallet:
		
		```
		{
		  "addresses": [
		    "002",
		    "001"
		  ]
		}
		```
		
		For wallet excepts BNB and EOS:
		
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
		--- AND ---
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
			-  `memo`: Memo on blockchain (This memo will be sent to blockchain)
			-  `user_id`: Specify certain user (optional)
			-  `message`: Message (This message only savced on CYBAVO, not sent to blockchain
)
		-  Sample:

		```
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
			-  `valid`: Activated API token
			-  `inactivated `: Not active API token
			-	`api_code `: API token for querying wallet
			-	`exp`: API token expiration unix time in UTC
		-	Sample:

		```
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
	PS: 1. 未啟用的API-CODE會在第一次使用(填在**X-API-CODE**)時自動生效，並讓當前已啟用的API-CODE失效
	    2. 若使用已失效的API-CODE查詢會得到403 Forbidden

#### Query notification callback history

- **GET** /v1/sofa/wallets/`WALLET_ID`/notifications?from\_time=`from`&to\_time=`to`&type=`type`
	- Request
		-	Params
			- 	`from`: Start date (unix time in UTC)
			-  `to`: End date (unix time in UTC)
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


#### Query vault/batch wallet transaction history

- **GET** /v1/sofa/wallets/`WALLET_ID`/transactions?from\_time=`from`&to\_time=`to`&start_index=`start`&request_number=`count`&state=`state`
	- Request
		-	Params
			- 	`from`: Start date
			-  `to`: End date
			-  `start`: Index of starting transaction record
			-  `count`: Count of returning transaction record
			-  `state`: Transaction state filter (optional, default: -1)
				-	-1: All states
				-	0: WaitApproval
				- 	1: Rejected
				-  2: Approved
				-  3: Failed
				-  4: NextLevel
				-  5: Cancelled
				-  6: BatchDone 
		-  Sample:
		
		```
		/v1/sofa/wallets/48/transactions?from_time=1559664000&to_time=1562255999&start_index=0&request_number=3
		```

	- Response
		-	Params
			-	`transaction_count `: Total transactions in specified date duration
			-	`transaction_item`: Array of transaction record
		-	Sample:

		```
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

#### Query wallet block info

- **GET** /v1/sofa/wallets/`WALLET_ID`/blocks

	- Response
		-	Params
			-	`latest_block_height `: latest block height on blockchain
			-	`synced_block_height`: current synced block height
		-	Sample:

		```
		{
		  "latest_block_height": 29317651,
		  "synced_block_height": 28529203
		}
		```
		
#### Query invalid deposit addresses

- **GET** /v1/sofa/wallets/`WALLET_ID`/addresses/invalid-deposit

	- Response
		-	Params
			-	`addresses `: array of invalid deposit address
		-	Sample:

		```
		{
		  "addresses": ["0x5dB3d8C70dAa9C919F9962221c2fDDbe8EBAa5F2"]
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
		-	`state`:
			-	0: Enqueue
			-	1: Processing batch in KMS
			-	2: TXID in pool
			-	3: TXID in chain
			-	**(DEPRECATED)** 4: TXID confirmed in N blocks
			-	5: Failed
			-	6: Resent
			-	7: Blocked due to risk controlled
			-	8: Cancelled
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
  "confirm_blocks": 1,
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
curl -X POST -d '{"requests":[{"order_id":"1","address":"0x60589A749AAC632e9A830c8aBE042D1899d8Dd15","amount":"0.0001","memo":"memo-001","user_id":"USER01","message":"message-001"},{"order_id":"2","address":"0xf16B7B8900F0d2f682e0FFe207a553F52B6C7015","amount":"0.0002","memo":"memo-002","user_id":"USER01","message":"message-002"}]}' \
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

### Query vault/batch wallet transaction history

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/transactions?start_index=0&from_time=1559664000&to_time=1562255999&request_number=8'
```

### Query deposit/withdraw wallet block info

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/blocks'
```

#### Query invalid deposit addresses

```
curl -X GET 'http://localhost:8889/v1/mock/wallets/{WALLET-ID}/addresses/invalid-deposit'
```