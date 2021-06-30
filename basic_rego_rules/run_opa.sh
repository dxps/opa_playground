#!/bin/sh

## Note: If you don't have OPA already on your PATH env var, go to
##       https://www.openpolicyagent.org/docs/latest/#running-opa

opa run -s -c ./opa_config.yml

