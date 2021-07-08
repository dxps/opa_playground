#!/bin/sh

curl -X PUT localhost:8181/v1/policies/products_access/dashboard \
     --data-binary @subject_has_access_to_product__rule.rego
