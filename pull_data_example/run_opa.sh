#!/bin/sh

## Either download OPA (from https://www.openpolicyagent.org/docs/latest/#1-download-opa)
## here and use ./opa run --server OR just include it in your PATH to be used as below.

## Info: To get the decision reasoning, you can use: opa run -s --set decision_logs.console=true
# opa run -s --set decision_logs.console=true

opa run --server 

