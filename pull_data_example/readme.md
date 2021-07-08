## Pull Data Example

This example showcases the 5th option of OPA, [Pull Data during Evaluation](https://www.openpolicyagent.org/docs/latest/external-data/#option-5-pull-data-during-evaluation), a nice feature that can simply be based on the built-in [http.send](https://www.openpolicyagent.org/docs/latest/policy-reference/#http) function.

<br/>

### Usage

Use the included scripts to start OPA, upload the policy and query for the decision.

1. Start OPA in server mode using `./run_opa.sh`.
1. Upload the policy using `./subject_has_access_to_product__upload_rule.sh`.
1. Query for decision using `./subject_has_access_to_product__query.sh`.

While evaluating the referred policy during the query processing, it should pull additional required data from an external endpoint, provided by the included `iam_svc` service.

<br/>
