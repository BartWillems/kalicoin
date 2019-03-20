# KaliCoin specifications

v1.0.1 - 20/03/2019

## Summary

Kalicoin (kc) is de voorgestelde in-chat munteenheid. Deze kan verdiend worden, uitgegeven, en geruild worden.

Het hoofddoel is een in-chat economie op te zetten zodat het spammen van bepaalde commando's kan verminderd worden, alsook dat users mekaar kunnen belonen wanneer hulp wordt verstrekt aan elkaar.
Zogenaamde "IRL" transacties kunnen hiermee ook gemaakt worden, bv het ruilen van een afgesproken aantal kc tegen een 3D geprint object.

## Configuration

`DATABASE_URI`

This is the URI used for connecting to the database

`KALI_ID`

This is the id of the telegram room containing the kali chat

## Non Functional Requirements

- Iedereen's wallet (hoeveelheid kc in bezit) moet opgeslagen worden als een record in een tabel
- Alle transacties moeten gelogged worden
  - datum transactie
  - sender
  - receiver
  - amount

## Voorgestelde nieuwe commando's

### /wallet

_Krijg eigen kc balans terug_

`usage: /wallet`

### /pay

_schrijf geld over naar een andere gebruiker in chat_

`usage: /pay @user amount`

- `@user`: tag van gebruiker
- `amount`: hoeveelheid uitgedrukt in kc

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
