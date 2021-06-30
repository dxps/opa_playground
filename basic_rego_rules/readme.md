## Basic Rego Rules

This illustrates the basic capabilities of creating and testing Rego rules.

Note that `input.json` file is needed if you want to use `OPA: Evaluate Package` option of VSCode OPA plugin.

This setup is based on _CarInfoStore_ sample, part of Styra's OPA Policy Authoring course.

<br/>

### Usage

#### Using VSCode

Using VSCode, having [OPA extension](https://marketplace.visualstudio.com/items?itemName=tsandall.opa) installed, you can just open either the policy file (`basic_rego_rules___policy.rego`) or test file (`test_basic_rego_rules___policy.rego`) and run the VSCode command _OPA: Evaluate package_.

The result should look as follows:

```json
[
  [
    {
      "allow": false,
      "test___car_create___anyone": true,
      "test___car_create___employee": true,
      "test___car_create___manager": true,
      "test___car_status___anyone": true,
      "test___car_status___employee": true,
      "test___car_status___manager": true,
      "test___car_update___anyone": true,
      "test___car_update___employee": true,
      "test___car_update___manager": true,
      "test_car_read_negative": true,
      "test_car_read_positive": true,
      "user_is_employee": false,
      "users": {
        "alice": {
          "manager": "charlie",
          "title": "salesperson"
        },
        "bob": {
          "manager": "charlie",
          "title": "salesperson"
        },
        "charlie": {
          "manager": "dave",
          "title": "manager"
        },
        "dave": {
          "manager": null,
          "title": "ceo"
        }
      }
    }
  ]
]
```

#### Using local OPA

As OPA can load policies (and data) as bundles (see details [here](https://www.openpolicyagent.org/docs/latest/management-bundles/)), the current policy has been packaged as a gzipped tarbal and it can be exposed through HTTP, and instruct OPA to fetch it at startup.

[http-server](https://www.npmjs.com/package/http-server) is the choice here to expose and serve this bundle.

Steps:

1. Run `./run_http.sh` to start the local HTTP server that serves the bundle.
1. Run `./run_opa.sh` to start OPA.

Then using OPA's REST API, you can query for some authz decisions using something like:

```shell
❯ curl localhost:8181/v1/data/rules/allow -d '{"input": {"method": "POST", "path": ["cars"], "user": "dave"}}'
{"result":true}
❯
```
