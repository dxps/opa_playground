#!/bin/sh

## Info: Passing ?explain=full&pretty will give in the "explanation" response key
## those valuable (for troubleshooting) `Note` entries representing `trace(_)` output.
curl -X POST 'localhost:8181/v1/data/products_access/dashboard?explain=full&pretty' \
     -d@subject_has_access_to_product__query_input.json
