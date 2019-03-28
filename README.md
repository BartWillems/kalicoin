# KaliCoin specifications

[![pipeline status](https://gitlab.com/bartwillems/kalicoin/badges/master/pipeline.svg)](https://gitlab.com/bartwillems/kalicoin/commits/master)

![alt text](kalicoin.png "Kalicoin Logo")

Current status: pre-alpha

## Summary

Kalicoin (kc) is the proposed in-chat currency. This can be earned, spent, and exchanged.

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

`DATABASE_URI`

This is the URI used for connecting to the database

## Non Functional Requirements

- Everyone's wallet (amount of kc in possession) must be stored as a record in a table
- All transactions must be logged
   - transaction date
   - sender
   - receiver
   - amount

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

### /transactions

#### POST Transaction{}

    Creates a new transaction and returns the transaction with the resulting status

Request

```bash
$ curl -X POST \
  http://localhost:8000/transactions \
  -H 'Content-Type: application/json' \
  -d '{
        "type": "trade",
        "sender": 69,
        "receiver": 420,
        "amount": 10
    }'
```

Result:

```json
{
  "id": 14,
  "type": "trade",
  "status": "succeeded",
  "sender": 69,
  "receiver": 420,
  "amount": 10,
  "failure_reason": "",
  "created_at": "2019-03-22T18:40:24.128529274Z",
  "updated_at": "2019-03-22T18:40:24.150446326Z"
}
```

#### GET

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

_Coins must be able to both be earned and spent to keep inflation or devaluation under control_

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
