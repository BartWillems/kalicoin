# KaliCoin specifications

[![pipeline status](https://gitlab.com/bartwillems/kalicoin/badges/master/pipeline.svg)](https://gitlab.com/bartwillems/kalicoin/commits/master)

![alt text](kalicoin.png "Kalicoin Logo")

Current status: public-beta

## Summary

Kalicoin (kc) is the proposed in-chat currency. This can be earned, spent, and exchanged.
This project is in no way affiliated with [kalicoin.io](https://kalicoin.io/).

The main objective is to set up an in-chat economy so that the spamming of certain commands can be reduced, and that users can reward each other when assistance is provided to each other.
So-called "IRL" transactions can also be made with this, for example exchanging an agreed number of kc for a 3D printed object.

## Developing

Build

```bash
go build kalicoin
```

Test

```bash
$ make test

/bin/bash -c "go test -mod vendor ./pkg/api && go test -mod vendor ./pkg/models"
ok      kalicoin/pkg/api        0.082s
ok      kalicoin/pkg/models     0.188s
docker stop postgres-kalicoin
postgres-kalicoin
```

### Configuration

| ENV                 | Description                               | Example                                                        |
| ------------------- | ----------------------------------------- | -------------------------------------------------------------- |
| DATABASE_URI        | URI Used for connecting to the database   | `postgres://user:pass@127.0.0.1:5432/kalicoin?sslmode=disable` |
| JAEGER_AGENT_HOST   | Jaeger host for tracing purposes          | `jaeger`                                                       |
| JAEGER_AGENT_PORT   | UDP Port to send the traces to            | `6831`                                                         |
| JAEGER_SERVICE_NAME | Service name that will be shown on jaeger | `kalicoin`                                                     |
| AUTH_USERNAME       | Basic Auth username for API access        | `octaaf`                                                       |
| AUTH_PASSWORD       | Basic Auth password for API access        | `secret`                                                       |
| API_PORT            | Port for the kalicoin api to bind on      | `:8000` or `0.0.0.0:8000` or `127.0.0.1:8000`                  |

## Non Functional Requirements

- Everyone's wallet (amount of kc in possession) must be stored as a record in a table
- All transactions must be logged
  - transaction date
  - update date (is set when succeeded/failed)
  - sender
  - receiver
  - amount
  - type
  - failure reason (if present)

## API

### /wallets

#### GET

    Returns an array of wallets

```json
[
  {
    "id": 3,
    "owner_id": 69,
    "capital": 40,
    "created_at": "2019-03-21T18:14:22.915032Z",
    "updated_at": "2019-03-22T18:40:24.148301Z"
  },
  {
    "id": 4,
    "owner_id": 420,
    "capital": 150,
    "created_at": "2019-03-21T18:14:22.953754Z",
    "updated_at": "2019-03-22T18:40:24.149513Z"
  }
]
```

### /{trades, rewards, payments}

#### POST Transaction{}

    Creates a new transaction and returns the transaction with the resulting status

Trades Request

```bash
$ curl -X POST \
  http://localhost:8000/trades \
  -H 'Authorization: Basic b2N0YWFmOnNlY3JldA==' \
  -H 'Content-Type: application/json' \
  -d '{
        "sender": 69,
        "receiver": 420,
        "amount": 1000
    }'
```

Result:

```json
{
  "id": 2,
  "type": "trade",
  "status": "failed",
  "sender": 10,
  "receiver": 69,
  "amount": 1000,
  "cause": null,
  "failure_reason": "Not enough money in your wallet",
  "created_at": "2019-03-22T18:40:24.128529274Z",
  "updated_at": "2019-03-22T18:40:24.150446326Z"
}
```

Reward Request

```bash
$ curl -X POST \
  http://localhost:8000/rewards \
  -H 'Authorization: Basic b2N0YWFmOnNlY3JldA==' \
  -H 'Content-Type: application/json' \
  -d '{
        "receiver": 420,
        "cause": "kalivent"
    }'
```

Result:

```json
{
  "id": 1,
  "type": "reward",
  "status": "succeeded",
  "sender": null,
  "receiver": 20,
  "amount": 20,
  "cause": "kalivent",
  "failure_reason": null,
  "created_at": "2019-03-22T18:40:24.128529274Z",
  "updated_at": "2019-03-22T18:40:24.150446326Z"
}
```

Payment Request

```bash
$ curl -X POST \
  http://localhost:8000/payments \
  -H 'Authorization: Basic b2N0YWFmOnNlY3JldA==' \
  -H 'Content-Type: application/json' \
  -d '{
    "sender": 69,
    "cause": "quote"
  }'
```

Result:

```json
{
  "id": 3,
  "type": "payment",
  "status": "succeeded",
  "sender": 69,
  "receiver": null,
  "amount": 10,
  "cause": "quote",
  "failure_reason": null,
  "created_at": "2019-03-22T18:40:24.128529274Z",
  "updated_at": "2019-03-22T18:40:24.150446326Z"
}
```

#### GET /transactions

    Returns the transactions

```json
[
  {
    "id": 8,
    "type": "trade",
    "status": "succeeded",
    "sender": 69,
    "receiver": 420,
    "amount": 10,
    "failure_reason": "",
    "created_at": "2019-03-21T13:14:22.933576Z",
    "updated_at": "2019-03-21T13:14:22.955853Z"
  },
  {
    "id": 9,
    "type": "trade",
    "status": "succeeded",
    "sender": 69,
    "receiver": 420,
    "amount": 10,
    "failure_reason": "",
    "created_at": "2019-03-21T13:14:29.676697Z",
    "updated_at": "2019-03-21T13:14:29.694235Z"
  },
  {
    "id": 10,
    "type": "trade",
    "status": "failed",
    "sender": 69,
    "receiver": 420,
    "amount": 100,
    "failure_reason": "Not enough money in your wallet",
    "created_at": "2019-03-21T13:14:34.568206Z",
    "updated_at": "2019-03-21T13:14:34.585406Z"
  }
]
```

### /roulette

_Slot machine-style gambling with kc, decided by RNG.
Could behave like the slots in pokemon where higher bets can produce higher win multipliers.
See: <https://bulbapedia.bulbagarden.net/wiki/Slot_machine#Payouts>_

`usage: / roulette amount`

- `amount`: quantity expressed in kc

## Proposed Premium Commands

Coins must be able to both be earned and spent to keep inflation or devaluation under control

### /{img,vod,audio}quote

- Creating new quotes could cost money, for example 10kc.
- Quotes retrieval remains free, unless `/ presidential_quote` can also cost money.

## Earning coins

### Check-in

_Users who respond to the daily check-in can be rewarded with a nice amount of kc. e.g. 1000 kc.
420 and 1337 could be returned without a leader board, as extra check-in moments. These would yield less kc than the random. (50kc e.g.)_

## Spending Coins

Multiple ways to spend coins (in the form of premium commands) should be thought of, to prevent devaluation.

_ In the future, for example, doomba remote control via telegram commands could work with kalicoins_

## Proposed initial prices

- random check-in: +100kc
- 420/1337: +20kc
- /all: -100kc
- /quote (maken): -5kc
- /presidential_quote: -5kc
- /img\*: -5kc

## Starting Budget

_Anyone could start with 100kc, so that we do not immediately fall without it, but it is still interesting to start earning immediately._
