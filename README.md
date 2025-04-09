# Vault

## Setup

```sh
export VAULT_PRIVATE_KEY_PATH="id_rsa"
```

## Add new command

```sh
cobra-cli add timezone
```

## Run command

```sh
vault timezone EST
```

## Clone command

```sh
cobra-cli add mfa
cobra-cli add getPasswd

vault mfa <encrypted data>
vault getPasswd <encrypted data>
```

## Build

```sh
go build -o dist/vault main.go
```
