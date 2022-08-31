# Buy BTC From Me

This is demo project to demonstrate how to integrate Bitcoin Core with Omise Payment Gateway in Golang backend.

## Omise Keys
You need go through registration process on https://omise.co and optain keys to succesfully run this project.

## Requirements
 - Mysql database (https://www.mysql.com/downloads)
 - Bitcoin Core (https://bitcoincore.org/en/download)
 - Go development environment https://go.dev/

## Environment Variables
Application uses set of environment variables that must be defined before running application
 - `OMISE_PKEY` - Omise Public Key
 - `OMISE_SKEY` - Omise Secret Key
 - `BITCOIN_CLI` - path to Bitcoin-Cli
 - `BITCOIN_CLI_WALLET_PASS` - password to Bitcoin Wallet

## Installation
  - Clone repository
  - Run `go get -d ./...` to download all dependencies
  - Start application by executing `go run cmd/web/main.go cmd/web/middleware.go cmd/web/routes.go`

### TODO
 - Write `Docker` scripts to install and run application
 - Cover Application with more unit tests
