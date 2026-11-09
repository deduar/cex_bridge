#!/usr/bin/env bash
set -e
BASE=http://localhost:8080
curl -s -X POST $BASE/register -H 'Content-Type: application/json' -d '{"username":"alice"}' >/dev/null
curl -s -X POST $BASE/register -H 'Content-Type: application/json' -d '{"username":"bob"}' >/dev/null
TOK=$(curl -s -X POST $BASE/login -H 'Content-Type: application/json' -d '{"username":"alice"}' | sed -n 's/.*"token":"\([^"]*\)".*/\1/p')
echo "alice token: $TOK"
echo "seed done (balances are ledger-side; demo flow uses /orders)"
