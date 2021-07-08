#!/bin/sh

# Info: Passing ?explain=full will give lots of Trace Events.
# curl -X POST localhost:8181/v1/data/products_access/dashboard?explain=full&pretty \
curl -X POST localhost:8181/v1/data/products_access/dashboard \
     -d@subject_has_access_to_product__query_input.json