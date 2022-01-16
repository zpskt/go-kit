#!/bin/bash
curl \
      --request PUT \
      --data @hello.json \
      localhost:8500/v1/agent/service/register
