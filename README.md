# KaliCoin specifications

![alt text](kalicoin.png "Kalicoin Logo")

v1.0.1 - 20/03/2019

## Summary

Kalicoin (kc) is de voorgestelde in-chat munteenheid. Deze kan verdiend worden, uitgegeven, en geruild worden.

Het hoofddoel is een in-chat economie op te zetten zodat het spammen van bepaalde commando's kan verminderd worden, alsook dat users mekaar kunnen belonen wanneer hulp wordt verstrekt aan elkaar.
Zogenaamde "IRL" transacties kunnen hiermee ook gemaakt worden, bv het ruilen van een afgesproken aantal kc tegen een 3D geprint object.

## Developing

Build

```bash
go build kalicoin
```

Test

```bash
$ docker run -e POSTGRES_DB=kalicoin_test -e POSTGRES_USER=user -e POSTGRES_PASSWORD=pass -p5432:5432 -it postgres:11

# Other terminal
$ DATABASE_URI="postgres://user:pass@localhost:5432/kalicoin_test?sslmode=disable" go test -mod vendor ./...
?       kalicoin                [no test files]
?       kalicoin/pkg/api        [no test files]
?       kalicoin/pkg/db         [no test files]
ok      kalicoin/pkg/models     0.605s
```

### Configuration

`DATABASE_URI`

This is the URI used for connecting to the database

## Non Functional Requirements

- Iedereen's wallet (hoeveelheid kc in bezit) moet opgeslagen worden als een record in een tabel
- Alle transacties moeten gelogged worden
  - datum transactie
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

_slotmachine-style gokken met kc, beslist door RNG.
Zou zich kunnen gedragen zoals de slots in pokemon waar hogere inzet hogere win multipliers kan opleveren.
Zie: https://bulbapedia.bulbagarden.net/wiki/Slot_machine#Payouts_

`usage: /roulette amount`

- `amount`: hoeveelheid uitgedrukt in kc

## Voorgestelde premium commando's

_Coins moeten zowel kunnen verdiend als uitgegeven worden om inflatie of devaluatie binnen de perken te houden_

### /{img,vod,audio}quote

- Nieuwe quotes maken zou geld kunnen kosten, bv 10kc.
- Quotes opvragen blijft gratis, tenzij `/presidential_quote` kan ook geld kosten.

## Coins verdienen

### Check-in

_Users die op de dagelijkse check-in reageren kunnen beloond worden met een mooi bedrag kc. bv 1000 kc.
420 and 1337 zouden kunnen terug gebracht worden zonder leaderbord, als extra check-in momenten. Deze zouden wel minder kc opleveren dan de random. (50kc bv.)_

## Coins uitgeven

Meerdere manieren om coins te spenderen (in de vorm van premium commando's) zouden moeten worden bedacht, om devaluatie tegen te gaan.

_In de toekomst zou bv doomba remote controle via telegram commando's via kalicoins kunnen werken_

## Voorgestelde prijzen

- random check-in: +100kc
- 420/1337: +20kc
- /all: -100kc
- /quote (maken): -5kc
- /presidential_quote: -5kc
- /img\*: -5kc

## Startersbudget

_Iedereen zou kunnen beginnen met 100kc, zodat we niet meteen zonder vallen, maar het toch interessant is om er meteen te beginnen verdienen._
