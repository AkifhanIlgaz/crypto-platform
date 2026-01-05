#!/bin/sh

set -e

vault login "${VAULT_DEV_ROOT_TOKEN_ID}"

# Verileri .env.example'dan al
vault kv put secret/database/postgres \
  username="${DB_USER}" \
  password="${DB_PASSWORD}" \
  database="${DB_NAME}" \
  host="${DB_HOST}" \
  port="${DB_PORT}" \
  ssl_mode="${DB_SSL_MODE}"

vault kv put secret/exchange/kucoin \
    api_key="${KUCOIN_API_KEY}" \
    api_secret="${KUCOIN_API_SECRET}" \
    passphrase="${KUCOIN_PASSPHRASE}"

vault kv put secret/exchange/binance \
    api_key="${BINANCE_API_KEY}" \
    api_secret="${BINANCE_API_SECRET}"


echo "Baslangic verileri yüklendi!"
